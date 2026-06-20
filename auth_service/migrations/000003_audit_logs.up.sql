CREATE TABLE audit_logs (
    id UUID PRIMARY KEY,
    event_type TEXT NOT NULL,
    user_id TEXT,
    correlation_id TEXT,
    metadata JSONB,
    created_at TIMESTAMP DEFAULT NOW()
);