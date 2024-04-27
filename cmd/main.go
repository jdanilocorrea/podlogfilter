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
	kubeconfig := config.ChooseKubeConfig(reader)

	namespace := util.ReadInput(reader, "Digite Namespace: ")
	podName := util.ReadInput(reader, "Digite nome do Pod: ")

	clientset := kubernetes.NewClient(kubeconfig)
	logContent, err := logs.GetPodLogsForAllContainers(clientset, namespace, podName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar logs: %v\n", err)
		os.Exit(1)
	}

	patterns := util.ReadInput(reader, "Enter log patterns to filter (comma-separated, e.g., ERROR,404,Warning): ")
	filteredLogs := logs.FilterLogs(logContent, patterns)
	fmt.Println("Filtered Logs:\n", filteredLogs)
}
