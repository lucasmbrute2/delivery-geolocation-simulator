package kafka

import (	
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lucasmbrute2/infra/kafka"
	route2 "github.com/lucasmbrute2/application/route"
	"os"
	"log"
	"encoding/json"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}