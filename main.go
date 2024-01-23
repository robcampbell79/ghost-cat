package main

import(
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
)

type Server struct {
	host 		string `toml:"server"`
	port 		string `toml:"port"`
	indexDir 	string `toml:"indexDir"`
}

func main() {
	var srv Server

	_, err := toml.DecodeFile("conf.toml", &srv)
	if err != nil {
		panic(err.Error())
	}

	fs := http.FileServer(http.Dir("www"))

	fmt.Println("GhostCat is running ... BOO!")

	http.Handle("/", fs)

	http.ListenAndServe(":8081", nil)
}