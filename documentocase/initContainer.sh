#!/usr/bin/env sh
set -e

echo "🔧 [init] Iniciando container para serviço: ${SERVICE_TYPE}"

#
# Aguarda dependências externas (RabbitMQ, DB, etc.)
#
wait_for_host() {
    HOST=$1
    PORT=$2
    NAME=$3

    echo "⏳ [init] Aguardando ${NAME} em ${HOST}:${PORT}..."

    while ! nc -z "$HOST" "$PORT"; do
        sleep 1
    done

    echo "✅ [init] ${NAME} disponível!"
}

# ---- DEPENDÊNCIAS COMUNS ----
if [ ! -z "$RABBIT_HOST" ]; then
    wait_for_host "$RABBIT_HOST" "$RABBIT_PORT" "RabbitMQ"
fi

if [ ! -z "$POSTGRES_HOST" ]; then
    wait_for_host "$POSTGRES_HOST" "$POSTGRES_PORT" "PostgreSQL"
fi

echo "🔧 [init] Dependências OK!"

#
# Executar o serviço correto
#
case "$SERVICE_TYPE" in
    api)
        echo "🚀 [init] Iniciando API com Air..."
        exec air -c .air.api.toml
        ;;
    worker)
        echo "🚀 [init] Iniciando Worker com Air..."
        exec air -c .air.worker.toml
        ;;
    prod-api)
        echo "🚀 [init] Iniciando API (produção)..."
        exec ./tmp/app
        ;;
    prod-worker)
        echo "🚀 [init] Iniciando Worker (produção)..."
        exec ./tmp/worker
        ;;
    *)
        echo "❌ [init] ERRO: SERVICE_TYPE inválido."
        echo "Valores esperados: api | worker | prod-api | prod-worker"
        exit 1
        ;;
esac