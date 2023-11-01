SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS`think_photo` DEFAULT CHARACTER SET utf8mb4;
USE `think_photo`;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
    `username` varchar(32) NOT NULL COMMENT '账号',
    `password` varchar(128) NOT NULL COMMENT '密码，加密存储',
    `salt` varchar(255)  NOT NULL COMMENT 'uuid，密码MD5加密',
    `nick_name` varchar(32) NOT NULL COMMENT '昵称',
    `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '头像url',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `username` (`username`) USING BTREE,
    UNIQUE KEY `nick_name` (`nick_name`) USING BTREE,
) ENGINE = InnoDB AUTO_INCREMENT = 1010 DEFAULT CHARSET=utf8 COMMENT='用户表';
