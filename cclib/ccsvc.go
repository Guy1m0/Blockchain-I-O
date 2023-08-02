package cclib

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Guy1m0/Blockchain-I-O/examples/ecomm"
	"github.com/Shopify/sarama"
	kb "github.com/philipjkim/kafka-brokers-go"
	"github.com/wvanbergen/kafka/consumergroup"
)

const (
	logFile = "log.json"
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
// @todo Log each events
func (svc *CCService) Publish(event string, payload []byte) error {
	message := &sarama.ProducerMessage{Topic: event}
	message.Value = sarama.ByteEncoder(payload)
	_, _, err := svc.kafkaProducer.SendMessage(message)
	if err != nil {
		return err
	}

	LogEventToFile(event, payload)
	//data, err := unmarshalPayload(event, payload)

	log.Printf("Published event: %s\n", event)
	return nil
}

func unmarshalPayload(event string, payload []byte) (interface{}, error) {
	var data interface{}

	switch event {
	case "auction_ending":
		data = new(ecomm.Auction)
	case "auction_creating":
		data = new(ecomm.Auction)
	case "add_asset":
		data = new(ecomm.Asset)
	case "bid_auc":
		data = new(ecomm.Bid)
	case "tx_mined":
		data = new(ecomm.Tx)
	// more cases...
	default:
		return nil, fmt.Errorf("unknown event: %s", event)
	}

	err := json.Unmarshal(payload, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func LogEventToFile(event string, payload []byte) {
	// Create new file or append to existing one
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse payload into a map
	// var payloadMap interface{}
	// err = json.Unmarshal(payload, &payloadMap)
	// if err != nil {
	// 	panic(err)
	// }

	data, err := unmarshalPayload(event, payload)
	if err != nil {
		panic(err)
	}

	// Create new event
	newEvent := map[string]interface{}{
		"Event": event,
		"Time":  time.Now().Format(time.RFC3339),
		"Data":  data,
	}

	// Convert the event to JSON
	eventJSON, err := json.MarshalIndent(newEvent, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the event to the file, followed by a newline
	if _, err := file.Write(append(eventJSON, '\n')); err != nil {
		panic(err)
	}
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
