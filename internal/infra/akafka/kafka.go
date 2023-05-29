package akafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,     //servidor de comunicacao
		"group.id":          "sergio-go", //grupo de consumidores
		"auto.offset.reset": "earliest",  //pegar a mensagem sempre do inicio
	})
	if err != nil {
		panic(err)
	}
	kafkaConsumer.SubscribeTopics(topics, nil) // consumir os topicos
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err != nil {
			msgChan <- msg //toda vez que receber mensagem do kafka vai enviar para o canal msgChan
		}
	}
}
