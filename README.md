#Description

Desafio com Multithreading - Go Expert

Requisitar dados de endereço em duas APIs distintas.

Retornar o resultado da API que responder mais rapido.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

    https://brasilapi.com.br/

    http://viacep.com.br/ws/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

##How to run the project

Informar o cep que deseja buscar na variavel cep.
executar no terminal go run main.go