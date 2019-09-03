package controllers

import (
	"net/http"

	"lenslocked.com/ch_07/07_3/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

// handler or "action"
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

type Users struct {
	NewView *views.View
}
