package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ChooseKubeConfig(reader *bufio.Reader) string {
	fmt.Println("Escolher contexto de configuração:")
	fmt.Println("[1] ~/.kube/config")
	fmt.Println("[2] ~/.k3s/k3s.yaml")
	fmt.Println("[3] Outro")
	fmt.Print("Escolha uma opção: ")

	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return os.ExpandEnv("$HOME/.kube/config")
	case "2":
		return os.ExpandEnv("$HOME/.k3s/k3s.yaml")
	case "3":
		fmt.Print("Digite o caminho completo do arquivo kubeconfig: ")
		path, _ := reader.ReadString('\n')
		return strings.TrimSpace(path)
	default:
		fmt.Println("Opção inválida, usando configuração padrão ~/.kube/config")
		return os.ExpandEnv("$HOME/.kube/config")
	}
}
