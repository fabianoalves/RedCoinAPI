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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela que armazena as transações das operações de compra, venda e devolucao de BitCoins';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `Perfil`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Perfil` (
  `idPerfil` tinyint(4) NOT NULL AUTO_INCREMENT,
  `perfil` varchar(20) CHARACTER SET utf8 NOT NULL,
  `registroApagado` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idPerfil`),
  UNIQUE KEY `perfil_UNIQUE` (`perfil`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela para armazenar os tipos de perfis de usuários no sistema (1- Comprador, 2- Vendedor)';
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `ClienteApi`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ClienteApi` (
  `idClienteApi` int(11) NOT NULL AUTO_INCREMENT,
  `usuario` varchar(20) CHARACTER SET utf8 NOT NULL,
  `senha` MEDIUMTEXT CHARACTER SET utf8 NOT NULL,
  PRIMARY KEY (`idClienteApi`),
  UNIQUE KEY `clienteApi_UNIQUE` (`usuario`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela para armazenar os clientes habilitados a acessarem a API';
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela responsável pelo relacionamento entre o Usuario e o seu Perfil. O usuario pode comprar, vender ou ambos.';
/*!40101 SET character_set_client = @saved_cs_client */;

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
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela de dominio para armazenar os tipos de transações possíveis no sistema';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `Usuario`
--

DROP TABLE IF EXISTS `Usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `Usuario` (
  `idUsuario` int(11) NOT NULL AUTO_INCREMENT,
  `email` varchar(50) NOT NULL,
  `senha` MEDIUMTEXT NOT NULL,
  `nome` varchar(30) NOT NULL,
  `ultimoNome` varchar(30) NOT NULL,
  `dataNascimento` date NOT NULL,
  `quantidadeMoeda` float DEFAULT NULL,
  `registroApagado` bit(1) DEFAULT b'0',
  PRIMARY KEY (`idUsuario`),
  UNIQUE KEY `email_UNIQUE` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1 COMMENT='Tabela que armazena os dados do usuário do sistema';
/*!40101 SET character_set_client = @saved_cs_client */;

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
