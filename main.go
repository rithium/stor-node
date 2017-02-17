package main

import (
	"./model"
	"./config"
	"flag"
	"log"
	"os"
	"github.com/rithium/version"
	"github.com/gorilla/mux"
	"net/http"
)

type Env struct {
	db model.Datastore
}

func init() {
	versionFlag := flag.Bool("v", false, "prints version")
	configFlag := flag.Bool("c", false, "dumps configuration")

	flag.Parse()

	if *versionFlag {
		log.Println("Stor Data", version.GetVersion())
		os.Exit(0)
	}

	config.LoadConfig()

	if *configFlag {
		log.Printf("HTTP:\t%+v\n", config.HttpServer)
		log.Printf("MySQL:\t%+v\n", config.MySQL)

		os.Exit(0)
	}
}

func main() {
	db, err := model.NewDb(config.MySQL)

	if err != nil {
		log.Panic(err)
	}

	if db == nil {
		log.Panic("mysql", err)
	}

	env := &Env{db}

	log.Println("Stor Node", version.GetVersion())

	router := mux.NewRouter()

	router.HandleFunc("/node", env.handleGetNodeByKey).Queries("key", "").Methods("GET")
	router.HandleFunc("/node/{nodeId:[0-9]+}", env.handleNodeGet).Methods("GET")
	router.HandleFunc("/node", env.handleNodePost).Methods("POST")

}

func (env *Env) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (env *Env)handleNodeGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (env *Env)handleGetNodeByKey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (env *Env)handleNodePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}