package cclib

import (
	"log"
	"strings"
	"time"

	"github.com/Shopify/sarama"
	kb "github.com/philipjkim/kafka-brokers-go"
	"github.com/wvanbergen/kafka/consumergroup"
)

type EventHandler func(payload []byte)

type CCService struct {
	serviceID string
	zkNodes   []string

	kafkaProducer sarama.SyncProducer

	handlers map[string]EventHandler
}

// creates a new instance of CCService, initializes the Kafka producer, and returns it.
func NewEventService(zkNodes []string, serviceID string) (*CCService, error) {
	svc := &CCService{
		serviceID: serviceID,
		zkNodes:   zkNodes,
		handlers:  make(map[string]EventHandler),
	}
	err := svc.setupKafkaProducer()
	if err != nil {
		return nil, err
	}
	return svc, nil
}

// adds a new event handler to the CCService's map of handlers.
func (svc *CCService) Register(event string, handler EventHandler) {
	svc.handlers[event] = handler
}

// sends a message (event and its payload) to the Kafka cluster.
func (svc *CCService) Publish(event string, payload []byte) error {
	message := &sarama.ProducerMessage{Topic: event}
	message.Value = sarama.ByteEncoder(payload)
	_, _, err := svc.kafkaProducer.SendMessage(message)
	if err != nil {
		return err
	}
	log.Printf("Published event: %s\n", event)
	return nil
}

func (svc *CCService) setupKafkaProducer() error {
	kbConn, err := kb.NewConn(svc.zkNodes)
	if err != nil {
		return err
	}
	brokerList, _, err := kbConn.GetW()
	if err != nil {
		return err
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewManualPartitioner

	svc.kafkaProducer, err = sarama.NewSyncProducer(brokerList, config)
	return err
}

// starts the CCService by creating a consumer that subscribes to the
// Kafka topics associated with the registered events. If the createTopic
// parameter is true, it attempts to create the topics in the Kafka cluster.

func (svc *CCService) Start(createTopic bool) error {
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetNewest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	topics := svc.topics()
	if createTopic {
		err := svc.publishEmptyEventsToCreateTopics(topics)
		if err != nil {
			return err
		}
	}

	consumer, err := consumergroup.JoinConsumerGroup(
		svc.serviceID, topics, svc.zkNodes, config,
	)
	if err != nil {
		return err
	}

	if createTopic {
		svc.listenAndCommitEmptyEvents(consumer)
	}

	go svc.listenKafkaConsumer(consumer)
	log.Printf("Subscribed topics: %s\n", strings.Join(topics, ","))

	return nil
}

// create new topics in the Kafka cluster by sending an empty message to each topic.
// Kafka automatically creates a topic if a message is produced to a non-existent
// topic and the server is configured to allow it (which is the default configuration).
func (svc *CCService) publishEmptyEventsToCreateTopics(topics []string) error {
	for _, t := range topics {
		// ignore error for the first time
		svc.kafkaProducer.SendMessage(&sarama.ProducerMessage{Topic: t})
	}
	for _, t := range topics {
		_, _, err := svc.kafkaProducer.SendMessage(&sarama.ProducerMessage{Topic: t})
		if err != nil {
			return err
		}
	}
	time.Sleep(3 * time.Second)
	return nil
}

// listens for the empty events sent by publishEmptyEventsToCreateTopics and commits them
func (svc *CCService) listenAndCommitEmptyEvents(consumer *consumergroup.ConsumerGroup) {
	timer := time.NewTimer(3 * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return
		case message := <-consumer.Messages():
			consumer.CommitUpto(message)
		}
	}
}

// continuously listens for incoming messages from Kafka, handles each message by
// calling the appropriate event handler, and commits the message.
func (svc *CCService) listenKafkaConsumer(consumer *consumergroup.ConsumerGroup) {
	for message := range consumer.Messages() {
		handler, ok := svc.handlers[message.Topic]
		if !ok {
			continue
		}
		log.Printf("Received event: %s\n", message.Topic)
		go handler(message.Value)
		consumer.CommitUpto(message)
	}
}

func (svc *CCService) topics() []string {
	result := make([]string, 0, len(svc.handlers))
	for event := range svc.handlers {
		result = append(result, event)
	}
	return result
}
