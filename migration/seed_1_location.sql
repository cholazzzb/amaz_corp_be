-- +migrate Up
INSERT INTO buildings(id, name)
VALUES ("bc133e57-df08-407e-b1e5-8e10c653ad3c", "Ice Office");
INSERT INTO buildings(id, name)
VALUES ("3806c2c4-a881-4242-b84a-dc04611cfc39", "Flame Office");

INSERT INTO rooms(id, name, building_id)
VALUES (UUID(), "Ice Cream", "bc133e57-df08-407e-b1e5-8e10c653ad3c");
INSERT INTO rooms(id, name, building_id)
VALUES (UUID(), "Ice Lemon", "bc133e57-df08-407e-b1e5-8e10c653ad3c");
INSERT INTO rooms(id, name, building_id)
VALUES (UUID(), "Ice Tea", "bc133e57-df08-407e-b1e5-8e10c653ad3c");
INSERT INTO rooms(id, name, building_id)
VALUES (UUID(), "Ice Coffee", "bc133e57-df08-407e-b1e5-8e10c653ad3c");
INSERT INTO rooms(id, name, building_id)
VALUES (UUID(),"Burn Meat", "3806c2c4-a881-4242-b84a-dc04611cfc39");
INSERT INTO rooms(id, name, building_id)
VALUES (UUID(),"Hot Tea", "3806c2c4-a881-4242-b84a-dc04611cfc39");
