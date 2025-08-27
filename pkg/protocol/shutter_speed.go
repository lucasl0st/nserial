package protocol

import (
	"time"

	"github.com/lucasl0st/nserial/pkg/model"
)

var byteToShutterSpeed = map[byte]model.ShutterSpeed{
	0xfd: model.ShutterSpeed(time.Minute), // nice overflow bro
	0xfe: model.ShutterSpeed(time.Second * 50),

	0x00: model.ShutterSpeed(time.Second * 30),
	0x04: model.ShutterSpeed(time.Second * 20),
	0x0a: model.ShutterSpeed(time.Second * 10),

	0x1e: model.ShutterSpeed(time.Second),
	0x1f: model.ShutterSpeed(time.Second),            // not used
	0x20: model.ShutterSpeed(769 * time.Millisecond), // 1/1.3 ≈ 769ms
	0x21: model.ShutterSpeed(769 * time.Millisecond), // 1/1.3 ≈ 769ms // not used
	0x22: model.ShutterSpeed(625 * time.Millisecond), // 1/1.6 ≈ 625ms
	0x23: model.ShutterSpeed(625 * time.Millisecond), // 1/1.6 ≈ 625ms not used
	0x24: model.ShutterSpeed(time.Second / 2),
	0x25: model.ShutterSpeed(time.Second / 2),        // not used
	0x26: model.ShutterSpeed(400 * time.Millisecond), // 1/2.5 = 400ms
	0x27: model.ShutterSpeed(400 * time.Millisecond), // 1/2.5 = 400ms not used
	0x28: model.ShutterSpeed(time.Second / 3),
	0x29: model.ShutterSpeed(time.Second / 3), // not used
	0x2a: model.ShutterSpeed(time.Second / 4),
	0x2b: model.ShutterSpeed(time.Second / 4), // not used
	0x2c: model.ShutterSpeed(time.Second / 5),
	0x2d: model.ShutterSpeed(time.Second / 5), // not used
	0x2e: model.ShutterSpeed(time.Second / 6),
	0x2f: model.ShutterSpeed(time.Second / 6), // not used
	0x30: model.ShutterSpeed(time.Second / 8),
	0x31: model.ShutterSpeed(time.Second / 8), // not used
	0x32: model.ShutterSpeed(time.Second / 10),
	0x33: model.ShutterSpeed(time.Second / 10), // not used
	0x34: model.ShutterSpeed(time.Second / 13),
	0x35: model.ShutterSpeed(time.Second / 13), // not used
	0x36: model.ShutterSpeed(time.Second / 15),
	0x37: model.ShutterSpeed(time.Second / 15), // not used
	0x38: model.ShutterSpeed(time.Second / 20),
	0x39: model.ShutterSpeed(time.Second / 20), // not used
	0x3a: model.ShutterSpeed(time.Second / 25),
	0x3b: model.ShutterSpeed(time.Second / 25), // not used
	0x3c: model.ShutterSpeed(time.Second / 30),
	0x3d: model.ShutterSpeed(time.Second / 30), // not used
	0x3e: model.ShutterSpeed(time.Second / 40),
	0x3f: model.ShutterSpeed(time.Second / 40), // not used
	0x40: model.ShutterSpeed(time.Second / 50),
	0x41: model.ShutterSpeed(time.Second / 50), // not used
	0x42: model.ShutterSpeed(time.Second / 60),
	0x43: model.ShutterSpeed(time.Second / 60), // not used
	0x44: model.ShutterSpeed(time.Second / 80),
	0x45: model.ShutterSpeed(time.Second / 80), // not used
	0x46: model.ShutterSpeed(time.Second / 100),
	0x47: model.ShutterSpeed(time.Second / 100), // not used
	0x48: model.ShutterSpeed(time.Second / 125),
	0x49: model.ShutterSpeed(time.Second / 25), // not used
	0x4a: model.ShutterSpeed(time.Second / 160),
	0x4b: model.ShutterSpeed(time.Second / 160), // not used
	0x4c: model.ShutterSpeed(time.Second / 200),
	0x4d: model.ShutterSpeed(time.Second / 200), // not used
	0x4e: model.ShutterSpeed(time.Second / 250),
	0x4f: model.ShutterSpeed(time.Second / 250), // not used
	0x50: model.ShutterSpeed(time.Second / 320),
	0x51: model.ShutterSpeed(time.Second / 320), // not used
	0x52: model.ShutterSpeed(time.Second / 400),
	0x53: model.ShutterSpeed(time.Second / 400), // not used
	0x54: model.ShutterSpeed(time.Second / 500),
	0x55: model.ShutterSpeed(time.Second / 500), // not used
	0x56: model.ShutterSpeed(time.Second / 640),
	0x57: model.ShutterSpeed(time.Second / 640), // not used
	0x58: model.ShutterSpeed(time.Second / 800),
	0x59: model.ShutterSpeed(time.Second / 800), // not used
	0x5a: model.ShutterSpeed(time.Second / 1000),
	0x5b: model.ShutterSpeed(time.Second / 1000), // not used
	0x5c: model.ShutterSpeed(time.Second / 1250),
	0x5d: model.ShutterSpeed(time.Second / 1250), // not used
	0x5e: model.ShutterSpeed(time.Second / 1600),
	0x5f: model.ShutterSpeed(time.Second / 1600), // not used
	0x60: model.ShutterSpeed(time.Second / 2000),
	0x61: model.ShutterSpeed(time.Second / 2000), // not used
	0x62: model.ShutterSpeed(time.Second / 2500),
	0x63: model.ShutterSpeed(time.Second / 2500), // not used
	0x64: model.ShutterSpeed(time.Second / 3200),
	0x65: model.ShutterSpeed(time.Second / 3200), // not used
	0x66: model.ShutterSpeed(time.Second / 4000),
	0x67: model.ShutterSpeed(time.Second / 4000), // not used
	0x68: model.ShutterSpeed(time.Second / 5000),
	0x69: model.ShutterSpeed(time.Second / 5000), // not used
	0x6a: model.ShutterSpeed(time.Second / 6400),
	0x6b: model.ShutterSpeed(time.Second / 6400), // not used
	0x6c: model.ShutterSpeed(time.Second / 8000),
}

func GetShutterSpeed(b byte) model.ShutterSpeed {
	s, ok := byteToShutterSpeed[b]
	if !ok {
		return model.ShutterSpeed(0)
	}

	return s
}

func ShutterSpeedFromString(s string) model.ShutterSpeed {
	for _, speed := range byteToShutterSpeed {
		if speed.String() != s {
			continue
		}

		return speed
	}

	return model.ShutterSpeed(0)
}
