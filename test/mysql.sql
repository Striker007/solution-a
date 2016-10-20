CREATE DATABASE IF NOT EXISTS `test`;

CREATE TABLE `test`.`test`
(
id int NOT NULL AUTO_INCREMENT,
LastName varchar(255),
FirstName varchar(255),
gender ENUM('Man', 'Woman'),
PRIMARY KEY (id)
);

INSERT INTO `test`.`test` (`LastName`, `FirstName`, `gender`) VALUES ('Ivan', 'Ivanov', 'Man');
