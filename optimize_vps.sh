#!/bin/bash

echo "=============================="
echo " 🚀 VPS Optimization"
echo "=============================="

# Ativar BBR (otimização de TCP)
echo "net.core.default_qdisc = fq" | sudo tee -a /etc/sysctl.conf
echo "net.ipv4.tcp_congestion_control = bbr" | sudo tee -a /etc/sysctl.conf
sudo sysctl -p

# Aumentar limites de arquivos
echo "* soft nofile 65535" | sudo tee -a /etc/security/limits.conf
echo "* hard nofile 65535" | sudo tee -a /etc/security/limits.conf

# Desativar serviços desnecessários (exemplos)
sudo systemctl disable --now ufw
sudo systemctl disable --now apache2
sudo systemctl disable --now mysql

# Ativar firewall básico (exemplo)
sudo ufw allow ssh
sudo ufw enable

echo "[+] VPS otimizada com sucesso!"
