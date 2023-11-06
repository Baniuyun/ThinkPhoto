SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS`think_photo` DEFAULT CHARACTER SET utf8mb4;
USE `think_photo`;

-- ----------------------------
-- Table structure for t_follow
-- ----------------------------
DROP TABLE IF EXISTS `t_follow`;
CREATE TABLE `t_follow` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '关注关系id，自增主键',
    `follower_id` bigint(20) UNSIGNED NOT NULL COMMENT '粉丝id',
    `following_id` bigint(20) UNSIGNED NOT NULL COMMENT '博主id',
    `is_follow` tinyint(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否关注：0未关注，1已关注',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `relation_follower_following` (`follower_id`,`following_id`) USING BTREE,
    KEY `update_time` (`update_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1177 DEFAULT CHARSET=utf8 COMMENT='关注表';

-- ----------------------------
-- Table structure for t_follow_count
-- ----------------------------
DROP TABLE IF EXISTS `t_follow_count`;
CREATE TABLE `t_follow_count` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '关注计数id，自增主键',
    `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
    `following_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '关注数',
    `follower_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '粉丝数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_id` (`user_id`) USING BTREE,
    KEY `update_time` (`update_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT = 1010 DEFAULT CHARSET=utf8 COMMENT '关注计数表';