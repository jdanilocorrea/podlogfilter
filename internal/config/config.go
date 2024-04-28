package config

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ChooseKubeConfig(reader *bufio.Reader) string {
	fmt.Println("Escolher contexto de configuração:")
	fmt.Println("[1] ~/.kube/config")
	fmt.Println("[2] ~/.k3s/k3s.yaml")
	fmt.Println("[3] Outro")
	fmt.Print("Escolha uma opção: ")

	choice, _ := reader.ReadString('\n')
	switch strings.TrimSpace(choice) {
	case "1":
		return filepath.Join(os.Getenv("HOME"), ".kube", "config")
	case "2":
		return filepath.Join(os.Getenv("HOME"), ".k3s", "k3s.yaml")
	case "3":
		fmt.Print("Digite o caminho completo do arquivo kubeconfig: ")
		path, _ := reader.ReadString('\n')
		return strings.TrimSpace(path)
	default:
		fmt.Println("Opção inválida, usando configuração padrão ~/.kube/config")
		return filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}
}
