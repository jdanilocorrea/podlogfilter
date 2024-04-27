package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	kubeconfig := chooseKubeConfig(reader)

	fmt.Print("Digite Namespace: ")
	namespace, _ := reader.ReadString('\n')
	namespace = strings.TrimSpace(namespace)

	fmt.Print("Digite nome do Pod: ")
	podName, _ := reader.ReadString('\n')
	podName = strings.TrimSpace(podName)

	clientset := configureKubeClient(kubeconfig)

	logs, err := getPodLogsForAllContainers(clientset, namespace, podName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar logs: %v\n", err)
		return
	}

	fmt.Print("Enter log patterns to filter (comma-separated, e.g., ERROR,404,Warning): ")
	input, _ := reader.ReadString('\n')
	patterns := strings.Split(strings.TrimSpace(input), ",")
	regexPattern := buildRegexPattern(patterns)

	filterAndDisplayLogs(strings.NewReader(logs), regexPattern)
}

func chooseKubeConfig(reader *bufio.Reader) string {
	fmt.Println("Escolher contexto de configuração:")
	fmt.Println("[1] ~/.kube/config")
	fmt.Println("[2] ~/.k3s/k3s.yaml")
	fmt.Println("[3] Outro")
	fmt.Print("Escolha uma opção: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return os.ExpandEnv("$HOME/.kube/config")
	case "2":
		return os.ExpandEnv("$HOME/.k3s/k3s.yaml")
	case "3":
		fmt.Print("Digite o caminho completo do arquivo kubeconfig: ")
		path, _ := reader.ReadString('\n')
		return strings.TrimSpace(path)
	default:
		fmt.Println("Opção inválida, usando configuração padrão ~/.kube/config")
		return os.ExpandEnv("$HOME/.kube/config")
	}
}

func configureKubeClient(kubeconfig string) *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falha ao construir kubeconfig: %v\n", err)
		os.Exit(2)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Falha ao criar cliente Kubernetes: %v\n", err)
		os.Exit(2)
	}
	return clientset
}

func getPodLogsForAllContainers(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return "", err
	}

	var allLogs strings.Builder
	for _, container := range pod.Spec.Containers {
		logOptions := &corev1.PodLogOptions{Container: container.Name}
		req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)
		podLogs, err := req.Stream(context.TODO())
		if err != nil {
			allLogs.WriteString(fmt.Sprintf("Error retrieving logs for container %s: %v\n", container.Name, err))
			continue
		}
		defer podLogs.Close() // Close the stream after log collection is complete

		scanner := bufio.NewScanner(podLogs)
		allLogs.WriteString(fmt.Sprintf("Logs for container %s:\n", container.Name))
		for scanner.Scan() {
			allLogs.WriteString(scanner.Text() + "\n")
		}
	}
	return allLogs.String(), nil
}

func buildRegexPattern(patterns []string) string {
	var regexParts []string
	for _, pattern := range patterns {
		regexParts = append(regexParts, regexp.QuoteMeta(pattern))
	}
	return strings.Join(regexParts, "|")
}

func filterAndDisplayLogs(logStream io.Reader, pattern string) {
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid regex pattern: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(logStream)
	fmt.Println("Filtered Logs:")
	for scanner.Scan() {
		line := scanner.Text()
		if regex.MatchString(line) {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading log stream: %v\n", err)
	}
}
