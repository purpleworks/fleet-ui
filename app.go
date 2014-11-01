package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"gopkg.in/unrolled/render.v1"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	renderer    = render.New(render.Options{})
	fleetClient = NewClientCLIWithPeer("http://192.168.81.101:4001")
	tempDir     = "./tmp"
)

func main() {
	r := mux.NewRouter().StrictSlash(false)

	api := r.PathPrefix("/api/v1").Subrouter()

	// machines collection
	machines := api.Path("/machines").Subrouter()
	machines.Methods("GET").HandlerFunc(machineAllHandler)

	// Units collection
	units := api.Path("/units").Subrouter()
	units.Methods("GET").HandlerFunc(statusAllHandler)
	units.Methods("POST").HandlerFunc(submitUnitHandler)
	units.Path("upload").Methods("POST").HandlerFunc(uploadUnitHandler)

	// Units singular
	unit := api.PathPrefix("/units/{id}").Subrouter()
	unit.Methods("GET").HandlerFunc(statusHandler)
	unit.Methods("DELETE").HandlerFunc(destroyHandler)
	unit.Path("/start").Methods("POST").HandlerFunc(startHandler)
	unit.Path("/stop").Methods("POST").HandlerFunc(stopHandler)
	unit.Path("/load").Methods("POST").HandlerFunc(loadHandler)

	// websocket
	r.Path("/ws/journal/{id}").HandlerFunc(wsHandler)

	n := negroni.Classic()
	n.UseHandler(r)

	n.Run(":3000")
}

func destroyHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	log.Printf("destroy %s unit", key)
	if err := fleetClient.Destroy(key); err != nil {
		// log.Printf("unit destroy error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func startHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	log.Printf("start %s unit", key)
	if err := fleetClient.Start(key); err != nil {
		// log.Printf("unit start error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func stopHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	log.Printf("stop %s unit", key)
	if err := fleetClient.Stop(key); err != nil {
		// log.Printf("unit stop error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func loadHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	log.Printf("load %s unit", key)
	if err := fleetClient.Load(key); err != nil {
		// log.Printf("unit load error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func uploadUnitHandler(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("file")
	defer file.Close()

	serviceFile := fmt.Sprintf("%s/%s", tempDir, header.Filename)
	out, err := os.Create(serviceFile)
	if err != nil {
		log.Printf("Open file errpr: %s", err)
		renderer.JSON(w, http.StatusBadRequest, err)
		return
	}
	defer out.Close()

	// write the content from POST to the file
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	err = fleetClient.Submit(header.Filename, serviceFile)
	if err != nil {
		// log.Printf("Fleet submit error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})

}

func submitUnitHandler(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	service := req.FormValue("service")

	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		os.Mkdir(tempDir, 0755)
	}

	serviceFile := fmt.Sprintf("%s/%s", tempDir, name)
	lines := strings.Split(string(service), "\\n")

	fo, err := os.Create(serviceFile)
	if err != nil {
		log.Printf("Open file errpr: %s", err)
		renderer.JSON(w, http.StatusBadRequest, err)
		return
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for _, str := range lines {
		fmt.Fprintln(fo, str)
	}

	err = fleetClient.Submit(name, serviceFile)
	if err != nil {
		// log.Printf("Fleet submit error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func machineAllHandler(w http.ResponseWriter, req *http.Request) {
	status, _ := fleetClient.MachineAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusAllHandler(w http.ResponseWriter, req *http.Request) {
	status, _ := fleetClient.StatusAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
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
	output, _ := fleetClient.JournalF(key)
	for line := range output {
		conn.WriteMessage(websocket.TextMessage, []byte(line))
	}
}
