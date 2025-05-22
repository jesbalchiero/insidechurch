# Checklist de Testes - InsideChurch

## 1. Testes de Responsividade

### Mobile (375px)
- [ ] Header
  - [ ] Logo "InsideChurch" visível e legível
  - [ ] Menu hamburguer aparece corretamente
  - [ ] Drawer lateral abre e fecha suavemente
  - [ ] Links do drawer funcionam corretamente
  - [ ] Avatar e dropdown não aparecem

- [ ] Página de Login
  - [ ] Formulário ocupa largura adequada
  - [ ] Campos de input com tamanho confortável para toque
  - [ ] Botão de login com tamanho adequado
  - [ ] Link para registro visível e clicável
  - [ ] Mensagens de erro legíveis

- [ ] Página de Registro
  - [ ] Formulário ocupa largura adequada
  - [ ] Campos de input com tamanho confortável para toque
  - [ ] Botão de registro com tamanho adequado
  - [ ] Link para login visível e clicável
  - [ ] Mensagens de erro legíveis

- [ ] Dashboard
  - [ ] Cards e grids se ajustam à largura
  - [ ] Textos permanecem legíveis
  - [ ] Botões e controles com tamanho adequado para toque
  - [ ] Tabelas com scroll horizontal quando necessário

### Tablet (768px)
- [ ] Header
  - [ ] Layout se ajusta para versão tablet
  - [ ] Menu hamburguer ainda presente
  - [ ] Espaçamento adequado entre elementos

- [ ] Páginas de Autenticação
  - [ ] Formulários com largura otimizada
  - [ ] Melhor aproveitamento do espaço
  - [ ] Elementos bem distribuídos

- [ ] Dashboard
  - [ ] Grid se ajusta para 2 colunas
  - [ ] Cards com tamanho adequado
  - [ ] Tabelas com melhor visualização

### Desktop (1920px)
- [ ] Header
  - [ ] Menu horizontal visível
  - [ ] Avatar e dropdown funcionando
  - [ ] Espaçamento adequado entre elementos

- [ ] Páginas de Autenticação
  - [ ] Formulários centralizados
  - [ ] Largura máxima adequada
  - [ ] Elementos bem distribuídos

- [ ] Dashboard
  - [ ] Grid com múltiplas colunas
  - [ ] Cards com tamanho otimizado
  - [ ] Tabelas com visualização completa

## 2. Testes de Funcionalidade

### Autenticação
- [ ] Login
  - [ ] Validação de campos obrigatórios
  - [ ] Mensagem de erro para credenciais inválidas
  - [ ] Redirecionamento após login
  - [ ] Persistência do token
  - [ ] Toast de sucesso/erro

- [ ] Registro
  - [ ] Validação de campos obrigatórios
  - [ ] Validação de formato de email
  - [ ] Validação de força da senha
  - [ ] Confirmação de senha
  - [ ] Toast de sucesso/erro

- [ ] Logout
  - [ ] Limpeza do token
  - [ ] Redirecionamento para login
  - [ ] Toast de despedida

### Dashboard
- [ ] Carregamento
  - [ ] Loading state
  - [ ] Tratamento de erros
  - [ ] Dados atualizados

- [ ] Interações
  - [ ] Filtros funcionando
  - [ ] Ordenação de tabelas
  - [ ] Paginação
  - [ ] Ações em lote

## 3. Testes de Performance

- [ ] Carregamento Inicial
  - [ ] Tempo de carregamento < 3s
  - [ ] First Contentful Paint < 1.5s
  - [ ] Time to Interactive < 3.5s

- [ ] Navegação
  - [ ] Transições suaves
  - [ ] Sem lag em interações
  - [ ] Carregamento lazy de imagens

- [ ] Responsividade
  - [ ] Adaptação suave entre breakpoints
  - [ ] Sem quebras de layout
  - [ ] Imagens otimizadas

## 4. Testes de Acessibilidade

- [ ] Navegação
  - [ ] Suporte a teclado
  - [ ] Foco visível
  - [ ] Ordem lógica de tabulação

- [ ] Conteúdo
  - [ ] Contraste adequado
  - [ ] Textos alternativos
  - [ ] Estrutura semântica

- [ ] Formulários
  - [ ] Labels associados
  - [ ] Mensagens de erro acessíveis
  - [ ] Validação em tempo real

## 5. Testes de Compatibilidade

### Navegadores
- [ ] Chrome (última versão)
- [ ] Firefox (última versão)
- [ ] Safari (última versão)
- [ ] Edge (última versão)

### Dispositivos
- [ ] iOS (Safari)
- [ ] Android (Chrome)
- [ ] Tablets (iOS/Android)

## Como Executar os Testes

1. **Testes de Responsividade**
```bash
# Usar DevTools do navegador
# Chrome: F12 > Toggle Device Toolbar
# Firefox: F12 > Responsive Design Mode
```

2. **Testes de Funcionalidade**
```bash
# Executar testes unitários
npm run test

# Executar testes e2e
npm run test:e2e
```

3. **Testes de Performance**
```bash
# Usar Lighthouse no Chrome DevTools
# Performance > Lighthouse
```

4. **Testes de Acessibilidade**
```bash
# Usar Lighthouse no Chrome DevTools
# Accessibility > Lighthouse
```

## Ferramentas Recomendadas

- Chrome DevTools
- Lighthouse
- axe DevTools
- Responsively App
- BrowserStack
- Jest
- Cypress 