package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

var (
	cep string = "01153000"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go CallBrasilApi(c1)
	go CallViacepApi(c2)

	select {
	case res := <-c1:
		println("received:", res)

	case res := <-c2:
		println("received:", res)

	case <-time.After(time.Second):
		println("timeout")
	}
}

func CallBrasilApi(data chan<- string) {
	res, err := CallApi("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao acessar brasilapi: %v\n", err)
	}
	data <- "Dados do endereço: " + res + "\n API envio: brasilapi"
}

func CallViacepApi(data chan<- string) {
	res, err := CallApi("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao acessar viacep: %v\n", err)
	}
	data <- "Dados do endereço: " + res + "\n API envio: viacep"
}

func CallApi(url string) (string, error) {
	req, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %w\n", err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
	}
	return string(res), nil
}
