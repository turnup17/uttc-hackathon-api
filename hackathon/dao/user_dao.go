package dao

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func init() {
	// Load environment variables from .env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env_mysql file: %v\n", err)
	}

	// Get MySQL connection details from environment variables
	dbUser := os.Getenv("MYSQL_USER")
	dbPwd := os.Getenv("MYSQL_PASSWORD")
	instanceConnectionName := os.Getenv("MYSQL_HOST")
	dbName := os.Getenv("MYSQL_DATABASE")

	fmt.Printf("TEST %s, %s, %s\n", dbUser, instanceConnectionName, dbName)

	dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", dbUser, dbPwd, instanceConnectionName, dbName)
	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return
	}
	Db = dbPool
}

// CloseDBWithSysCall closes the database connection on Ctrl+C or SIGTERM
func CloseDBWithSysCall() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-sig
		log.Printf("received syscall, %v", s)
		if err := Db.Close(); err != nil {
			log.Fatal(err)
		}
		log.Printf("success: db.Close()")
		os.Exit(0)
	}()
}
