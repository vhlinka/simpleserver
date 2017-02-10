package handlers

import (
	"gateway"
	"log"
	"net/http"
	"serial"
	"strconv"
	"time"
	//	"sync"
)

func openDeviceSerialPort() (*serial.Port, error) {

	var s *serial.Port
	var err error

	//	errorFlag := false
	portnumber := 0
	doneFlag := false
	ttyNameBase := "/dev/ttyUSB" // Linus tty device - without the trailing port digit
	ttyName := ttyNameBase + strconv.Itoa(portnumber)

	for !doneFlag {

		c := &serial.Config{Name: ttyName, Baud: 9600, ReadTimeout: time.Second * 5}
		s, err = serial.OpenPort(c)
		if err != nil {
			if portnumber >= 9 {
				//				errorFlag = true // flag that there was an error opening the device serial port
				doneFlag = true
				log.Fatal(err)
			} else {
				portnumber++
				ttyName = ttyNameBase + strconv.Itoa(portnumber)

			}
		} else {
			doneFlag = true
			err = nil
		}

	}

	return s, err

}

var gatewayInChan chan []byte
var done chan byte

func EnterGatewayMode(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

	s, err := openDeviceSerialPort()
	if err != nil {
		log.Fatal(err)
	}

	// channels used by the gateway
	gatewayInChan = make(chan []byte, 512)
	done = make(chan byte)

	go gateway.ReadGateway(gatewayInChan, s)
	go gateway.DisplayGateway(gatewayInChan, done)

	log.Println("----------------<DEVICE IS NOW SET TO GATEWAY MODE>-----------------")
}

func ExitGatewayMode(w http.ResponseWriter, r *http.Request) {

	mu.Lock() // log visit
	count++
	mu.Unlock()

}
