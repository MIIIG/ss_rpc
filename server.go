package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Result struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}
type Request struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type Params struct {
	Name string `json:"name"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req Request

	switch r.Method {
	case "GET":
		w.WriteHeader(200)
		io.WriteString(w, "This is RPC server!\n")
	case "POST":
		reqbody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("server: could not read request body: %s\n", err)
		}
		jbody := json.Unmarshal(reqbody, &req)
		if jbody != nil {
			fmt.Printf("Can't unmarshall data: %s\n", jbody)
		}
		MethodHandler(req.Method, req, w)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func MethodHandler(method string, request Request, w http.ResponseWriter) {
	switch method {
	case "greeting":
		Greeting(request, w)
	default:
		io.WriteString(w, "Nothing to do!\n")
	}
}

func Greeting(request Request, w http.ResponseWriter) {
	var output Result
	result := fmt.Sprintf("Hello, %s", request.Params.Name)
	output = Result{request.Id, request.Jsonrpc, result}
	js, err := json.Marshal(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
