package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    // Nome do arquivo a ser lido
    filename := "/etc/passwd"

    // Lê o conteúdo do arquivo
    content, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatalf("Erro ao ler o arquivo %s: %v", filename, err)
    }

    // Converte o conteúdo para string e imprime
    fmt.Println(string(content))
}
