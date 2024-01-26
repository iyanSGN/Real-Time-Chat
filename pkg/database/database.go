package database

import (
	"fmt"
	"log"
	"os"
	"realtime/models"
	"realtime/pkg/file"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

type credential struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Name     string
	Timezone string
}


func Init(driver string) {
	credential := &credential{
		Host: Get("DB_HOST"),
		Port: Get("DB_PORT"),
		User: Get("DB_USER"),
		Pass: Get("DB_PASSWORD"),
		Name: Get("DB_NAME"),
		Timezone: Get("TIMEZONE_TZ"),
	}

	switch driver {
	case "postgres":
		credential.getPostgres()
	default:
		credential.getPostgres()
	}
}

func DBManager() *gorm.DB {
	return db
}

func (c *credential) getPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s TimeZone=%s", c.Host, c.Port, c.User, c.Pass, c.Name, c.Timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Print(c.Host)
}

func Migrate() {
	db.AutoMigrate(
		models.User{},
		models.Message{},
	)
}


func EnvInit() {
	var err error

	os.Setenv("PROJECT_DIR", "serelo-backend")
	rootPath := file.GetRootDirectory()

	envFilePath := rootPath + ".env"

	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatal("Errorrr loading .env file : ", err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}