CREATE DATABASE `steradian`;
USE `steradian`;
CREATE TABLE `article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(250) NOT NULL DEFAULT '',
  `body` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `users` (
  `userId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(250) NOT NULL,
  `phoneNumber` varchar(250),
  `city` varchar(250),
  `zip` varchar(250),
  `message` varchar(250),
  `password` varchar(250) NOT NULL,
  `username` varchar(250) NOT NULL,
  `address` varchar(250),
  PRIMARY KEY (`userId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `orders` (
  `orderId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `pickUpLoc` varchar(250) NOT NULL,
  `dropOffLoc` varchar(250),
  `pickupDate` text,
  `dropOffDate` text,
  `pickUpTime` text,
  `carId` int(11),
  `userId` int(11),
  `adminId` int(11),
  PRIMARY KEY (`orderId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(250) NOT NULL,
  `password` varchar(250) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
CREATE TABLE `cars` (
  `carId` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `carType` varchar(250),
  `rating` varchar(250),
  `fuel` varchar(250),
  `image` text,
  `hourRate` varchar(250),
  `dayRate` varchar(250),
  `monthRate` varchar(250),
  PRIMARY KEY (`carId`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;