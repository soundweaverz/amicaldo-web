CREATE TABLE IF NOT EXISTS users (
    PRIMARY KEY (id),
    id         INT UNSIGNED NOT NULL AUTO_INCREMENT,
    first_name TINYTEXT     NOT NULL,
    last_name  TINYTEXT     NOT NULL
)
