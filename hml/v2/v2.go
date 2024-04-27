package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Definindo as flags para os argumentos de entrada
	podName := flag.String("pod", "", "Name of the pod to fetch logs from")
	namespace := flag.String("namespace", "", "Namespace of the pod")
	kubeconfig := flag.String("kubeconfig", os.Getenv("KUBECONFIG"), "Path to the Kubernetes config file")

	// Parseia as flags
	flag.Parse()

	// Verifica se o nome do pod e o namespace foram fornecidos
	if *podName == "" || *namespace == "" {
		fmt.Println("You must specify both the pod name and the namespace.")
		os.Exit(1)
	}

	// Cria a configuração do cliente do Kubernetes
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Cria um cliente do Kubernetes
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Configura as opções de log
	logOptions := &corev1.PodLogOptions{
		TailLines: int64Ptr(100), // Limita a quantidade de linhas para evitar logs muito extensos
	}

	// Obtém os logs do pod especificado
	req := clientset.CoreV1().Pods(*namespace).GetLogs(*podName, logOptions)
	logs, err := req.Stream(context.TODO())
	if err != nil {
		fmt.Printf("Error in opening stream: %v\n", err)
		os.Exit(1)
	}
	defer logs.Close()

	// Lê e filtra os logs
	buf := new(bytes.Buffer)
	buf.ReadFrom(logs)
	filterLogs(buf.String())
}

// Função para filtrar mensagens de erro ou status codes específicos nos logs
func filterLogs(logContent string) {
	// Regex para filtrar mensagens de erro ou status codes específicos
	errorRegex := regexp.MustCompile(`ERROR|Error|error|404|500|503`)
	lines := bytes.Split([]byte(logContent), []byte("\n"))

	fmt.Println("Filtered Logs:")
	for _, line := range lines {
		if errorRegex.Match(line) {
			fmt.Println(string(line))
		}
	}
}

// Função auxiliar para converter um int64 em um ponteiro para int64
func int64Ptr(i int64) *int64 {
	return &i
}
