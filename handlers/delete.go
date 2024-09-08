package handlers

import (
    "net/http"
)

func DeletePage(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "Title":  "Delete - Simple CRUD",
        "Header": "Delete an Item",
    }
    Tmpl.ExecuteTemplate(w, "delete.html", data)  // Use global 'Tmpl'
}
