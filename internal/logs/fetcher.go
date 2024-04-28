package logs

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1" // Uso de corev1 para tipos principais
	// Uso de metav1 para metadados
	"k8s.io/client-go/kubernetes"
)

// fetchLogForContainer obtém os logs de um contêiner específico em um pod.
func fetchLogForContainer(clientset *kubernetes.Clientset, namespace, podName, containerName string) (string, error) {
	logOptions := &corev1.PodLogOptions{Container: containerName} // Usando corev1.PodLogOptions
	logRequest := clientset.CoreV1().Pods(namespace).GetLogs(podName, logOptions)

	logStream, err := logRequest.Stream(context.Background())
	if err != nil {
		return "", fmt.Errorf("error opening log stream for container %s in pod %s: %v", containerName, podName, err)
	}
	defer logStream.Close()

	scanner := bufio.NewScanner(logStream)
	var logBuilder strings.Builder
	for scanner.Scan() {
		logBuilder.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading log for container %s in pod %s: %v", containerName, podName, err)
	}

	return logBuilder.String(), nil
}
