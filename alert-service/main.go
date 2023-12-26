package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	database "alert-service/database/sqlc"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// graceful shutdown
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	postgres, conn, err := database.NewPostresDB(context.TODO(), os.Getenv("POSTGRES_ADDRESS"))
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer conn.Close(context.TODO())

	token, err := NewPasetoMaker(os.Getenv("TOKEN_SYMMETRIC_KEY"))
	if err != nil {
		log.Fatal("Error creating token maker:", err)
	}

	validator := validator.New()

	authSvc := NewAuthSvc(postgres, token, 1*time.Hour)

	redis, err := NewRedis(os.Getenv("REDIS_ADDRESS"))
	if err != nil {
		log.Fatal("Error connecting to redis:", err)
	}

	alertSvc := NewAlertService(redis, postgres)

	api := NewAPI(":3000", token, authSvc, validator, alertSvc).Run(mainCtx)

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		log.Println("starting server on port", "3000")
		return api.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Println("shutting down server...")
		return api.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

	// producer, err := NewKafkaProducer(
	// 	[]string{
	// 		os.Getenv("KAFKA_ADDRESS"),
	// 	},
	// 	os.Getenv("KAFKA_TOPIC"),
	// )
	// if err != nil {
	// 	log.Fatal("Error setting up kafka:", err)
	// }

	// for i := 0; i < 10; i++ {
	// 	producer.Send(i, 100*i)
	// 	time.Sleep(10 * time.Second)
	// }
}
