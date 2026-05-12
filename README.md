# Email Worker

Worker responsável pelo processamento e envio de emails assíncronos.

O serviço executa tarefas agendadas para verificar mensagens pendentes na base de dados, processá-las e atualizar seus respectivos estados.

---

# 📌 Objetivo do Projeto

Este projeto foi desenvolvido para:

- Desacoplar envio de emails da API principal
- Melhorar performance da aplicação principal
- Evitar timeout em requests HTTP
- Garantir reprocessamento em caso de falhas
- Permitir escalabilidade independente
- Centralizar lógica de envio de emails

---

# 🧱 Arquitetura

```txt
API Principal
      │
      │ grava mensagens
      ▼
Base de Dados
      │
      ▼
Email Worker (Go)
      │
      ├── verifica mensagens pendentes
      ├── processa envio
      ├── atualiza status
      └── registra logs
```

---

# ⚙️ Tecnologias Utilizadas

- Go (Golang)
- PostgreSQL
- Cron Jobs (`robfig/cron`)
- SMTP
- Docker (opcional)

---

# 📂 Estrutura do Projeto

```txt
email-worker/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── configs/
│   ├── db/
│   ├── job/
│   ├── mailer/
│   ├── models/
│   ├── repository/
│   └── service/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

# 🚀 Como Executar

## 1. Clonar projeto

```bash
git clone <repo-url>
```

---

## 2. Entrar na pasta

```bash
cd email-worker
```

---

## 3. Instalar dependências

```bash
go mod tidy
```

---

## 4. Configurar variáveis ambiente

Criar arquivo `.env`:

```env
DATABASE_URL=postgres://user:password@localhost:5432/database

SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_EMAIL=example@gmail.com
SMTP_PASSWORD=password

CRON_EXPRESSION=0 */30 * * * *
```

---

## 5. Executar projeto

```bash
go run cmd/main.go
```

---

# 🕒 Cron Job

O worker utiliza `robfig/cron` com suporte a segundos:

```go
cron.New(cron.WithSeconds())
```

Isso significa que as expressões cron possuem 6 campos:

```txt
segundo minuto hora dia mês dia-da-semana
```

## Exemplos

| Frequência | Expressão |
|---|---|
| A cada 5 segundos | `*/5 * * * * *` |
| A cada 5 minutos | `0 */5 * * * *` |
| A cada 30 minutos | `0 */30 * * * *` |
| A cada 1 hora | `0 0 */1 * * *` |

---

# 🔄 Fluxo do Worker

## 1. Worker inicia

- Carrega variáveis ambiente
- Conecta na base de dados
- Inicializa scheduler

---

## 2. Cron executa tarefa

A cada intervalo configurado:

- Busca mensagens pendentes
- Valida mensagens
- Envia email
- Atualiza status

---

## 3. Atualização de status

Possíveis status:

| Status | Descrição |
|---|---|
| PENDING | Aguardando envio |
| PROCESSING | Em processamento |
| SENT | Enviado com sucesso |
| FAILED | Falha no envio |

---

# 🧠 Estratégia de Reprocessamento

Mensagens com falha podem ser:

- reprocessadas automaticamente
- reenviadas manualmente
- movidas para dead-letter futuramente

---

# 🛡️ Tratamento de Falhas

O projeto possui:

- logs de erro
- persistência de mensagens
- retry natural via cron
- isolamento do serviço principal

---

# 🔐 Variáveis de Ambiente

| Variável | Descrição |
|---|---|
| DATABASE_URL | URL conexão PostgreSQL |
| SMTP_HOST | Host SMTP |
| SMTP_PORT | Porta SMTP |
| SMTP_EMAIL | Email remetente |
| SMTP_PASSWORD | Senha SMTP |
| CRON_EXPRESSION | Frequência execução job |

---

# 📦 Banco de Dados

Exemplo simplificado tabela `messages`:

```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    receiver VARCHAR(255),
    subject VARCHAR(255),
    body TEXT,
    status VARCHAR(50) DEFAULT 'PENDING',
    error_message TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

---

# 📧 Processo de Envio

O worker:

1. Busca mensagens `PENDING`
2. Atualiza para `PROCESSING`
3. Tenta enviar
4. Atualiza:
   - `SENT`
   - ou `FAILED`

---

# 🧵 Graceful Shutdown

O projeto utiliza:

- `context.Context`
- captura de sinais do sistema
- encerramento controlado do cron

Exemplo:

```go
signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
```

---

# 📈 Melhorias Futuras

- RabbitMQ / Kafka
- Retry exponencial
- Dead Letter Queue
- Métricas Prometheus
- Health Check
- Dashboard administrativo
- Concorrência controlada
- Worker pools
- Observabilidade
- Distributed tracing

---

# 🐳 Docker (Opcional)

## Build

```bash
docker build -t email-worker .
```

## Run

```bash
docker run --env-file .env email-worker
```

---

# 🧪 Testes

Executar testes:

```bash
go test ./...
```

---

# 📋 Logs

Exemplos:

```txt
Iniciando worker...
CRONJOB configurado com sucesso
Verificando mensagens pendentes...
Email enviado com sucesso
```

---

# ⚠️ Cuidados Importantes

- Nunca salvar senha SMTP no código
- Utilizar `.env`
- Garantir índices na tabela
- Evitar múltiplos workers sem locking
- Validar retry para evitar duplicidade

---

# 🏗️ Escalabilidade

O worker pode evoluir para:

- múltiplas instâncias
- processamento distribuído
- filas de mensagens
- arquitetura orientada a eventos

---

# 👨‍💻 Autor

Desenvolvido por Felici.