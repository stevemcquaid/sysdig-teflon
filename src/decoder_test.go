package main

import (
	//"bytes"
	"errors"
	"fmt"
	//"io/ioutil"
	//"log"
	"math/big"
	//"reflect"
	//"strings"
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

func TestCreateFalco(t *testing.T) {
	falcoJson := `{"output": "18:37:22.181204909: Notice A shell was spawned in a container with an attached terminal (user=root k8s.pod=falco-6htpw container=5dea0c14015a shell=bash parent=<NA> cmdline=bash  terminal=34818)","priority":"Notice","rule":"Terminal shell in container","time":"2018-03-28T18:37:22.181204909Z", "output_fields": {"container.id":"5dea0c14015a","evt.time":1522262242181204909,"k8s.pod.name":"falco-6htpw","proc.cmdline":"bash ","proc.name":"bash","proc.pname":null,"proc.tty":34818,"user_name":"root"}}`
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

	// convert falcoJson to []bytes
	falcoJsonBytes := []byte(falcoJson)
	// @TODO - handle errors here
	got, err := createFalco(falcoJsonBytes)
	if err != nil {
		t.Errorf(err.Error())
	}
	result, err := compareFR(got, want)
	if result != true {
		t.Errorf(err.Error())
	}
}

func compareFR(one *Falco_Response, two *Falco_Response) (bool, error) {
	if one.Output != two.Output {
		msg := fmt.Sprintf("Outputs do not match: %v || %v", one.Output, two.Output)
		return false, errors.New(msg)
	}
	if one.Priority != two.Priority {
		msg := fmt.Sprintf("Prioritys do not match: %v || %v", one.Priority, two.Priority)
		return false, errors.New(msg)
	}
	if one.Rule != two.Rule {
		msg := fmt.Sprintf("Rules do not match: %v || %v", one.Rule, two.Rule)
		return false, errors.New(msg)
	}
	if one.Time != two.Time {
		msg := fmt.Sprintf("Times do not match: %v || %v", one.Time, two.Time)
		return false, errors.New(msg)
	}
	return compareFRO(one.Fields, two.Fields)
}

func compareFRO(one Falco_Output_Fields, two Falco_Output_Fields) (bool, error) {
	if one.Container_id != two.Container_id {
		msg := fmt.Sprintf("Container_ids do not match: %v || %v", one.Container_id, two.Container_id)
		return false, errors.New(msg)
	}
	if one.Evt_time.Cmp(&two.Evt_time) != 0 {
		msg := fmt.Sprintf("Evt_times do not match: %v || %v", one.Evt_time.String(), two.Evt_time.String())
		return false, errors.New(msg)
	}
	if one.K8s_pod_name != two.K8s_pod_name {
		msg := fmt.Sprintf("K8s_pod_names do not match: %v || %v", one.K8s_pod_name, two.K8s_pod_name)
		return false, errors.New(msg)
	}
	if one.Proc_cmdline != two.Proc_cmdline {
		msg := fmt.Sprintf("Proc_cmdlines do not match: %v || %v", one.Proc_cmdline, two.Proc_cmdline)
		return false, errors.New(msg)
	}
	if one.Proc_name != two.Proc_name {
		msg := fmt.Sprintf("Proc_names do not match: %v || %v", one.Proc_name, two.Proc_name)
		return false, errors.New(msg)
	}
	if one.Proc_pname != two.Proc_pname {
		msg := fmt.Sprintf("Proc_pnames do not match: %v || %v", one.Proc_pname, two.Proc_pname)
		return false, errors.New(msg)
	}
	if one.Proc_tty != two.Proc_tty {
		msg := fmt.Sprintf("Proc_ttys do not match: %v || %v", one.Proc_tty, two.Proc_tty)
		return false, errors.New(msg)
	}
	if one.User_name != two.User_name {
		msg := fmt.Sprintf("User_names do not match: %v || %v", one.User_name, two.User_name)
		return false, errors.New(msg)
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
