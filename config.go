package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	Addrport   string // http address / port
	WSAddrport string // Web socket address / port

	// found in mqtt.go
	MQTTConfiguration

	// found in serial.go
	SerialConfiguration

	Dir       string
	NoService bool
	Fetch     bool
	Cache     bool
	Plugdir   string // Directory that contains plugin dirs

	ListPlugins bool
	ListRoutes  bool

	LogLevel  string
	LogOutput string
	LogFormat string
}

func init() {
	config = Configuration{
		Addrport:  ":3333",
		LogLevel:  "warn",
		LogOutput: "stdout",
		LogFormat: "json",
	}

	flag.StringVar(&config.LogLevel, "level", "warn", "set log level")
	flag.StringVar(&config.LogFormat, "format", "json", "set log format")
	flag.StringVar(&config.LogOutput, "logfile", "stdout", "logfile, stdout or stderr")
	flag.StringVar(&config.Plugdir, "plugdir", ".", "the dir to look for plugins")

	flag.StringVar(&config.Addrport, "addr", ":4433", "address and port to listen on")
	flag.StringVar(&config.WSAddrport, "wsaddr", ":4434", "websocket address to listen on")

	flag.StringVar(&config.MQTTAddr, "mqttaddr", "10.24.2.13:1883", "address of MQTT broker")
	flag.StringVar(&config.MQTTTopic, "mqttSubjects", "#", "mqtt subject to listen to")
	flag.StringVar(&config.MQTTUsername, "mqtt-user", "", "username for mqtt broker")
	flag.StringVar(&config.MQTTPassword, "mqtt-password", "", "password for mqtt broker")

	flag.BoolVar(&config.ListRoutes, "routes", false, "Walk the routes after they have been added")
	flag.BoolVar(&config.ListPlugins, "plugins", false, "List the plugins we are using")

	flag.BoolVar(&config.NoService, "no-daemon", false, "run in background as a service")
}

// Save we can start using store for this, correct?
func (c *Configuration) Save(fname string) (err error) {
	var jbuf []byte

	jbuf, err = json.Marshal(c)
	checkError(err)

	err = ioutil.WriteFile(fname, jbuf, 0644)
	checkError(err)

	return
}

func checkError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
