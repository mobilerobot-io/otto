package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/tarm/serial"
)

type SerialConfiguration struct {
	SerialPort  string
	SerialSpeed int
}

type SerialPort struct {
	*serial.Port
	msgnum int
	err    error

	incoming int
	outgoing int

	listen bool
}

var (
	msgnum  int
	portmap map[string]*SerialPort
)

func init() {
	portmap = make(map[string]*SerialPort, 5)
}

func GetSerialPort(name string) (s *SerialPort, err error) {
	var ex bool

	if s, ex = portmap[name]; ex {
		return s, nil
	}

	c := &serial.Config{Name: config.SerialPort, Baud: 115200}
	s = &SerialPort{}
	s.Port, err = serial.OpenPort(c)
	if err != nil {
		log.Errorf("failed to open port %s ~> %v", name, err)
		return nil, err
	}
	portmap[name] = s
	log.Infof("starting port %s", name)
	return s, nil
}

// Send a message on the serial port
func (s *SerialPort) Send(msg string) (err error) {
	n, err := s.Write([]byte(msg))
	if err != nil {
		log.Errorf("serial port send failed %v", err)
		return err
	}
	log.Info("serial port sent %d bytes", n)
	return nil
}

// Listen for incoming data on the serial port
func (s *SerialPort) Listen() (err error) {

	buf := make([]byte, 128)
	for s.listen {

		log.Debugln("Waiting to read from serial")
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
			return err
		}

		log.Infoln("incoming buffer send to incoming message channel")
		Incoming(string(buf))

		log.Debugf("  read [%d] => %s", n, string(buf))
		log.Debugf("  %q", buf[:n])
	}
	return nil
}

func (s *SerialPort) Close() {
	s.Port.Close()
}
