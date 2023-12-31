package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

/*
source: https://github.com/kubernetes/client-go/blob/master/examples/out-of-cluster-client-configuration/main.go
*/
func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

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
	// watch for changes to pods
	watcher, err := clientset.CoreV1().Pods("").Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// loop through events from the watcher
	for event := range watcher.ResultChan() {
		pod := event.Object.(*corev1.Pod)
		switch event.Type {
		case watch.Added:
			fmt.Printf("Pod %s added\n", pod.Name)
			// todo: reconcile logic goes here
		case watch.Modified:
			fmt.Printf("Pod %s modified\n", pod.Name)
			// todo: reconcile logic goes here
		case watch.Deleted:
			fmt.Printf("Pod %s deleted\n", pod.Name)
			// todo: reconcile logic goes here
		}
	}
}
