package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Printf("This is main.\n")
	startServer()
}

type Information struct {
	XFoo string
	NODE string
	POD  string
}

//
// shamelessly ripped off from github.com/kube-tools/kportal/pkg/{server,handler}
//

func startServer() {
	fmt.Printf(">> startServer()\n")
	s := http.Server{
		Addr:         ":8085",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      createServeMux(),
	}
	err := s.ListenAndServe()
	fmt.Printf("startServer() is back from ListenAndServe()\n")
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}

}

func createServeMux() *http.ServeMux {
	fmt.Printf(">> createServeMux()\n")
	serveMux := http.NewServeMux()
	configureMux(serveMux)
	return serveMux
}

func configureMux(serveMux *http.ServeMux) {

	// processa pedido de autenticação
	serveMux.Handle("/", http.HandlerFunc(mainHandler))

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(">> mainHandler")
	information, err := getInformation()
	if err != nil {
		log.Printf("ERROR getting information: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		out, err := json.Marshal(information)
		if err != nil {
			log.Printf("ERROR marshalling information: %s", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write(out)
		}
	}
}

func getInformation() (*Information, error) {
	info := new(Information)

	nodeName, ok := os.LookupEnv("MY_NODE_NAME")

	if ok {
		info.NODE = nodeName
	} else {
		info.NODE = "please define MY_NODE_NAME using downward api"
	}

	return info, nil
}
