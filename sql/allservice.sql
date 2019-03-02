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

 Date: 02/03/2019 18:27:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for allservice
-- ----------------------------
DROP TABLE IF EXISTS `allservice`;
CREATE TABLE `allservice`  (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `route` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的路由',
  `ip` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的ip地址',
  `port` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的端口',
  `srvname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的名字，这里是唯一的',
  `srv` varchar(300) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '服务的所有配置全在这里',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of allservice
-- ----------------------------
INSERT INTO `allservice` VALUES (1, '/root', '192.168.56.11', '8080', '', 'ser:xx');
INSERT INTO `allservice` VALUES (2, '/root', '192.168.56.11', '8080', '', 'ser:xx');
INSERT INTO `allservice` VALUES (3, '/root', '192.168.56.11', '8080', '', 'ser:xx');
INSERT INTO `allservice` VALUES (5, '/root', '192.168.56.11', '8080', 'aaa', 'ser:xx');
INSERT INTO `allservice` VALUES (6, '/root', '192.168.56.11', '8080', 'aaaa', 'ser:xx');
INSERT INTO `allservice` VALUES (7, '/root', '192.168.56.11', '8080', 'aaaaa', 'ser:xx');
INSERT INTO `allservice` VALUES (8, '/root', 'localhost', '8080', 'ab', 'ser:xxx');
INSERT INTO `allservice` VALUES (9, '/root', 'localhost', '8080', 'ac', 'ser:xxx');

SET FOREIGN_KEY_CHECKS = 1;
