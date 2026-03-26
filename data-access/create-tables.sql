DROP TABLE IF EXISTS event;
CREATE TABLE event (
    id      INT AUTO_INCREMENT NOT NULL,
    name    VARCHAR(128) NOT NULL,
    age     INT NOT NULL,
    PRIMARY KEY (`id`)
);

INSERT INTO event
    (id, name, age)
VALUES
    (1, 'Yunara', 28),
    (2, 'Thimmy', 23);