#!/bin/bash

echo "##########################################";
echo "# local Kafka Docker 环境管理            "
echo "# 1 状态                                "
echo "# 2 创建 topic                          "
echo "# 3 向 topic 中插入消息                  "
echo "# 4 消费 topic 中的消息 - from-beginning "
echo "# 5 zookeeper shell"
echo "##########################################";
read -p "请输入对应操作编号：" op

#!/bin/bash
containerId=kafka
topicName=guestbook
zookeeperIP=zookeeper:2181

function getContainerId()
{
    read -p "请输入容器ID：" containerId
    export containerId
}

function getTopicName()
{
    read -p "请输入Topic名称：" topicName
    export topicName
}

function getContainerIp()
{
    read -p "请输入容器IP：" containerIp
    export containerIp
}

function status()
{
    getContainerId
    docker exec -it ${containerId} kafka-topics.sh --list --zookeeper ${zookeeperIP}
}

function createTopic()
{
    getContainerId
    getTopicName
    docker exec -it ${containerId} kafka-topics.sh --create --zookeeper ${zookeeperIP} --replication-factor 1 --partitions 1 --topic ${topicName}
}

function insertMessage()
{
    getContainerId
    getContainerIp
    getTopicName
    docker exec -it ${containerId} kafka-console-producer.sh --broker-list ${containerIp}:9092 --topic ${topicName}
}

function consumeMessageFromBeginning()
{
    getContainerId
    getContainerIp
    getTopicName
    docker exec -it ${containerId} kafka-console-consumer.sh --bootstrap-server ${containerIp}:9092 --topic ${topicName} --from-beginning
}

function zkShell()
{
    getContainerId
    docker exec -it ${containerId} zookeeper-shell.sh ${zookeeperIP}
}

case $op in
    1)
	status
	exit 0
	;;
    2)
	createTopic
	exit 0
	;;
    3)
	insertMessage
	exit 0
	;;
    4)
	consumeMessageFromBeginning
	exit 0
	;;
    5)
	zkShell
	exit 0
	;;
esac