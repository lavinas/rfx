package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"conversorCadoc9011/internal/builder"
	"conversorCadoc9011/internal/config"
	"conversorCadoc9011/internal/excel"
	"conversorCadoc9011/internal/validator"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("uso: conversor90x1 <arquivo.xlsx>")
		os.Exit(1)
	}

	input := os.Args[1]

	meta, rows, periods, err := excel.ParseFile(input)
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Load("app.yml")
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
