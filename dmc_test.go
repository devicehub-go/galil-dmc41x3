package galildmc41x3_test

import (
	"fmt"
	"log"
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
			Host:         "10.0.4.134",
			Port:         23,
			ReadTimeout:  2 * time.Second,
			WriteTimeout: 2 * time.Second,
		},
	})
	if err := DMC41x3.Connect(); err != nil {
		t.Errorf("Error on connect to galil: %v", err)
	}
	defer DMC41x3.Disconnect()

	axis := 'A'
	DMC41x3.WriteCommand("CN 1")
	if err := DMC41x3.Disable(string(axis)); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetMotorType(axis, -2.5); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetAmplifierGain(axis, 0); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetLowCurrent(axis, 1); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetSpeed(axis, 20480/2); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetAcceleration(axis, 20480); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.SetDeceleration(axis, 20480); err != nil {
		log.Fatal(err)
	}
	if err := DMC41x3.Enable(string(axis)); err != nil {
		log.Fatal(err)
	}

	if err := DMC41x3.SetTargetRelPosition(axis, 20480*10); err != nil {
		message, _ := DMC41x3.GetErrorCode(1)
		log.Fatalf("Error on set relative position: %s", message)
	}
	if err := DMC41x3.Begin(string(axis)); err != nil {
		message, _ := DMC41x3.GetErrorCode(1)
		log.Fatalf("Error on begin: %s", message)
	}

	for {
		fwd, err := DMC41x3.IsForwardSwitchOn(axis)
		if err != nil {
			message, _ := DMC41x3.GetErrorCode(1)
			log.Printf("Error on get forward limit: %s", message)
		}
		rev, err := DMC41x3.IsReverseSwitchOn(axis)
		if err != nil {
			message, _ := DMC41x3.GetErrorCode(1)
			log.Printf("Error on get reverse limit: %s", message)
		}
		position, err := DMC41x3.GetPosition(axis)
		if err != nil {
			message, _ := DMC41x3.GetErrorCode(1)
			log.Printf("Error on get position: %s", message)
		}

		fmt.Printf("Forward: %v | Reverse %v | Position: %.4f\r", fwd, rev, position)
		time.Sleep(1 * time.Second)
	}
}
