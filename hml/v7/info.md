### Aprimoramentos Propostos:

1. **Tratamento de Erros Aprimorado**: Melhorar como os erros são tratados e reportados ao usuário.
2. **Funcionalidade de Fechamento Automático de Streams**: Usar o defer de maneira correta para garantir que todos os recursos sejam liberados apropriadamente.
3. **Claridade nos Comentários**: Incluir comentários mais descritivos para cada função, explicando o que elas fazem.
4. **Validação de Entrada do Usuário**: Assegurar que todas as entradas do usuário sejam validadas para evitar erros de execução.
5. **Uso Eficiente de Recursos**: Otimizar a coleta de logs para evitar abrir múltiplos streams simultaneamente.


### Melhorias Implementadas:

- **Gerenciamento de Recursos**: Utilizei o `defer` para garantir que cada stream de log seja fechado após o término da coleta de logs para cada contêiner, evitando deixar recursos abertos inadvertidamente.
- **Validação de Entrada**: As entradas do usuário para a escolha de configuração agora são validadas, e mensagens de erro claras são exibidas para entradas inválidas.
- **Organização e Comentários**: O código foi organizado para melhor fluidez e compreensão, e comentários pertinentes foram adicionados para clarificar a função de cada bloco de código.
