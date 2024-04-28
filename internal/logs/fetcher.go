package logs

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetPodLogsForAllContainers obtém os logs de todos os contêineres de um pod específico.
func GetPodLogsForAllContainers(clientset *kubernetes.Clientset, namespace, podName string) ([]string, error) {
	// Pega as informações do pod especificado.
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error fetching pod info: %w", err)
	}

	var logs []string
	for _, container := range pod.Spec.Containers {
		log, err := fetchLogForContainer(clientset, namespace, podName, container.Name)
		if err != nil {
			// Se não conseguir pegar os logs de um contêiner, adiciona a mensagem de erro na lista de logs
			logs = append(logs, fmt.Sprintf("Error retrieving logs for container %s: %v", container.Name, err))
			continue
		}
		logs = append(logs, log)
	}

	return logs, nil
}

// fetchLogForContainer obtém os logs de um contêiner específico.
func fetchLogForContainer(clientset *kubernetes.Clientset, namespace, podName, containerName string) (string, error) {
	// Configuração das opções de logs.
	logOptions := &corev1.PodLogOptions{Container: containerName}

	// Cria a requisição de logs para o contêiner.
	req := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)

	// Captura o stream de logs.
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", fmt.Errorf("error in opening stream for container %s: %w", containerName, err)
	}
	defer podLogs.Close()

	// Lê os logs e os concatena em uma string.
	scanner := bufio.NewScanner(podLogs)
	var sb strings.Builder
	for scanner.Scan() {
		sb.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading log for container %s: %w", containerName, err)
	}

	return sb.String(), nil
}
