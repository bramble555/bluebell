CREATE DATABASE IF NOT EXISTS bluebell;
use bluebell;
CREATE TABLE  IF NOT exists `user`(
	`id` BIGINT ( 20 ) NOT NULL auto_increment,-- 防止用户知道我这个项目有多少人注册了
	`user_id` BIGINT ( 20 ) NOT NULL,
	`username` VARCHAR ( 64 ) NOT NULL,
	`password` VARCHAR ( 64 ) NOT NULL,
	`email` VARCHAR ( 64 ),
	`gender` TINYINT ( 1 ) NOT NULL DEFAULT 0, -- 0表示女生
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY ( `id` ),
	UNIQUE KEY `idx_username` ( `username` ) USING BTREE,
	UNIQUE KEY `idx_user_id` ( `user_id` ) USING BTREE
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

CREATE TABLE IF NOT exists `community`(
	`id` BIGINT ( 20 ) NOT NULL auto_increment,-- 防止用户知道我这个项目有多少人注册了
	`community_id` BIGINT ( 20 ) NOT NULL,
	`community_name` VARCHAR ( 128 ) NOT NULL,
	`introduction` VARCHAR ( 64 ) NOT NULL,
	`create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
	`update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY ( `id` ),
	UNIQUE KEY `idx_community_id`(`community_id`),
	UNIQUE KEY `idx_community_name`(`community_name`)
)ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;

INSERT INTO `community` (`community_id`,community_name,introduction)
VALUES (1,'Go','Golang');
INSERT INTO `community` (`community_id`,community_name,introduction)
VALUES (2,'leetcode','要刷题咯');
INSERT INTO `community` (`community_id`,community_name,introduction)
VALUES (3,'LOL','欢迎来到影响联盟');
INSERT INTO `community` (`community_id`,community_name,introduction)
VALUES (4,'CF','A小有人');

CREATE TABLE IF NOT exists `post` (
	`id` BIGINT ( 20 ) NOT NULL auto_increment,
	`post_id` BIGINT(20) not null,
	`title` VARCHAR(128) not null,
	`content` TEXT not null,
	`user_id` BIGINT(20) not null,
	`community_id` BIGINT(20) not null,
	`status` TINYINT(1) not null DEFAULT 1 ,
	`create_time` TIMESTAMP null DEFAULT CURRENT_TIMESTAMP,
	`update_time` TIMESTAMP null DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `idx_post_id`(`post_id`),
	KEY `idx_user_id` (`user_id`)
)ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci;
