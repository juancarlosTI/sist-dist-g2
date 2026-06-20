#!/bin/bash

URL="https://localhost/api/v1/documento/documento/add"

USER1_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMTllY2JlMS01ZDJiLTc3MWItODgxOC0xZDkyZTk0NjI0NTQiLCJyb2xlcyI6IlBST0ZJU1NJT05BTCIsImF1dG9yIjp7ImlkIjoiMDE5ZWNiZTEtNWQyYi03NzFiLTg4MTgtMWQ5MmU5NDYyNDU0IiwidGlwbyI6IiJ9LCJvcmlnZW0iOnsiY2FuYWwiOiIiLCJzaXN0ZW1hIjoiIn0sImV4cCI6MTc4MTkwMzQxN30.68F4a7iwfKwSINvNwTlRT02eKTy17MiqXS15cPy8EjU"
USER2_TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIwMTllY2JlMS01ZDJiLTc3MWItODgxOC0xZDkyZTk0NjI0NTQiLCJyb2xlcyI6IlBST0ZJU1NJT05BTCIsImF1dG9yIjp7ImlkIjoiMDE5ZWNiZTEtNWQyYi03NzFiLTg4MTgtMWQ5MmU5NDYyNDU0IiwidGlwbyI6IiJ9LCJvcmlnZW0iOnsiY2FuYWwiOiIiLCJzaXN0ZW1hIjoiIn0sImV4cCI6MTc4MTkwMzQ1Mn0.eoMaxls1MFfYYcvwpLGtgqpqZuLSYGsL41omv3BZKG0"

FILE1="Building Microservices - Designing Fine-Grained Systems.pdf"
FILE2="Domain-Driven Design Distilled PDF.pdf"

request () {
  TOKEN=$1
  USER=$2

  echo "[$USER] enviando request..."

  curl -k -s -o /dev/null -w "[$USER] HTTP:%{http_code} TIME:%{time_total}\n" \
    -X POST "$URL" \
    -H "Authorization: Bearer $TOKEN" \
    -F "arquivos=@$FILE1" \
    -F "arquivos=@$FILE2"
}

echo "Iniciando teste paralelo..."

for i in {1..5}
do
  echo "================ ITERAÇÃO $i ================"

  request "$USER1_TOKEN" "USER1" &
  request "$USER2_TOKEN" "USER2" &

  wait
done
