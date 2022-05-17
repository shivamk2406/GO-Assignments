package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/shivamk2406/newsletter-subscriptions/internal/models"
	pb "github.com/shivamk2406/newsletter-subscriptions/internal/proto"
	"github.com/shivamk2406/newsletter-subscriptions/internal/service/user"
)

type ProducerConsumerService struct {
	serv      user.UserManagementServer
	Topic     string
	GroupId   string
	BrokerUrl []string
}
type UserNews struct {
	user    models.User
	allnews models.News
}

func NewProducerConsumerService(serv user.UserManagementServer, topic string, Groupid string,
	brokerUrl []string) *ProducerConsumerService {
	return &ProducerConsumerService{serv: serv, Topic: topic, GroupId: Groupid, BrokerUrl: brokerUrl}
}

func (pcs *ProducerConsumerService) GetNewKafkaReader(kafkaUrl []string, topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{Brokers: kafkaUrl, Topic: topic, GroupID: groupID})
}

func (pcs *ProducerConsumerService) GetNewKafkaWriter(kafkaUrl []string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{Brokers: kafkaUrl,
		Topic:        topic,
		RequiredAcks: -1,
		MaxAttempts:  3})
}

func (pcs *ProducerConsumerService) ConsumeNews(ctx context.Context, topic string, groupID string) {
	r := pcs.GetNewKafkaReader(pcs.BrokerUrl, topic, groupID)
	defer func() {
		if err := r.Close(); err != nil {
			log.Println(err)

		}
	}()
	log.Println("Starting consumer group")
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go pcs.createNewsWorker(ctx, r, wg)
	wg.Wait()
}

func (pcs *ProducerConsumerService) NewsProducer(ctx context.Context, topic string) {
	w := pcs.GetNewKafkaWriter(pcs.BrokerUrl, topic)
	defer func() {
		if err := w.Close(); err != nil {
			log.Println(err)

		}
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go pcs.ProduceNews(ctx, w, wg)
	wg.Wait()

}

func (pcs *ProducerConsumerService) createNewsWorker(ctx context.Context, r *kafka.Reader, wg *sync.WaitGroup) {
	defer r.Close()
	defer wg.Done()

	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Printf("FetchMessage %v", err)
			return
		}

		var news pb.News

		if err := json.Unmarshal(m.Value, &news); err != nil {
			log.Printf("json.Unmarshal %v", err)
			continue
		}
		if err := r.CommitMessages(ctx, m); err != nil {
			log.Printf("FetchMessage %v", err)
			continue
		}
		fmt.Printf("News Receieved %v \n", news.News)
	}
}

func (pcs *ProducerConsumerService) ProduceNews(ctx context.Context, w *kafka.Writer, wg *sync.WaitGroup) {
	topic := "my-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	defer wg.Done()

	// conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	// w := &kafka.Writer{
	// 	Addr:         kafka.TCP("localhost:9092"),
	// 	Topic:        topic,
	// 	Balancer:     &kafka.LeastBytes{},
	// 	RequiredAcks: -1,
	// 	MaxAttempts:  3,
	// }

	newss, err := pcs.serv.ListNews(ctx, &pb.ListNewsRequest{Subsid: 1})
	if err != nil {
		fmt.Println(err)
	}

	newsBytes, err := json.Marshal(&newss)
	if err != nil {
		log.Fatal(err)
	}

	msg := kafka.Message{
		Value: newsBytes,
	}
	err = w.WriteMessages(ctx, msg)
	if err != nil {
		fmt.Println(err)
	}
	// _, err = conn.WriteMessages(
	// 	kafka.Message{Value: []byte("one!")},
	// 	kafka.Message{Value: []byte("two!")},
	// 	kafka.Message{Value: []byte("three!")},
	// )
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
