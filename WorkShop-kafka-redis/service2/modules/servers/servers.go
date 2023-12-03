package servers

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"

	"service2/configs"
	"service2/modules/entities/events"
	"service2/pkg/utils"
)

type server struct {
	App                  *fiber.App
	Db                   *gorm.DB
	Cfg                  *configs.Config
	ConsumerGroup        sarama.ConsumerGroup
	SyncProducer         sarama.SyncProducer
	consumerGroupHandler sarama.ConsumerGroupHandler
	Redis                *redis.Client
}

func NewServer(
	cfg *configs.Config,
	db *gorm.DB,
	consumerGroup sarama.ConsumerGroup,
	syncProducer sarama.SyncProducer,
	redis *redis.Client,
) *server {
	return &server{
		App:           fiber.New(),
		Db:            db,
		Cfg:           cfg,
		ConsumerGroup: consumerGroup,
		SyncProducer:  syncProducer,
		Redis:         redis,
	}
}

func (s *server) Start() {
	if err := s.Handlers(); err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.UrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	// Start consumer
	go func() {
		log.Info("Connect to kafa server:", s.Cfg.Kafkas.Hosts, ",Group:", s.Cfg.Kafkas.Group)
		log.Info("Subscribed topics:", events.SubscribedTopics)
		for {
			s.ConsumerGroup.Consume(context.Background(), events.SubscribedTopics, s.consumerGroupHandler)
		}
	}()

	port := s.Cfg.App.Port
	log.Info("server  started on localhost: ", port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}
}
