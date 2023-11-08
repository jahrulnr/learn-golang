CREATE TABLE 
    `user`(
        `id` BIGINT(11) UNSIGNED NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(255) NOT NULL,
        `username` VARCHAR(30) NOT NULL,
        `password` VARCHAR(128) NOT NULL,
        `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NULL DEFAULT NULL, 
        PRIMARY KEY(`id`), 
        INDEX `username`(`username`)
    ) 
ENGINE = InnoDB;

INSERT INTO `user` (`id`,
    `name`,
    `username`,
    `password`,
    `created_at`,
    `updated_at`) VALUES (NULL, 'Yang disana', 'Jahrulnr', 'disini', current_timestamp(), NULL);