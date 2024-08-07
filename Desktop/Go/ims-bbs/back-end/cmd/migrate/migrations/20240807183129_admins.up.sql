CREATE TABLE IF NOT EXISTS admins (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `firstName` VARCHAR(255) NOT NULL,
    `lastName` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `password` VARCHAR(255) NOT NULL,
	`contact`     VARCHAR(255) NOT NULL,
	`address`     TEXT NOT NULL,
	`age`         INT NOT NULL,
	`joiningDate` VARCHAR(255) NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY(id),
    UNIQUE(email)
)