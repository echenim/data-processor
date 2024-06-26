package config

type Config struct {
	DBConnectionString string
	PubSubSubscription string
}

func Load(conn_string, scan_topic string) (*Config, error) {
	return &Config{
		DBConnectionString: conn_string,
		PubSubSubscription: scan_topic,
	}, nil
}
