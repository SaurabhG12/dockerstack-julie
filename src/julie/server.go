package main

import (
	"filefetch"
	"fmt"
	"github.com/bmizerany/pat"
	"github.com/pelletier/go-toml"
	"log"
	"logger"
	"net/http"
)

func Info(w http.ResponseWriter, req *http.Request) {

	log.Print("Welcome to Julie Orchestration Engine 1.0!")

	w.Write([]byte("############################################\n" +
		"############################################\n" +
		"Welcome to Julie Orchestration Engine 1.0!\n"))

}

func init() {

	config, err := toml.LoadFile("../../config/julie.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Print("Working")
	}

	devport := config.Get("development.port").(int64)

	security := config.Get("security.accept").(string)

	fmt.Print(security)

	logger.AppStart(devport)

	//global_devport := config.Get("development.port").(string)

	//fmt.Println(dev)

}

func main() {

	m := pat.New()
	m.Get("/api/info", http.HandlerFunc(Info))
	m.Get("/api/filefetch/readdir/:folder", http.HandlerFunc(filefetch.ReadDir))
	http.Handle("/", m)

	err := http.ListenAndServe(":1987", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
