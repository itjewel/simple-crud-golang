package handlers

import (
    "net/http"
)

func ReadPage(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
        "Title":  "Read - Simple CRUD",
        "Header": "Read Items",
    }
    Tmpl.ExecuteTemplate(w, "read.html", data)  // Use global 'Tmpl'
}
