package main

import(
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("www"))
	fmt.Println("GhostCat is running ... BOO!")

	http.Handle("/", fs)

	http.ListenAndServe(":8081", nil)
}