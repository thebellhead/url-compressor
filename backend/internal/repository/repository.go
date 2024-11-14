package repository

import (
	"context"

	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/redis/go-redis/v9"

	"log"

	"github.com/lithammer/shortuuid"
)

type Config struct {
	RedisDSN           string
	ClickHouseDSN      string
	ClickHouseDBName   string
	ClickHouseUsername string
	ClickHousePassword string
}

type CompressorRepository struct {
	cache *redis.Client
	db    driver.Conn
}

func New(cfg Config) (*CompressorRepository, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisDSN,
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("cannot ping redis: %v", err)
	}
	ch, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ClickHouseDSN},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouseDBName,
			Username: cfg.ClickHouseUsername,
			Password: cfg.ClickHousePassword,
		},
	})
	if err != nil {
		log.Fatalf("cannot initialize clickhouse connection: %v", err)
	}
	if err = ch.Ping(ctx); err != nil {
		log.Fatalf("cannot ping clickhouse: %v", err)
	}
	return &CompressorRepository{
		cache: rdb,
		db:    ch,
	}, nil
}

func (cr *CompressorRepository) GetURL(shortURL string) (string, error) {
	ctx := context.Background()
	longURL, err := cr.getFromCache(ctx, shortURL)
	if err == redis.Nil {
		log.Printf("compressed url not found in cache")
		longURL, err = cr.getFromDB(ctx, shortURL)
		if err != nil {
			return "", err
		}
	}
	return longURL, nil
}

func (cr *CompressorRepository) getFromCache(ctx context.Context, shortURL string) (string, error) {
	longURL, err := cr.cache.Get(ctx, shortURL).Result()
	if longURL != "" {
		return longURL, nil
	}
	return "", err
}

func (cr *CompressorRepository) getFromDB(ctx context.Context, shortURL string) (string, error) {
	row := cr.db.QueryRow(
		ctx,
		`SELECT long_url FROM url_compressor.url_map WHERE short_url == $1`,
		shortURL,
	)
	var longURL string
	if err := row.Scan(&longURL); err != nil {
		return "", fmt.Errorf("cannot parse clickhouse resulting row for long_url")
	}
	return longURL, nil
}

func (cr *CompressorRepository) PostURL(longURL string) (string, error) {
	ctx := context.Background()
	shortURL := shortuuid.NewWithNamespace(longURL)
	gotLongURL, _ := cr.getFromCache(ctx, shortURL)
	if gotLongURL != "" {
		log.Printf("compressed url %s already exists in cache", shortURL)
		return shortURL, nil
	}
	_, err := cr.insertCache(ctx, shortURL, longURL)
	if err != nil {
		return "", err
	}
	gotLongURL, _ = cr.getFromDB(ctx, shortURL)
	if gotLongURL != "" {
		log.Printf("compressed url %s already exists in database", shortURL)
		return shortURL, nil
	}
	_, err = cr.insertDB(ctx, shortURL, longURL)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (cr *CompressorRepository) insertCache(ctx context.Context, shortURL, longURL string) (string, error) {
	err := cr.cache.Set(ctx, shortURL, longURL, 0).Err()
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (cr *CompressorRepository) insertDB(ctx context.Context, shortURL, longURL string) (string, error) {
	log.Printf("insertDB %s %s", shortURL, longURL)
	_ = cr.db.QueryRow(
		ctx,
		`INSERT INTO url_compressor.url_map (short_url, long_url) VALUES ($1, $2)`,
		shortURL,
		longURL,
	)
	return shortURL, nil
}
