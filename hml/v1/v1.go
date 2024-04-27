package main

import (
	"bytes"         // Pacote para manipulação de buffers de bytes
	"context"       // Pacote para gerenciamento de contextos em Go, usado em operações que podem ser canceladas ou expiram
	"flag"          // Pacote para parsing de flags de linha de comando
	"fmt"           // Pacote para operações de entrada/saída formatadas
	"path/filepath" // Pacote para manipulação de caminhos de arquivos em sistemas operacionais

	corev1 "k8s.io/api/core/v1"                   // Core V1 API do Kubernetes para operações com objetos core
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1" // Metadados da API do Kubernetes utilizados em seleções e filtros
	"k8s.io/client-go/kubernetes"                 // Cliente da API do Kubernetes para interação programática com clusters
	"k8s.io/client-go/tools/clientcmd"            // Ferramentas para configurar o cliente do Kubernetes a partir de um arquivo kubeconfig
	"k8s.io/client-go/util/homedir"               // Utilitário para encontrar o diretório home do usuário em diversas plataformas
)

func main() {
	// Declarando uma variável para armazenar o caminho do arquivo de configuração do Kubernetes
	var kubeconfig *string

	// Determinando o diretório home do usuário e configurando o caminho padrão do kubeconfig
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".k3s", "k3s.yaml"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	// Parseando as flags de linha de comando para permitir a customização do caminho do arquivo kubeconfig via flag
	flag.Parse()

	// Construindo a configuração do cliente do Kubernetes a partir do arquivo de configuração fornecido
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error()) // Em caso de erro, termina a execução com uma mensagem explicativa
	}

	// Criando um cliente do Kubernetes a partir da configuração para interagir com o cluster
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error()) // Em caso de erro, termina a execução com uma mensagem explicativa
	}

	// Listando todos os pods do cluster utilizando o cliente do Kubernetes
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error()) // Em caso de erro, termina a execução com uma mensagem explicativa
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items)) // Imprime o número de pods no cluster

	// Iterando sobre cada pod encontrado para recuperar e imprimir seus logs
	for _, pod := range pods.Items {
		logOptions := &corev1.PodLogOptions{}                                       // Configuração de opções de log (vazias por padrão)
		req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, logOptions) // Criando uma requisição de logs para cada pod
		logs, err := req.Stream(context.TODO())
		if err != nil {
			fmt.Printf("error in opening stream: %v\n", err)
			continue // Em caso de erro ao abrir o stream, pula para o próximo pod
		}
		defer logs.Close() // Garante que o stream dos logs seja fechado ao sair do bloco de código

		buf := new(bytes.Buffer)                                 // Cria um buffer para armazenar os logs
		buf.ReadFrom(logs)                                       // Lê os logs do stream para o buffer
		fmt.Printf("Logs for %s:\n%s\n", pod.Name, buf.String()) // Imprime os logs do pod
	}
}
