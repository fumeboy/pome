package conf

import "github.com/Shopify/sarama"

type kafkaConf struct {
	SwitchOn bool
	Addrs   []string
	Version sarama.KafkaVersion `yaml:"-"`
}

func (this *kafkaConf) init() (err error) {
	if !this.SwitchOn {
		return
	}
	this.Version, _ = sarama.ParseKafkaVersion("2.4.1")
	return
}
