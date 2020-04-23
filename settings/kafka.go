package settings

import "github.com/Shopify/sarama"

type kafkaConf struct {
	SwitchOn bool
	Addrs   []string
	Version sarama.KafkaVersion `yaml:"-"`
}
