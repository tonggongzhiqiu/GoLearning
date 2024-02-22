package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var ProductId = 1
var Conn = "127.0.0.1:9092"
var Topic = "carpool_service"

type ProducerClient struct {
	Producer  sarama.SyncProducer
	Topic     string
	ProductId int
	MessageId int
}

func (p *ProducerClient) InitProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{Conn}, config)
	if err != nil {
		fmt.Errorf("producer closed, err:%v", err)
		return
	}

	p.Producer = client
	p.Topic = Topic
	p.ProductId = ProductId
	p.MessageId = 1

	ProductId++
}

func (p *ProducerClient) SendMessage() {
	// 构造消息
	message := &sarama.ProducerMessage{}
	message.Topic = p.Topic
	message.Value = sarama.StringEncoder(fmt.Sprintf("message%d", p.MessageId))
	// 发送消息
	partition, offset, err := p.Producer.SendMessage(message)
	if err != nil {
		fmt.Printf("send message err=%+v", err)
	}
	fmt.Println("send message success! partition=%v, offset=%v", partition, offset)
	// 消息 id ++
	p.MessageId++
}

func (p *ProducerClient) close() {
	p.Producer.Close()
}
