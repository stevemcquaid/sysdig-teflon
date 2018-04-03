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

	//for {
	//pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	//if err != nil {
	//panic(err.Error())
	//}
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	//// Examples for error handling:
	//// - Use helper functions like e.g. errors.IsNotFound()
	//// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//namespace := "default"
	//_, err = clientset.CoreV1().Pods(namespace).Get(podname, metav1.GetOptions{})
	//if errors.IsNotFound(err) {
	//fmt.Printf("Pod %s in namespace %s not found\n", podname, namespace)
	//} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//podname, namespace, statusError.ErrStatus.Message)
	//} else if err != nil {
	//panic(err.Error())
	//} else {
	//fmt.Printf("Found pod %s in namespace %s\n", podname, namespace)
	//}

	//}
}
