CREATE DATABASE IF NOT EXISTS `htz_sutra` DEFAULT CHARACTER SET utf8;

USE `htz_sutra`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` VARCHAR(36) NOT NULL,
  `avatar` VARCHAR(36) NOT NULL DEFAULT '', -- 头像文件id
  `mobile` VARCHAR(12) NOT NULL DEFAULT '', 
  `name` VARCHAR(36) NOT NULL DEFAULT '',
  `gender` VARCHAR(36) NOT NULL DEFAULT '', 
  `birthday_year` INT(5) UNSIGNED NOT NULL DEFAULT 0,
  `birthday_month` INT(4) UNSIGNED NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `recommendations`;
CREATE TABLE `recommendations` (
  `sutra_id` VARCHAR(36) NOT NULL,
  `sutra_name` VARCHAR(36) NOT NULL,
  `sutra_cover` VARCHAR(36) NOT NULL, -- 封面文件id
  `sutra_desc` VARCHAR(500) NOT NULL,
  `sort` INT(10) UNSIGNED NOT NULL,-- 根据此排序
  PRIMARY KEY(`sutra_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `sutras`;
CREATE TABLE `sutras`(
  `id` VARCHAR(36) NOT NULL,
  `name` VARCHAR(36) NOT NULL, -- 经典名称
  `cover` VARCHAR(36) NOT NULL, -- 封面文件id
  `description` VARCHAR(500) DEFAULT '', -- 经典简介
  `played_count` INT(20) UNSIGNED DEFAULT '0', -- 播放次数
  `item_total` INT(10) UNSIGNED DEFAULT 0,-- 整个经典多少集，不包括歌词数量
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  UNIQUE `sutras_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `sutra_items`;
CREATE TABLE `sutra_items`(
  `id` VARCHAR(36) NOT NULL,
  `sutra_id` VARCHAR(36) NOT NULL,
  `title` VARCHAR(100) NOT NULL, -- 本集标题
  `description` VARCHAR(500) NOT NULL, -- 本集简介
  `original` TEXT NOT NULL, -- 本集原文
  `interpretation` TEXT NOT NULL, -- 讲师讲解
  `audio_id` VARCHAR(36) NOT NULL, -- 音频文件id
  `lyric_id` VARCHAR(36) NOT NULL, -- 音频所对应的歌词文件id
  `lesson` INT(10) UNSIGNED NOT NULL,-- 集数，根据此排序
  `played_count` INT(20) UNSIGNED DEFAULT '0',-- 播放次数
  `duration` INT(20) UNSIGNED NOT NULL, -- 时长
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY(`id`),
  KEY `sutra_items_sutra_id_lesson` (`sutra_id`,`lesson`),
  UNIQUE `sutra_items_lesson` (`lesson`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`(
  `id` VARCHAR(36) NOT NULL,
  `ref_count` INT(10) UNSIGNED NOT NULL DEFAULT 0,
  `file_size` BIGINT(20) NOT NULL,
  `mime` VARCHAR(50) NOT NULL,
  `sha1_hash` varchar(40) NOT NULL,
  `status` TINYINT(1) NOT NULL DEFAULT 0,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `push`;
CREATE TABLE `push`(
	`id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
	`reg_id` VARCHAR(36) NOT NULL,
	`user_id` VARCHAR(36) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE `push_reg_id` (`reg_id`), 
	KEY `push_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `notifications`;
CREATE TABLE `notifications`(
	`id` VARCHAR(36) NOT NULL, 
	`user_id` VARCHAR(36) NOT NULL,
	`title` VARCHAR(100) NOT NULL,
	`msg` VARCHAR(512) NOT NULL,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `notifications_user_id_created_at` (`user_id`,`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `histories`;
CREATE TABLE `histories`(
	`user_id` VARCHAR(36) NOT NULL,
	`sutra_id` VARCHAR(36) NOT NULL,
	`sutra_name` VARCHAR(36) NOT NULL,
	`sutra_cover` VARCHAR(36) NOT NULL,
	`sutra_item_id` VARCHAR(36) NOT NULL,
	`sutra_item_title` VARCHAR(100) NOT NULL,
	`last_position` INT(10) NOT NULL, -- 上次听到哪里，单位：秒。小于零表示已经听完。
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`user_id`,`sutra_id`),
	KEY `histories_user_id_created_at` (`user_id`,`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

DROP TABLE IF EXISTS `donation_logs`;
CREATE TABLE `donation_logs`(
	`id` VARCHAR(36) NOT NULL, 
	`user_id` VARCHAR(36) NOT NULL,
	`amount` INT(20) UNSIGNED NOT NULL, -- 捐款金额，单位：
	`status` TINYINT(1) NOT NULL, -- 状态；0：发起捐款，1:捐款支付确认中，2:捐款成功 
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	KEY `donation_logs_user_id` (`user_id`),
	KEY `donation_logs_updated_at` (`updated_at`),
	KEY `donation_logs_created_at` (`created_at`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;