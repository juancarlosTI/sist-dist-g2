CREATE TABLE snapshot_documento (
    id TEXT PRIMARY KEY,

    estado INT,

    origem_canal TEXT,
    origem_sistema TEXT,

    autor_tipo TEXT,
    autor_id TEXT,

    versao INT NOT NULL,

    processos_ids JSONB,

    criado_as TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_snapshot_documento_versao
ON snapshot_documento (versao);