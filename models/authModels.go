package models

import (
	"time"
	"user_services/base"

	"github.com/Shopify/sarama"
	"github.com/dgrijalva/jwt-go"
)

type AccountData struct {
	UserId      string
	Username    string
	Phone       string
	Password    string
	Email       string
	FullName    string
	Source      string
	FbId        string
	GoogleId    string
	ProfilePict string
	Token       string
	OtpAuth     string
}

type CreatorBalance struct {
	IdCreatorBalance string
	UserId           string
	Balance          float32
	pin              string
	IsBlocked        int
	IsActive         int
	Point            int
}

type CreatorDetails struct {
	UserId     string
	About      string
	UrlLink    string
	BirthDate  string
	IsActive   string
	Profession string
	Gender     int
	IsPrivate  int
}

type LogBalance struct {
	Id               string
	IdCreatorBalance string
	Timestamps       string
	Content          string
	Status           string
	BalanceNow       string
	BalanceBefore    string
}

type Token struct {
	UserId     string
	FbId       string
	GoogleId   string
	Username   string
	Timestamps time.Time
	jwt.StandardClaims
}
type KafkaProducer struct {
	Producer sarama.SyncProducer
}

// create table creator (userId integer primary key, username varchar(20) unique NOT NULL, phone varchar(13) unique NOT NULL, password varchar(32) NOT NULL, email varchar(50) unique NOT NULL, fullName varchar(40) NOT NULL, source varchar(30) NOT NULL, fbId varchar(20) null, googleId varchar(20) null, profilePict varchar(40) null, token varchar(500) null, otpAuth varchar(6) null)
// ALTER TABLE creatordetails (userid varchar(20) primary key, about varchar(200) null, urllink varchar(100) null, birtdate varchar(50) null, isactive varchar(1) null, isactive varchar(1) null, gender varchar(1) null, isprivate(1) null) index

// Insert data to mongodb for log and details user
func LoggingAddDetails(user_id string) {
	// Add to Creator Balance and Creator Details
	var isBlocked, isActive, balance, point, isprivate = 0, 0, 0, 0, 0
	var about = "This is about you, so describe who are you?"

	idCBalance := GetRandomString()
	_, err := base.GetDB().Query("insert into creatorbalance (idcreatorbalance, userid, balance, pin, isblocked, isactive, point) values($1, $2, $3, $4, $5, $6)",
		idCBalance, user_id, balance, "", isBlocked, isActive, point)

	if err != nil {
		panic(err)
	}

	_, err = base.GetDB().Query("insert into creatordetails (userid, about, isactive, isprivate) values($1, $2, $3, $4)",
		user_id, about, isActive, isprivate)

	if err != nil {
		panic(err)
	}

	// Add to creator log using mongodb

}
