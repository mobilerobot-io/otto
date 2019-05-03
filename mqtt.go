package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\n\tMessage: %s\n", message.Topic(), message.Payload())
}

func mqtt_run() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://10.24.0.112:1883", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topic := flag.String("topic", "#", "Topic to subscribe to")
	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")
	username := flag.String("username", "", "A username to authenticate to the MQTT server")
	password := flag.String("password", "", "Password to match username")
	flag.Parse()

	connOpts := mqtt.NewClientOptions().AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)
	if *username != "" {
		connOpts.SetUsername(*username)
		if *password != "" {
			connOpts.SetPassword(*password)
		}
	}
	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived); token.Wait() && token.Error() != nil {
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

/*
import (
	"fmt"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

var (
	mqttCli    mqtt.Client
	msgHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", msg.Topic())
		fmt.Printf("MSG: %s\n", msg.Payload())
	}
)

// Handle MQTT messages
func mqtt_init() {
	//mqtt.DEBUG = log.New()
	mqtt.ERROR = log.New()

	opts := mqtt.NewClientOptions().AddBroker("tcp://10.24.0.112:1883").SetClientID("otto")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(msgHandler)
	opts.SetPingTimeout(1 * time.Second)

	mqttCli := mqtt.NewClient(opts)
	if mqttCli == nil {
		log.Fatal("Failed to create mqtt client ")
	}

	if token := mqttCli.Connect(); token.Wait() && token.Error() != nil {
		log.Error(token.Error().Error())
		return
	}
}

func mqtt_subscribe(subj string) {
	if token := mqttCli.Subscribe(subj, 0, nil); token.Wait() && token.Error() != nil {
		log.Error(token.Error().Error())
		return
	}
}

func mqtt_publish_loop(subj string) {
	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("message #%d", i)
		token := mqttCli.Publish(subj, 0, false, text)
		token.Wait()
	}

	time.Sleep(6 * time.Second)

	if token := mqttCli.Unsubscribe(subj); token.Wait() && token.Error() != nil {
		log.Error(token.Error())
		return
	}
	mqttCli.Disconnect(250)
	time.Sleep(1 * time.Second)
}
*/
