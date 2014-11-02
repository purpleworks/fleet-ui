package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
)

var (
	renderer    *render.Render
	fleetClient FleetClient
	tempDir     string
)

func init() {
	renderer = render.New(render.Options{})
	fleetClient = NewClientCLIWithPeer("http://192.168.81.101:4001")
	tempDir = "./tmp"
}

func main() {
	r := mux.NewRouter().StrictSlash(false)

	api := r.PathPrefix("/api/v1").Subrouter()

	// routing machines collection
	machines := api.Path("/machines").Subrouter()
	machines.Methods("GET").HandlerFunc(machineAllHandler)

	// routing units collection
	units := api.Path("/units").Subrouter()
	units.Methods("GET").HandlerFunc(statusAllHandler)
	units.Methods("POST").HandlerFunc(submitUnitHandler)
	units.Path("upload").Methods("POST").HandlerFunc(uploadUnitHandler)

	// routing units singular
	unit := api.PathPrefix("/units/{id}").Subrouter()
	unit.Methods("GET").HandlerFunc(statusHandler)
	unit.Methods("DELETE").HandlerFunc(destroyHandler)
	unit.Path("/start").Methods("POST").HandlerFunc(startHandler)
	unit.Path("/stop").Methods("POST").HandlerFunc(stopHandler)
	unit.Path("/load").Methods("POST").HandlerFunc(loadHandler)

	// routing websocket
	r.Path("/ws/journal/{id}").HandlerFunc(wsHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":3000")
}
