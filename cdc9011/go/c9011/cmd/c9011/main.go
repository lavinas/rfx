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
	// Get the input file from command line arguments
	if len(os.Args) < 2 {
		fmt.Println("uso: c9011 <arquivo.xlsx>")
		os.Exit(1)
	}
	input := os.Args[1]

	// Load the configuration
	cfg, err := config.Load("c9011.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Parse the Excel file
	meta, rows, periods, err := excel.ParseFile(input)
	if err != nil {
		log.Fatal(err)
	}

	// Validate the hierarchy of the data
	skipMap := cfg.Validator.SkipMap()
	err = validator.ValidateHierarchy(rows, periods, skipMap)
	if err != nil {
		log.Fatal(err)
	}

	// Build the JSON document
	doc, err := builder.BuildDocument(meta, rows, periods)
	if err != nil {
		log.Fatal(err)
	}

	// Validate the structure of the JSON document
	err = validator.ValidateStructure(doc)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON document to a file
	ext := filepath.Ext(input)
	output := input[:len(input)-len(ext)] + ".json"
	err = builder.WriteJSON(output, doc)
	if err != nil {
		log.Fatal(err)
	}

	// Print success message
	log.Println("Arquivo gerado com sucesso em:", output)
}
