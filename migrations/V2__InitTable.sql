CREATE TABLE IF NOT EXISTS universes
(
    id UUID PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS galaxies
(
    id          UUID PRIMARY KEY,
    universe_id UUID REFERENCES universes (id)
);

CREATE TABLE IF NOT EXISTS systems
(
    id        UUID PRIMARY KEY,
    galaxy_id UUID REFERENCES galaxies (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id          UUID PRIMARY KEY,
    universe_id UUID,
    user_name   VARCHAR(30)
);

CREATE TABLE IF NOT EXISTS users
(
    id       UUID PRIMARY KEY REFERENCES users (id),
    password UUID,
    login    VARCHAR(30)
);


CREATE TABLE IF NOT EXISTS planets
(
    id        UUID PRIMARY KEY,
    system_id UUID REFERENCES systems (id),
    user_id   UUID REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS storages
(
    id               UUID PRIMARY KEY,
    planet_id        UUID REFERENCES planets (id),
    resource_type    VARCHAR(30),
    building_level   NUMERIC,
    base_capacity    NUMERIC,
    current_capacity NUMERIC,
    base_update_cost NUMERIC
);

CREATE TABLE IF NOT EXISTS factories
(
    id               UUID PRIMARY KEY,
    planet_id        UUID REFERENCES planets (id),
    resource_type    VARCHAR(30),
    building_level   NUMERIC,
    base_performance NUMERIC,
    base_update_cost NUMERIC
);
