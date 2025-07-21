package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

// Configuração de webhook
type WebhookConfig struct {
	WebhookURL string `json:"webhook_url"`
}

var (
	rateLimiter = time.Tick(5 * time.Second) // ⏱️ 1 mensagem a cada 5 segundos
	wg          sync.WaitGroup
)

// Carrega configuração do webhook
func loadWebhookConfig() WebhookConfig {
	file, err := os.ReadFile("config/webhook.json")
	if err != nil {
		log.Println("Webhook config não encontrado, usando placeholder.")
		return WebhookConfig{WebhookURL: "https://discord.com/api/webhooks/xxx/xxx"}
	}
	var config WebhookConfig
	_ = json.Unmarshal(file, &config)
	return config
}

// Envia alerta para webhook com rate limit
func sendWebhook(webhookURL, message string) {
	<-rateLimiter // ⏳ Aguarda o tick (5 segundos entre mensagens)

	wg.Add(1)
	go func() {
		defer wg.Done()

		payload := map[string]string{"content": message}
		data, _ := json.Marshal(payload)

		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			log.Println("Falha ao enviar webhook:", err)
			return
		}
		defer resp.Body.Close()

		log.Printf("Webhook enviado. Status: %s", resp.Status)
	}()
}








// Função para executar comandos shell

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Erro ao executar %s: %v", name, err)
	}
}

// Cria pastas de output
func createFolders() {
	folders := []string{
		"output/subdomains",
		"output/resolved",
		"output/alive",
		"output/ports",
		"output/crawl",
		"output/vulnerabilities",
		"output/screenshots",
	}

	for _, folder := range folders {
		_ = os.MkdirAll(folder, os.ModePerm)
	}
}

// Pipeline de Recon
func recon(domain string, webhook WebhookConfig) {
	start := time.Now()

	sendWebhook(webhook.WebhookURL, fmt.Sprintf("🚀 Iniciando Recon para %s", domain))

	// Subdomain Enum
	fmt.Println("\n[+] Buscando subdomínios...")
	runCommand("subfinder", "-d", domain, "-silent", "-o", "output/subdomains/"+domain+".txt")

	// DNS Resolve
	fmt.Println("\n[+] Resolvendo DNS...")
	runCommand("dnsx", "-l", "output/subdomains/"+domain+".txt", "-silent", "-o", "output/resolved/"+domain+".txt")

	// Check Hosts Alive
	fmt.Println("\n[+] Verificando hosts vivos...")
	runCommand("httpx", "-l", "output/resolved/"+domain+".txt", "-silent", "-o", "output/alive/"+domain+".txt")

	//Screenshots
	fmt.Println("\n[+] Capturando screenshots...")

	runCommand( "gowitness", "scan" ,"file" , "-f", "output/alive/"+domain+".txt", "--screenshot-path",  "output/screenshots")

	// Port Scan
	fmt.Println("\n[+] Escaneando portas abertas...")
	runCommand("naabu", "-l", "output/resolved/"+domain+".txt", "-silent", "-o", "output/ports/"+domain+".txt")

	// Crawling
	fmt.Println("\n[+] Executando crawling...")
	runCommand("katana", "-list", "output/alive/"+domain+".txt", "-o", "output/crawl/"+domain+".txt")

		// Nuclei Scan
	fmt.Println("\n[+] Executando Nuclei scan...")
	runCommand("nuclei", "-l", "output/alive/"+domain+".txt", "-o", "output/vulnerabilities/"+domain+"_nuclei.txt")

	// Dalfox Scan
	fmt.Println("\n[+] Executando Dalfox scan...")
	runCommand("dalfox", "file", "output/alive/"+domain+".txt", "--output", "output/vulnerabilities/"+domain+"_dalfox.txt")

	// SQLMap (Exemplo básico)
	fmt.Println("\n[+] Executando SQLMap (manual review recomendado)...")
	runCommand("sqlmap", "-m", "output/alive/"+domain+".txt", "--batch", "-o")

	sendWebhook(webhook.WebhookURL, fmt.Sprintf("✅ Recon concluído para %s em %s", domain, time.Since(start).String()))
	fmt.Println("\n[+] Recon finalizado!")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Uso: ./recon <dominio>")
		os.Exit(1)
	}
	domain := os.Args[1]

	// Cria pastas
	createFolders()

	// Carrega webhook
	webhook := loadWebhookConfig()

	// Executa Recon
	recon(domain, webhook)
}
