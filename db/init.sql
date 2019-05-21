-- MySQL dump 10.13  Distrib 5.7.26, for Linux (x86_64)
--
-- Host: 192.168.3.18    Database: k8s
-- ------------------------------------------------------
-- Server version	5.7.21

CREATE DATABASE `k8s`;

USE  k8s;

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `dep`
--

DROP TABLE IF EXISTS `dep`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `dep` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `dep` varchar(50) NOT NULL,
  `image` varchar(50) NOT NULL,
  `ver` varchar(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `dep`
--

LOCK TABLES `dep` WRITE;
/*!40000 ALTER TABLE `dep` DISABLE KEYS */;
/*!40000 ALTER TABLE `dep` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `log`
--

DROP TABLE IF EXISTS `log`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `log` (
  `time` varchar(100) NOT NULL,
  `etype` varchar(10) NOT NULL,
  `event` varchar(50) NOT NULL,
  PRIMARY KEY (`time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `log`
--

LOCK TABLES `log` WRITE;
/*!40000 ALTER TABLE `log` DISABLE KEYS */;
/*!40000 ALTER TABLE `log` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(5) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-21 11:24:48

INSERT INTO k8s.user (id, name, password) VALUES (1, 'root', 'password');

INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:05:03 CST 2018', 'info', 'root Login from  [::1]:57531');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:24:35 CST 2018', 'info', 'Delete pod course-55b97d78f4-v8v5f success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:24:42 CST 2018', 'info', 'Delete pod course-55b97d78f4-f6cnl success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:26:18 CST 2018', 'info', 'Delete pod course-55b97d78f4-q8wr7 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:26:36 CST 2018', 'info', 'Delete pod course-55b97d78f4-q8l5d success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:26:50 CST 2018', 'info', 'Delete pod course-55b97d78f4-ss7l9 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:27:30 CST 2018', 'info', 'Delete pod course-55b97d78f4-cjmvl success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:30:11 CST 2018', 'info', 'Delete pod course-55b97d78f4-w6qtc success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:30:26 CST 2018', 'info', 'Delete pod course-55b97d78f4-wbtcz success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:31:56 CST 2018', 'info', 'Delete pod course-55b97d78f4-ktfk2 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:32:08 CST 2018', 'info', 'Delete pod course-55b97d78f4-jlfb5 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:36:14 CST 2018', 'info', 'Delete pod course-55b97d78f4-tmj6n success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:36:55 CST 2018', 'info', 'Delete pod course-55b97d78f4-k5hts success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:37:27 CST 2018', 'info', 'Delete pod course-55b97d78f4-5m99x success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:39:32 CST 2018', 'info', 'Delete pod course-55b97d78f4-4lcbm success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:39:42 CST 2018', 'error', 'Failed to delete pod course-55b97d78f4-4lcbm pods "course-55b97d78f4-4lcbm" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:39:52 CST 2018', 'info', 'Delete pod course-55b97d78f4-zcznc success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:45:31 CST 2018', 'info', 'Delete pod course-55b97d78f4-f2l58 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:46:34 CST 2018', 'info', 'Delete pod course-55b97d78f4-q29q7 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:46:55 CST 2018', 'info', 'Delete pod course-55b97d78f4-mxgxl success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:47:12 CST 2018', 'info', 'Delete pod course-55b97d78f4-ss9ns success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:50:25 CST 2018', 'info', 'Delete pod course-55b97d78f4-zjpw9 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:51:25 CST 2018', 'info', 'Delete pod course-55b97d78f4-xfg9w success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:52:19 CST 2018', 'info', 'Delete pod course-55b97d78f4-5b7rc success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 14:58:17 CST 2018', 'info', 'root Login from  [::1]:61384');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 20 15:14:02 CST 2018', 'info', 'root Login from  192.168.0.32:46347');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 11:40:45 CST 2018', 'info', 'root Login from  [::1]:57523');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 11:42:14 CST 2018', 'info', 'root Login from  [::1]:57708');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:15:01 CST 2018', 'info', 'Created pv test success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:15:03 CST 2018', 'error', 'Failed to create  PV test persistentvolumes "test" already exists');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:15:20 CST 2018', 'info', 'Created pv tt success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:19:23 CST 2018', 'info', 'Created StatefulSets test success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:37:07 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:40:44 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:40:45 CST 2018', 'error', 'Failed to get latest version of Statefulset statefulsets.apps "mysql" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:42:17 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:43:55 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 14:48:33 CST 2018', 'info', 'Created StatefulSets mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:05:53 CST 2018', 'error', 'Failed to get latest version of Statefulset statefulsets.apps "mysql" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:06:20 CST 2018', 'info', 'Delete Statefulset mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:22:04 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:27:35 CST 2018', 'error', 'Created StatefulSets mysql error');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:27:53 CST 2018', 'info', 'Created StatefulSets mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:28:00 CST 2018', 'error', 'Failed to Delete for PVC mysql invalid resource name "%!s(func() string=0x108c1f0)-mysql-0": [may not contain ''%'']');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:28:53 CST 2018', 'info', 'Delete Statefulset mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:52:18 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 15:53:07 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 17:18:59 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 17:59:22 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:01:55 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:02:43 CST 2018', 'info', 'Created StatefulSets mysqlfinaltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:13:03 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:13:04 CST 2018', 'error', 'Created StatefulSets mysql error');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:13:19 CST 2018', 'info', 'Created StatefulSets mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:14:37 CST 2018', 'error', 'Failed to delete pod mysqlfinaltest pods "mysqlfinaltest-1" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:14:38 CST 2018', 'error', 'Failed to get latest version of Statefulset statefulsets.apps "mysqlfinaltest" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:15:13 CST 2018', 'error', 'Failed to get latest version of Statefulset statefulsets.apps "mysqlfinaltest" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:15:18 CST 2018', 'error', 'Failed to get latest version of Statefulset statefulsets.apps "mysqlfinaltest" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:15:49 CST 2018', 'error', 'Failed to delete pod mysqltest pods "mysqltest-0" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:16:38 CST 2018', 'error', 'root login with wrong password ip:[::1]:4824');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:16:41 CST 2018', 'info', 'root Login from  [::1]:4824');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:20:37 CST 2018', 'info', 'Created pv pv1 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:21:13 CST 2018', 'info', 'Created StatefulSets mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:22:57 CST 2018', 'info', 'Created StatefulSets mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:25:52 CST 2018', 'info', 'Created pv pv2 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:26:22 CST 2018', 'info', 'Created StatefulSets mq success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:27:40 CST 2018', 'info', 'Delete Statefulset mysqltest success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Fri Apr 27 18:27:44 CST 2018', 'info', 'Delete Statefulset mq success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:29:53 CST 2018', 'info', 'root Login from  [::1]:13281');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:53:39 CST 2018', 'info', 'root Login from  [::1]:15295');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:54:08 CST 2018', 'info', 'root Login from  [::1]:15375');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:56:13 CST 2018', 'error', 'root login with wrong password ip:[::1]:15453');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:56:16 CST 2018', 'info', 'root Login from  [::1]:15453');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:56:30 CST 2018', 'info', 'root Login from  [::1]:15401');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 10:57:50 CST 2018', 'info', 'root Login from  [::1]:15505');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 11:01:47 CST 2018', 'info', 'root Login from  [::1]:15673');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 11:03:26 CST 2018', 'info', 'root Login from  [::1]:15797');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 11:03:47 CST 2018', 'info', 'root Login from  [::1]:15856');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 23 11:04:20 CST 2018', 'info', 'root Login from  [::1]:15868');
INSERT INTO k8s.log (time, etype, event) VALUES ('Mon Apr 30 17:43:27 CST 2018', 'info', 'root Login from  [::1]:14461');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 09:20:14 CST 2018', 'info', 'root Login from  [::1]:6430');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 09:29:44 CST 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 10:18:59 CST 2018', 'info', 'Created pod mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 10:32:35 CST 2018', 'info', 'root Login from  [::1]:8353');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 10:47:23 CST 2018', 'info', 'root Login from  [::1]:8917');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 10:47:51 CST 2018', 'info', 'root Login from  [::1]:8943');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:38:33 CST 2018', 'error', 'Failed to create  PV test PersistentVolume "test" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:39:12 CST 2018', 'error', 'Failed to create  PV test PersistentVolume "test" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:39:38 CST 2018', 'error', 'Failed to create  PV test PersistentVolume "test" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:40:06 CST 2018', 'error', 'Failed to ParseQuantity for PV asd quantities must match the regular expression ''^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$''');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:40:15 CST 2018', 'error', 'Failed to create  PV asd PersistentVolume "asd" is invalid: [spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce", spec.nfs.server: Required value]');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:47:48 CST 2018', 'error', 'Failed to create  PV asd PersistentVolume "asd" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:48:03 CST 2018', 'error', 'Failed to create  PV asd PersistentVolume "asd" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:48:32 CST 2018', 'error', 'Failed to create  PV asd PersistentVolume "asd" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:48:44 CST 2018', 'error', 'Failed to ParseQuantity for PV asd quantities must match the regular expression ''^([+-]?[0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$''');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:49:19 CST 2018', 'error', 'Failed to create  PV test PersistentVolume "test" is invalid: spec.accessModes: Unsupported value: "": supported values: "ReadOnlyMany", "ReadWriteMany", "ReadWriteOnce"');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:50:09 CST 2018', 'info', 'Created pv test success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:52:44 CST 2018', 'info', 'root Login from  [::1]:11004');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:55:14 CST 2018', 'info', 'Created deployment sad success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 11:55:37 CST 2018', 'info', 'Delete deployment sad success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 12:02:07 CST 2018', 'info', 'root Login from  [::1]:11257');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 12:46:43 CST 2018', 'info', 'root Login from  127.0.0.1:15774');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 13:49:24 CST 2018', 'info', 'Created pv asd success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 13:59:14 CST 2018', 'info', 'root Login from  [::1]:20215');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:06:23 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:07:42 CST 2018', 'info', 'Created pv test success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:20:19 CST 2018', 'info', 'Created pv tt success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:26:10 CST 2018', 'info', 'Created StatefulSets tt success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:37:59 CST 2018', 'error', 'Failed to delete pod tt pods "tt-0" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:38:07 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:43:46 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:45:52 CST 2018', 'error', 'Failed to delete pod mysql pods "mysql-0" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:47:48 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:49:37 CST 2018', 'error', 'Failed to delete pod mysql pods "mysql-0" not found');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:50:29 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:50:51 CST 2018', 'info', 'Delete pod mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:52:20 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:53:40 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 14:55:15 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:10:21 CST 2018', 'info', 'Created pv mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:20:09 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:23:38 CST 2018', 'info', 'root Login from  [::1]:22877');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:23:43 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:24:23 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:33:02 CST 2018', 'info', 'root Login from  [::1]:23416');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:33:05 CST 2018', 'info', 'Delete pod mysql-0 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:33:12 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:34:07 CST 2018', 'info', 'Created pv mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:35:36 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:36:03 CST 2018', 'info', 'Delete service mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:36:46 CST 2018', 'info', 'Created service mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:38:55 CST 2018', 'info', 'Delete pod mysql-0 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:39:11 CST 2018', 'info', 'Delete pod mysql-0 success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:41:40 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:45:06 CST 2018', 'info', 'Created pv mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:46:51 CST 2018', 'info', 'Created pv mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:47:15 CST 2018', 'info', 'root Login from  [::1]:23995');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:48:47 CST 2018', 'info', 'Created StatefulSets mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:54:45 CST 2018', 'info', 'Delete Statefulset mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Sat Apr 28 15:55:30 CST 2018', 'info', 'Created pv ttt success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Tue May  1 12:11:44 CST 2018', 'info', 'root Login from  [::1]:4385');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:19:16 UTC 2018', 'info', 'root Login from  192.168.11.254:27534');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:19:55 UTC 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:22:35 UTC 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:22:36 UTC 2018', 'info', 'kompose  up success 
Your application has been deployed to Kubernetes. You can run ''kubectl get deployment,svc,pods,pvc'' for details.
');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:27:40 UTC 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:28:08 UTC 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:28:09 UTC 2018', 'info', 'kompose  up success 
Your application has been deployed to Kubernetes. You can run ''kubectl get deployment,svc,pods,pvc'' for details.
');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:28:26 UTC 2018', 'info', 'Delete service mysql success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:40:08 UTC 2018', 'info', 'root Login from  192.168.11.254:28236');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:52:43 UTC 2018', 'info', 'Update deployment course success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 05:55:24 UTC 2018', 'info', 'Delete pod course-5466655d55-bnplb success');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 06:10:51 UTC 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 14:13:24 CST 2018', 'info', 'root Login from  192.168.11.254:29314');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 14:24:02 CST 2018', 'info', 'Already truncate docker-compose.yaml');
INSERT INTO k8s.log (time, etype, event) VALUES ('Wed May  2 14:24:03 CST 2018', 'info', 'kompose  up success 
Your application has been deployed to Kubernetes. You can run ''kubectl get deployment,svc,pods,pvc'' for details.
');