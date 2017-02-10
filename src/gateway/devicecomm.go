package gateway

import (
	"bytes"
	//	"fmt"
	//	"log"
	"os"
	"serial"
	//	"time"
)

//
// reads the incoming data from the serial port - assumed to be the gateway - and
// places the content on a buffered channel
//
//
func ReadGateway(gatewayInChan chan []byte, s *serial.Port) {

	buf := make([]byte, 2048)
	quit := false

	for quit == false {
		n, err := s.Read(buf)
		if err != nil {
			//			quit = true			// TODO: need to add error code -
		} else {
			gatewayInChan <- buf[:n]
		}
	}

}

//
// pull content off of the inbound getway channel - display for now
//
func DisplayGateway(gatewayInChan chan []byte, done chan byte) {

	var b bytes.Buffer
	chanClosed := false

	for !chanClosed {
		buf, chan_ok := <-gatewayInChan
		if !chan_ok {
			chanClosed = true
		}
		b.Write(buf)
		b.WriteTo(os.Stdout)

	}

	done <- 0x01
}
