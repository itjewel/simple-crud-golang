package handlers

import (
    "net/http"
)

func CreatePage(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "Title":  "Create - Simple CRUD",
        "Header": "Create an Item",
    }
    Tmpl.ExecuteTemplate(w, "create.html", data)  // Use global 'Tmpl'
}
