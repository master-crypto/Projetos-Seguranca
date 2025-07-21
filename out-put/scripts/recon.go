package main

import (
    "fmt"
    "log"
    "os"
    "os/exec"
    "path/filepath"
)

func checkAndRun(name string, cmd *exec.Cmd) {
    fmt.Printf("🚀 Executando %s...\n", name)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        log.Printf("❌ Erro ao executar %s: %v\n", name, err)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("❗ Uso: go run recon.go <domínio>")
        return
    }

    domain := os.Args[1]
    baseOutput := filepath.Join("../output", domain)

    // Criar estrutura de diretórios
    dirs := []string{
        "subdomains", "resolved", "alive",
        "ports", "crawl", "vulnerabilities", "screenshots",
    }

    for _, dir := range dirs {
        path := filepath.Join(baseOutput, dir)
        os.MkdirAll(path, 0755)
    }

    fmt.Printf("📂 Diretórios criados em: %s\n", baseOutput)

    // Etapas de Reconhecimento
    checkAndRun("Subfinder",
        exec.Command("subfinder", "-d", domain, "-o", filepath.Join(baseOutput, "subdomains", "subdomains.txt")))

    checkAndRun("Httpx",
        exec.Command("httpx", "-l", filepath.Join(baseOutput, "subdomains", "subdomains.txt"),
            "-o", filepath.Join(baseOutput, "alive", "alive.txt")))

    checkAndRun("Naabu",
        exec.Command("naabu", "-list", filepath.Join(baseOutput, "alive", "alive.txt"),
            "-o", filepath.Join(baseOutput, "ports", "ports.txt")))

    checkAndRun("Nuclei",
        exec.Command("nuclei", "-l", filepath.Join(baseOutput, "alive", "alive.txt"),
            "-o", filepath.Join(baseOutput, "vulnerabilities", "vulns.txt")))

    fmt.Println("✅ Reconhecimento concluído!")
}
