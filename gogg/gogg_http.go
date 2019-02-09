package gogg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type httpGGTable struct {
	nodes []*Node

	host         string
	port         int
	readTimeout  int
	writeTimeout int
}

func NewHTTPGGTable(host string, port int, readTimeout int, writeTimeout int) GoroutineGroupTable {
	return httpGGTable{
		nodes: make([]*Node, 0),

		host:         host,
		port:         port,
		readTimeout:  readTimeout,
		writeTimeout: writeTimeout,
	}
}

func (hggt httpGGTable) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/linkMe", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.linkMe(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	router.HandleFunc("/createGroup/{groupName}", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.createGroup(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	router.HandleFunc("/deleteGroup/{groupName}", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.deleteGroup(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	router.HandleFunc("/joinGroup/{groupName}/{goroutineName}", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.joinGroup(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	router.HandleFunc("/leaveGroup/{groupName}/{goroutineName}", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.leaveGroup(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	router.HandleFunc("/transferToGroup/{groupName}", func(w http.ResponseWriter, r *http.Request) {
		err := hggt.transferToGroup(r)
		if err != nil {
			fmt.Fprintf(w, "Unsuccessful")
		} else {
			fmt.Fprintf(w, "Successful")
		}
	})

	address := hggt.host + ":" + strconv.Itoa(hggt.port)

	srv := &http.Server{
		Handler: router,
		Addr:    address,

		WriteTimeout: time.Duration(hggt.writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(hggt.readTimeout) * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func (hggt httpGGTable) linkMe(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) createGroup(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) deleteGroup(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) joinGroup(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) leaveGroup(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) transferToGroup(request *http.Request) error {
	return nil
}

func (hggt httpGGTable) GetNodes() []*Node {
	return hggt.nodes
}

func (hggt httpGGTable) LinkNode(node Node) error {
	err := hggt.sendLinkNodeRequest(node)
	if err != nil {
		return err
	}

	return nil
}

func (hggt httpGGTable) sendLinkNodeRequest(node Node) error {
	address := node.Host + ":" + strconv.Itoa(node.Port)
	resp, err := http.Post("http://"+address, "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) != "Successful" {
		return errors.New("Response of LinkNode is not Successful")
	}
	return nil
}

func (hggt httpGGTable) Create(groupName string) {

}

func (hggt httpGGTable) Delete(groupName string) {

}

func (hggt httpGGTable) Join(groupName string, goroutineName string) error {
	return nil
}

func (hggt httpGGTable) Leave(groupName string, goroutineName string) error {
	return nil
}

func (hggt httpGGTable) GetMembers(groupName string) ([]*DistributedGoroutine, error) {
	return nil, nil
}

func (hggt httpGGTable) GetLocalMembers(groupName string) ([]*DistributedGoroutine, error) {
	return nil, nil
}

func (hggt httpGGTable) WhichGroups() []string {
	return nil
}

func (hggt httpGGTable) TransferToGroup(message interface{}, groupName string) error {
	return nil
}
