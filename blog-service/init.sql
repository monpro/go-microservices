create DATABASE
IF NOT EXISTS blog_service DEFAULT CHARACTER
SET utf8mb4 DEFAULT collate utf8mb4_general_ci;

CREATE TABLE `blog_auth` (
     `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
     `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
     `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',
     `created_on` int(10) unsigned DEFAULT '0',
     `created_by` varchar(100) DEFAULT '',
     `modified_on` int(10) unsigned DEFAULT '0',
     `modified_by` varchar(100) DEFAULT '',
     `deleted_on` int(10) unsigned DEFAULT '0',
     `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT '0 not deleted, 1 deleted',
     PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='auth management';

CREATE TABLE `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` varchar(100) DEFAULT '',
    `created_on` int(10) unsigned DEFAULT '0',
    `created_by` varchar(100) DEFAULT '',
    `modified_on` int(10) unsigned DEFAULT '0',
    `modified_by` varchar(100) DEFAULT '',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `is_del` tinyint(3) unsigned DEFAULT '0', # 0 not deleted, 1 deleted
    `state` tinyint(3) unsigned DEFAULT '1' # 0 not in use, 1 in use
) ENGINE =InnoDB DEFAULT CHARSET = utf8mb4;


CREATE TABLE `blog_article` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `title` varchar(100) DEFAULT '',
    `desc` varchar(255) DEFAULT '',
    `cover_image_url` varchar(255) DEFAULT '',
    `content` longtext,
    `created_on` int(10) unsigned DEFAULT '0',
    `created_by` varchar(100) DEFAULT '',
    `modified_on` int(10) unsigned DEFAULT '0',
    `modified_by` varchar(100) DEFAULT '',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `is_del` tinyint(3) unsigned DEFAULT '0', # 0 not deleted, 1 deleted
    `state` tinyint(3) unsigned DEFAULT '1' # 0 not in use, 1 in use
) ENGINE =InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `blog_article_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `article_id` int(11) NOT NULL,
    `tag_id` int(10) unsigned NOT NULL DEFAULT '0',
    `created_on` int(10) unsigned DEFAULT '0',
    `created_by` varchar(100) DEFAULT '',
    `modified_on` int(10) unsigned DEFAULT '0',
    `modified_by` varchar(100) DEFAULT '',
    `deleted_on` int(10) unsigned DEFAULT '0',
    `is_del` tinyint(3) unsigned DEFAULT '0' # 0 not deleted, 1 deleted
) ENGINE =InnoDB DEFAULT CHARSET = utf8mb4;
