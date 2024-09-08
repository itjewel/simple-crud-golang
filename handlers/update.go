package handlers

import (
    "net/http"
)

func UpdatePage(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "Title":  "Update - Simple CRUD",
        "Header": "Update an Item",
    }
    Tmpl.ExecuteTemplate(w, "update.html", data)  // Use global 'Tmpl'
}
