package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Escolha do arquivo de configuração do Kubernetes
	kubeconfig := chooseKubeConfig(reader)

	// Entrada do namespace
	fmt.Print("Digite Namespace: ")
	namespace, _ := reader.ReadString('\n')
	namespace = strings.TrimSpace(namespace)

	// Entrada do nome do pod
	fmt.Print("Digite nome do Pod: ")
	podName, _ := reader.ReadString('\n')
	podName = strings.TrimSpace(podName)

	clientset := configureKubeClient(kubeconfig)

	// Obter e exibir os logs
	logs, err := getPodLogs(clientset, namespace, podName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar logs: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(logs)
}

func chooseKubeConfig(reader *bufio.Reader) string {
	fmt.Println("Escolher contexto de configuração:")
	fmt.Println("[1] ~/.kube/config")
	fmt.Println("[2] ~/k3s/k3s.yaml")
	fmt.Println("[3] Outro")
	fmt.Print("Escolha uma opção: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return os.ExpandEnv("$HOME/.kube/config")
	case "2":
		return os.ExpandEnv("$HOME/k3s/k3s.yaml")
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

func getPodLogs(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
	logOptions := &corev1.PodLogOptions{}
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	var logs strings.Builder
	scanner := bufio.NewScanner(podLogs)
	for scanner.Scan() {
		logs.WriteString(scanner.Text() + "\n")
	}
	return logs.String(), scanner.Err()
}
