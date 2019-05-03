package main

import (
	"crypto/tls"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Message ~ topic %s ~ %s\n", message.Topic(), message.Payload())

	// Now we need to write this message to the websocket channel, letting the
	// websocket pick it up and run with it.
}

func mqtt_run() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	/*
		server := flag.String("server", config.MQTTAddr, "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
		topic := flag.String("topic", "#", "Topic to subscribe to")
		qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
		clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
		username := flag.String("username", "", "A username to authenticate to the MQTT server")
		password := flag.String("password", "", "Password to match username")
		flag.Parse()
	*/

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
		fmt.Printf("Connected to %s\n", *server)
	}

	<-c
}
