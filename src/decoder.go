package main

import (
	"encoding/json"
	"fmt"
	//"net/http"
	"math/big"
)

type Falco_Output_Fields struct {
	Container_id string `json:"container.id"`
	Evt_time big.Int `json:"evt.time"`
	K8s_pod_name string `json:"k8s.pod.name"`
	Proc_cmdline string `json:"proc.cmdline"`
	Proc_name string `json:"proc.name"`
	Proc_pname string `json:"proc.pname"`
	Proc_tty int `json:"proc.tty"`
	User_name string `json:"user_name"`
}

type Falco_Response struct {
    Output   string   `json:"output"`
    Priority   string   `json:"priority"`
    Rule   string   `json:"rule"`
    Time   string `json:"time"`
    Fields Falco_Output_Fields  `json:"output_fields"`
}


func foo(){
	falcoJson := `{"output":"18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user_name":"root"}}`
	var fr Falco_Response
	json.Unmarshal([]byte(falcoJson), &fr)
	fmt.Println(fr)
}

func main() {
	//http.HandleFunc("/", parseGhPost)
	//http.ListenAndServe(":80", nil)
	foo()
}

//func parseGhPost(rw http.ResponseWriter, request *http.Request) {
	//	//decoder := json.NewDecoder(request.Body)
	//
	//	falcoJson := `{"output":"18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user.name":"root"}}`
	//	var fr Falco_Response
	//	json.Unmarshal([]byte(falcoJson), &fr)
	//	fmt.Println(fr.Fields.K8s_pod_name)
	//
	//	//
	//	//var t test_struct
	//	//err := decoder.Decode(&t)
	//	//
	//	//if err != nil {
	//	//	panic(err)
	//	//}
	//	//
	//	//fmt.Println(t.Test)
	//}
