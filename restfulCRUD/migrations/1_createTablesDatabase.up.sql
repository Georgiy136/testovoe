CREATE SCHEMA IF NOT EXISTS currency;

/* Цены койнов в момент времени */
CREATE TABLE IF NOT EXISTS currency.watch_list (
    coin_id INTEGER NOT NULL,
    price FLOAT NOT NULL,
    currency CHAR(3) NOT NULL, /* Валюта */
    ch_dt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_currency_watch_list_coin_id_ch_dt
    ON currency.watch_list (coin_id, ch_dt);

/* Список койнов */
CREATE TABLE IF NOT EXISTS currency.coins (
    ID SERIAL PRIMARY KEY,
    coin_name CHAR(50) NOT NULL,
    added_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_currency_coins_coin_name
    ON currency.coins (coin_name);

CREATE INDEX IF NOT EXISTS idx_currency_coins_id
    ON currency.coins (id);
