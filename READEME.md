
# Kubernetes Pod Log Extractor

Este programa em Go foi desenvolvido para facilitar a coleta e filtragem de logs de todos os contêineres dentro de um pod específico em um cluster Kubernetes. Ele é ideal para desenvolvedores e operadores de sistemas que precisam de uma maneira rápida e eficiente de acessar logs relevantes para diagnóstico e monitoramento.

## Funcionalidades

- **Extração Automática de Logs**: Coleta logs de todos os contêineres dentro de um pod especificado sem a necessidade de seleção manual.
- **Filtragem de Logs**: Permite aos usuários especificar padrões via expressões regulares para filtrar as linhas de log desejadas.
- **Configuração Flexível**: Suporta múltiplas configurações de kubeconfig para se conectar a diferentes clusters Kubernetes.

## Pré-requisitos

Antes de executar o programa, certifique-se de que você tem o seguinte instalado em seu sistema:

- Go (Golang) - [Instruções de instalação](https://golang.org/doc/install)
- Acesso a um cluster Kubernetes e o arquivo kubeconfig correspondente.

## Configuração

1. **Clone o Repositório**:
   ```bash
   git clone [URL_DO_REPOSITÓRIO]
   cd [NOME_DO_REPOSITÓRIO]
   ```

2. **Construção**:
   Use o seguinte comando para compilar o programa:
   ```bash
   go build -o podLogFilter
   ```

## Uso

Para iniciar o programa, execute o binário compilado:

```bash
./podLogFilter
```

Siga as instruções interativas para escolher a configuração do kubeconfig, especificar o namespace e o nome do pod. O programa então perguntará por padrões de regex para filtrar os logs. Aqui estão os passos detalhados:

1. **Escolha do arquivo kubeconfig**: Selecione uma das configurações pré-definidas ou forneça um caminho customizado.
2. **Namespace**: Forneça o namespace onde o pod está localizado.
3. **Nome do Pod**: Forneça o nome do pod do qual os logs devem ser extraídos.
4. **Padrões de Filtragem**: Insira qualquer padrão de regex, separado por vírgulas, para filtrar os logs.

### Exemplo de Entrada

```plaintext
Digite Namespace: default
Digite nome do Pod: meu-pod
Enter log patterns to filter (comma-separated, e.g., ERROR,404,Warning): ERROR,Failed
```

## Formato de Saída

Os logs serão apresentados no console filtrados de acordo com os padrões especificados. Cada conjunto de logs de um contêiner começará com um cabeçalho indicando de qual contêiner os logs foram extraídos.

## Solução de Problemas

- **Permissões de Acesso**: Verifique se o arquivo kubeconfig tem as permissões corretas e se você tem permissões suficientes para acessar os logs dos pods.
- **Validação de Entrada**: Certifique-se de que todas as entradas estão corretas e que os nomes dos namespaces e pods existem no cluster.

### Estrutura do Projeto

```
/
├── cmd
│   └── main.go
├── internal
│   ├── config
│   │   └── config.go
│   ├── kubernetes
│   │   └── client.go
│   ├── logs
│   │   ├── fetcher.go
│   │   └── filter.go
│   └── util
│       └── regex.go
├── go.mod
├── go.sum
└── README.md
```

## Contribuindo

Contribuições para o projeto são bem-vindas! Sinta-se livre para clonar, modificar e enviar pull requests. Para bugs, questões ou sugestões, por favor abra uma issue no repositório do GitHub.


