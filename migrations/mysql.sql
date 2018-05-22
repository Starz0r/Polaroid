CREATE TABLE `users` (
    `id` INT(8) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NOT NULL,
    `email` VARCHAR(96) NOT NULL,
    `keyword` BLOB NOT NULL,
    `suspended` TINYINT(1) NOT NULL DEFAULT '0',
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=12
;