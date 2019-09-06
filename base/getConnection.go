package base

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("user")
	password := os.Getenv("password")
	dbName := os.Getenv("dbname")
	dbHost := os.Getenv("hostname")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=5432", dbHost, username, dbName, password)

	conn, err := sql.Open("postgres", dbUri)

	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

func GetDB() *sql.DB {
	return db
}
