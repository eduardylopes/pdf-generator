package htmlParser

import (
	"html/template"
	"os"

	"github.com/google/uuid"
)

type htmlStruct struct {
	rootPath string
}

func New(rootPath string) HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (hs *htmlStruct) Create(templateName string, data interface{}) (string, error) {
	template, err := template.ParseFiles(templateName)

	if err != nil {
		return "", err
	}

	fileName := hs.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)

	if err != nil {
		return "", err
	}

	if err := template.Execute(fileWriter, data); err != nil {
		return "", err
	}

	return fileName, nil
}
