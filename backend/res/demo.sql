-- MariaDB dump 10.19  Distrib 10.5.10-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: lectureplan
-- ------------------------------------------------------
-- Server version	10.5.10-MariaDB

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `courseSemesters`
--

DROP TABLE IF EXISTS `courseSemesters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `courseSemesters` (
  `idCourse` int(11) NOT NULL,
  `idSemester` int(11) NOT NULL,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`idCourse`,`idSemester`),
  UNIQUE KEY `UQ` (`idCourse`,`idSemester`),
  KEY `fk_courseSemesters_semesters1_idx` (`idSemester`),
  CONSTRAINT `fk_courseSemesters_courses1` FOREIGN KEY (`idCourse`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_courseSemesters_semesters1` FOREIGN KEY (`idSemester`) REFERENCES `semesters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courseSemesters`
--

LOCK TABLES `courseSemesters` WRITE;
/*!40000 ALTER TABLE `courseSemesters` DISABLE KEYS */;
INSERT INTO `courseSemesters` VALUES (1,1,'2021-05-31 13:40:54'),(1,2,'2021-05-31 13:40:57');
/*!40000 ALTER TABLE `courseSemesters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `courses`
--

DROP TABLE IF EXISTS `courses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `courses` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `year` year(4) NOT NULL,
  `faculty` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `courses`
--

LOCK TABLES `courses` WRITE;
/*!40000 ALTER TABLE `courses` DISABLE KEYS */;
INSERT INTO `courses` VALUES (1,'2021-05-31 13:17:09','TIF20A',2020,'Technik');
/*!40000 ALTER TABLE `courses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `eventView`
--

DROP TABLE IF EXISTS `eventView`;
/*!50001 DROP VIEW IF EXISTS `eventView`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `eventView` (
  `id` tinyint NOT NULL,
  `courseName` tinyint NOT NULL,
  `lecturerSurname` tinyint NOT NULL,
  `date` tinyint NOT NULL,
  `startTime` tinyint NOT NULL,
  `endTime` tinyint NOT NULL,
  `room` tinyint NOT NULL,
  `confirmed` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `events`
--

DROP TABLE IF EXISTS `events`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `events` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `lecture` int(11) NOT NULL,
  `date` date NOT NULL,
  `startTime` time NOT NULL,
  `endTime` time NOT NULL,
  `room` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `confirmed` tinyint(1) NOT NULL,
  `comment` text COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_events_lecture1_idx` (`lecture`),
  CONSTRAINT `fk_events_lecture1` FOREIGN KEY (`lecture`) REFERENCES `lectures` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `events`
--

LOCK TABLES `events` WRITE;
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` VALUES (2,'2021-05-31 13:56:15',1,'2021-05-31','09:00:00','16:15:00','BBB',1,'Testkommentar');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary table structure for view `lectureView`
--

DROP TABLE IF EXISTS `lectureView`;
/*!50001 DROP VIEW IF EXISTS `lectureView`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE TABLE `lectureView` (
  `id` tinyint NOT NULL,
  `lectureName` tinyint NOT NULL,
  `lecturerName` tinyint NOT NULL,
  `lecturerSurname` tinyint NOT NULL,
  `courseName` tinyint NOT NULL
) ENGINE=MyISAM */;
SET character_set_client = @saved_cs_client;

--
-- Table structure for table `lecturers`
--

DROP TABLE IF EXISTS `lecturers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lecturers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `surname` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `givenName` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `samlUID` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_lecturers_userRights1_idx` (`samlUID`),
  CONSTRAINT `fk_lecturers_userRights1` FOREIGN KEY (`samlUID`) REFERENCES `userRights` (`samlUID`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lecturers`
--

LOCK TABLES `lecturers` WRITE;
/*!40000 ALTER TABLE `lecturers` DISABLE KEYS */;
INSERT INTO `lecturers` VALUES (2,'2021-05-31 13:10:31','Birn','Kristina','birnk@dhbw-loerrach.de',NULL),(3,'2021-06-01 19:12:31','Zaremba','Fabian','behrends@dhbw-loerrach.de',NULL),(4,'2021-06-01 19:12:38','Zaremba','Fabian','behrends@dhbw-loerrach.de',NULL);
/*!40000 ALTER TABLE `lecturers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `lectures`
--

DROP TABLE IF EXISTS `lectures`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `lectures` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `lecturer` int(11) NOT NULL,
  `semester` int(11) NOT NULL,
  `course` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_lecture_lecturer_idx` (`lecturer`),
  KEY `fk_lectures_semester1_idx` (`semester`),
  KEY `fk_lectures_courses1_idx` (`course`),
  CONSTRAINT `fk_lecture_lecturer` FOREIGN KEY (`lecturer`) REFERENCES `lecturers` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_lectures_courses1` FOREIGN KEY (`course`) REFERENCES `courses` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_lectures_semester1` FOREIGN KEY (`semester`) REFERENCES `semesters` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lectures`
--

LOCK TABLES `lectures` WRITE;
/*!40000 ALTER TABLE `lectures` DISABLE KEYS */;
INSERT INTO `lectures` VALUES (1,'2021-05-31 13:55:53','Projektmanagement',2,2,1);
/*!40000 ALTER TABLE `lectures` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `preferredDays`
--

DROP TABLE IF EXISTS `preferredDays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `preferredDays` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `day` date NOT NULL,
  `lecturer` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk__lecturers1_idx` (`lecturer`),
  CONSTRAINT `fk__lecturers10` FOREIGN KEY (`lecturer`) REFERENCES `lecturers` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `preferredDays`
--

LOCK TABLES `preferredDays` WRITE;
/*!40000 ALTER TABLE `preferredDays` DISABLE KEYS */;
/*!40000 ALTER TABLE `preferredDays` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `preferredWeekdays`
--

DROP TABLE IF EXISTS `preferredWeekdays`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `preferredWeekdays` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `lecturer` int(11) NOT NULL,
  `monday` tinyint(1) NOT NULL,
  `tuesday` tinyint(1) NOT NULL,
  `wednesday` tinyint(1) NOT NULL,
  `thursday` tinyint(1) NOT NULL,
  `friday` tinyint(1) NOT NULL,
  `saturday` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`,`lecturer`),
  KEY `fk_preferredWeekdays_lecturers1_idx` (`lecturer`),
  CONSTRAINT `fk_preferredWeekdays_lecturers1` FOREIGN KEY (`lecturer`) REFERENCES `lecturers` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `preferredWeekdays`
--

LOCK TABLES `preferredWeekdays` WRITE;
/*!40000 ALTER TABLE `preferredWeekdays` DISABLE KEYS */;
/*!40000 ALTER TABLE `preferredWeekdays` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `semesters`
--

DROP TABLE IF EXISTS `semesters`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `semesters` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `start` date NOT NULL,
  `end` date NOT NULL,
  `name` varchar(45) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `semesters`
--

LOCK TABLES `semesters` WRITE;
/*!40000 ALTER TABLE `semesters` DISABLE KEYS */;
INSERT INTO `semesters` VALUES (1,'2021-05-31 13:15:42','2020-10-01','2021-03-21','Wintersemester 2020/21'),(2,'2021-05-31 13:16:38','2021-03-29','2021-10-03','Sommersemester 2021');
/*!40000 ALTER TABLE `semesters` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `unavailableTimes`
--

DROP TABLE IF EXISTS `unavailableTimes`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `unavailableTimes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT current_timestamp(),
  `startTime` datetime NOT NULL,
  `endTime` datetime NOT NULL,
  `lecturer` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk__lecturers1_idx` (`lecturer`),
  CONSTRAINT `fk__lecturers1` FOREIGN KEY (`lecturer`) REFERENCES `lecturers` (`id`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `unavailableTimes`
--

LOCK TABLES `unavailableTimes` WRITE;
/*!40000 ALTER TABLE `unavailableTimes` DISABLE KEYS */;
/*!40000 ALTER TABLE `unavailableTimes` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `userRights`
--

DROP TABLE IF EXISTS `userRights`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `userRights` (
  `samlUID` int(10) unsigned NOT NULL,
  `email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `isAdmin` tinyint(1) NOT NULL,
  `addCourse` tinyint(1) NOT NULL,
  `editEvent` tinyint(1) NOT NULL,
  `addEvent` tinyint(1) NOT NULL,
  PRIMARY KEY (`samlUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `userRights`
--

LOCK TABLES `userRights` WRITE;
/*!40000 ALTER TABLE `userRights` DISABLE KEYS */;
INSERT INTO `userRights` VALUES (1234,'zarembaf@dhbw-loerrach.de',1,1,1,1);
/*!40000 ALTER TABLE `userRights` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Final view structure for view `eventView`
--

/*!50001 DROP TABLE IF EXISTS `eventView`*/;
/*!50001 DROP VIEW IF EXISTS `eventView`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`mysql`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `eventView` AS select `events`.`id` AS `id`,`lectureView`.`courseName` AS `courseName`,`lectureView`.`lecturerSurname` AS `lecturerSurname`,`events`.`date` AS `date`,`events`.`startTime` AS `startTime`,`events`.`endTime` AS `endTime`,`events`.`room` AS `room`,`events`.`confirmed` AS `confirmed` from (`events` join `lectureView`) where `events`.`lecture` = `lectureView`.`id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `lectureView`
--

/*!50001 DROP TABLE IF EXISTS `lectureView`*/;
/*!50001 DROP VIEW IF EXISTS `lectureView`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`mysql`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `lectureView` AS select `lectures`.`id` AS `id`,`lectures`.`name` AS `lectureName`,`lecturers`.`givenName` AS `lecturerName`,`lecturers`.`surname` AS `lecturerSurname`,`courses`.`name` AS `courseName` from ((`lectures` join `lecturers`) join `courses`) where `lectures`.`course` = `courses`.`id` and `lectures`.`lecturer` = `lecturers`.`id` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-06-10 15:47:56
