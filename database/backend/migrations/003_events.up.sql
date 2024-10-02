CREATE TYPE cryptogate.events_status AS ENUM (
    'new',
    'inProcessing',
    'success',
    'failed',
    'unknown'
);

CREATE TABLE cryptogate.events (
    id                  UUID                     NOT NULL PRIMARY KEY,
    log_index           NUMERIC                  NOT NULL,
    amount              TEXT,
    height              NUMERIC,
    status              cryptogate.events_status NOT NULL DEFAULT 'new'::cryptogate.events_status,
    "from"              TEXT,
    smart_contract      TEXT,
    smart_contract_type TEXT,
    type                TEXT NOT NULL,
    hash                TEXT NOT NULL,
    coin_slug           TEXT,
    gas_used            TEXT,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL
);