CREATE DATABASE  IF NOT EXISTS `redcoin` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `redcoin`;
-- MySQL dump 10.13  Distrib 8.0.18, for Win64 (x86_64)
--
-- Host: localhost    Database: redcoin
-- ------------------------------------------------------
-- Server version	5.7.28

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
-- Table structure for table `Operacao`
--

DROP TABLE IF EXISTS `Operacao`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Operacao` (
  `idOperacao` int(11) NOT NULL AUTO_INCREMENT,
  `idTipoOperacao` tinyint(4) NOT NULL,
  `idVendedor` int(11) NOT NULL,
  `idComprador` int(11) NOT NULL,
  `dataOperacao` datetime NOT NULL,
  `valorMoeda` float NOT NULL,
  `valorBitCoin` float NOT NULL,
  PRIMARY KEY (`idOperacao`),
  KEY `FK_Operacao_I_idx` (`idTipoOperacao`),
  KEY `FK_Operacao_II_idx` (`idVendedor`),
  KEY `FK_Operacao_III_idx` (`idComprador`),
  CONSTRAINT `FK_Operacao_I` FOREIGN KEY (`idTipoOperacao`) REFERENCES `TipoOperacao` (`idTipoOperacao`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_Operacao_II` FOREIGN KEY (`idVendedor`) REFERENCES `Usuario` (`idUsuario`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_Operacao_III` FOREIGN KEY (`idComprador`) REFERENCES `Usuario` (`idUsuario`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=latin1 COMMENT='Tabela que armazena as transações das operações de compra, venda e devolucao de BitCoins';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Operacao`
--

LOCK TABLES `Operacao` WRITE;
/*!40000 ALTER TABLE `Operacao` DISABLE KEYS */;
INSERT INTO `Operacao` VALUES (8,1,2,1,'2019-12-16 14:34:12',4369.34,0.15),(9,1,2,1,'2019-12-16 14:35:10',2039.03,0.07),(10,1,2,1,'2019-12-16 14:35:53',1019.51,0.035),(11,1,2,1,'2019-12-16 14:36:31',509.76,0.0175),(12,1,1,3,'2019-12-15 10:25:15',509.76,0.0175),(13,1,1,3,'2019-12-15 08:33:29',509.76,0.0175),(14,1,1,3,'2019-12-02 07:42:29',509.76,0.0175);
/*!40000 ALTER TABLE `Operacao` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Perfil`
--

DROP TABLE IF EXISTS `Perfil`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Perfil` (
  `idPerfil` tinyint(4) NOT NULL AUTO_INCREMENT,
  `perfil` varchar(20) CHARACTER SET utf8 NOT NULL,
  `registroApagado` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idPerfil`),
  UNIQUE KEY `perfil_UNIQUE` (`perfil`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1 COMMENT='Tabela para armazenar os tipos de perfis de usuários no sistema (1- Comprador, 2- Vendedor)';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Perfil`
--

LOCK TABLES `Perfil` WRITE;
/*!40000 ALTER TABLE `Perfil` DISABLE KEYS */;
INSERT INTO `Perfil` VALUES (1,'Comprador',_binary '\0'),(2,'Vendedor',_binary '\0');
/*!40000 ALTER TABLE `Perfil` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `PerfilUsuario`
--

DROP TABLE IF EXISTS `PerfilUsuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `PerfilUsuario` (
  `idPerfilUsuario` int(11) NOT NULL AUTO_INCREMENT,
  `idPerfil` tinyint(4) NOT NULL,
  `idUsuario` int(11) NOT NULL,
  PRIMARY KEY (`idPerfilUsuario`),
  KEY `FK_PerfilUsuario_I_idx` (`idPerfil`),
  KEY `FK_PerfilUsuario_II_idx` (`idUsuario`),
  CONSTRAINT `FK_PerfilUsuario_I` FOREIGN KEY (`idPerfil`) REFERENCES `Perfil` (`idPerfil`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `FK_PerfilUsuario_II` FOREIGN KEY (`idUsuario`) REFERENCES `Usuario` (`idUsuario`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=latin1 COMMENT='Tabela responsável pelo relacionamento entre o Usuario e o seu Perfil. O usuario pode comprar, vender ou ambos.';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `PerfilUsuario`
--

LOCK TABLES `PerfilUsuario` WRITE;
/*!40000 ALTER TABLE `PerfilUsuario` DISABLE KEYS */;
INSERT INTO `PerfilUsuario` VALUES (1,1,1),(2,1,1),(3,1,1),(4,1,1),(5,1,1);
/*!40000 ALTER TABLE `PerfilUsuario` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `TipoOperacao`
--

DROP TABLE IF EXISTS `TipoOperacao`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `TipoOperacao` (
  `idTipoOperacao` tinyint(4) NOT NULL AUTO_INCREMENT,
  `Operacao` varchar(20) NOT NULL,
  `registroApagado` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idTipoOperacao`),
  UNIQUE KEY `Operacao_UNIQUE` (`Operacao`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1 COMMENT='Tabela de dominio para armazenar os tipos de transações possíveis no sistema';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `TipoOperacao`
--

LOCK TABLES `TipoOperacao` WRITE;
/*!40000 ALTER TABLE `TipoOperacao` DISABLE KEYS */;
INSERT INTO `TipoOperacao` VALUES (1,'Compra',_binary '\0'),(3,'Venda',_binary '\0'),(4,'',_binary '\0'),(11,'Teste API',_binary '\0');
/*!40000 ALTER TABLE `TipoOperacao` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `Usuario`
--

DROP TABLE IF EXISTS `Usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Usuario` (
  `idUsuario` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `senha` varchar(50) NOT NULL,
  `nome` varchar(30) NOT NULL,
  `ultimoNome` varchar(30) NOT NULL,
  `dataNascimento` date NOT NULL,
  `quantidadeMoeda` int(11) DEFAULT NULL,
  `registroApagado` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idUsuario`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COMMENT='Tabela que armazena os dados do usuário do sistema';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `Usuario`
--

LOCK TABLES `Usuario` WRITE;
/*!40000 ALTER TABLE `Usuario` DISABLE KEYS */;
INSERT INTO `Usuario` VALUES (1,'rteles@outlook.com','123Mudar','Rodolfo','Teles','1986-10-18',0,_binary '\0'),(2,'daiane_dantas2@outlook.com','123Mudar','Daiane Dantas','de Oliveira','1986-12-21',0,_binary '\0'),(3,'manu_teles@outlook.com','123Mudar','Manu Dantas','Teles','2015-05-18',0,_binary '\0');
/*!40000 ALTER TABLE `Usuario` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'goDb'
--

--
-- Dumping routines for database 'goDb'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-12-17 12:04:58
