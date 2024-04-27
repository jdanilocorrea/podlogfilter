Para tornar a interação mais amigável e permitir que o usuário escolha configurações e especifique detalhes durante a execução do programa, você pode adicionar uma interface de usuário interativa no terminal. Isso pode ser feito solicitando inputs do usuário após o programa iniciar, em vez de usar argumentos de linha de comando. Aqui está como você pode fazer isso em Go:

### Código com Interface Interativa


### Explicação das Mudanças

- **Interface de Usuário Interativa**: Usamos `bufio.Reader` para ler a entrada do usuário diretamente do terminal, permitindo escolhas e inputs personalizados.
- **Escolha do Arquivo de Configuração**: O usuário pode escolher entre opções pré-definidas de configuração ou especificar um caminho personalizado.
- **Flexibilidade e Usabilidade**: A interface interativa torna o programa mais flexível e amigável, especialmente útil para quem não está familiarizado com a linha de comando ou prefere uma interação mais guiada.

Este programa agora permite uma interação muito mais dinâmica e adaptável, ideal para cenários onde diferentes usuários ou diferentes configurações de cluster precisam ser rapidamente alternadas.