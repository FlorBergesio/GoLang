package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Go web server")
	/* Route detection */
	http.HandleFunc("/", home)

	/* Port listener - http://localhost:3000/ */
	http.ListenAndServe(":3000", nil)
}

/* Routes */
func home(w http.ResponseWriter, r *http.Request) {
	/* Serve file */
	http.ServeFile(w, r, "./index.html")
}
