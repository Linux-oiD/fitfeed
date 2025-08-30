package v1

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (h *V1) getAuthCallbackFunction(w http.ResponseWriter, r *http.Request) {

	provider := chi.URLParam(r, "provider")
	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	fmt.Println(user)

	http.Redirect(w, r, "http://localhost:5173/", http.StatusFound)
}
