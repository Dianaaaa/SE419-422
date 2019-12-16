DROP TABLE IF EXISTS `urls`;
CREATE TABLE `urls` (
    shortlink VARCHAR(8) NOT NULL PRIMARY KEY,
    expiration_length_in_minutes INT NOT NULL,
    created_at DATETIME NOT NULL,
    paste_path VARCHAR(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;