/*
 Navicat Premium Data Transfer

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 50717
 Source Host           : localhost:3306
 Source Schema         : confcenter

 Target Server Type    : MySQL
 Target Server Version : 50717
 File Encoding         : 65001

 Date: 02/03/2019 18:27:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for configuration
-- ----------------------------
DROP TABLE IF EXISTS `configuration`;
CREATE TABLE `configuration`  (
  `id` bigint(255) NOT NULL AUTO_INCREMENT,
  `ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网关的外网ip地址',
  `port` varbinary(10) NOT NULL COMMENT '网关的外网端口',
  `timeout` int(32) NULL DEFAULT NULL COMMENT '网关的超时时间',
  `loglevel` varchar(8) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网关日志级别',
  `logpath` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '网关日志的路径',
  `modification` int(8) NULL DEFAULT 0 COMMENT '是否重复如果重复是1，不是重复的为0',
  `bufpool` int(64) NULL DEFAULT 0 COMMENT '池子的数量',
  `intranetip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内网的ip',
  `intranetport` varchar(10) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内网的启动端口',
  `maxheader` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '请求头接收最大数据',
  `managerroute` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内网网管配置路由',
  `serviceroute` varchar(128) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '内网服务路由',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 29 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of configuration
-- ----------------------------
INSERT INTO `configuration` VALUES (1, '127.0.0.1', 0x38303830, 15, 'debug', '/root/xxx/xxx/xxx', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (2, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (3, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (4, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (5, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (6, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (7, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (8, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (9, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (10, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (11, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (12, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (13, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (14, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (15, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (16, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (17, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (18, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (19, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (20, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (21, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, NULL, '', '', '', '', '');
INSERT INTO `configuration` VALUES (22, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, 1, '', '', '', '', '');
INSERT INTO `configuration` VALUES (23, 'localhost1', 0x38303830, 15, 'debug', '/root/', 1, 1, '', '', '', '', '');
INSERT INTO `configuration` VALUES (24, 'localhost1', 0x38303830, 15, 'debug', '/root/xx', 1, 1, '', '', '', '', '');
INSERT INTO `configuration` VALUES (25, 'localhost1', 0x38303830, 15, 'debug', '/root/xx', 1, 10, '', '', '', '', '');
INSERT INTO `configuration` VALUES (26, 'localhost1', 0x38303830, 15, 'debug', '/root/xx', 1, 11, '', '', '', '', '');
INSERT INTO `configuration` VALUES (27, 'localhost', 0x38303830, 15, 'debug', 'D:/project/src/quick/logcatlog', 1, 500, '127.0.0.1', '6060', '/rand', '/manager_configuration', '/service_configuration');
INSERT INTO `configuration` VALUES (28, '127.0.0.1', 0x38303930, 15, 'debug', 'D:/project/src/quick/logcatlog', 0, 0, '127.0.0.1', '6060', '', '/manager', '/service');

SET FOREIGN_KEY_CHECKS = 1;
