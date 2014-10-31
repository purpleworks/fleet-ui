package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gopkg.in/unrolled/render.v1"
	"log"
	"net/http"
)

var renderer = render.New(render.Options{})

func main() {
	r := mux.NewRouter().StrictSlash(false)

	api := r.PathPrefix("/api/v1").Subrouter()

	// machines collection
	machines := api.Path("/machines").Subrouter()
	machines.Methods("GET").HandlerFunc(machineAllHandler)

	// Units collection
	units := api.Path("/units").Subrouter()
	units.Methods("GET").HandlerFunc(statusAllHandler)

	// Units singular
	unit := api.PathPrefix("/units/{id}").Subrouter()
	unit.Methods("GET").HandlerFunc(statusHandler)

	// websocket
	r.Path("/ws/journal/{id}").HandlerFunc(wsHandler)

	n := negroni.New()
	n.UseHandler(r)

	n.Run(":3000")
}

func machineAllHandler(w http.ResponseWriter, req *http.Request) {
	fleetClient := NewClientCLIWithPeer("http://192.168.81.101:4001")
	status, _ := fleetClient.MachineAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusAllHandler(w http.ResponseWriter, req *http.Request) {
	fleetClient := NewClientCLIWithPeer("http://192.168.81.101:4001")
	status, _ := fleetClient.StatusAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	fleetClient := NewClientCLIWithPeer("http://192.168.81.101:4001")
	status, _ := fleetClient.StatusUnit(key)
	renderer.JSON(w, http.StatusOK, status)
}

// websocket handler

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// write journal message
	key := mux.Vars(r)["id"]
	fleetClient := NewClientCLIWithPeer("http://192.168.81.101:4001")
	output, _ := fleetClient.JournalF(key)
	for line := range output {
		conn.WriteMessage(websocket.TextMessage, []byte(line))
	}
}
