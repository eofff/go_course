package controllers

import (
	"html/template"
	"net/http"

	"week4/models"
)

type IndexController struct {
}

func (controller *IndexController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	var users [3]models.User
	users[0] = models.User{"Some name", "test@test.com", "fdscxv"}
	users[1] = models.User{"Another name", "hz@hz.hz", "aasd"}
	users[2] = models.User{"Afdsi", "fff@asczx.net", "7asfgvsd"}
	t, err := template.ParseFiles("templates/table.html")
	if err == nil {
		t.ExecuteTemplate(w, "table", users)
	} else {
		w.Write([]byte(err.Error()))
	}
}
