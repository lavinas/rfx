package main

import (
	"github.com/lavinas/cadoc6334/internal/adapter"
	"github.com/lavinas/cadoc6334/internal/usecase"
)

// main function to run the ReconcileIntercam function
func main() {
	repo, err := adapter.NewPostgresGormAdapter(adapter.PostgresConfig{
		Host:     "localhost",
		Port:     5434,
		User:     "root",
		Password: "root",
		DBName:   "reg",
		SSLMode:  "disable",
	})
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	usecase.NewReconciliateCase(repo).ExecuteAll()
}
