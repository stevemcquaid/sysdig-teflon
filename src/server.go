package main

import (
	"encoding/json"
	"net/http"
)

type Person struct {
	ID        int    `json:"id"`
	Firstname string `json:"firstname"`
}

var people []Person

func main2() {
	people = append(people, Person{ID: 1, Firstname: "Colyn"})
	people = append(people, Person{ID: 2, Firstname: "April"})
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(people)
	})

	http.ListenAndServe(":80", nil)
}

// {"output":"18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user.name":"root"}}
