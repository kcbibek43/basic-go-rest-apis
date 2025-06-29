package main

import (
	"go-apis/internal/env"
	"log"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	// os.LookupEnv();

	mux := app.mount()
	log.Fatal(app.runApi(mux))
}
