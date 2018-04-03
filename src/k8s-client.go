package main

import (
	//"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"os"
	//"path/filepath"
	"time"
)

func deleteK8SPod(kubeconfig *string, podname string, namespace string) {
	fmt.Println("Deleting podname: ", podname, "namespace: ", namespace, "...")

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Delete the pod
	err := clientset.CoreV1().Pods(namespace).Delete(podname)
	if err != nil {
		fmt.Println("")
	}

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

	//time.Sleep(10 * time.Second)
	//}
}
