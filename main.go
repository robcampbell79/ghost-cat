package main

import(
	"fmt"
	"net/http"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Host 	string `toml:"host"`
	Port 	string `toml:"port"`
	Index 	string `toml:"index"`
}

func main() {
	var conf Config

	_, err := toml.DecodeFile("./conf.toml", &conf)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("index: ", conf.Index)
	fmt.Println("port: ", conf.Port)

	fs := http.FileServer(http.Dir(conf.Index))

	fmt.Println("GhostCat is running ... BOO!")

	http.Handle("/", fs)

	http.ListenAndServe(conf.Port, nil)
}