package config

type (
	PostgresConfig struct {
		Timeout  int
		DBName   string
		Username string
		Password string
		Host     string
		Port     string
	}

	RedisConfig struct {
		Host string
		Port string
	}

	AWSConfig struct {
		AccessKey string
		SecretKey string
	}

	EnvConfig struct {
		Host     string
		Port     int
		Postgres PostgresConfig
		Redis    RedisConfig
		AWS		 AWSConfig
	}
)
