CREATE TABLE outbox (
    id UUID PRIMARY KEY,

    evento_id UUID NOT NULL,
    evento_versao INTEGER NOT NULL,
    evento_nome VARCHAR(150) NOT NULL,

    correlacao_id UUID,
    causalidade_id UUID,

    payload JSONB NOT NULL,
    routing_key VARCHAR(100) NOT NULL,

    autor_tipo VARCHAR(100) NOT NULL,
    autor_id VARCHAR(100) NOT NULL,
    origem_canal VARCHAR(100) NOT NULL,
    origem_sistema VARCHAR(100) NOT NULL,

    ocorreu_as TIMESTAMP WITH TIME ZONE NOT NULL,
    criado_as TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

    publicado_as TIMESTAMP WITH TIME ZONE,
    tentativas INTEGER NOT NULL DEFAULT 0,
    ultimo_erro TEXT
);

CREATE UNIQUE INDEX idx_outbox_evento_id_versao
ON outbox (evento_id, evento_versao);

CREATE INDEX idx_outbox_publicado_as
ON outbox (publicado_as);