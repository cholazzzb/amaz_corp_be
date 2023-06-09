-- +migrate Up
INSERT INTO buildings(name)
VALUES ("Ice Office");
INSERT INTO buildings(name)
VALUES ("Flame Office");

INSERT INTO rooms(name, building_id)
VALUES ("Ice Cream", 1);
INSERT INTO rooms(name, building_id)
VALUES ("Ice Lemon", 1);
INSERT INTO rooms(name, building_id)
VALUES ("Ice Tea", 1);
INSERT INTO rooms(name, building_id)
VALUES ("Ice Coffee", 1);
INSERT INTO rooms(name, building_id)
VALUES ("Burn Meat", 2);
INSERT INTO rooms(name, building_id)
VALUES ("Hot Tea", 2);
