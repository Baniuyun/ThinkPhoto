SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE IF NOT EXISTS`think_photo` DEFAULT CHARACTER SET utf8mb4;
USE `think_photo`;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `t_video`;
CREATE TABLE `t_video`
(
    `id`           bigint(20) NOT NULL AUTO_INCREMENT COMMENT '视频id，自增主键',
    `author_id`    bigint(20) NOT NULL COMMENT '视频作者id',
    `author_name`  varchar(255) NOT NULL COMMENT '作者名字',
    `play_url`     varchar(255) NOT NULL COMMENT '播放url',
    `cover_url`    varchar(255)  COMMENT '封面url',
    `avatar`       varchar(255) NOT NULL COMMENT '头像地址',
    `favorite_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '视频的点赞数量',
    `comment_count` bigint(20)  DEFAULT 0 COMMENT '视频的评论数量',
--   `is_favorite` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否点赞',
--   `is_follow` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否关注',
--    `title`        varchar(255) DEFAULT NULL COMMENT '视频标题',
    `information` varchar(255) NOT NULL COMMENT '视频简介',
    `tag` bigint(20) NOT NULL COMMENT '视频标签',
    `time` bigint(20) NOT NULL COMMENT '时间戳',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=115  COMMENT='视频表';