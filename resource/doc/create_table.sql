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