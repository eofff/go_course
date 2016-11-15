package controllers

import (
	"html/template"
	"net/http"

	"week4/models"
)

type IndexController struct {
}

func (controller *IndexController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/table.html")
	if err == nil {
		t.ExecuteTemplate(w, "table", models.GetUsers())
	} else {
		w.Write([]byte(err.Error()))
	}
}
