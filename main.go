package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/coomp/ccs-app/handler"
	ccssdk "github.com/coomp/ccs-sdk"
)

type App struct {
	Sdk *ccssdk.CcsSdk
}

type MovieTicketBookRequest struct {
	Username        string // 用户名
	MoiveId         string // 电影ID
	MovieSequenceId string // 电影场次
	Timestamp       string // 请求时间戳
}

func (a *App) echo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Write([]byte("HTTP POST required"))
	}

	d := json.NewDecoder(req.Body)
	ticketReq := &MovieTicketBookRequest{}
	err := d.Decode(ticketReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// check if tickt param is valid

	// send msg to ccs gateway using sdk through grpc
	bytes, _ := json.Marshal(ticketReq)
	a.Sdk.HandleMessage(string(bytes), true, 3)
}

var app App

func main() {

	requestHandlers := append(ccssdk.NewEmptyRequestHandlers(), handler.NewExampleRequestHandler())
	responseHandlers := append(ccssdk.NewEmptyResponseHandlers(), handler.NewExampleResponseHandler())
	// TODO add your handles here

	// Init CCS SDK
	sdk, err := ccssdk.NewCcsSdk(requestHandlers, responseHandlers)
	if err != nil {
		log.Fatal(err)
	}

	app.Sdk = sdk

	// Business code here, if pure CCS service, ignore and remove code blow
	http.HandleFunc("/echo", app.echo)
	http.ListenAndServe(":8090", nil)
}
