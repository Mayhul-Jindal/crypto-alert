package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// ogux ynug bnjb bkad

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	gmail := NewGmailSender(
		os.Getenv("GAMIL_NAME"),
		os.Getenv("GMAIL_ADDRESS"),
		os.Getenv("GMAIL_PASSWORD"),
	)

	consumer, err := NewKafkaConsumer(
		gmail,
		[]string{
			os.Getenv("KAFKA_ADDRESS"),
		},
		os.Getenv("KAFKA_GROUP"),
		[]string{
			os.Getenv("KAFKA_TOPIC"),
		},
	)
	if err != nil {
		log.Fatal("Error setting up kafka:", err)
	}

	ctx := context.Background()
	log.Fatal(consumer.Process(ctx))
}
