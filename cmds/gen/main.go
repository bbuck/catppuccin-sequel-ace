package main

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"text/template"

	catppuccin "github.com/catppuccin/go"
)

var filePath = flag.String("template", "", "Provide the file path of the spTheme template")
var outputPath = flag.String("output", "./themes/", "The destination for the generated theme files.")

func main() {
	flag.Parse()

	spThemeTemplate, err := template.ParseFiles(*filePath)
	if err != nil {
		panic(err)
	}
	themes := map[string]catppuccin.Theme{
		"latte":     catppuccin.Latte,
		"frappe":    catppuccin.Frappe,
		"macchiato": catppuccin.Macchiato,
		"mocha":     catppuccin.Mocha,
	}

	err = os.Mkdir(*outputPath, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}

	templateName := filepath.Base(*filePath)
	themeTemplate := spThemeTemplate.Lookup(templateName)
	for themeName, theme := range themes {
		themePath := filepath.Join(*outputPath, "catppuccin-"+themeName+".spTheme")
		err := writeFile(themeTemplate, themePath, theme)
		if err != nil {
			panic(err)
		}
	}
}

func writeFile(tmpl *template.Template, themePath string, theme catppuccin.Theme) error {
	themeFile, err := os.OpenFile(themePath, os.O_CREATE|os.O_WRONLY, 0777)
	if errors.Is(err, os.ErrExist) {
		err = os.Remove(themePath)
		if err != nil {
			return err
		}
		themeFile, err = os.OpenFile(themePath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	defer themeFile.Close()
	err = tmpl.Execute(themeFile, theme)

	return err
}
