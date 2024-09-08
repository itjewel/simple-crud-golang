package handlers

import (
	"fmt"
	"net/http"
)

func CreatePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		item := r.FormValue("item")
		fmt.Fprintf(w, "Item '%s' created successfully!", item)
	} else {
		data := map[string]string{
			"Title":  "Create - Simple CRUD",
			"Header": "Create an Item",
		}
		Tmpl.ExecuteTemplate(w, "create.html", data)
	}
}
