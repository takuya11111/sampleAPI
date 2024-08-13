CREATE DATABASE IF NOT EXISTS simpleapi;

DROP TABLE IF EXISTS `todo_items`;
CREATE TABLE `todo_items` (
  `id` varchar(50) NOT NULL DEFAULT '',
  `todo` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE DATABASE IF NOT EXISTS simpleapi;


