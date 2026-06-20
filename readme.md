


Conta:

login:teste@teste.com
senha:teste123

Caso não exista, crie uma:
role: PROFISSIONAL
login: email
nome: nome
password: senha



-----
Para inicializar:

docker compose -f docker-compose-doc.yml up

Entre no endpoint do swagger:
localhost/api/v1/auth-service/

1. Faça login (endpoint /login)
2. Copie o refresh_token

-----

Para cada usuario no script abaixo, crie um refresh_token (pode ser com a mesma conta; logue 2x, vai funcionar)



Substitua por tokens válidos:


USER1_TOKEN="ey..."
USER2_TOKEN="eyJhb..."



bash stress_test.sh : simula dois usuarios que enviam 2 pdf's cada. Cada mensagem processa dois pdf's.

-----

DocumentoProcessorWorker

- Declare o número de workers na variavel:

documentocase/cmd/documento-processor-worker/main.go
consumer :=
		consumers.NewDocumentoProcessorConsumer(
			rabbit.ConsumerChannel,
			"documento.processor.queue",
			dispatcher,
			2, // <-Workers
		)


Pode atualizar o numero com o docker aberto, o .air atualiza o build.

