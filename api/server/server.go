package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MudassirDev/all-in-one-business-tool/api/controllers"
	"github.com/MudassirDev/all-in-one-business-tool/internal/database"
	"github.com/MudassirDev/all-in-one-business-tool/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	JWTEXPIRINGTIME = 1 * time.Hour
)

func CreateServer() *http.Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	dbUrl := os.Getenv("DB_URL")
	secretKey := os.Getenv("JWT_SECRET_KEY")

	if dbUrl == "" {
		log.Fatal("DB url must not be empty")
	}

	mux := http.NewServeMux()

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("failed to establish a connection with database")
	}

	queries := database.New(conn)

	apiCfg := models.APICfg{
		DB:              queries,
		JWTSecretKey:    secretKey,
		JWTExpiringTime: JWTEXPIRINGTIME,
	}

	mux.HandleFunc("POST /register", controllers.CreateUserRegisterHandler(&apiCfg))
	mux.HandleFunc("POST /login", controllers.CreateUserLoginHandler(&apiCfg))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
