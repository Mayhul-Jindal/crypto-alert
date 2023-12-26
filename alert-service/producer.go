package main

import (
	"log"
	"strconv"

	"github.com/IBM/sarama"
)

type Producer interface {
	Send(id string, price int) error
}

type kafkaProducer struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaProducer(addr []string, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		return nil, err
	}

	return &kafkaProducer{
		producer: producer,
		topic:    topic,
	}, nil
}

func (k *kafkaProducer) Send(id string, price int) error {
	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.StringEncoder(id),
		Value: sarama.StringEncoder(strconv.Itoa(price)),
	}

	partition, offset, err := k.producer.SendMessage(msg)
	log.Println(partition, offset, msg, err)
	
	return err
}