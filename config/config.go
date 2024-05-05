package config

import (
	"os"
	"strconv"
	"sync"

	// "github.com/bhagas/go-svc-echo/config/cache"
	"github.com/bhagas/go-svc-echo/config/database"
)

type config struct {
	dbGorm database.GormDatabase
	port   int

	// redis  cache.Redis
}

type Config interface {
	ServiceName() string
	DB() database.GormDatabase
	// ES() elastic.ElasticSearch
	// Redis() cache.Redis
	Port() int
	ENV() string
}

func NewConfig() Config {
	cfg := new(config)
	cfg.connectDB()
	// cfg.connectES()
	// cfg.connectRedis()
	return cfg
}

func (c *config) ServiceName() string {
	return os.Getenv(`SERVICE_NAME`)
	// return ""
}

func (c *config) connectDB() {
	var loadonce sync.Once
	loadonce.Do(func() {
		c.dbGorm = database.InitGorm()
	})
}

func (c *config) DB() database.GormDatabase {
	return c.dbGorm
}

// func (c *config) connectES() {
// 	var loadonce sync.Once
// 	loadonce.Do(func() {
// 		c.es = elastic.InitElasticSearch()
// 	})
// }

// func (c *config) ES() elastic.ElasticSearch {
// 	return c.es
// }

// func (c *config) connectRedis() {
// 	var loadonce sync.Once
// 	loadonce.Do(func() {
// 		c.redis = cache.InitRedis()
// 	})
// }

// func (c *config) Redis() cache.Redis {
// 	return c.redis
// }

func (c *config) Port() int {
	v := os.Getenv("PORT")
	c.port, _ = strconv.Atoi(v)

	return c.port
}

func (c *config) ENV() string {
	return os.Getenv(`ENVIRONTMENT`)
}
