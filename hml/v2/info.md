### Explicações Adicionais

- **Flags de Linha de Comando**: As flags agora não possuem valores padrão para `namespace` e `pod`, e o programa verificará se esses valores foram realmente fornecidos antes de continuar.
- **Verificações de Argumentos**: Se `namespace` ou `podName` não forem especificados, o programa exibirá uma mensagem de erro e terminará. Isso garante que o usuário forneça os parâmetros necessários para a execução do script.
- **Execução**: Para rodar o programa, você precisará usar comandos como o seguinte no terminal, substituindo `<namespace>`, `<podname>`, e `<path_to_kubeconfig>` pelos valores apropriados:
  ```bash
  go run main.go -namespace=<namespace> -pod=<podname> -kubeconfig=<path_to_kubeconfig>
  ```
