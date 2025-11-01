package main

import (
	"database/sql"
	"fmt"
	"gintugas/database"
	routers "gintugas/modules"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		panic("Error loading .env file")
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	db, err = sql.Open("postgres", psqlInfo)
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Berhasil Koneksi Ke Database")

	database.DBMigrate(db)

	InitiateRouter(db)
}

func InitiateRouter(db *sql.DB) {
	router := gin.Default()

	routers.Initiator(router, db)

	router.Run(":8080")

}
