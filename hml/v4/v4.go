package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Configure and parse command-line flags
	kubeconfig, podName, namespace, filter := setupFlags()
	clientset := configureKubeClient(kubeconfig)

	// Get pod logs
	logs, err := getPodLogs(clientset, *namespace, *podName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving logs: %v\n", err)
		os.Exit(1)
	}

	// Filter and display logs
	if *filter != "" {
		filterAndDisplayLogs(logs, *filter)
	} else {
		fmt.Println(logs)
	}
}

func setupFlags() (*string, *string, *string, *string) {
	kubeconfig := flag.String("kubeconfig", os.Getenv("KUBECONFIG"), "path to the Kubernetes config file")
	podName := flag.String("pod", "", "name of the pod to fetch logs from")
	namespace := flag.String("namespace", "", "namespace of the pod")
	filter := flag.String("filter", "", "regex pattern to filter log messages")
	flag.Parse()

	if *podName == "" || *namespace == "" {
		fmt.Println("Both pod and namespace must be specified.")
		flag.Usage()
		os.Exit(1)
	}
	return kubeconfig, podName, namespace, filter
}

func configureKubeClient(kubeconfig *string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to build kubeconfig: %v\n", err)
		os.Exit(2)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create Kubernetes client: %v\n", err)
		os.Exit(2)
	}
	return clientset
}

func getPodLogs(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	logOptions := &corev1.PodLogOptions{}
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	scanner := bufio.NewScanner(podLogs)
	var logs strings.Builder
	for scanner.Scan() {
		logs.WriteString(scanner.Text() + "\n")
	}
	return logs.String(), scanner.Err()
}

func filterAndDisplayLogs(logContent, pattern string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex pattern: %s\n", err)
		os.Exit(3)
	}
	lines := strings.Split(logContent, "\n")
	for _, line := range lines {
		if regex.MatchString(line) {
			fmt.Println(line)
		}
	}
}
