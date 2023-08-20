CREATE TABLE `t_user` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `email` varchar(64) DEFAULT NULL COMMENT '邮箱',
      `password` varchar(64) DEFAULT NULL COMMENT '密码',
      `avatar` varchar(512) DEFAULT NULL COMMENT '头像',
      `nick_name` varchar(64) DEFAULT NULL COMMENT '昵称',
      `description` varchar(32) DEFAULT NULL COMMENT '描述',
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `deleted_at` datetime DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_note` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `user_id` int(11) NOT NULL COMMENT '用户ID',
      `content` longtext DEFAULT NULL COMMENT '内容',
      `images` text DEFAULT NULL COMMENT '图片',
      `location` varchar(64) DEFAULT NULL COMMENT '头像',
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `deleted_at` datetime DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_chat` (
     `id` int(11) NOT NULL AUTO_INCREMENT,
     `user_id` varchar(16) DEFAULT '' COMMENT '用户ID，逗号分割',
     `sender_id` int(11) NOT NULL COMMENT '发送者ID,最新的一条',
     `content` text DEFAULT NULL COMMENT '内容，最新的一条',
     `is_liked` int DEFAULT 2 COMMENT '是否链接 1是 2不是',
     `unread`int DEFAULT 1 COMMENT '未读 1未读 2已读',
     `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
     `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     `deleted_at` datetime DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;

CREATE TABLE `t_message` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `sender_id` int(11) NOT NULL COMMENT '发送者ID',
      `receiver_id` int(11) NOT NULL COMMENT '接收者ID',
      `content` text DEFAULT NULL COMMENT '内容',
      `is_liked` int DEFAULT 2 COMMENT '是否链接 1是 2不是',
      `unread`int DEFAULT 1 COMMENT '未读 1未读 2已读',
      `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
      `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      `deleted_at` datetime DEFAULT NULL,
      PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;