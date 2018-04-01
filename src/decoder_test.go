package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
)

func TestGetFilter(t *testing.T) {
	got := getFilter()
	want := "filter"
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}

type superSub struct {
	superStr string
	subStr   string
}

func TestShouldDeletePod(t *testing.T) {
	superStr := "racecar"
	subStr := "car"
	got := shouldDeletePod(superStr, subStr)
	want := true

	if want != got {
		t.Errorf("got %v want %v, given superStr: %v, subStr %v", got, want, superStr, subStr)
	}

	// for superStr, subStr in []
	positiveTestArr := [3]superSub{
		superSub{superStr: "funtsunami", subStr: "fun"},
		superSub{superStr: "1234", subStr: "23"},
		superSub{superStr: "booyeah123", subStr: "oyeah1"},
	}

	for _, superSub := range positiveTestArr {
		got := shouldDeletePod(superSub.superStr, superSub.subStr)
		want := true

		if want != got {
			t.Errorf("got %v want %v, given superStr: %v, subStr %v", got, want, superSub.superStr, superSub.subStr)
		}
	}

	negativeTestArr := [3]superSub{
		superSub{superStr: "funtsunami", subStr: "boring"},
		superSub{superStr: "1234", subStr: "45"},
		superSub{superStr: "booyeah123", subStr: "boobooyea"},
	}

	for _, superSub := range negativeTestArr {
		got := shouldDeletePod(superSub.superStr, superSub.subStr)
		want := false

		if want != got {
			t.Errorf("got %v want %v, given superStr: %v, subStr %v", got, want, superSub.superStr, superSub.subStr)
		}
	}
}

func testCreateFalco(t *testing.T) {
	falcoJson := `{"output": "18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)", ,"priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user_name":"root"}}`

	want := &Falco_Response{
		Output:   "18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)",
		Priority: "Notice",
		Rule:     "Terminal shell in container",
		Time:     "2018-03-28T18:37:22.181204909Z",
		Fields: Falco_Output_Fields{
			Container_id: "5dea0c14015a",
			Evt_time:     *big.NewInt(1522262242181204909),
			K8s_pod_name: "falco-6htpw",
			Proc_cmdline: "bash ",
			Proc_name:    "bash",
			Proc_pname:   "",
			Proc_tty:     34818,
			User_name:    "root",
		},
	}

	got, nil := createFalco(falcoJson)
	result, err := compareFR(got, want)
	if result != true {
		t.Errorf(err.Error())
	}
}

func compareFR(one *Falco_Response, two *Falco_Response) (bool, error) {
	if &one.Output != &two.Output {
		fmt.Println("Outputs do not match: %v || %v", &one.Output, &two.Output)
		return false, error.Error()
	}
	return compareFRO(&one.Fields, &two.Fields)
}

func compareFRO(one *Falco_Output_Fields, two *Falco_Output_Fields) (bool, error) {
	if &one.Container_id != &two.Container_id {
		fmt.Println("Container_ids do not match: %v || %v", &one.Container_id, &two.Container_id)
		return false, error.Error()
	}
	return true, nil
}

func TestMarshall(t *testing.T) {
	// Test idempotence
	// Create Falco struct
	// marshal
	// Unmarshal
	// compare to original

	// Test byte array input
	// define/create byte array
	// create desired falco struct
	// Unmarshal
	// compare to created struct

	// test ioReader type
	// define/create struct of the ioreader type.. likely from the opposide function
	// create desired falco struct
	// Unmarshal
	// compare to created struct
}
