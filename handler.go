package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

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
	log.Printf("start unit %s", key)
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
	log.Printf("load unit - %s", key)
	if err := fleetClient.Load(key); err != nil {
		// log.Printf("unit load error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func uploadUnitHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("upload unit")
	if _, err := os.Stat(tempDir); err != nil {
		log.Printf("err: %s", err)
	} else if os.IsNotExist(err) {
		os.Mkdir(tempDir, 0755)
	}

	file, header, err := req.FormFile("file")
	if err != nil {
		log.Printf("read file error : %s", err)
		renderer.JSON(w, http.StatusBadRequest, err)
		return
	}
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
	if err := fleetClient.Start(header.Filename); err != nil {
		// log.Printf("unit start error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})

}

func submitUnitHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("submit unit")
	type UnitForm struct {
		Name    string `json:"name"`
		Service string `json:"service"`
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("read request body error : %s", err)
		renderer.JSON(w, http.StatusBadRequest, err)
		return
	}

	var t UnitForm
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.Decode(&t)
	if err != nil {
		log.Printf("json decode error : %s", err)
	}

	if _, err := os.Stat(tempDir); err != nil {
		log.Printf("err: %s", err)
	} else if os.IsNotExist(err) {
		os.Mkdir(tempDir, 0755)
	}

	serviceFile := fmt.Sprintf("%s/%s", tempDir, t.Name)
	lines := strings.Split(string(t.Service), "\\n")

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

	err = fleetClient.Submit(t.Name, serviceFile)
	if err != nil {
		// log.Printf("Fleet submit error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}

	if err := fleetClient.Start(t.Name); err != nil {
		// log.Printf("unit start error: %s", err)
		// renderer.JSON(w, http.StatusBadRequest, err)
		// return
	}
	renderer.JSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func machineAllHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("machine all")
	status, _ := fleetClient.MachineAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusAllHandler(w http.ResponseWriter, req *http.Request) {
	log.Println("unit all")
	status, _ := fleetClient.StatusAll()
	renderer.JSON(w, http.StatusOK, status)
}

func statusHandler(w http.ResponseWriter, req *http.Request) {
	key := mux.Vars(req)["id"]
	log.Printf("unit detail - %s", key)
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
	stdout, errout, err := fleetClient.JournalF(key)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	go func() {
		for line := range stdout {
			conn.WriteMessage(websocket.TextMessage, []byte(line))
		}
	}()
	go func() {
		for line := range errout {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Stderr: %s", line)))
		}
	}()
}
