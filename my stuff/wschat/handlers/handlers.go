package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

func renderPage(w http.ResponseWriter, template string, data jet.VarMap) error {
	view, err := views.GetTemplate(template)
	if err != nil {
		log.Println(err)
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
