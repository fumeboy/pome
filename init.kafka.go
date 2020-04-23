package main

import (
	"github.com/Shopify/sarama"
	"github.com/fumeboy/pome/courier/mq"
	"github.com/fumeboy/pome/settings"
)

func init_kafka() (err error) {
	var this = settings.Kafka
	if !this.SwitchOn {
		return
	}
	this.Version, _ = sarama.ParseKafkaVersion("2.4.1")
	wg.Add(2)
	go func() {
		mq.InitConsumer()
		wg.Done()
	}()
	go func() {
		mq.InitProducer()
		wg.Done()
	}()
	return
}
