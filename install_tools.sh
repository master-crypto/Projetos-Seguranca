#!/bin/bash

# ==============================
# ⚙️ VPSRecon - Install Tools
# ==============================

set -e

echo "🔧 Iniciando instalação das ferramentas do VPSRecon..."

# Update e pacotes básicos
sudo apt update -y
sudo apt install -y git curl wget unzip python3-pip build-essential jq chromium-browser

# Instalar Go
GO_VERSION="1.22.3"
cd ~
wget https://golang.org/dl/go$GO_VERSION.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go$GO_VERSION.linux-amd64.tar.gz
echo "export PATH=\$PATH:/usr/local/go/bin:\$HOME/go/bin" >> ~/.bashrc
source ~/.bashrc

# Criar diretório de ferramentas
mkdir -p ~/tools
cd ~/tools

# Instalar ferramentas Go
echo "📦 Instalando ferramentas em Go..."

go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest
go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
go install -v github.com/projectdiscovery/dnsx/cmd/dnsx@latest
go install -v github.com/projectdiscovery/katana/cmd/katana@latest
go install -v github.com/projectdiscovery/chaos-client/cmd/chaos@latest
go install -v github.com/lc/gau/v2/cmd/gau@latest
go install -v github.com/ffuf/ffuf/v2@latest
go install -v github.com/tomnomnom/assetfinder@latest
go install -v github.com/tomnomnom/anew@latest
go install -v github.com/tomnomnom/unfurl@latest
go install -v github.com/tomnomnom/gf@latest
go install -v github.com/hakluke/hakrawler@latest
go install -v github.com/hakluke/haklistgen@latest
go install -v github.com/hakluke/haktldextract@latest
go install -v github.com/projectdiscovery/interactsh/cmd/interactsh-client@latest

# Instalando Amass
sudo snap install amass

# Instalando Dalfox
go install github.com/hahwul/dalfox/v2@latest

# Instalando Gowitness
go install github.com/sensepost/gowitness@latest

# Instalando GetJS
git clone https://github.com/003random/getJS.git ~/tools/getJS

# Instalando Paramspider
git clone https://github.com/devanshbatham/paramspider.git ~/tools/paramspider
pip3 install -r ~/tools/paramspider/requirements.txt

# Instalando LinkFinder
git clone https://github.com/GerbenJavado/LinkFinder.git ~/tools/LinkFinder
pip3 install -r ~/tools/LinkFinder/requirements.txt

# Instalando Log4j-scan
git clone https://github.com/fullhunt/log4j-scan.git ~/tools/log4j-scan

# Instalando XSStrike
git clone https://github.com/s0md3v/XSStrike.git ~/tools/XSStrike
pip3 install -r ~/tools/XSStrike/requirements.txt

# Instalando SQLMap
git clone --depth 1 https://github.com/sqlmapproject/sqlmap.git ~/tools/sqlmap

echo "✅ Instalação concluída com sucesso!"
echo "📂 Ferramentas salvas em ~/tools"
echo "🔁 Execute 'source ~/.bashrc' para atualizar as variáveis de ambiente."
