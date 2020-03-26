package mq

import (
	"fmt"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/fumeboy/pome/conf"
	"github.com/gogo/protobuf/proto"
	_ "math/rand"
)

func InitConsumer(handle func(v *MqMsg) error) {
	groupId := conf.NodeName()
	brokers := conf.Kafka.Addrs
	topics := []string{conf.NodeName()}
	config := cluster.NewConfig()
	config.Version = conf.Kafka.Version
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.CommitInterval = 1
	client, _ := cluster.NewClient(brokers, config)
	consumer, err := cluster.NewConsumerFromClient(client, groupId, topics)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	//  consume errors
	go func() {
		for err := range consumer.Errors() {
			fmt.Println("kafka consumer Error:" + err.Error())
		}
	}()

	//  consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			fmt.Println("Rebalanced:", ntf)
		}
	}()

	//  consume messages, watch signals
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {

				var m MqMsg
				if proto.Unmarshal(msg.Value, &m) == nil{
					handle(&m)
				}
				// fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") //  mark message as processed
			}
		case err := <-consumer.Errors():
			fmt.Println(err,123)
			return
		}
	}
}
