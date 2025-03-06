-- MySQL dump 10.13  Distrib 8.0.34, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: oneiot_server
-- ------------------------------------------------------
-- Server version	8.2.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `buyers`
--

DROP TABLE IF EXISTS `buyers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `buyers` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `FullName` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `PhoneNumber` varchar(255) NOT NULL,
  `FullAddress` varchar(255) NOT NULL,
  `AdditionalNotes` text,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `buyers`
--

LOCK TABLES `buyers` WRITE;
/*!40000 ALTER TABLE `buyers` DISABLE KEYS */;
INSERT INTO `buyers` VALUES (1,'vincent','vincent@gmail.com','08232413212','jonggol','yang penting cepat'),(2,'Vincent Kenutama','vincent@gmail.com','072112123','Jonggol Barat, Jakarta','Yang penting nyoba dulu hehe'),(3,'eko','eko@gmail.com','092832983','Jakarta',''),(4,'Vincent Kenutama','','','',''),(5,'Vincent Kenutama','','','',''),(6,'Vincent Kenutama','','','',''),(7,'Vincent Kenutama','','','',''),(8,'Vincent Kenutama','','','',''),(9,'Vincent Kenutama','','','',''),(10,'Vincent Kenutama','','','',''),(11,'','','','',''),(12,'Erlangga Satrya','erlangga@gmail.com','072112123','Testing, Yogyakarta','Yang penting nyoba dulu hehe'),(13,'Erlangga Satrya','erlangga@gmail.com','072112123','Testing, Yogyakarta','Yang penting nyoba dulu hehe'),(14,'Erlangga Satrya','erlangga@gmail.com','072112123','Testing, Yogyakarta','Yang penting nyoba dulu hehe'),(15,'Erlangga Satrya','erlangga@gmail.com','072112123','Testing, Yogyakarta','Yang penting nyoba dulu hehe'),(16,'LOLOLO Satrya','erlangga@gmail.com','072112123','Testing, Yogyakarta','Yang penting nyoba dulu hehe');
/*!40000 ALTER TABLE `buyers` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `deliverystatuses`
--

DROP TABLE IF EXISTS `deliverystatuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `deliverystatuses` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `DeliveryDate` date DEFAULT NULL,
  `ArriveEstimation` date DEFAULT NULL,
  `RecipientName` varchar(255) DEFAULT NULL,
  `Courier` varchar(255) DEFAULT NULL,
  `Address` varchar(255) DEFAULT NULL,
  `TrackingNumber` varchar(255) DEFAULT NULL,
  `DeliveryCourier` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `deliverystatuses`
--

LOCK TABLES `deliverystatuses` WRITE;
/*!40000 ALTER TABLE `deliverystatuses` DISABLE KEYS */;
INSERT INTO `deliverystatuses` VALUES (2,'2025-02-11','2025-02-11','John Doe','FedEx','123 Main St, City, Country','TRK123456','FedEx Express'),(3,'2025-02-13','2025-02-13','Erlangga','JNT','Jakarta','1231','JNT'),(9,'2025-02-13','2025-02-13','Erlangga','JNT','Jakarta','1231','JNT'),(10,'2025-02-13','2025-02-13','Erlangga','JNT','Jakarta','1231','JNT'),(11,'2025-02-13','2025-02-13','Erlangga','JNT','Jakarta','1231','JNT'),(12,'2025-02-13','2025-02-13','Erlangga','JNT','Jakarta','1231','JNT');
/*!40000 ALTER TABLE `deliverystatuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orderdetails`
--

DROP TABLE IF EXISTS `orderdetails`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orderdetails` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `OrderName` varchar(255) NOT NULL,
  `ServicesId` int NOT NULL,
  `Deadline` date DEFAULT NULL,
  `Speed` enum('Regular','Express','Full Speed') DEFAULT NULL,
  `BriefFile` varchar(255) NOT NULL,
  `ImportantPoint` text,
  `AdditionalNotes` text,
  `OrderSummaryFile` varchar(255) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `ServicesId` (`ServicesId`),
  CONSTRAINT `orderdetails_ibfk_1` FOREIGN KEY (`ServicesId`) REFERENCES `services` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orderdetails`
--

LOCK TABLES `orderdetails` WRITE;
/*!40000 ALTER TABLE `orderdetails` DISABLE KEYS */;
INSERT INTO `orderdetails` VALUES (15,'Logo Design',1,'2025-02-15','Regular','127.0.0.1:8000/static/order_briefs/2_2025-02-06 15-55-05_Ubuntu Server CLI cheat sheet 2024 v6.pdf','Minimalist style, black and white','Client prefers a modern look','logo_summary.pdf'),(17,'PCB Design',2,'2025-02-15','Regular','pcb_design_brief.pdf','yang penting jadi','dan bagus','summary.pdf'),(20,'Alat Monitoring Kesehatan Pasien',1,'2025-02-04','Regular','alat_monitoring.pdf','','Dipercepat karena dibutuhkan segera','order_summary.pdf'),(21,'Alat Monitoring Kesehatan Pasien',1,'2025-02-04','Regular','alat_monitoring.pdf','','Dipercepat karena dibutuhkan segera','order_summary.pdf'),(22,'PCB Design',2,'2025-02-15','Regular','pcb_design_brief.pdf','yang penting jadi','dan bagus','summary.pdf'),(23,'PCB Design',2,'2025-02-15','Regular','pcb_design_brief.pdf','yang penting jadi','dan bagus','summary.pdf'),(24,'PCB Design',2,'2025-02-15','Regular','127.0.0.1:8000/static/order_briefs/13_2025-02-09 03-12-19_resistorsandcaps.pdf','yang penting jadi','dan bagus','summary.pdf'),(25,'PCB Design',2,'2025-02-15','Regular','pcb_design_brief.pdf','yang penting jadi','dan bagus','summary.pdf'),(26,'PCB Design',2,'2025-02-15','Regular','127.0.0.1:8000/static/order_briefs/15_2025-02-09 03-44-04_fiony cantik.jpg','yang penting jadi','dan bagus','summary.pdf'),(27,'PCB Design',2,'2025-02-15','Regular','127.0.0.1:8000/static/order_briefs/15_2025-02-09 03-39-22_fiony cantik.jpg','yang penting jadi','dan bagus','summary.pdf');
/*!40000 ALTER TABLE `orderdetails` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `UserId` int NOT NULL,
  `BuyerId` int NOT NULL,
  `OrderDetailId` int NOT NULL,
  `IsActive` tinyint(1) DEFAULT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  `Confirmed` tinyint(1) DEFAULT '0',
  `TransactionCreated` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`),
  KEY `OrderDetailId` (`OrderDetailId`),
  KEY `orders_ibfk_1` (`BuyerId`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`BuyerId`) REFERENCES `buyers` (`Id`),
  CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`BuyerId`) REFERENCES `buyers` (`Id`),
  CONSTRAINT `orders_ibfk_3` FOREIGN KEY (`OrderDetailId`) REFERENCES `orderdetails` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,2,2,15,1,'2025-02-03 00:00:00',0,0),(2,2,2,17,0,'2025-02-04 00:00:00',0,0),(3,2,2,15,1,'2025-02-04 14:51:56',0,0),(4,2,9,20,0,'0001-01-01 00:00:00',0,0),(5,2,10,21,0,'2025-02-04 15:47:08',0,0),(6,2,12,23,0,'2025-02-04 17:25:23',1,0),(7,13,13,24,0,'2025-02-09 00:15:03',0,0),(8,13,14,25,0,'2025-02-09 00:28:35',0,0),(9,15,15,26,0,'2025-02-09 03:21:13',0,0),(10,15,16,27,0,'2025-02-09 03:38:34',0,0);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `paymentmethods`
--

DROP TABLE IF EXISTS `paymentmethods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `paymentmethods` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Name` varchar(255) DEFAULT NULL,
  `Number` varchar(255) DEFAULT NULL,
  `Logo` text,
  `Acronym` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `paymentmethods`
--

LOCK TABLES `paymentmethods` WRITE;
/*!40000 ALTER TABLE `paymentmethods` DISABLE KEYS */;
INSERT INTO `paymentmethods` VALUES (2,'BCA Virtual Account','0823','logo','BCA VA'),(4,'BCA Virtual Account','0823','logo','BCA VA');
/*!40000 ALTER TABLE `paymentmethods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `PaymentProof` text,
  `Invoice` text,
  `Paid` tinyint(1) DEFAULT '0',
  `PaymentMethodsId` int DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `PaymentMethodsId` (`PaymentMethodsId`),
  CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`PaymentMethodsId`) REFERENCES `paymentmethods` (`Id`),
  CONSTRAINT `payments_ibfk_2` FOREIGN KEY (`PaymentMethodsId`) REFERENCES `paymentmethods` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (1,'Set test','aa',1,2),(3,'0','test_invoice.pdf',0,2),(6,'proof.jpg','invoice.pdf',0,2),(12,'proof.jpg','invoice.pdf',0,2),(13,'proof.jpg','invoice.pdf',0,2),(14,'proof.jpg','invoice.pdf',0,2),(15,'proof.jpg','invoice.pdf',0,2);
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pricings`
--

DROP TABLE IF EXISTS `pricings`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `pricings` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `BasePrice` decimal(10,0) DEFAULT NULL,
  `ServicePrice` decimal(10,0) DEFAULT NULL,
  `DeliveryFee` decimal(10,0) DEFAULT NULL,
  `Tax` decimal(10,0) DEFAULT NULL,
  `AdditionalPrice` decimal(10,0) DEFAULT NULL,
  `TotalPrice` decimal(10,0) DEFAULT NULL,
  `PaymentsId` int DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `PaymentsId` (`PaymentsId`),
  CONSTRAINT `pricings_ibfk_1` FOREIGN KEY (`PaymentsId`) REFERENCES `payments` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pricings`
--

LOCK TABLES `pricings` WRITE;
/*!40000 ALTER TABLE `pricings` DISABLE KEYS */;
INSERT INTO `pricings` VALUES (1,7000,8000,8000,12000,20000,100000,3),(3,1000,1000,1000,440,1000,4440,3),(4,2000,2000,2000,1100,4000,11100,6),(10,2000,2000,2000,1100,4000,11100,12),(11,2000,2000,2000,1100,4000,11100,13),(12,2000,2000,2000,1100,4000,11100,14),(13,2000,2000,2000,1100,4000,11100,15);
/*!40000 ALTER TABLE `pricings` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `productionstatuses`
--

DROP TABLE IF EXISTS `productionstatuses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `productionstatuses` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `ProductionDate` datetime DEFAULT NULL,
  `EstimatedDate` datetime DEFAULT NULL,
  `LatestStatus` varchar(255) DEFAULT NULL,
  `ProductionStages` text,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `productionstatuses`
--

LOCK TABLES `productionstatuses` WRITE;
/*!40000 ALTER TABLE `productionstatuses` DISABLE KEYS */;
INSERT INTO `productionstatuses` VALUES (1,'2025-02-11 16:27:32','2025-02-11 16:27:32','HihiJson','Haha'),(3,'2025-02-13 17:39:56','2025-02-13 17:39:56','Unpaid','{}'),(9,'2025-02-13 18:20:30','2025-02-13 18:20:30','Unpaid','{}'),(10,'2025-02-13 18:35:03','2025-02-13 18:35:03','Unpaid','{}'),(11,'2025-02-13 18:35:09','2025-02-13 18:35:09','Unpaid','{}'),(12,'2025-02-13 18:35:03','2025-02-13 18:35:03','Unpaid','{}');
/*!40000 ALTER TABLE `productionstatuses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reviews`
--

DROP TABLE IF EXISTS `reviews`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reviews` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `UserId` int DEFAULT NULL,
  `TransactionId` int DEFAULT NULL,
  `Rating` decimal(10,0) DEFAULT NULL,
  `Commentary` text,
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`),
  KEY `TransactionId` (`TransactionId`),
  CONSTRAINT `reviews_ibfk_1` FOREIGN KEY (`UserId`) REFERENCES `users` (`Id`),
  CONSTRAINT `reviews_ibfk_2` FOREIGN KEY (`TransactionId`) REFERENCES `transactions` (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reviews`
--

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;
/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `services`
--

DROP TABLE IF EXISTS `services`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `services` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Icon` varchar(255) DEFAULT NULL,
  `BgColor` varchar(255) DEFAULT NULL,
  `ServiceName` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `services`
--

LOCK TABLES `services` WRITE;
/*!40000 ALTER TABLE `services` DISABLE KEYS */;
INSERT INTO `services` VALUES (1,'sokfjskl.jpg','#0xFFFFF#','Internet of Things'),(2,'sokfjskl.jpg','#0xFFFFF#','Machine Learning');
/*!40000 ALTER TABLE `services` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `transactions`
--

DROP TABLE IF EXISTS `transactions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `transactions` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `UserId` int DEFAULT NULL,
  `OrderId` int DEFAULT NULL,
  `PricingId` int DEFAULT NULL,
  `ProductionStatusesId` int DEFAULT NULL,
  `DeliveryStatusesId` int DEFAULT NULL,
  `Status` varchar(255) DEFAULT NULL,
  `CreatedAt` datetime DEFAULT NULL,
  `Complained` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`),
  KEY `OrderId` (`OrderId`),
  KEY `PricingId` (`PricingId`),
  KEY `ProductionStatusesId` (`ProductionStatusesId`),
  KEY `DeliveryStatusesId` (`DeliveryStatusesId`),
  CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`UserId`) REFERENCES `users` (`Id`),
  CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`OrderId`) REFERENCES `orders` (`Id`),
  CONSTRAINT `transactions_ibfk_3` FOREIGN KEY (`PricingId`) REFERENCES `pricings` (`Id`),
  CONSTRAINT `transactions_ibfk_4` FOREIGN KEY (`ProductionStatusesId`) REFERENCES `productionstatuses` (`Id`),
  CONSTRAINT `transactions_ibfk_5` FOREIGN KEY (`DeliveryStatusesId`) REFERENCES `deliverystatuses` (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `transactions`
--

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;
INSERT INTO `transactions` VALUES (2,15,10,3,1,2,'Unpaid','2025-02-11 17:47:21',0),(3,15,10,3,1,2,'Paid','2025-02-11 17:47:35',0),(11,15,10,11,10,10,'Unpaid','2025-02-13 18:35:03',0),(12,15,10,12,11,11,'Unpaid','2025-02-13 18:35:09',0),(13,15,10,13,12,12,'Unpaid','2025-02-13 18:35:03',0);
/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `Id` int NOT NULL AUTO_INCREMENT,
  `Fullname` varchar(255) NOT NULL,
  `Email` varchar(255) NOT NULL,
  `Password` varchar(255) NOT NULL,
  `PhoneNumber` varchar(255) NOT NULL,
  `picture` varchar(255) DEFAULT NULL,
  `Address` varchar(255) NOT NULL,
  `Location` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'angga','angga@gmail.com','slfkjsf','62812938','dlkfjsljf','vietnam','{latitude: \'3943948\', longitude: \'2382038\'}'),(2,'testing','testing@gmail.com','testingpassword','081234','testingpic.jpg','testing jawa tengah','testing{}'),(6,'Vincent Kenutama','vincent@gmail.com','password','082131313','vincent.jpg','Yogyakarta','{}'),(7,'John Doe','john.doe@gmail.com','inipasswordaaa','08961930','picture.jpg','sana sini','Jonggol'),(10,'Erlangga Satrya','erlangga@gmail.com','erlanggaGanteng','0821','127.0.0.1:8000/static/user_pictures/10_fiony cantik.jpg','yogya',''),(11,'Erlangga Satrya','aaa@gmail.com','aaaGanteng','0821','','yogya',''),(12,'','test@gmail.com','test','','127.0.0.1:8000/static/user_pictures/0_fiony cantik.jpg','',''),(13,'','anggi@gmail.com','anggi','','127.0.0.1:8000/static/user_pictures/0_fiony cantik.jpg','',''),(14,'Erlangga Satrya','ast@gmail.com','sat','0821','127.0.0.1:8000/static/user_pictures/14_fiony cantik.jpg','yogya',''),(15,'Erlangga Satrya','abc@gmail.com','$2a$10$PpdqnGZeAhvlxhLnJb1Xl.c/PSSBa/QW8qi77onh1pQGY8fuWE0F6','0821','127.0.0.1:8000/static/user_pictures/15_gabriela.jpeg','yogya','');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-02-16 13:04:34
