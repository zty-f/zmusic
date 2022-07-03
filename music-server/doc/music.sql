-- 管理员表

DROP TABLE IF EXISTS `admins`;
CREATE TABLE `admins`
(
    `id`       int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name`     varchar(45)      NOT NULL,
    `password` varchar(45)      NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name_UNIQUE` (`name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8;

INSERT INTO `admins`
VALUES (1, 'zty', '123456'),
       (2, 'fmj', '123456');

# 用户表
DROP TABLE IF EXISTS `consumers`;
CREATE TABLE `consumers`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username`     varchar(255)     NOT NULL,
    `password`     varchar(100)     NOT NULL,
    `sex`          tinyint(4)                DEFAULT NULL, # 0 女 1 男 2 未知
    `phone_num`    char(15)                  DEFAULT NULL,
    `email`        char(30)                  DEFAULT NULL,
    `birth`        datetime                  DEFAULT NULL,
    `introduction` varchar(255)              DEFAULT NULL,
    `location`     varchar(45)               DEFAULT NULL,
    `avatar`       varchar(255)              DEFAULT NULL,
    `create_time`  datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    `update_time`  datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_UNIQUE` (`username`),
    UNIQUE KEY `phone_num_UNIQUE` (`phone_num`),
    UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 29
  DEFAULT CHARSET = utf8;
