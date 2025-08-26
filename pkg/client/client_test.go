//go:build unit

package client_test

import (
	"embed"
	"log/slog"
	"path"
	"testing"

	serial_test "github.com/lucasl0st/nserial/internal/mock/serial"
	"github.com/lucasl0st/nserial/internal/util"
	"github.com/lucasl0st/nserial/pkg/client"
	"github.com/lucasl0st/nserial/pkg/model"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

//go:embed test/test_dump.bin
var testDump []byte

//go:embed test/*.SDF
var sdfTestData embed.FS

var testDumpDataRolls = map[uint]model.Roll{}

func init() {
	files, err := sdfTestData.ReadDir("test")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		b, err := sdfTestData.ReadFile(path.Join("test", file.Name()))
		if err != nil {
			panic(err)
		}

		roll, err := util.ParseSDF(b)
		if err != nil {
			panic(err)
		}

		testDumpDataRolls[roll.Number] = *roll
	}
}

type ClientTestSuite struct {
	suite.Suite

	ctrl           *gomock.Controller
	serialPortMock *serial_test.MockPort

	testSubject client.Client
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.serialPortMock = serial_test.NewMockPort(suite.ctrl)

	testSubject, err := client.New(client.WithSerialPort(suite.serialPortMock))
	suite.Require().NoError(err)
	suite.testSubject = testSubject
}

func (suite *ClientTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *ClientTestSuite) TestConnect() {
	suite.serialPortMock.EXPECT().Write([]byte{0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f, 0x6f}).Return(8, nil).Times(1)
	suite.serialPortMock.EXPECT().Write([]byte{0x00, 0xff, 0x41, 0x00, 0x04, 0x01, 0x45}).Return(7, nil).Times(1)
	suite.serialPortMock.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
		p[1] = 0x03
		return 3, nil
	})

	err := suite.testSubject.Connect()
	suite.Require().NoError(err)
}

func (suite *ClientTestSuite) TestGetData() {
	i := 0

	suite.serialPortMock.EXPECT().Write([]byte{0xff, 0x41, 0x00, 0x04, 0x60, 0x26}).Return(6, nil).Times(1)
	suite.serialPortMock.EXPECT().Read(gomock.Any()).DoAndReturn(func(p []byte) (int, error) {
		n := 0

		for k := range p {
			if i >= len(testDump) {
				break
			}

			p[k] = testDump[i]

			i++
			n++
		}

		return n, nil
	}).AnyTimes()

	consumer := func(roll model.Roll) error {
		expected, ok := testDumpDataRolls[roll.Number]
		if !ok {
			slog.Warn("skipping roll, no test data found", "roll", roll.Number)
			return nil
		}

		suite.Equal(expected.ISO, roll.ISO)
		suite.Len(roll.Frames, len(expected.Frames))

		for l, expectedFrame := range expected.Frames {
			actualFrame := roll.Frames[l]
			suite.Equal(expectedFrame.Aperture, actualFrame.Aperture)
			// suite.Equal(expectedFrame.MaxAperture, actualFrame.MaxAperture)
			// suite.Equal(expectedFrame.Number, actualFrame.Number)
			// suite.Equal(expectedFrame.ShutterSpeed, actualFrame.ShutterSpeed)
			// suite.Equal(expectedFrame.MaxAperture, actualFrame.MaxAperture)
		}

		return nil
	}

	data, err := suite.testSubject.GetData(consumer)
	suite.Require().NoError(err)
	suite.NotNil(data)
}
