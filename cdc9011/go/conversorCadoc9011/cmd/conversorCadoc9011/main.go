package main

import (
	"fmt"
	"os"
	"path/filepath"

	"conversorCadoc9011/internal/builder"
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
		panic(err)
	}

	err = validator.ValidateHierarchy(rows, periods)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := builder.BuildDocument(meta, rows, periods)
	if err != nil {
		panic(err)
	}

	ext := filepath.Ext(input)
	output := input[:len(input)-len(ext)] + ".json"

	err = builder.WriteJSON(output, doc)
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo gerado com sucesso em:", output)
}
