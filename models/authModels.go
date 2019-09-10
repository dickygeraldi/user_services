package models

import (
	"time"

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
