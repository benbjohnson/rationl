package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/benbjohnson/rationl"
	"github.com/benbjohnson/rationl/handlers"
)

func main() {
	log.SetFlags(0)
	var (
		addr     = flag.String("addr", ":9000", "bind address")
		dataDir  = flag.String("data-dir", "", "data directory")
		clientID = flag.String("client-id", "", "GitHub application client id")
		secret   = flag.String("secret", "", "GitHub application client secret")
	)
	flag.Parse()

	// Validate parameters.
	if *dataDir == "" {
		log.Fatal("data directory required: -data-dir")
	} else if *clientID == "" {
		log.Fatal("client id required: -client-id")
	} else if *secret == "" {
		log.Fatal("secret required: -secret")
	}

	// Create data directory, if necessary.
	if err := os.Mkdir(*dataDir, 0700); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	// Open the database.
	db, err := rationl.Open(filepath.Join(*dataDir, "db"), 0600)
	if err != nil {
		log.Fatal("open: " + err.Error())
	}

	// Enable logging.
	log.SetFlags(log.LstdFlags)

	// Setup the HTTP handlers.
	http.Handle("/", handlers.NewHandler(db, *clientID, *secret))

	// Start the HTTP server.
	go func() { http.ListenAndServe(*addr, nil) }()

	fmt.Printf("Listening on http://localhost%s\n", *addr)
	select {}
}
