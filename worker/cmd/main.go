package main

import (
	sabbitmq "govision_worker/internal/services/rabbitmq"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	_ = godotenv.Load()

	rabbitMQConnection, err := sabbitmq.NewRabbittMQConnection(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Printf("[ERROR] - RabbitMQ connection error: %v", err)
		panic(err)
	}
	defer rabbitMQConnection.Close()

	ch, err := rabbitMQConnection.Channel()
	if err != nil {
		log.Printf("[ERROR] - RabbitMQ channel error: %v", err)
		panic(err)
	}
	defer ch.Close()
}