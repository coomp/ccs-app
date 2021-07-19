package main

import (
	"encoding/json"
	"net/http"

	"github.com/coomp/ccs-app/handler"
	ccssdk "github.com/coomp/ccs-sdk"
)

type MovieTicketBookRequest struct {
	Username        string // 用户名
	MoiveId         string // 电影ID
	MovieSequenceId string // 电影场次
	Timestamp       string // 请求时间戳
}

func echo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.Write([]byte("HTTP POST required"))
	}

	d := json.NewDecoder(req.Body)
	ticketReq := &MovieTicketBookRequest{}
	err := d.Decode(ticketReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// TODO send msg to ccs gateway using sdk through grpc
}

func main() {
	// Init CCS SDK
	handles := ccssdk.NewEmpty()
	handles = append(handles, handler.NewExampleHandler())
	// TODO add your handles here
	ccssdk.InitCcsSdk(handles)

	// Business code here, if pure CCS service, ignore and remove code blow
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8090", nil)
}
