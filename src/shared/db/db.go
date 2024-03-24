package db_provider

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	users_models "github.com/levysantiago/gotickets/src/modules/users/models"
	wallets_models "github.com/levysantiago/gotickets/src/modules/wallets/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func init(){
	var err = godotenv.Load(".env")
	if(err != nil){
		log.Fatal("Error loading .env file")
	}

	dsn := "host="+os.Getenv("DB_HOST")+" user="+os.Getenv("DB_USER")+" password="+os.Getenv("DB_PASSWORD")+" dbname="+os.Getenv("DB_NAME")+" port="+os.Getenv("DB_PORT")+" sslmode=disable TimeZone=Asia/Shanghai"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if(err != nil){
		panic(err)
	}

	db = d

	db.AutoMigrate(&users_models.User{})
	db.AutoMigrate(&wallets_models.Wallet{})
}

func GetDB() *gorm.DB{
	return db
}