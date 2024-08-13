package handlers

import (
	"net/http"

	"github.com/jeftavares/htmx-starter/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	//return render(w, r, views.Home())
	return render(w, r, views.Home())
}
