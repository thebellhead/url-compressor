package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	APPPort    string `env:"APP_PORT" envDefault:"1228"`
	RDHost     string `env:"REDIS_HOST" envDefault:"localhost"`
	RDPort     string `env:"REDIS_PORT" envDefault:"6379"`
	CHUser     string `env:"CLICKHOUSE_USER" envDefault:"keeley"`
	CHPassword string `env:"CLICKHOUSE_PASSWORD" envDefault:"electronics"`
	CHDBName   string `env:"CLICKHOUSE_DB" envDefault:"url_compressor"`
	CHHost     string `env:"CLICKHOUSE_HOST" envDefault:"localhost"`
	CHPort     string `env:"CLICKHOUSE_PORT" envDefault:"9000"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config from env")
	}
	return cfg, nil
}

func (cfg Config) GetAppAddress() string {
	return fmt.Sprintf(":%s", cfg.APPPort)
}

func (cfg Config) GetRedisDSN() string {
	return fmt.Sprintf("%s:%s", cfg.RDHost, cfg.RDPort)
}

func (cfg Config) GetClickHouseDSN() string {
	return fmt.Sprintf("%s:%s", cfg.CHHost, cfg.CHPort)
}
