package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jdanilocorrea/podlogfilter/internal/config"
	"github.com/jdanilocorrea/podlogfilter/internal/kubernetes"
	"github.com/jdanilocorrea/podlogfilter/internal/logs"
	"github.com/jdanilocorrea/podlogfilter/internal/util"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Iniciando cliente do Kubernetes...")
	kubeconfig, err := config.ChooseKubeConfig(reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escolher a configuração do kubeconfig: %v\n", err)
		os.Exit(1)
	}

	namespace := util.ReadInput(reader, "Digite Namespace: ")
	clientset, err := kubernetes.NewClient(kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar cliente Kubernetes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Escolha a origem dos logs:")
	fmt.Println("[1] Pod")
	fmt.Println("[2] Service")
	sourceOption := util.ReadInput(reader, "Digite a opção: ")

	var logEntries []logs.LogEntry

	switch sourceOption {
	case "1":
		podName := util.ReadInput(reader, "Digite nome do Pod: ")
		logEntries, err = logs.GetPodLogs(clientset, namespace, podName)
	case "2":
		serviceName := util.ReadInput(reader, "Digite nome do Service: ")
		logEntries, err = logs.GetServiceLogs(clientset, namespace, serviceName)
	default:
		fmt.Println("Opção inválida.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar logs: %v\n", err)
		os.Exit(1)
	}

	logs.DisplayLogs(logEntries)
}
