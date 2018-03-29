package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
)

type Falco_Output_Fields struct {
	Container_id string  `json:"container.id"`
	Evt_time     big.Int `json:"evt.time"`
	K8s_pod_name string  `json:"k8s.pod.name"`
	Proc_cmdline string  `json:"proc.cmdline"`
	Proc_name    string  `json:"proc.name"`
	Proc_pname   string  `json:"proc.pname"`
	Proc_tty     int     `json:"proc.tty"`
	User_name    string  `json:"user_name"`
}

type Falco_Response struct {
	Output   string              `json:"output"`
	Priority string              `json:"priority"`
	Rule     string              `json:"rule"`
	Time     string              `json:"time"`
	Fields   Falco_Output_Fields `json:"output_fields"`
}

func marshalFalco(falcoJson []byte) {
	//falcoJson := `{"output":"18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user_name":"root"}}`
	var fr Falco_Response
	json.Unmarshal(falcoJson, &fr)
	fmt.Println(fr)
}

func handleFalco(w http.ResponseWriter, r *http.Request) {
	// Read body of request
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	// Print raw byte array of request
	//fmt.Println(string(b))

	// Unmarshal into expected falco struct
	var fr Falco_Response
	err = json.Unmarshal(b, &fr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Alternate Unmarshal Method
	//decoder := json.NewDecoder(req.Body)
	//err := decoder.Decode(&fr)

	// Respond to request with Success
	s := "Success! "
	output, err := json.Marshal(s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	// Process metrics for this event
	processFalcoEventMetrics(&fr)

	// Now determine if we should delete the pod
	filter := getFilter()
	namespace := getNamespace()
	if shouldDeletePod(fr.Fields.K8s_pod_name, filter) {
		deletePod(fr.Fields.K8s_pod_name, namespace)
	}
}

func getFilter() string {
	return "filter"
}

func getNamespace() string {
	// Lookup namespace from the current teflon deployment
	// or from the falco deployment via the k8s api
	return "default"
}

func shouldDeletePod(podname string, filter string) bool {
	if contains(podname, filter) {
		return true
	}

	if contains(podname, "delete") {
		return true
	}

	return false
}

func contains(superStr string, subStr string) bool {
	return true
}

func deletePod(podname string, namespace string) {
	fmt.Println("Deleting podname: ", podname, "namespace: ", namespace, "...")
	// Do k8s things here
}

func processFalcoEventMetrics(fr *Falco_Response) error {
	// Increment counter of falco events
	// Is there any benefit to add event to histogram?
	return nil
}

func processDeletePodEventMetrics(podname string) error {
	// Increment counter of delete pod events
	// Increment counter for delete of this podname
	// Increment counter for delete in this namespace
	return nil
}

func main() {
	http.HandleFunc("/", handleFalco)
	http.ListenAndServe(":80", nil)
	//falcoJson := `{"output":"18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user_name":"root"}}`
	//marshalFalco([]byte(falcoJson))
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
