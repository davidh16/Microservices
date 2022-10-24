package pubsub

import (
	"context"
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
	"user-address-agg/cmd/initializers"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	if msg.Topic() == "user-microservices-final-practice-topic/1" {

		id, err := strconv.Atoi(string(msg.Payload()))
		if err != nil {
			log.Printf("convertion error: %v", err)
		}

		database := initializers.Client.Database("quickstart")
		usersCollection := database.Collection("users")
		_, err = usersCollection.InsertOne(context.Background(), bson.D{
			{"user_id", id},
		})
		if err != nil {
			log.Printf("mongodb error, user: %v", err)
		}
	}

	if msg.Topic() == "address-microservices-final-practice-topic/1" {
		id, err := strconv.Atoi(string(msg.Payload()))
		if err != nil {
			log.Printf("convertion error: %v", err)
		}

		database := initializers.Client.Database("quickstart")
		usersCollection := database.Collection("addresses")
		_, err = usersCollection.InsertOne(context.Background(), bson.D{
			{"address_id", id},
		})
		if err != nil {
			log.Printf("mongodb error, address: %v", err)
		}
	}

	if msg.Topic() == "user-registration-microservices-final-practice-topic/1" {
		var user any
		err := json.Unmarshal(msg.Payload(), &user)
		if err != nil {
			log.Printf("convertion error: %v", err)
		}

		database := initializers.Client.Database("quickstart")
		usersCollection := database.Collection("registered_users")
		_, err = usersCollection.InsertOne(context.Background(), bson.D{
			{"user", user},
		})
		if err != nil {
			log.Printf("mongodb error, registered_user: %v", err)
		}
	}
}

var c mqtt.Client

func Connect() error {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("emqx:1883").SetClientID(fmt.Sprintf("test-clinet-%d", rand.Intn(123)))

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	c = mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	//topics := map[string]byte{"user-microservices-final-practice-topic/1": 1, "address-microservices-final-practice-topic/1": 2, "user-registration-microservices-final-practice-topic/1": 3}

	token := c.Subscribe("user-microservices-final-practice-topic/1", 0, f)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	token = c.Subscribe("address-microservices-final-practice-topic/1", 0, nil)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	token = c.Subscribe("user-registration-microservices-final-practice-topic/1", 0, nil)
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
