INSERT INTO universes VALUES( uuid_generate_v4() );

INSERT INTO galaxies(galaxy_id, universe_id) VALUES(uuid_generate_v4(), (select universe_id FROM universes)::uuid );

INSERT INTO systems(system_id, galaxy_id) VALUES(uuid_generate_v4(), (select galaxy_id FROM galaxies)::uuid );

INSERT INTO users(user_id, universe_id, user_name) VALUES(uuid_generate_v4(), (select universe_id FROM universes)::uuid, 'first' );

INSERT INTO planets(PLANET_ID, SYSTEM_ID, USER_ID) VALUES(uuid_generate_v4(), (select system_id FROM systems)::uuid, (select user_id FROM users)::uuid);

INSERT INTO factories(factory_id, planet_id, resource_type, building_level, base_performance, base_update_cost)
 VALUES(uuid_generate_v4(), (select planet_id FROM planets)::uuid, 'BTC', 0, 1, 1);

INSERT INTO storages(storage_id, planet_id, resource_type, building_level, base_capacity, current_capacity, base_update_cost)
VALUES(uuid_generate_v4(), (select planet_id FROM planets)::uuid, 'BTC', 0, 10, 1, 1);