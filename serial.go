package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/tarm/serial"
)

type SerialConfiguration struct {
	SerialPort  string
	SerialSpeed int
}

var (
	msgnum int
)

func send_cmd(s *serial.Port, r, l int) {
	msgnum += 1
	n, err := fmt.Fprintf(s, "m:%d:%d:%d", msgnum, r, l)
	if err != nil {
		panic(err)
	}
	log.Printf("wrote %d bytes\n", n)
	s.Flush()
}

// serial_run starts an io process on the serial port
func serial_service() {
	// weight group (wg) is a global variable
	defer func() {
		wg.Done()
		log.Errorln("Exiting serial service")
	}()

	// XXX TODO: Make the serial
	c := &serial.Config{Name: config.SerialPort, Baud: 115200}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	var i int
	buf := make([]byte, 128)
	for true {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		i++
		if i%10 == 0 {
			log.Printf("sending data 150 150  ")
			send_cmd(s, 0, 0)
		}
		log.Printf("%q", buf[:n])
	}

}
