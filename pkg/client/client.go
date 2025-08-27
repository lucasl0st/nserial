package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	"go.bug.st/serial"

	"github.com/lucasl0st/nserial/pkg/model"
	"github.com/lucasl0st/nserial/pkg/protocol"
)

var SerialMode = &serial.Mode{
	BaudRate: 9600,
	DataBits: 8,
	Parity:   serial.EvenParity,
	StopBits: serial.TwoStopBits,
}

var ErrWroteIncompleteData = errors.New("wrote incomplete data")

// RollConsumer consumes a Roll
type RollConsumer func(roll model.Roll) error

type Client interface {
	io.Closer

	// Connect initiate a new connection with the camera
	// must be run before performing other actions
	Connect() error

	// GetData returns all data about stored films and shots saved in the camera
	// consume rolls one by one using the RollConsumer
	// unmodified/unparsed data is returned in the end
	GetData(consumer RollConsumer) ([]byte, error)
}

type client struct {
	conn serial.Port
	lock sync.Mutex
}

type Opt func(c *client) error

// WithPort specify a port using a path
// /dev/ttyX on linux
// COM1 etc on windows
func WithPort(port string) Opt {
	return func(c *client) error {
		s, err := serial.Open(port, SerialMode)
		if err != nil {
			return err
		}

		c.conn = s
		return nil
	}
}

// WithSerialPort specify a port directly from the serial library
func WithSerialPort(port serial.Port) Opt {
	return func(c *client) error {
		c.conn = port
		return nil
	}
}

// New returns a new Client with specified options
func New(opts ...Opt) (Client, error) {
	c := &client{
		lock: sync.Mutex{},
	}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *client) Close() error {
	return c.conn.Close()
}

// warning: I have no idea if these sequences actually perform the things I described them with,
// I threw around bytes until I got something working
var (
	resetSequence          = bytes.Repeat([]byte{0x6f}, 8)
	getCameraModelSequence = []byte{0x00, 0xff, 0x41, 0x00, 0x04, 0x01, 0x45}
	getDataSequence        = []byte{0xff, 0x41, 0x00, 0x04, 0x60, 0x26}
)

func (c *client) Connect() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	n, err := c.conn.Write(resetSequence)
	if err != nil {
		panic(err)
	}

	if n != len(resetSequence) {
		return ErrWroteIncompleteData
	}

	// my setup doesn't seem to like it if you hit it too fast
	time.Sleep(time.Second)

	n, err = c.conn.Write(getCameraModelSequence)
	if err != nil {
		panic(err)
	}

	if n != len(getCameraModelSequence) {
		return ErrWroteIncompleteData
	}

	buf := make([]byte, 100)
	n, err = c.conn.Read(buf)
	if err != nil {
		return err
	}

	// I would expect that 0x03 corresponds to the Nikon F5 model (which I am using for testing)
	// but can't be sure without other cameras at hand
	// someone want to send me an F6?
	if !bytes.Equal(buf[:n], []byte{0x00, 0x03, 0x00}) {
		return fmt.Errorf("unknown respond sequence: %q", buf[:n])
	}

	// my setup doesn't seem to like it if you hit it too fast
	time.Sleep(time.Second)
	return nil
}

func (c *client) GetData(consume RollConsumer) ([]byte, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	n, err := c.conn.Write(getDataSequence)
	if err != nil {
		panic(err)
	}

	if n != len(getDataSequence) {
		return nil, ErrWroteIncompleteData
	}

	// my setup doesn't seem to like it if you hit it too fast
	time.Sleep(time.Second)
	var data []byte
	var all []byte

	for {
		buf := make([]byte, 100)
		n, err = c.conn.Read(buf)
		if err != nil {
			return nil, err
		}

		data = append(data, buf[:n]...)
		all = append(all, buf[:n]...)

		// I think this is a start sequence, maybe total number of rolls
		if bytes.HasPrefix(data, []byte{0x3e, 0xf9}) {
			data = data[2:]
		}

		for {
			// 0x80, 0x81 is the exit sequence for a single roll
			i := bytes.Index(data, []byte{0x80, 0x81})
			if i == -1 {
				break
			}

			roll := protocol.ParseRoll(data[:i])
			err = consume(roll)
			if err != nil {
				return nil, err
			}

			data = data[i+2:]
		}

		// 0x80, 0x80 is the exit sequence for the whole thing
		if bytes.Contains(data, []byte{0x80, 0x80}) {
			break
		}
	}

	return all, nil
}
