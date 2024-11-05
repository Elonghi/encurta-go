package caching

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var RedisClient *redis.Client

func StartRedis() {
	ctx := context.Background()

	// Carrega as variáveis de ambiente do arquivo .env, se existir
	if err := godotenv.Load(); err != nil {
		log.Println("Nenhum arquivo .env encontrado ou erro ao carregar o arquivo .env")
	}

	// Recupera as variáveis de ambiente
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUseTLS, err := strconv.ParseBool(os.Getenv("REDIS_USE_TLS"))
	if err != nil {
		log.Fatalf("Erro ao converter REDIS_USE_TLS para booleano: %v", err)
	}

	// Configurações básicas do Redis
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Username: redisUsername, // Nome de usuário, se houver
		Password: redisPassword, // Senha, se houver
		DB:       0,             // Banco de dados padrão
	}

	// Configurações TLS, se necessário
	if redisUseTLS {
		options.TLSConfig = &tls.Config{
			InsecureSkipVerify: false, // Altere para true se estiver usando certificados autoassinados
		}
	}

	// Cria o cliente Redis
	RedisClient = redis.NewClient(options)

	// Testa a conexão
	pong, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Erro ao conectar ao Redis: %v", err)
	}

	log.Println("Conexão com o Redis estabelecida com sucesso:", pong)
}
