package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func produceMessages(producer *kafka.Producer, topic string) {
	time_stamp := time.Now().Format(time.RFC850)
	message := "'idSensor': '001',\n 'valor': '30',\n 'timestamp': '" + fmt.Sprintf("%d" ,time_stamp) + "',\n 'tipoPoluentes': 'PM2.5',\n 'nivel': '35.2'"
	fmt.Println("Sensor 1 - Produzindo mensagens...")

	for {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(message),
		}, nil)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Configurações do produtor
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"client.id":         "go-producer",
	})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	topic := "qualidadeAr"

	// Iniciar produtor em uma goroutine separada
	go produceMessages(producer, topic)

	// Manter o programa em execução
	select {}
}
