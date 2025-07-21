#!/bin/bash

# Script para otimização básica de VPS Ubuntu para pentest e recon

set -e

echo "🚀 Iniciando otimização da VPS..."

# Atualizar sistema
echo "🔄 Atualizando pacotes..."
sudo apt update && sudo apt upgrade -y

# Instalar pacotes básicos e ferramentas úteis
echo "📦 Instalando pacotes essenciais..."
sudo apt install -y git curl wget unzip build-essential net-tools htop jq nmap

# Otimizações de sysctl para rede
echo "⚙️ Configurando parâmetros do kernel (sysctl)..."
sudo tee /etc/sysctl.d/99-custom.conf > /dev/null <<EOF
# Aumentar número máximo de conexões abertas
fs.file-max = 1000000

# Ajustes de rede para melhorar desempenho de rede
net.core.somaxconn = 65535
net.core.netdev_max_backlog = 65535
net.ipv4.tcp_max_syn_backlog = 65535
net.ipv4.tcp_fin_timeout = 15
net.ipv4.tcp_tw_reuse = 1
net.ipv4.ip_local_port_range = 1024 65535
net.ipv4.tcp_max_tw_buckets = 1440000
net.ipv4.tcp_window_scaling = 1
net.ipv4.tcp_rmem = 4096 87380 6291456
net.ipv4.tcp_wmem = 4096 65536 6291456
net.ipv4.tcp_congestion_control = cubic
EOF

sudo sysctl --system

# Ajustar limites de arquivos abertos
echo "📈 Ajustando limites de arquivos abertos..."
sudo tee /etc/security/limits.d/99-custom.conf > /dev/null <<EOF
*         soft    nofile      1000000
*         hard    nofile      1000000
EOF

# Configurar shell para aplicar limites
echo "ulimit -n 1000000" | sudo tee -a /etc/profile

echo "✅ Otimização concluída! Reinicie a VPS para aplicar todas as mudanças."
