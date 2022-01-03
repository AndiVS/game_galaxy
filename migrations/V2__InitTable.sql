CREATE TABLE IF NOT EXISTS universes
(
    name VARCHAR(30) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS galaxies
(
    id            UUID PRIMARY KEY,
    universe_name VARCHAR(30) REFERENCES universes (name)
);

CREATE TABLE IF NOT EXISTS systems
(
    id        UUID PRIMARY KEY,
    galaxy_id UUID REFERENCES galaxies (id)
);

CREATE TABLE IF NOT EXISTS accounts
(
    login    VARCHAR(30) PRIMARY KEY,
    password VARCHAR(200)
);

CREATE TABLE IF NOT EXISTS users
(
    id            UUID PRIMARY KEY,
    account_login VARCHAR(30) REFERENCES accounts (login),
    universe_name VARCHAR(30) REFERENCES universes (name),
    username      VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS planets
(
    id        UUID PRIMARY KEY,
    system_id UUID REFERENCES systems (id),
    user_id   UUID REFERENCES users (id),
    name      VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS storages
(
    id               UUID PRIMARY KEY,
    planet_id        UUID REFERENCES planets (id),
    resource_type    VARCHAR(30),
    building_level   NUMERIC,
    base_capacity    NUMERIC,
    current_capacity NUMERIC,
    base_update_cost NUMERIC,
    base_update_time NUMERIC
);

CREATE TABLE IF NOT EXISTS factories
(
    id               UUID PRIMARY KEY,
    planet_id        UUID REFERENCES planets (id),
    resource_type    VARCHAR(30),
    building_level   NUMERIC,
    base_performance NUMERIC,
    base_update_cost NUMERIC,
    base_update_time NUMERIC
);
