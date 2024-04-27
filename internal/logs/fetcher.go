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

func GetPodLogsForAllContainers(clientset *kubernetes.Clientset, namespace, podName string) (string, error) {
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
		defer podLogs.Close()

		scanner := bufio.NewScanner(podLogs)
		allLogs.WriteString(fmt.Sprintf("Logs for container %s:\n", container.Name))
		for scanner.Scan() {
			allLogs.WriteString(scanner.Text() + "\n")
		}
	}
	return allLogs.String(), nil
}
