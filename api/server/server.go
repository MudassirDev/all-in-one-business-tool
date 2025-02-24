package server

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/MudassirDev/all-in-one-business-tool/api/controllers"
	"github.com/MudassirDev/all-in-one-business-tool/internal/database"
	"github.com/MudassirDev/all-in-one-business-tool/models"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func CreateServer() *http.Server {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	dbUrl := os.Getenv("DB_URL")

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
		DB: queries,
	}

	mux.HandleFunc("/register", controllers.CreateUserRegisterHandler(&apiCfg))
	mux.HandleFunc("/login", controllers.CreateUserLoginHandler(&apiCfg))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return srv
}
