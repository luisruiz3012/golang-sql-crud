package routes

import "net/http"

func IndexRotue(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world! :)"))
}
