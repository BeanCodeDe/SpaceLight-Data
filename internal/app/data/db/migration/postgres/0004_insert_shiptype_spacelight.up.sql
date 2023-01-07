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
        `[{"Id":0,"RoomBlockList": [{"Id":0,"PosX":3,"PosY":0},{"Id":1,"PosX":4,"PosY":0},{"Id":2,"PosX":3,"PosY":1},{"Id":3,"PosX":4,"PosY":1}]},{"Id":1,"RoomBlockList": [{"Id":4,"PosX":5,"PosY":0},{"Id":5,"PosX":6,"PosY":0},{"Id":6,"PosX":5,"PosY":1},{"Id":7,"PosX":6,"PosY":1}]},{"Id":2,"RoomBlockList": [{"Id":8,"PosX":7,"PosY":0},{"Id":9,"PosX":8,"PosY":0},{"Id":10,"PosX":7,"PosY":1},{"Id":11,"PosX":8,"PosY":1},{"Id":12,"PosX":9,"PosY":1}]},{"Id":3,"RoomBlockList": [{"Id":13,"PosX":9,"PosY":0},{"Id":14,"PosX":9,"PosY":1}]},{"Id":4,"RoomBlockList": [{"Id":15,"PosX":1,"PosY":1},{"Id":16,"PosX":2,"PosY":1}]},{"Id":5,"RoomBlockList": [{"Id":17,"PosX":11,"PosY":1},{"Id":18,"PosX":12,"PosY":1}]},{"Id":6,"RoomBlockList": [{"Id":19,"PosX":13,"PosY":1},{"Id":20,"PosX":14,"PosY":1}]},{"Id":7,"RoomBlockList": [{"Id":21,"PosX":0,"PosY":2},{"Id":22,"PosX":1,"PosY":2}]},{"Id":8,"RoomBlockList": [{"Id":23,"PosX":8,"PosY":2},{"Id":24,"PosX":9,"PosY":2}]},{"Id":9,"RoomBlockList": [{"Id":25,"PosX":10,"PosY":2},{"Id":26,"PosX":11,"PosY":2}]},{"Id":10,"RoomBlockList": [{"Id":27,"PosX":14,"PosY":2},{"Id":28,"PosX":15,"PosY":2}]},{"Id":11,"RoomBlockList": [{"Id":29,"PosX":1,"PosY":3},{"Id":30,"PosX":2,"PosY":3}]},{"Id":12,"RoomBlockList": [{"Id":31,"PosX":3,"PosY":3},{"Id":32,"PosX":4,"PosY":3},{"Id":33,"PosX":3,"PosY":4},{"Id":34,"PosX":4,"PosY":4}]},{"Id":13,"RoomBlockList": [{"Id":35,"PosX":5,"PosY":3},{"Id":36,"PosX":6,"PosY":3},{"Id":37,"PosX":5,"PosY":4},{"Id":38,"PosX":6,"PosY":4}]},{"Id":14,"RoomBlockList": [{"Id":39,"PosX":7,"PosY":3},{"Id":40,"PosX":8,"PosY":3},{"Id":41,"PosX":7,"PosY":4},{"Id":42,"PosX":8,"PosY":4}]},{"Id":15,"RoomBlockList": [{"Id":43,"PosX":9,"PosY":3},{"Id":44,"PosX":9,"PosY":4}]},{"Id":16,"RoomBlockList": [{"Id":45,"PosX":11,"PosY":3},{"Id":46,"PosX":12,"PosY":3}]},{"Id":17,"RoomBlockList": [{"Id":47,"PosX":13,"PosY":3},{"Id":48,"PosX":14,"PosY":3}]}]`
        ``
    );


CREATE TABLE spacelight_data.profil (
    id uuid PRIMARY KEY NOT NULL,
    name varchar NOT NULL
    rome_place_list json NOT NULL
    door_list json NOT NULL
    weapon_place_list json NOT NULL
);