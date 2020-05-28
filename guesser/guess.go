package guesser

import "net/http"

func Guess(w http.ResponseWriter, r *http.Request)  {
	render(w, "index.html", r)
}