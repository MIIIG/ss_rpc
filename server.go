package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Id      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}
type Request struct {
	Id      string         `json:"id"`
	Jsonrpc string         `json:"jsonrpc"`
	Method  string         `json:"method"`
	Params  map[string]any `json:"params"`
}

type Error struct {
	Id      string    `json:"id"`
	Jsonrpc string    `json:"jsonrpc"`
	Error   RpcErrors `json:"error"`
}

type RpcErrors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var myErr = map[string]RpcErrors{
	"InternalError":  {"-32603", "Internal error"},
	"ParseError":     {"-32700", "Parse error"},
	"MethodNotFound": {"-32601", "Method not found"},
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req Request

	reqbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler("InternalError", w)
	}
	jbody := json.Unmarshal(reqbody, &req)
	if jbody != nil {
		ErrorHandler("ParseError", w)
	}
	if req.Method != "" {
		switch req.Method {
		case "greeting":
			Greeting(req, w)
		default:
			NotFound(req, w)
		}
	}
}

func ErrorHandler(errtype string, w http.ResponseWriter) {
	var output = Error{"null", "2.0", myErr[errtype]}
	js, _ := json.Marshal(output)
	WriteJSON(js, w)
}

func Greeting(request Request, w http.ResponseWriter) {
	var output Result
	result := fmt.Sprintf("Hello, %s", request.Params["name"])
	output = Result{request.Id, request.Jsonrpc, result}
	js, err := json.Marshal(output)
	if err != nil {
		ErrorHandler("InternalError", w)
	}
	WriteJSON(js, w)
}

func NotFound(request Request, w http.ResponseWriter) {
	var output = Error{request.Id, request.Jsonrpc, myErr["MethodNotFound"]}
	js, err := json.Marshal(output)
	if err != nil {
		ErrorHandler("InternalError", w)
	}
	WriteJSON(js, w)
}

func WriteJSON(data []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Handler)
	srv := &http.Server{
		Handler: mux,
		Addr:    "127.0.0.1:8000",
	}

	log.Println("Air on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
