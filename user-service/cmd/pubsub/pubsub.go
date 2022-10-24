package pubsub

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"user-service/cmd/initializers"
	"user-service/cmd/model"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	time.Sleep(time.Second)
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	id, err := strconv.Atoi(string(msg.Payload()))
	if err != nil {
		log.Printf("convertion error: %v", err)
	}
	user := model.User{}
	result := initializers.DB.Model(&user).Where("id=?", id).Updates(model.User{Valid: true, ValidAt: time.Now()})
	if result.Error != nil {
		log.Printf("update error: %v", result.Error)
	}

	log.Println("User valid")

	c.Publish("user-microservices-final-practice-topic/1", 0, false, msg.Payload())

	log.Println("User published")
}

var c mqtt.Client

func Connect() error {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("emqx:1883").SetClientID(fmt.Sprintf("test-clinet-%d", rand.Intn(23)))

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	token := c.Subscribe("commit-microservices-final-practice-topic/1", 0, nil)
	if token.Wait() && token.Error() != nil {
		return token.Error()
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

	return nil
}
