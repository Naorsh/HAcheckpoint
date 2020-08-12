package main

import (
	"log"
	"os"

	"checkpointapp/internal/pkg/db"
	"checkpointapp/internal/pkg/server"
)

const (
	port          = "8123"
	awsConnString = "dbname=haCheckpoint user=postgres password=passw0rd host=database-1.cehqo07zdaiy.us-east-2.rds.amazonaws.com port=5432 sslmode=disable"
)

func main() {
	defer db.Disconnect()
	log.Printf("Connecting to database...")
	connStr := os.Getenv("PG_URL")
	if connStr == "" {
		log.Println("Could not connect to local database. Error: missing PG_URL env variable connecting to fallback AWS RDS instance")
		connStr = awsConnString
	}
	err := db.Connect(connStr)
	if err != nil {
		log.Fatalf("Could not connect to database. Error: %s", err)
	}
	log.Println("Datababse connected")

	log.Println("Starting HTTP server...")
	log.Fatal(server.Start(port))
}
