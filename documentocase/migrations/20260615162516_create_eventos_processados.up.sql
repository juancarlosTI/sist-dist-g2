CREATE TABLE eventos_processados (
    evento_id UUID PRIMARY KEY,
    processado_em TIMESTAMPTZ NOT NULL DEFAULT NOW()
);