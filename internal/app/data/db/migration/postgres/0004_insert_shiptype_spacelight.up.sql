INSERT INTO spacelight_data.ship_type 
    (
        id,
        name,
        rome_place_list,
        door_list,
        weapon_place_list
    )
VALUES
    (
        `f457cf8b-7116-4b8b-a077-16ebc2e5a291`,
        `SpaceLight`,
        `[{"PosX":3,"PosY":0,"SizeX":0,"SizeY":0},{"PosX":5,"PosY":0,"SizeX":0,"SizeY":0},{"PosX":7,"PosY":0,"SizeX":0,"SizeY":0},{"PosX":9,"PosY":0,"SizeX":0,"SizeY":0},{"PosX":1,"PosY":1,"SizeX":0,"SizeY":0},{"PosX":11,"PosY":1,"SizeX":0,"SizeY":0},{"PosX":13,"PosY":1,"SizeX":0,"SizeY":0},{"PosX":0,"PosY":2,"SizeX":0,"SizeY":0},{"PosX":8,"PosY":2,"SizeX":0,"SizeY":0},{"PosX":10,"PosY":2,"SizeX":0,"SizeY":0},{"PosX":14,"PosY":2,"SizeX":0,"SizeY":0},{"PosX":1,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":3,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":5,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":7,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":9,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":11,"PosY":3,"SizeX":0,"SizeY":0},{"PosX":13,"PosY":3,"SizeX":0,"SizeY":0}]`
        ``
    );


CREATE TABLE spacelight_data.profil (
    id uuid PRIMARY KEY NOT NULL,
    name varchar NOT NULL
    rome_place_list json NOT NULL
    door_list json NOT NULL
    weapon_place_list json NOT NULL
);