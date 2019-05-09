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
	s      *serial.Port
	msgnum int
)

func serial_send(s *serial.Port, cmd []byte) {
	n, err := s.Write(cmd) //fmt.Fprint(s, cmd)
	if err != nil {
		log.Errorf("failed to write to serial port")
	}
	log.Infof("serial sent %d bytes", n)
}

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

	var err error
	c := &serial.Config{Name: config.SerialPort, Baud: 115200}
	s, err = serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	for true {

		log.Println("Waiting to read from serial")
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("  read [%d] => %s\n", n, string(buf))
		log.Printf("%q", buf[:n])
	}

}
