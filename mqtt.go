package main

import (
	"crypto/tls"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// MQTTConfiguration allows us to comment to the approriate MQTT channel(s)
type MQTTConfiguration struct {
	MQTTAddr     string // the address and port for the MQTT broker
	MQTTTopic    string
	MQTTUsername string
	MQTTPassword string
}

// onMessageReceived is called for every message the arrives under one
// of the topics we are subscribed to.
func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	msg := message.Payload()
	topic := message.Topic()
	log.Infof("Message ~ topic %s ~ %s\n", topic, msg)

	var err error
	switch topic {
	case "mot":
		if Serial == nil {
			Serial, err = GetSerialPort(config.SerialPort)
			if err != nil {
				log.Errorf("error getting serial port %s ~> %v", config.SerialPort, err)
				return
			}
		}
		if err = Serial.Send(string(msg)); err != nil {
			log.Errorf("Error sending motor msg %v", err)
		}
	default:
		log.Errorf("MQTT Do not know how to handle ", topic)
	}
}

// mqtt subscribes and responds to the channels we are interested in
func mqtt_service() {
	// wg is a global
	defer func() {
		wg.Done()
		log.Errorln("Exiting MQTT service")
	}()

	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	log.Println("Start MQTT client service")

	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()
	connOpts := mqtt.NewClientOptions()
	connOpts.AddBroker(config.MQTTAddr)
	connOpts.SetClientID("otto-" + hostname)
	connOpts.SetCleanSession(true)

	if config.MQTTUsername != "" {
		connOpts.SetUsername(config.MQTTUsername)
		if config.MQTTPassword != "" {
			connOpts.SetPassword(config.MQTTPassword)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(config.MQTTTopic, 0, onMessageReceived); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}
	}

	client := mqtt.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		log.Infof("Connected to %s\n", server.Addr)
	}
	<-c
}
