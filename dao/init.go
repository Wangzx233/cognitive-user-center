package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

var (
	remoteUsername = "root"
	remotePassword = "root"
	remoteHost     = "mariadb"
	remotePort     = "3306"
	remoteDbname   = "cognitive_user_center"
	username       = "root"
	password       = "root"
	host           = "localhost"
	port           = "3306"
	dbname         = "cognitive_user_center"
)

func InitMysql() {
	dsn := username + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn1 := remoteUsername + ":" + remotePassword + "@(" + remoteHost + ":" + remotePort + ")/" + remoteDbname + "?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}

	db = Db

	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println(err)
		return
	}

}

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		log.Println(err)
	}
}
