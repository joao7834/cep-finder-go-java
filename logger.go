package main

import (
	"log"
	"net/http"
)

func lBlock(r *http.Request, ip string) { //Recebe a requisição e o ip de quem tentou atingir o server
	log.Printf("ALERTA 429: IP %s FOI BLOQUEADO POR RATE LIMIT! ROTA %s\n", ip, r.URL.Path) //r.URL.Path -> R guarda tudo que veio do usuário (IP); URL cuida do servidor web; Path pega o caminho apenas do server
}

func cF(r *http.Request, cep string) {
	log.Printf("ALERTA 400: CEP INSERIDO INVÁLIDO: '%s' ORIGEM: %s", cep, r.URL.Path)
}
