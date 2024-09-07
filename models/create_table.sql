CREATE TABLE `user` (
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