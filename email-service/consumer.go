package main

import (
	database "email-service/database/sqlc"
	"context"
	"log"

	"github.com/IBM/sarama"
)

// kafka consumer code here
type Consumer interface {
	Process(context.Context) error
}

type kafkaConsumer struct {
	db database.Querier
	email Emailer
	cg    sarama.ConsumerGroup
	topics []string
}

func NewKafkaConsumer(db database.Querier, email Emailer, addr []string, group string, topics []string) (Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	
	consumerGroup, err := sarama.NewConsumerGroup(addr, group, config)
	if err != nil {
		return nil, err
	}

	return &kafkaConsumer{
		db: db,
		email: email,
		cg:    consumerGroup,
		topics: topics,
	}, nil
}

func (*kafkaConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*kafkaConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (k *kafkaConsumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Println(string(msg.Key), string(msg.Value))
		k.email.send(
			"Crypto Alert", 
			string(msg.Value), 
			[]string{"mayhuljindal@gmail.com"}, 
			nil, 
			nil, 
			nil,
		)
		sess.MarkMessage(msg, "")
	}

	return nil
}

func (k *kafkaConsumer) Process(ctx context.Context) error {
	for {
		err := k.cg.Consume(ctx, k.topics, k)
		if err != nil {
			log.Println("Error from consumer: ", err)
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}
