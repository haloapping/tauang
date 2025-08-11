-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    phone TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE wallets(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,
    user_id TEXT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    description TEXT DEFAULT NULL,
    currency TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE categories(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE transactions(
    id TEXT NOT NULL UNIQUE PRIMARY KEY,
    user_id TEXT NOT NULL,
    wallet_id TEXT NOT NULL,
    category_id TEXT NOT NULL,
    name TEXT NOT NULL UNIQUE,
    note TEXT DEFAULT NULL,
    amount BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (wallet_id) REFERENCES wallets(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
