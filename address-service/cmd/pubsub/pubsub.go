package pubsub

import (
	"address-service/cmd/initializers"
	"address-service/cmd/model"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	id, err := strconv.Atoi(string(msg.Payload()))
	if err != nil {
		log.Printf("convertion error: %v", err)
	}
	address := model.Address{}
	result := initializers.DB.Model(&address).Where("id=?", id).Updates(model.Address{Valid: true, ValidAt: time.Now()})
	if result.Error != nil {
		log.Printf("update error: %v", result.Error)
	}

	log.Println("Address valid")

	c.Publish("address-microservices-final-practice-topic/1", 0, false, msg.Payload())

	log.Println("Address published")
}

var c mqtt.Client

func Connect() (mqtt.Client, error) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("emqx:1883").SetClientID(fmt.Sprintf("test-clinet-%d", rand.Intn(50)))

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	token := c.Subscribe("commit-microservices-final-practice-topic/1", 0, nil)
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	for {
		token.Wait()
		i := c.IsConnected()
		if !i {
			Connect()
		}
	}

	// Disconnect
	//c.Disconnect(250)
	//time.Sleep(1 * time.Second)

	return c, nil
}
