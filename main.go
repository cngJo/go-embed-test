package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed assets/*
var assetData embed.FS

func main() {
	t, err := template.ParseFS(assetData, "assets/report.html")
	if err != nil {
		fmt.Println(err)
	}

	templateData := struct {
		Title string
	}{
		Title: "File inside go",
	}

	t.Execute(os.Stdout, templateData)
}
