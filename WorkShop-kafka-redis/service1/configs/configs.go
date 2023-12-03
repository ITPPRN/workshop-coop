package configs

type Config struct {
	App      Fiber
	Postgres PostgresSql
	Kafkas     Kafka
}

type Fiber struct {
	Port string
}

type PostgresSql struct {
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	SslMode      string
}

type Kafka struct {
	Hosts []string
	Group string
}
