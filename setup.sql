DROP TABLE if exists locations;

CREATE TABLE locations
(
    id        INT PRIMARY KEY AUTO_INCREMENT,
    bin       INT not null,
    owner     VARCHAR(255) not null,
    latitude  FLOAT not null,
    longitude FLOAT not null,
    accuracy  INT default 0,
    timestamp DATETIME not null
);