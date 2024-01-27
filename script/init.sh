
# 打开 MySQL 控制台
docker exec -it mysql bash

# 登录 MySQL
mysql -uroot -p

# 部署 MySQL 脚本
SHOW DATABASE;

CREATE DATABASE IF NOT EXISTS `mall` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `mall`;

DROP TABLE IF EXISTS `user`;

DROP TABLE IF EXISTS `order`;

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL DEFAULT '0',
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `order`(
                        `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
                        `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `create_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '创建者',
                        `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
                        `update_by` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '更新者',
                        `version` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
                        `is_del` tinyint(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除：0正常1删除',

                        `user_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '用户id',
                        `order_id` BIGINT(20) UNSIGNED NOT NULL COMMENT '订单id',
                        `trade_id` VARCHAR(128) NOT NULL DEFAULT '' COMMENT '交易单号',
                        `pay_channel` tinyint(4) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付方式',
                        `status` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单状态:100创建订单/待支付 200已支付 300交易关闭 400完成',
                        `pay_amount` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '支付金额（分）',
                        `pay_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '支付时间',

                        INDEX (user_id),
                        INDEX (order_id),
                        INDEX (trade_id),
                        INDEX (is_del)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT = '订单表';


# 检查 Redis 脚本
docker-compose -f ./docker-compose.yml exec redis7 sh -c 'redis-cli'
