/*
DROP DATABASE snai_test
;
*/
CREATE DATABASE snai_test CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci
;

USE snai_test
;
CREATE TABLE `user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `number` varchar(255) NOT NULL DEFAULT '' COMMENT '学号',
    `name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `gender` char(5)  NOT NULL COMMENT '男｜女｜未公开',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `number_unique` (`number`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;

INSERT INTO `user` (number,name,password,gender)values ('666','小明','123456','男');