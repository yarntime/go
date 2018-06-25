package main

import (
	"bytes"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var conn *amqp.Connection
var channel *amqp.Channel
var count = 0

const (
	queueName = "pysix"
	exchange  = "pysix"
	mqurl     = "amqp://guest:password@192.168.254.44:30159/"
)

func main() {
	go func() {
		for {
			err := push()
			if err != nil {
				fmt.Printf("error: %s\n", err.Error())
				mqConnect()
			}
			time.Sleep(1 * time.Second)
		}
	}()
	select {}
	fmt.Println("end")
	close()
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func mqConnect() {
	var err error
	conn, err = amqp.Dial(mqurl)
	failOnErr(err, "failed to connect tp rabbitmq")

	channel, err = conn.Channel()
	failOnErr(err, "failed to open a channel")
}

func close() {
	channel.Close()
	conn.Close()
}

func push() error {

	if channel == nil {
		mqConnect()
	}
	msgContent := "hello world!"

	err := channel.Publish(exchange, queueName, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})

	return err
}

func BytesToString(b *[]byte) *string {
	s := bytes.NewBuffer(*b)
	r := s.String()
	return &r
}
