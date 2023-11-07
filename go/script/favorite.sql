SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS`think_photo` DEFAULT CHARACTER SET utf8mb4;
USE `think_photo`;

CREATE TABLE `like_record` (
                               `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `linking_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞对象ID',
                               `liker_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
                               `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                               `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB  COMMENT='点赞记录表';

CREATE TABLE `like_count` (
                              `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                              `linking_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞对象id',
                              `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
                              `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                              `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB  COMMENT='点赞计数表';