package main

import (
	"log"
	"mongodb-connection/internal/data"
	"mongodb-connection/internal/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Connect to MongoDB and get the client
	client, err := db.ConnectToMongo()
	if err != nil {
		log.Panic("MongoDB connection failed:", err)
	}
	defer db.DisconnectMongo()

	// Create a new instance of LogEntry using the client
	logEntry := data.New(client)

	// Insert a new log entry
	err = logEntry.LogEntry.Insert(data.LogEntry{
		Name: "New Log",
		Data: "This is a test data.",
	})
	if err != nil {
		log.Println("Error adding data:", err)
	} else {
		log.Println("Data inserted successfully")
	}

	// Try to retrieve the log entry by ObjectID
	entry, err := logEntry.LogEntry.GetOne("6761474b992b5c980fe8fcde") // Example ObjectID
	if err != nil {
		log.Println("Error retrieving data:", err)
	} else {
		log.Printf("Retrieved data: %+v\n", entry)
	}

	// Wait for application shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down the application...")
}

/*
	mongodb://admin:password@localhost:64001/logs?authSource=admin&readPreference=primary&appname=MongDB%20Compass&&directConnection=true&ssl=false
*/
