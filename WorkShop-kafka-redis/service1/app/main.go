package main

import (
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"

	"service1/configs"
	"service1/modules/servers"
	databases "service1/pkg/databases/postgres"
)

func main() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	cfg := new(configs.Config)

	//App env
	cfg.App.Port = os.Getenv("APP_PORT")

	//postgres env
	cfg.Postgres.Host = os.Getenv("DB_HOST")
	cfg.Postgres.Port = os.Getenv("DB_PORT")
	cfg.Postgres.Username = os.Getenv("DB_USER")
	cfg.Postgres.Password = os.Getenv("DB_PASSWORD")
	cfg.Postgres.DatabaseName = os.Getenv("DB_NAME")
	cfg.Postgres.SslMode = os.Getenv("DB_SSLMODE")

	// Kafka
	cfg.Kafkas.Hosts = []string{os.Getenv("KAFKA_BROKER_1")}
	cfg.Kafkas.Group = os.Getenv("KAFKA_GROUP")

	db, err := databases.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	producer, err := sarama.NewSyncProducer(cfg.Kafkas.Hosts, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumerGroup(cfg.Kafkas.Hosts, cfg.Kafkas.Group, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	server := servers.NewServer(cfg,db,consumer,producer)
	server.Start()
}
