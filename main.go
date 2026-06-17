package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

//r - Pega dados do usuario
//w - Entrega dados ao usuario

func pCEP(w http.ResponseWriter, r *http.Request) {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr) //Pega apenas o ip do usuario
	
	if !VLim(r) { //Se VLim for falso
		lBlock(r, ip) //Chama o log de bloqueio (loggerR.go)
		http.Error(w, "CÓDIGO 429: Muitas requisições")//Imprime o codigo de muitas requisicoes
		return
	}
	
	cep := r.URL.Query().Get("cep") //Pega o cep da url e armazena em uma var
	cep = strings.ReplaceAll(cep, "\r", "")
	cep = strings.ReplaceAll(cep, "\n", "") 
	cep = strings.TrimSpace(cep) //Deleta espacos vazios
	
	ras := regexp.MustCompile("[^0-9]") 
	cep = ras.ReplaceAllString(cep, "") //Tira tudo que nn for int de 0 a 9
	
	if !VCep(cep) { //Verifica o padrão do cep
		cF(r, cep) //Chama o padrao de cep falso (logger)
		http.Error(w, "CEP INVALIDO", http.StatusBadRequest)
		return
	} else {
		log.Printf("CEP [%s] enviado ao Java\n", cep)
	
		java := "http://localhost:9090/buscar?cep=" + cep //Cria uma var na api do java (9090)
		rJava, fx := http.Get(java) //Realiza uma requisição ao java, que armazema a resposta (rJava) e, se ocorrer, um erro (fx)
		
		if fx != nil {
			http.Error(w, "Não foi possivel se conectar ao Java", http.StatusInternalServerError)
			return
		} 
			defer rJava.Body.Close() //Desliga a conexao com java no final da funcao
		
		rJson, fx := io.ReadAll(rJava.Body) //Puxa todos os dados dentro da conexao rJava.Body, enviando para rJson
		if fx != nil {
			http.Error(w, "Não foi possivel ler o conteudo da conexão java")
			return
		} 
		w.Header().Set("Content-Type", "application/json") //Avisa ao navegador que tá chegando um json
		w.Write(rJson) //Imprime o json
	}	
}

func main() {
	http.HandleFunc("/cepR", pCEP)

	log.Println("Server Go Ligado!")

	server := &http.Server{
		Addr: "0.0.0.0:7070",
	}

	bB := server.ListenAndServe()

	log.Fatal(bB)
}
