package pdfGenerator

import (
	"fmt"
	"os"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/google/uuid"
)

type wkHTMLtoPDF struct {
	rootPath string
}

func NewWkHTLtoPDF(rootPath string) PDFGeneratorInterface {
	return &wkHTMLtoPDF{rootPath: rootPath}
}

func (w *wkHTMLtoPDF) Create(htmlFile string) (string, error) {
	f, err := os.Open(htmlFile)
	if err != nil {
		return "", err
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return "", err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

	if err := pdfg.Create(); err != nil {
		return "", err
	}

	fileName := w.rootPath + "/" + uuid.New().String() + ".pdf"

	fmt.Println(fileName)

	if err := pdfg.WriteFile(fileName); err != nil {
		return "", err
	}

	return fileName, nil
}
