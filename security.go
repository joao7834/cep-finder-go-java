package main

import (
	"net"
	"net/http"
	"regexp"
)

var ac = make(map[string]int) //Guarda o ip do usuario e a quantidade de vezes que ele fez requisições

func VLim(r *http.Request) bool {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr) //Pega apenas o ip do usuario

	ac[ip] = ac[ip] + 1

	if ac[ip] > 20 {
		return false
	}
	return true
}

func VCep(cep string) bool {
	// Expressão Regular (Regex):
	// ^       -> Garante que o padrao começa com o .
	// [0-9]   -> Aceita apenas int de 0 a 9.
	// {8}     -> Exige que existam apenas 8 caracteres.
	// $       -> Garante que o texto TERMINA ali.
	p := `^[0-9]{8}$`

	v, _ := regexp.MatchString(p, cep)

	return v
}
