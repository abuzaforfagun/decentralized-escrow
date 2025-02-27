CREATE TABLE products (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
	seller TEXT NOT NULL,
    description TEXT,
    price TEXT NOT NULL,
    stock TEXT NOT NULL,
    transaction_hash TEXT NOT NULL,
    on_chain_id TEXT NOT NULL,
    status INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    creation_completed_at TIMESTAMPTZ
);
