CREATE TABLE IF NOT EXISTS `problem`
(
    `id`          int UNSIGNED NOT NULL AUTO_INCREMENT,
    `identity`    varchar(36) NULL COMMENT '唯一标识',
    `category_id` varchar(255) NULL COMMENT '以逗号分隔的分类id',
    `title`       varchar(255) NULL COMMENT '问题题目',
    `content`     text NULL COMMENT '问题正文描述',
    `total_num`   int(11) NULL DEFAULT 0 COMMENT '总共提交个数',
    `created_at`  datetime NULL default null COMMENT '创建时间',
    `updated_at`  datetime NULL default null COMMENT '操作时间',
    `deleted_at`  datetime NULL default null COMMENT '删除时间，软删除',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `identity`(`identity`) USING BTREE
);

create table IF NOT EXISTS `user`
(
    `id`         int unsigned not null auto_increment,
    `identity`   varchar(36) NULL COMMENT '唯一标识',
    `name`       varchar(100) null comment '用户名称',
    `password`   varchar(32) null comment '密码',
    `phone`      varchar(100) null comment '手机号',
    `email`      varchar(100) null comment '用户邮箱',
    `created_at` datetime NULL default null COMMENT '创建时间',
    `updated_at` datetime NULL default null COMMENT '操作时间',
    `deleted_at` datetime NULL default null COMMENT '删除时间，软删除',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `identity`(`identity`) USING BTREE

);

create table IF NOT EXISTS `category`
(
    `id`         int unsigned not null auto_increment,
    `identity`   varchar(36) NULL COMMENT '唯一标识',
    `name`       varchar(100) null comment '名称',
    `parent_id`  int unsigned default 0 comment '上级id',
    `created_at` datetime NULL default null COMMENT '创建时间',
    `updated_at` datetime NULL default null COMMENT '操作时间',
    `deleted_at` datetime NULL default null COMMENT '删除时间，软删除',
    PRIMARY KEY (`id`),
    INDEX        `parent_id`(`parent_id`) USING BTREE,
    UNIQUE INDEX `identity`(`identity`) USING BTREE

);

create table IF NOT EXISTS `submit`
(
    `id`       int unsigned not null auto_increment,
    `identity` varchar(36) NULL COMMENT '唯一标识',
    `problem_identity` varchar(36) null comment '问题标识',
    `user_identity` varchar(36) null comment '用户标识',
    `path` varchar(150) null comment '提交的路径',
    `status` tinyint unsigned default 0 comment '状态，0:待判断，1:正确,2:错误',
    PRIMARY KEY (`id`),
    INDEX `problem_identity`(`problem_identity`) USING BTREE,
    INDEX `user_identity`(`user_identity`) USING BTREE,
    UNIQUE INDEX `identity`(`identity`) USING BTREE
);
