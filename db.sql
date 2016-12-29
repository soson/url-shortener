-- Adminer 4.2.4 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `url`;
CREATE TABLE `url` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `short_url` varchar(32) NOT NULL,
  `long_url` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `short` (`short_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 2016-12-29 11:46:41
