package db

// TODO: $ go mod tidy
import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB への接続
func NewDB() *gorm.DB {
	// .env を使用可能に
	err := godotenv.Load(".env")

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// URL をもとにDBへ接続
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	fmt.Println(url)
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil { log.Fatalln(err) }
	fmt.Println("Connected")
	return db
}

// DB との接続解除
func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil { log.Fatalln(err) }
}