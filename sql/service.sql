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

 Date: 02/03/2019 18:26:53
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for service
-- ----------------------------
DROP TABLE IF EXISTS `service`;
CREATE TABLE `service`  (
  `id` bigint(255) NOT NULL AUTO_INCREMENT,
  `route` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的路径',
  `service` varchar(600) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的全部配置',
  `servicename` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '服务的名字',
  PRIMARY KEY (`id`, `servicename`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of service
-- ----------------------------
INSERT INTO `service` VALUES (1, '/service_operation', '{\"serviceaddr\":[\"127.0.0.1:8081\"],\"registertime\":\"2019-2-17\",\"altreason\":\"\",\"balance\":\"random\"}', '\"0\"');
INSERT INTO `service` VALUES (2, '/test', '{\"serviceaddr\":[\"127.0.0.1:6061\"],\"registertime\":\"2019-2-17\",\"altreason\":\"\",\"balance\":\"polling\"}', 'apigetway');
INSERT INTO `service` VALUES (3, '/test2', '{\"serviceaddr\":[\"127.0.0.1:6061\"],\"registertime\":\"2019-2-17\",\"altreason\":\"\",\"balance\":\"polling\"}', 'apigetway');
INSERT INTO `service` VALUES (4, '/xx', '{\"serviceaddr\":[\"127.0.0.1:6062\"],\"registertime\":\"2019.2.22\",\"altreason\":\"xx\",\"servicename\":\"\",\"Balance\":\"random\"}', 'aa');
INSERT INTO `service` VALUES (5, '/xxx', '{\"serviceaddr\":[\"127.0.0.1:6068\"],\"registertime\":\"2019.2.22\",\"altreason\":\"patch test\",\"servicename\":\"\",\"Balance\":\"random\"}', 'liantiao');
INSERT INTO `service` VALUES (7, '/', '{\"serviceaddr\":[\"127.0.0.1:6066\"],\"registertime\":\"2019.2.22\",\"altreason\":\"patch test\",\"servicename\":\"\",\"Balance\":\"random\"}', 'liantiao1');

SET FOREIGN_KEY_CHECKS = 1;
