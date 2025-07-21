# VPSRecon 🚀

Framework de Reconhecimento Automático para Bug Bounty e Pentest.

---

## 📌 Sobre o Projeto

O **VPSRecon** é uma estrutura modular projetada para automação de tarefas de reconhecimento em segurança ofensiva, bug bounty e pentest. Ele permite desde a enumeração de subdomínios até varredura de portas, verificação de serviços ativos, crawling e identificação de vulnerabilidades.

Desenvolvido para rodar em VPS de alta performance, otimizando tempo e resultados.

# 📁 Estrutura de Diretórios

├── output/ # Resultados das execuções
│ ├── subdomains/ # Subdomínios encontrados
│ ├── resolved/ # Subdomínios resolvidos (DNS)
│ ├── alive/ # Hosts ativos (HTTP/HTTPS)
│ ├── ports/ # Varredura de portas
│ ├── crawl/ # Dados de crawling
│ ├── vulnerabilities/ # Relatórios de vulnerabilidades
│ └── screenshots/ # Capturas de tela dos hosts
├── config/ # Arquivos de configuração
│ └── webhook.json # Webhooks (Discord, Slack, Telegram)
├── scripts/ # Scripts e automações
│ └── recon.go # Script principal em Go
├── install.sh # Script de instalação de dependências
├── optimize_vps.sh # Script de otimização da VPS
└── README.md # Documentação do projeto


## ⚙️ Instalação ##

### Pré-requisitos

- VPS com **Ubuntu 22.04+**
- Acesso **root** ou via **sudo**
- Go instalado (`sudo apt install golang`)

### Instalação completa

```bash
chmod +x install.sh optimize_vps.sh
./install.sh
./optimize_vps.sh



### Pré-requisitos ##

- VPS com **Ubuntu 22.04+**
- Acesso **root** ou via **sudo**
- Go instalado (`sudo apt install golang`)

### Instalação completa

bash chmod +x install.sh optimize_vps.sh

./install.sh
./optimize_vps.sh

🚀 Como Usar

Executar Reconhecimento

cd scripts
go run recon.go alvo.com
Os resultados serão salvos automaticamente na pasta output/.

🛠️ Ferramentas Integradas
Subfinder – Enumeração de subdomínios

Httpx – Verificação de hosts ativos

Naabu – Varredura de portas

Masscan – Scanner de portas ultra rápido

Nuclei – Scans de vulnerabilidades

Chromium – Captura de screenshots via automação

Configuração de Webhook

Edite o arquivo:

config/webhook.json

Exemplo de configuração:

{
  "discord_webhook": "https://discord.com/api/webhooks/SEU_WEBHOOK",
  "slack_webhook": "https://hooks.slack.com/services/SEU_WEBHOOK",
  "telegram_bot_token": "BOT_TOKEN",
  "telegram_chat_id": "CHAT_ID"
}

🧠 To-Do | Melhorias Futuras
Integração com APIs de ASN e Shodan

Coleta de arquivos JS e análise de endpoints

Identificação de tokens e credenciais expostas

Integração com Burp Suite para scans automatizados

Dashboard Web (em desenvolvimento)

🤝 Contribuições
Contribuições são bem-vindas!
Sinta-se livre para abrir issues ou enviar pull requests.

🛡️ Licença
Distribuído sob a Licença MIT.
Veja o arquivo LICENSE para mais informações.

