# 🛡️ VPSRecon - Reconhecimento Ofensivo Automatizado

**VPSRecon** é uma automação poderosa e modular para reconhecimento ofensivo, voltada para profissionais de **Bug Bounty**, **Red Team**, **CTFs** e **Pentest autorizado**. O projeto executa uma série de etapas com base nas melhores ferramentas open-source de segurança ofensiva e é totalmente otimizado para execução em ambientes VPS.

---

## 🚀 Funcionalidades

- ✅ Enumeração de subdomínios (Subfinder, Amass, Chaos, Assetfinder)
- ✅ Resolução DNS com dnsx
- ✅ Verificação de hosts vivos com httpx
- ✅ Scan de portas com Naabu
- ✅ Coleta de URLs com GAU, WaybackURLs e Katana
- ✅ Extração e análise de JavaScript (getJS, LinkFinder)
- ✅ Coleta de parâmetros e fuzzing
- ✅ Testes de XSS, SSTI, SQLi, Log4j, secrets
- ✅ Captura de screenshots com Gowitness
- ✅ Geração de relatório automatizado (.md)
- ✅ Log de erros com timestamp

---

## 🏗️ Estrutura de Diretórios

Após a execução, o projeto gerará a seguinte estrutura:

```
output/
└── alvo.com/
    ├── subdomains/
    │   ├── amass.txt
    │   ├── assetfinder.txt
    │   ├── chaos.txt
    │   ├── subfinder.txt
    │   └── all_subs.txt
    ├── resolved/
    │   └── resolved.txt
    ├── alive/
    │   └── alive.txt
    ├── ports/
    │   └── ports.txt
    ├── urls/
    │   ├── gau.txt
    │   ├── wayback.txt
    │   └── all_urls.txt
    ├── js/
    │   └── jsfiles.txt
    ├── parameters/
    │   └── params.txt
    ├── screenshots/
    │   └── *.png
    ├── vulnerabilities/
    │   ├── xss.txt
    │   ├── ssti.txt
    │   ├── secrets.txt
    │   ├── log4j.txt
    │   └── sqlmap.txt
    ├── report.md
    └── execution.log
```

---

## 💻 Requisitos Mínimos e Recomendados

### VPS Recomendada:

| Recurso | Recomendado        |
| ------- | ------------------ |
| CPU     | 8 vCores           |
| RAM     | 24 GB              |
| Disco   | SSD NVMe           |
| Tráfego | 32 TB ou ilimitado |
| SO      | Ubuntu 22.04+      |

### Também suporta VPS com:

- 2 GB, 4 GB, 8 GB, 16 GB de RAM (com ajustes nos módulos)

---

## 🛠️ Como Usar

```bash
chmod +x vpsrecon.sh
./vpsrecon.sh alvo.com
```

Todos os dados ficarão salvos na pasta `output/alvo.com/`.

---

## 📦 Instalar Ferramentas

Execute o script de instalação automática:

```bash
chmod +x install_tools.sh
./install_tools.sh
```

Isso instalará todas as dependências e ferramentas necessárias.

---

## 📑 Gerar Relatório Final

```bash
python3 generate_report.py alvo.com
```

Isso cria um `report.md` com os principais achados e erros.

---

## ⚙️ Arquivo de Configuração YAML (vpsrecon.yaml)

```yaml
project:
  name: VPSRecon
  version: 1.0
  description: Reconhecimento ofensivo automatizado.
  author: HexaSec Consultoria

vps:
  recommended:
    cpu: 8
    ram: 24GB
    storage: SSD NVMe
    traffic: 32TB+

structure:
  base_output_dir: output/
  subdirs:
    - subdomains
    - resolved
    - alive
    - ports
    - urls
    - js
    - parameters
    - vulnerabilities
    - screenshots
    - logs

tools:
  - name: subfinder
    command: subfinder -d {domain} -all -silent
    output: subdomains/subfinder.txt
  - name: naabu
    command: naabu -list {input} -silent -o {output}
  - name: httpx
    command: httpx -silent
  - name: dalfox
    command: dalfox file {input} --only-poc
  - name: gowitness
    command: gowitness file -f {input} -P screenshots

report:
  generate_with: python3 generate_report.py {domain}
  output: report.md

logging:
  execution_log: execution.log
  error_capture: true
```

---

## 📦 Tecnologias Utilizadas

- Go Tools: Subfinder, Naabu, Httpx, DNSx, Katana
- Recon: Amass, Assetfinder, Chaos, GAU, WaybackURLs
- Vulnerabilidades: Dalfox, Xray, Log4j-scan, SQLMap, Paramspider
- Captura: Gowitness
- Scripts: Bash + Python (report)

---

## ⚠️ Termos de Uso

> Esta ferramenta é para fins educacionais e uso ético apenas. O uso em sistemas sem permissão é ilegal. Utilize apenas em alvos autorizados.

---

## 📡 Conecte-se com a HexaSec

- 📸 Instagram: [@hexasec\_consultoria](https://www.instagram.com/hexasec_consultoria)
- 📧- 🌐 Projeto: VPSRecon by HexaSec

---

## ⭐ Contribua

Achou útil? Dê uma estrela ⭐ neste repositório e compartilhe! Sugestões de melhoria, issues e PRs são muito bem-vindos.

