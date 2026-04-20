package galildmc41x3_test

import (
	"fmt"
	"testing"
	"time"

	galil "github.com/devicehub-go/galil-dmc41x3"
	"github.com/devicehub-go/unicomm"
	"github.com/devicehub-go/unicomm/protocol/unicommtcp"
)

func TestConnection(t *testing.T) {
	DMC41x3 := galil.New(unicomm.Options{
		Protocol: unicomm.TCP,
		TCP: unicommtcp.TCPOptions{
			Host:         "10.0.4.194",
			Port:         23,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := DMC41x3.Connect(); err != nil {
		t.Errorf("Error on connect to galil: %v", err)
	}
	defer DMC41x3.Disconnect()

	fmt.Println(DMC41x3.QueryFloat64("MG _LRA"))

}
