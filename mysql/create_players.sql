DROP TABLE IF EXISTS `players`;
CREATE TABLE `players` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `balance` float(10,2) unsigned DEFAULT '0.0' NOT NULL ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;