CREATE EXTENSION pgcrypto;

CREATE SCHEMA spacelight

CREATE TABLE spacelight.user (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR ( 50 ) UNIQUE NOT NULL,
    created_on TIMESTAMP NOT NULL,
    last_login TIMESTAMP NOT NULL
);

GRANT ALL PRIVILEGES ON DATABASE spacelight TO spacelight;