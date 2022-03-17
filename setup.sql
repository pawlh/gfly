CREATE TABLE locations(
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          owner VARCHAR(255),
                          latitude FLOAT,
                          longitude FLOAT,
                          accuracy INT,
                          timestamp DATE
);