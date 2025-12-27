package controller

import (
	"net/http"
	"text/template"
)

func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}
