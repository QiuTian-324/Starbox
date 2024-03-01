create database buzzbox_user;
use buzzbox_user;

CREATE TABLE `user` (
                        `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                        `username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
                        `avatar` varchar(256) NOT NULL DEFAULT '' COMMENT '头像',
                        `mobile` varchar(128) NOT NULL DEFAULT '' COMMENT '手机号',
                        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                        PRIMARY KEY (`id`),
                        KEY `ix_update_time` (`update_time`),
                        UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户表';

insert into user(username, avatar, mobile) values ('张三', 'http://avatar.jpg', '13800138000');