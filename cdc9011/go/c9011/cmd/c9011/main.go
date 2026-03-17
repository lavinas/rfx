package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"c9011/internal/builder"
	"c9011/internal/config"
	"c9011/internal/excel"
	"c9011/internal/validator"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("uso: c9011 <arquivo.xlsx>")
		os.Exit(1)
	}

	input := os.Args[1]

	meta, rows, periods, err := excel.ParseFile(input)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Load("c9011.yml")
	if err != nil {
		log.Fatal(err)
	}

	skipMap := cfg.Validator.SkipMap()

	err = validator.ValidateHierarchy(rows, periods, skipMap)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := builder.BuildDocument(meta, rows, periods)
	if err != nil {
		log.Fatal(err)
	}

	ext := filepath.Ext(input)
	output := input[:len(input)-len(ext)] + ".json"

	err = builder.WriteJSON(output, doc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Arquivo gerado com sucesso em:", output)
}
