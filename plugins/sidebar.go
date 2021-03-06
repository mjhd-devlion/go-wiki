package plugins

import (
	"html/template"
	"os"

	"../models"
	"../utils"
)

func Sidebar() template.HTML {

	var page models.Page

	err := page.LoadFromTitle("_sidebar")
	if err != nil {
		utils.PromulgateFatal(os.Stdout, err)
		panic(err.Error())
	}

	return template.HTML(string(page.Markdown()))
}
