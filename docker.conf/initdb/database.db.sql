
-- Ensure the user exists before granting privileges
CREATE USER IF NOT EXISTS 'appuser'@'%' IDENTIFIED BY 'letmein';

-- Give all privileges to appuser
GRANT ALL PRIVILEGES ON *.* TO 'appuser'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;

-- Create a permissions table for users
CREATE TABLE IF NOT EXISTS `permissions` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `permission` VARCHAR(255) NOT NULL,
    PRIMARY KEY(`id`)
);

-- Insert permissions in the table
INSERT INTO `permissions` (`permission`) VALUES
('admin'),
('moderator'),
('user');

-- Create a users table
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `permissionId` INT NOT NULL,
    `username` VARCHAR(255) NOT NULL UNIQUE,
    `password` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `phone` VARCHAR(15) UNIQUE,
    `creationDate` DATETIME NOT NULL,
    `lastConnection` DATETIME NOT NULL,
    `lastIP` VARCHAR(45) NOT NULL,
    `session_id` VARCHAR(512),
    PRIMARY KEY(`id`),
    FOREIGN KEY (`permissionId`) REFERENCES `permissions`(`id`) ON DELETE CASCADE
);

-- Insert a sample user into the users table
INSERT INTO `users` (permissionId, username, password, email, phone, creationDate, lastConnection, lastIP) VALUES 
(1, 'onhlt', '21232f297a57a5a743894a0e4a801fc3', 'onhlt@nihilism.moe', NULL, NOW(), NOW(), '127.0.0.1'),
(2, 'modo', '21232f297a57a5a743894a0e4a801fc3', 'modo@nihilism.moe', "0606060606", NOW(), NOW(), '127.0.0.1'),
(3, 'user', '21232f297a57a5a743894a0e4a801fc3', 'user@nihilism.moe', "0707070707", NOW(), NOW(), '127.0.0.1');

-- Create a table to store schema information
CREATE TABLE IF NOT EXISTS `maria_schema` (
    `table_name` VARCHAR(255),
    `table_type` VARCHAR(255),
    `engine` VARCHAR(255),
    `version` INT,
    `row_format` VARCHAR(255),
    `table_rows` BIGINT,
    `avg_row_length` BIGINT,
    `data_length` BIGINT,
    `max_data_length` BIGINT,
    `index_length` BIGINT,
    `data_free` BIGINT,
    `auto_increment` BIGINT,
    `create_time` DATETIME,
    `update_time` DATETIME,
    `check_time` DATETIME,
    `table_collation` VARCHAR(255),
    `checksum` BIGINT,
    `create_options` VARCHAR(255),
    `table_comment` VARCHAR(255)
);

-- Insert schema information from information_schema.tables
INSERT INTO `maria_schema` (table_name, table_type, engine, version, row_format, table_rows, avg_row_length, data_length, max_data_length, index_length, data_free, auto_increment, create_time, update_time, check_time, table_collation, checksum, create_options, table_comment)
SELECT 
    table_name,
    table_type,
    engine,
    version,
    row_format,
    table_rows,
    avg_row_length,
    data_length,
    max_data_length,
    index_length,
    data_free,
    auto_increment,
    create_time,
    update_time,
    check_time,
    table_collation,
    checksum,
    create_options,
    table_comment
FROM information_schema.tables
WHERE table_schema = 'appdb';