package handlers

import (
    "html/template"
    "log"
)

var Tmpl *template.Template  // Make 'Tmpl' global and accessible

func InitTemplates() {
    var err error
    Tmpl, err = template.ParseGlob("templates/*.html")
    if err != nil {
        log.Fatal("Error parsing templates:", err)
    }
}
