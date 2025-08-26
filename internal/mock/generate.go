//go:build generate

package mock

//go:generate go tool go.uber.org/mock/mockgen -destination=serial/port_mock.go -package=serial_test go.bug.st/serial Port
