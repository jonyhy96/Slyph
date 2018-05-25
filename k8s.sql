CREATE TABLE `user` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=UTF-8;
CREATE TABLE `dep` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `dep` varchar(50) NOT NULL,
  `image` varchar(50) NOT NULL,
  `ver` varchar(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=UTF-8;
CREATE TABLE `logs`(
  `time` varchar(20) NOT NULL,
  `etype` varchar(10) NOT NULL,
  `event` varchar(50) NOT NULL,
  PRIMARY KEY (`time`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=UTF-8;