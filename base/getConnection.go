package base

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *sql.DB

var dbMongo *mongo.Client

func init() {

	// Postgresql connect
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

	// // MongoDb Connect
	mongoUser := base64.URLEncoding.EncodeToString([]byte(os.Getenv("mongoUser")))
	mongoPass := base64.StdEncoding.EncodeToString([]byte(os.Getenv("mongoPwd")))
	dbConn := fmt.Sprintf("mongodb://%s:%s@127.0.0.1:27017/?authSource=admin", mongoUser, mongoPass)
	fmt.Println(dbConn)

	clientOption := options.Client().ApplyURI(dbConn)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	dbMongo = client
	fmt.Println("Connected to MongoDB!")
}

func GetDB() *sql.DB {
	return db
}

func GetMongo() *mongo.Client {
	return dbMongo
}
