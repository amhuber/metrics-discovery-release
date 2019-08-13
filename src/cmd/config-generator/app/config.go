package app

import (
	"code.cloudfoundry.org/go-envstruct"
	"log"
	"time"
)

type Config struct {
	NatsHosts                []string      `env:"NATS_HOSTS,              required, report"`
	ScrapeConfigFilePath     string        `env:"SCRAPE_CONFIG_FILE_PATH, required, report"`
	ConfigExpirationInterval time.Duration `env:"CONFIG_EXPIRATION_INTERVAL,        report"`
	ConfigTimeToLive         time.Duration `env:"CONFIG_TTL,                        report"`

	ScrapeCAPath   string `env:"SCRAPE_CA_PATH"`
	ScrapeCertPath string `env:"SCRAPE_CERT_PATH"`
	ScrapeKeyPath  string `env:"SCRAPE_KEY_PATH"`

	MetricsPort     int    `env:"METRICS_PORT, report"`
	MetricsCAPath   string `env:"METRICS_CA_PATH"`
	MetricsCertPath string `env:"METRICS_CERT_PATH"`
	MetricsKeyPath  string `env:"METRICS_KEY_PATH"`
}

func LoadConfig(log *log.Logger) Config {
	cfg := Config{
		ConfigExpirationInterval: 15 * time.Second,
		ConfigTimeToLive:         45 * time.Second,
	}

	if err := envstruct.Load(&cfg); err != nil {
		log.Fatal(err)
	}

	envstruct.WriteReport(&cfg)

	return cfg
}
