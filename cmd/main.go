package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

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

	clientset, err := kubernetes.NewClient(kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar cliente Kubernetes: %v\n", err)
		os.Exit(1)
	}

	rawLogs, err := logs.GetPodLogsForAllContainers(clientset, namespace, podName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao recuperar logs: %v\n", err)
		os.Exit(1)
	}

	pattern := util.ReadInput(reader, "Enter log patterns to filter (comma-separated, e.g., ERROR,404,Warning): ")
	filteredLogs := logs.FilterLogs(rawLogs, pattern)

	// Simulando dados estruturados de logs - isto deve ser ajustado para corresponder à sua fonte de dados real
	logEntries := make([]logs.LogEntry, len(filteredLogs))
	for i, logLine := range filteredLogs {
		logEntries[i] = logs.LogEntry{
			CreatedAt:  time.Now(), // Você precisará ajustar com o tempo real do log
			StatusCode: 200,        // Substitua com o status code real do log
			Method:     "GET",      // Substitua com o método HTTP real do log
			Path:       "/api",     // Substitua com o caminho real do log
			Details:    logLine,    // Este é o log textual
		}
	}

	logs.DisplayLogs(logEntries)
}
