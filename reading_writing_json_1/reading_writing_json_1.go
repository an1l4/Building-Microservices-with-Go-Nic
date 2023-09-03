package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
    Message string
}

func main()  {
    port:= 8080

    http.HandleFunc("/helloworld",helloWorldHandler)

    log.Printf("server starting on %v\n",port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port),nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
response:=helloWorldResponse{
    Message: "Hi Hello World!",
}
data,err:=json.Marshal(&response)

if err!=nil{
    panic("oops : error while marshalling")
}
fmt.Fprint(w,string(data))
}