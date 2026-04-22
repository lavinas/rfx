package main

import (
	"github.com/lavinas/cadoc6334/internal/adapter"
	"github.com/lavinas/cadoc6334/internal/usecase"
)

// main function to run the ReconcileIntercam function
func main() {
	/*
	repo, err := adapter.NewPostgresGormAdapter(adapter.PostgresConfig{
		Host:     "localhost",
		Port:     5434,
		User:     "root",
		Password: "root",
		DBName:   "reg",
		SSLMode:  "disable",
	})
	*/
	repo, err := adapter.NewPostgresGormAdapter(adapter.PostgresConfig{
		Host:     "192.168.100.78",
		Port:     5436,
		User:     "sys_flexcon",
		Password: "Wgkjsjjja8872Xl",
		DBName:   "dev_regulat",
		SSLMode:  "disable",
	})
	
	if err != nil {
		panic(err)
	}
	defer repo.Close()
	usecase.NewReconciliateCase(repo).ExecuteAll()
}
