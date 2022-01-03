INSERT INTO accounts(login, password) VALUES ('admin','admin');

INSERT INTO universes(name) VALUES( 'first_universe' );

INSERT INTO users(id, account_login, universe_name, username)
VALUES(uuid_generate_v4(), (select login FROM accounts), (select name FROM universes), 'first_user' );

INSERT INTO galaxies(id, universe_name) VALUES(uuid_generate_v4(), (select name FROM universes));

INSERT INTO systems(id, galaxy_id) VALUES(uuid_generate_v4(), (select id FROM galaxies)::uuid );

INSERT INTO planets(id, system_id, user_id, name) VALUES(uuid_generate_v4(), (select id FROM systems)::uuid, (select id FROM users)::uuid, 'first_planet');

INSERT INTO factories(id, planet_id, resource_type, building_level, base_performance, base_update_cost, base_update_time)
 VALUES(uuid_generate_v4(), (select id FROM planets)::uuid, 'BTC', 0, 1, 1, 1);

INSERT INTO storages(id, planet_id, resource_type, building_level, base_capacity, current_capacity, base_update_cost, base_update_time)
VALUES(uuid_generate_v4(), (select id FROM planets)::uuid, 'BTC', 0, 10, 1, 1, 1);