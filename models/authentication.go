package models

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"user_services/base"

	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/minio/highwayhash"
)

// Initializar .env variable
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// Kafka Configuration
func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}

// Kafka producer
func SendMessage(topic, msg string) map[string]interface{} {

	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"localhost:9200"}, kafkaConfig)

	if err != nil {
		return gin.H{
			"error": err,
		}
	}

	defer func() {
		if err := producers.Close(); err != nil {
			panic(err)
		}
	}()

	kafka := &KafkaProducer{
		Producer: producers,
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := kafka.Producer.SendMessage(kafkaMsg)
	if err != nil {
		return gin.H{
			"error": err,
		}
	}

	fmt.Println("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}

// Func get random unique data
func GetRandomString() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 15)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

// validasi data
func (data *AccountData) Validate() (map[string]interface{}, bool) {
	var countPhone, countEmail, countUsername int
	if !strings.Contains(data.Email, "@") {
		return gin.H{"status": false, "message": "Email address is required"}, false
	}

	if len(data.Password) < 6 {
		return gin.H{"status": false, "message": "Password is required"}, false
	}

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	if re.MatchString(data.Phone) == false {
		return gin.H{"status": false, "message": "Phone number must be number"}, false
	}

	base.GetDB().QueryRow("select count(*) from creator where phone = $1", data.Phone).Scan(&countPhone)
	if countPhone > 0 {
		return gin.H{"status": false, "message": "Phone number has been used"}, false
	}

	base.GetDB().QueryRow("select count(*) from creator where username = $1", data.Username).Scan(&countUsername)
	if countUsername > 0 {
		return gin.H{"status": false, "message": "Username has been used"}, false
	}

	base.GetDB().QueryRow("select count(*) from creator where email = $1", data.Email).Scan(&countEmail)
	if countEmail > 0 {
		return gin.H{"status": false, "message": "Email has been used"}, false
	}

	return gin.H{"status": false, "message": "Requirement passed"}, true
}

// Authentikasi untuk create account
func (data *AccountData) CreateCreator() map[string]interface{} {

	if resp, ok := data.Validate(); !ok {
		return resp
	}

	// Send to kafka
	accountByte, _ := json.Marshal(data)
	accountJSONString := string(accountByte)

	go SendMessage("mailing_service", accountJSONString)

	// Add parameter
	data.UserId = GetRandomString()
	// go LoggingAddDetails(data.UserId)

	// Hashing Password
	key := os.Getenv("password_token")

	hashPassword := highwayhash.Sum([]byte(data.Password), []byte(key))
	data.Password = hex.EncodeToString(hashPassword[:])

	// Create an sql and insert
	_, err := base.GetDB().Query("insert into creator (userId, username, phone, password, email, fullname, source) values ($1, $2, $3, $4, $5, $6, $7)",
		data.UserId, data.Username, data.Phone, data.Password, data.Email, data.FullName, data.Source)

	if err != nil {
		return gin.H{
			"error": err,
		}
	}

	// Create response
	response := gin.H{"Status": true, "Message": "Data berhasil masuk"}

	return response
}

func (data *AccountData) Data() map[string]interface{} {
	response := gin.H{"Status": true}
	response["data"] = data
	return response
}
