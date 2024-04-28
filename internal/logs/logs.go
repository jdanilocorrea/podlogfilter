package logs

import (
	"context"
	"fmt"
	"strings"
	"time" // Importando o pacote time para manipulação de datas

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	// Garantir que este import está correto
)

// LogEntry representa uma entrada de log com os campos esperados.

type LogEntry struct {
	CreatedAt string // Armazena a data e hora no formato string RFC 3339
	Status    int    // Código de status HTTP para simplificação
	Method    string // Método HTTP utilizado na requisição
	Path      string // Caminho da requisição
	Details   string // Detalhes adicionais ou o corpo do log
}

// GetPodLogs obtém os logs de todos os contêineres de um pod específico.
func GetPodLogs(clientset *kubernetes.Clientset, namespace, podName string) ([]LogEntry, error) {
	pod, err := clientset.CoreV1().Pods(namespace).Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting pod %s: %v", podName, err)
	}

	var logEntries []LogEntry
	for _, container := range pod.Spec.Containers {
		log, err := fetchLogForContainer(clientset, namespace, podName, container.Name)
		if err != nil {
			logEntries = append(logEntries, LogEntry{
				CreatedAt: time.Now().Format(time.RFC3339), // Usando o timestamp atual
				Status:    500,
				Method:    "GET",
				Path:      fmt.Sprintf("Pod: %s, Container: %s", podName, container.Name),
				Details:   fmt.Sprintf("Error retrieving logs: %v", err),
			})
			continue
		}
		logEntries = append(logEntries, LogEntry{
			CreatedAt: time.Now().Format(time.RFC3339), // Usando o timestamp atual
			Status:    200,
			Method:    "GET",
			Path:      fmt.Sprintf("Pod: %s, Container: %s", podName, container.Name),
			Details:   log,
		})
	}

	return logEntries, nil
}

// GetServiceLogs obtém os logs dos pods que estão sob um serviço específico.
func GetServiceLogs(clientset *kubernetes.Clientset, namespace, serviceName string) ([]LogEntry, error) {
	service, err := clientset.CoreV1().Services(namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("error getting service %s: %v", serviceName, err)
	}

	labelSelector := labelsToSelector(service.Spec.Selector)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{LabelSelector: labelSelector})
	if err != nil {
		return nil, fmt.Errorf("error listing pods for service %s with selector %s: %v", serviceName, labelSelector, err)
	}

	var logEntries []LogEntry
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			log, err := fetchLogForContainer(clientset, namespace, pod.Name, container.Name)
			if err != nil {
				logEntries = append(logEntries, LogEntry{
					CreatedAt: time.Now().Format(time.RFC3339), // Usando o timestamp atual
					Status:    500,
					Method:    "GET",
					Path:      fmt.Sprintf("Service: %s, Pod: %s, Container: %s", serviceName, pod.Name, container.Name),
					Details:   fmt.Sprintf("Error retrieving logs: %v", err),
				})
				continue
			}
			logEntries = append(logEntries, LogEntry{
				CreatedAt: time.Now().Format(time.RFC3339), // Usando o timestamp atual
				Status:    200,
				Method:    "GET",
				Path:      fmt.Sprintf("Service: %s, Pod: %s, Container: %s", serviceName, pod.Name, container.Name),
				Details:   log,
			})
		}
	}

	return logEntries, nil
}

// Helper function to convert label selector map to a string.
func labelsToSelector(labels map[string]string) string {
	var selectorParts []string
	for key, val := range labels {
		selectorParts = append(selectorParts, fmt.Sprintf("%s=%s", key, val))
	}
	return strings.Join(selectorParts, ",")
}
