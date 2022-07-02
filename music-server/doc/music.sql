-- 管理员表

CREATE TABLE `admin`
(
    `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`     varchar(45) NOT NULL,
    `password` varchar(45) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

INSERT INTO `admin` VALUES (1,'zty','123456'),(2,'fmj','123456');
