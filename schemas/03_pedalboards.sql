CREATE DATABASE pedalboard_db;
USE pedalboard_db;

SET FOREIGN_KEY_CHECKS=0;

CREATE TABLE `pedalboards` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `brand` varchar(50) NOT NULL DEFAULT '',
  `name` varchar(50) NOT NULL DEFAULT '',
  `width` decimal(5,2) NOT NULL DEFAULT 0.0,
  `height` decimal(5,2) NOT NULL DEFAULT 0.0,
  `image` varchar(1024) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS=1;
