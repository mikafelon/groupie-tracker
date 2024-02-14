package main

import (
	"fmt"
	"net/http"
)

const port = ":8888"

func main() {
	http.HandleFunc("/", Temp)
	http.HandleFunc("/artiste", ArtistHandler)
	fmt.Println("(http://localhost:8888) - Server started on port", port)
	http.ListenAndServe(port, nil)
}

//moka
