# CCS example app

> ccs based app example

## Cenerate project

+ `go get github.com/coomp/ccs-gen`
+ `ccs-gen myproject example.com/org/repo`

## Init project

+ `cd myproject && go mod tidy`

## Development

+ follow handler/example.go to create your own req/resp handlers
+ follow main.go/echo/HandleMessage to send your own messages
