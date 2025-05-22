# Guia de Contribuição

Obrigado por considerar contribuir com o InsideChurch! Este documento fornece diretrizes e instruções para contribuir com o projeto.

## 🚀 Como Começar

1. Faça um fork do repositório
2. Clone seu fork:
```bash
git clone https://github.com/seu-usuario/insidechurch.git
cd insidechurch
```

3. Configure o ambiente de desenvolvimento:
```bash
# Instalar dependências do backend
cd backend
go mod download

# Instalar dependências do frontend
cd ../frontend
npm install
```

4. Crie uma branch para sua feature:
```bash
git checkout -b feature/nome-da-feature
```

## 📝 Padrões de Código

### Backend (Go)

- Siga o [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` para formatação
- Execute `go vet` antes de commitar
- Escreva testes para novas funcionalidades
- Documente funções e pacotes

### Frontend (Vue/TypeScript)

- Siga o [Vue Style Guide](https://vuejs.org/style-guide/)
- Use TypeScript para tipagem
- Siga os padrões de nomenclatura do projeto
- Escreva testes unitários e E2E
- Documente componentes e funções

## 🔄 Padrões de Commit

Siga o padrão [Conventional Commits](https://www.conventionalcommits.org/):

```
<tipo>(<escopo>): <descrição>

[corpo opcional]

[rodapé opcional]
```

Tipos:
- `feat`: Nova feature
- `fix`: Correção de bug
- `docs`: Documentação
- `style`: Formatação
- `refactor`: Refatoração
- `test`: Testes
- `chore`: Tarefas de manutenção

Exemplos:
```
feat(auth): adiciona autenticação JWT
fix(api): corrige validação de email
docs(readme): atualiza instruções de instalação
```

## 🔍 Processo de Pull Request

1. **Preparação**
   - Atualize sua branch com a main
   - Execute testes localmente
   - Verifique linting e formatação

2. **Criação do PR**
   - Use o template de PR
   - Descreva as mudanças
   - Referencie issues relacionadas
   - Adicione screenshots se aplicável

3. **Revisão**
   - Responda aos comentários
   - Faça ajustes necessários
   - Mantenha o PR atualizado

4. **Merge**
   - Aguarde aprovação
   - Resolva conflitos se houver
   - Mantenha o histórico limpo

## 🧪 Testes

### Backend
```bash
cd backend
go test ./...
```

### Frontend
```bash
cd frontend
# Testes unitários
npm run test

# Testes E2E
npm run test:e2e
```

## 📚 Documentação

- Mantenha a documentação atualizada
- Documente novas features
- Atualize exemplos de código
- Adicione comentários quando necessário

## 🤝 Código de Conduta

- Seja respeitoso
- Mantenha discussões construtivas
- Aceite críticas construtivas
- Foque no bem do projeto

## 📫 Suporte

- Abra uma issue para bugs
- Use discussions para ideias
- Entre em contato com mantenedores

## 🎉 Agradecimentos

Obrigado por contribuir com o InsideChurch! Sua ajuda é fundamental para o sucesso do projeto. 