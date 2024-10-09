package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type RedisConfig struct {
	Host string
	Port string
}

type Server struct {
	echo         *echo.Echo
	queryService *QueryService
	cacheService *CacheService
}

func NewServer(dbConfig DBConfig, redisConfig RedisConfig) (*Server, error) {
	e := echo.New()

	queryService, err := NewQueryService(dbConfig)
	if err != nil {
		return nil, err
	}

	cacheService, err := NewCacheService(redisConfig)
	if err != nil {
		return nil, err
	}

	s := &Server{
		echo:         e,
		queryService: queryService,
		cacheService: cacheService,
	}

	s.registerRoutes()

	return s, nil
}

func (s *Server) registerRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Website Monitor!")
	})
}

func (s *Server) Start(address string) error {
	return s.echo.Start(address)
}

func (s *Server) Close() {
	if err := s.queryService.Close(); err != nil {
		log.Printf("Error closing query service: %v", err)
	}
	if err := s.cacheService.Close(); err != nil {
		log.Printf("Error closing cache service: %v", err)
	}
}

func WaitForDB(config DBConfig) error {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Name)

	var db *sqlx.DB
	var err error

	for i := 0; i < 30; i++ {
		db, err = sqlx.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return nil
			}
		}
		log.Printf("Waiting for database... (%d/30): %v", i+1, err)
		time.Sleep(1 * time.Second)
	}
	return fmt.Errorf("database connection failed after 30 attempts: %v", err)
}
