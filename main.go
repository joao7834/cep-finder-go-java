package main

import (
	"fmt"      // Biblioteca pra printar txt
	"log"      //Se algo ser merda
	"net/http" //Criar server web
)

func proCEP(w http.ResponseWriter, r *http.Request) { //Processador de cep. R pede o CEP & W dá o CEP
	fmt.Fprint(w, "Funcao que vai processar o CEP")
}

func main() {
	http.HandleFunc("/cep", proCEP)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
