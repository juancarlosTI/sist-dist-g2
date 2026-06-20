CREATE TABLE evento_store (
    evento_id UUID NOT NULL,
    agregado_id TEXT NOT NULL,
    agregado_tipo TEXT NOT NULL,
    evento_versao INT NOT NULL,
    evento_nome TEXT NOT NULL,

    correlacao_id UUID,
    causalidade_id UUID,

    payload JSONB NOT NULL,

    autor_tipo TEXT,
    autor_id TEXT,

    origem_canal TEXT,
    origem_sistema TEXT,
    role_tipo TEXT,
    
    ocorreu_as TIMESTAMPTZ NOT NULL,
    criado_as TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    PRIMARY KEY (evento_id)
);

-- Garantia de ordenação de eventos por agregado
CREATE UNIQUE INDEX idx_event_store_agregado_versao
ON evento_store (agregado_id, evento_versao);

-- Query principal do event sourcing
CREATE INDEX idx_event_store_agregado
ON evento_store (agregado_id);

-- Para debugging / auditoria
CREATE INDEX idx_event_store_evento_nome
ON evento_store (evento_nome);