package main

import (
	"fmt"
	"internal/service"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/defaultdb?sslmode=disable"
	port        = 3000
)

func main() {
	fmt.Println("Hello World")

	db, err := sql.Open("pgx", databaseURL)

	if err != nil {
		log.Fatalf("Could not open db connection: %v\n", err)
		return
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not ping to db: %v\n", err)
		return
	}

	codec := branca.NewBranca("asupersecretkeyyoushouldnotcommit")

	s := service.Service{
		db:    db,
		codec: codec,
	}

}
