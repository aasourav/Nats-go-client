package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Build the configuration from the kubeconfig file.
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %s\n", err.Error())
		return
	}

	// Create the Kubernetes client.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client: %s\n", err.Error())
		return
	}

	// Define the namespace, pod name, and container name.
	namespace := "core-ns"
	podName := "ingress-nginx-controller-6f6cf945d-lfsrc"
	// containerName := "my-container" // Optional: if the pod has a single container, you can omit this.

	// Create a request to get the pod logs.
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, &corev1.PodLogOptions{
		Follow: true, // Optional: use true if you want to follow the logs.
	})

	// Get the logs as a stream.
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		fmt.Printf("Error in opening stream: %s\n", err.Error())
		return
	}
	defer podLogs.Close()

	// Copy the logs to stdout.
	_, err = io.Copy(os.Stdout, podLogs)
	if err != nil {
		fmt.Printf("Error in copying logs to stdout: %s\n", err.Error())
	}
}
