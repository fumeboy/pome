package mq

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fumeboy/pome/conf"
	"github.com/golang/protobuf/proto"
)

var producer_ sarama.AsyncProducer

func InitProducer() {
	//if true {
	//	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)
	//}
	config := sarama.NewConfig()
	config.Version = conf.Kafka.Version
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	//使用配置,新建一个异步生产者
	producer, e := sarama.NewAsyncProducer(conf.Kafka.Addrs, config)
	if e != nil {
		fmt.Println(e)
		return
	}
	defer producer.AsyncClose()
	producer_ = producer
	//循环判断哪个通道发送过来数据.
	func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				if fail != nil{
					fmt.Println("1 err: ", fail.Err)
				}else{
					fmt.Println("2 errr")
				}
			}
		}
	}(producer)
}

func Send(TargetTopic string, Value *MqMsg) {
	buf, _ := proto.Marshal(Value)
	msg := &sarama.ProducerMessage{
		Topic: TargetTopic,
		Value: sarama.ByteEncoder(buf),
	}
	//使用通道发送
	producer_.Input() <- msg
}
