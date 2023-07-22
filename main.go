package main

import (
	"fmt"
	"os"

	"github.com/eduardylopes/pdf-generator/htmlParser"
	"github.com/eduardylopes/pdf-generator/pdfGenerator"
)

type Data struct {
	Name string
}

func main() {
	h := htmlParser.New("tmp")
	wk := pdfGenerator.NewWkHTLtoPDF("tmp")

	dataHTML := Data{
		Name: "Eduardy",
	}

	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	defer os.Remove(htmlGenerated)

	if err != nil {
		fmt.Println(err)
		return
	}

	pdfGenerated, err := wk.Create(htmlGenerated)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Generated template: ", htmlGenerated)
	fmt.Println("Generated pdf: ", pdfGenerated)
}
