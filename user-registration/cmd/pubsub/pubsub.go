package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"os"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func Connect() mqtt.Client {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("emqx:1883").SetClientID(fmt.Sprintf("test-clinet-%d", rand.Intn(200)))

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c
}

func PublishUser(message []byte) {
	c := Connect()
	token := c.Publish("user-registration-microservices-final-practice-topic/1", 0, false, message)
	if token.Error() != nil {
		log.Println(token.Error())
	}
	log.Println("test")

	token.Wait()

	time.Sleep(6 * time.Second)

	c.Disconnect(250)

	time.Sleep(1 * time.Second)

}

func PublishCommit(id string) {
	c := Connect()
	token := c.Publish("commit-microservices-final-practice-topic/1", 0, false, id)
	if token.Error() != nil {
		log.Println(token.Error())
	}

	token.Wait()

	time.Sleep(6 * time.Second)

	c.Disconnect(250)

	time.Sleep(1 * time.Second)

}
