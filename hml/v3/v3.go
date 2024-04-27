package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", os.Getenv("KUBECONFIG"), "path to the Kubernetes config file")
	podName    = flag.String("pod", "", "Name of the pod to fetch logs from")
	namespace  = flag.String("namespace", "", "Namespace of the pod")
	filter     = flag.String("filter", "", "Regex pattern to filter log messages")
)

func main() {
	flag.Parse()
	if *podName == "" || *namespace == "" {
		log.Fatalf("Both pod and namespace must be specified.")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Failed to build kubeconfig: %s", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create kubernetes client: %s", err)
	}

	logOptions := &corev1.PodLogOptions{}
	podLogs, err := clientset.CoreV1().Pods(*namespace).GetLogs(*podName, logOptions).Stream(context.TODO())
	if err != nil {
		log.Fatalf("Failed to get logs for pod: %s", err)
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(podLogs)
	if *filter != "" {
		applyFilter(buf.String(), *filter)
	} else {
		fmt.Println(buf.String())
	}
}

func applyFilter(logs, pattern string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatalf("Invalid regex pattern: %s", err)
	}

	lines := strings.Split(logs, "\n")
	for _, line := range lines {
		if regex.MatchString(line) {
			fmt.Println(line)
		}
	}
}
