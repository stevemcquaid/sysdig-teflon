package main

import (
	//"flag"
	"errors"
	"fmt"
	//"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
	//"os"
	//"path/filepath"
	//"time"
)

func handleDeleteK8SPod(w http.ResponseWriter, r *http.Request) {
	podname := r.URL.Query().Get("pod")
	namespace := r.URL.Query().Get("namespace")
	kubeconfig := "/src/kube.config"

	//fmt.Printf("podname: %v, namespace: %v", podname, namespace)

	// Add metrics handler here to provide scoped metrics with pod, namespace, etc.
	infections.Inc()

	msg, err := deleteK8SPod(&kubeconfig, podname, namespace)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
	} else {
		w.Header().Set("content-type", "text/plain")
		w.Write([]byte(msg))
	}

}

func deleteK8SPod(kubeconfig *string, podname string, namespace string) (string, error) {
	log.Printf("Attempting to delete pod: %v in namespace: %v...", podname, namespace)

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		msg := fmt.Sprintf("Unable to create config from kubeconfig")
		return "", errors.New(msg)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		msg := fmt.Sprintf("Unable to create clientset from kubeconfig")
		return "", errors.New(msg)
	}

	// Delete the pod
	err2 := clientset.CoreV1().Pods(namespace).Delete(podname, &metav1.DeleteOptions{})
	if err2 != nil {
		msg := fmt.Sprintf("Unable to delete pod: %v in namespace: %v", podname, namespace)
		return "", errors.New(msg)
	}

	msg := fmt.Sprintf("Successfully deleted pod: %v in namespace: %v", podname, namespace)
	log.Println(msg)
	return msg, nil
}
