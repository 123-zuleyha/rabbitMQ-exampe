package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Succesfly Connected To Our RabbitMQ Instance")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue", //Queue name
		false,       // durable
		false,       //autoDelete
		false,       //exclusive
		false,       //noWait
		nil,         //arguments
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)
	err = ch.Publish(
		"",          //exchange
		"TestQueue", //routing key
		false,       // mandatory
		false,       //immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World"),
		},
	)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Succesfly Published Message to Queue")

}
