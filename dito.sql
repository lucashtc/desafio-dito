USE `dito`
DROP TABLE IF EXISTS `events`;
CREATE TABLE IF NOT EXISTS `events` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `EVENT` varchar(15) NOT NULL DEFAULT '',
  `TIMESTAMP` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8;

-- Copiando dados para a tabela dito.events: ~59 rows (aproximadamente)
/*!40000 ALTER TABLE `events` DISABLE KEYS */;
INSERT INTO `events` (`ID`, `EVENT`, `TIMESTAMP`) VALUES
	(2, 'comprou-produto', '2019-09-07 00:05:56'),
	(3, 'comprou', '2019-09-07 00:06:09'),
	(83, 'buy', '2019-09-07 11:38:27'),
	(84, 'buy', '2019-09-07 11:38:27'),
	(85, 'buy', '2019-09-07 11:38:27'),
	(86, 'buy', '2019-09-07 11:38:27'),
	(87, 'buy', '2019-09-07 11:38:28'),
	(88, 'buy', '2019-09-07 11:38:28'),
	(89, 'buy', '2019-09-07 11:38:28'),
	(90, 'buy', '2019-09-07 11:38:28'),
	(91, 'buy', '2019-09-07 11:38:28'),
	(92, 'buy', '2019-09-07 11:38:28'),
	(93, 'buy', '2019-09-07 11:38:28'),
	(94, 'buy', '2019-09-07 11:38:28'),
	(95, 'buy', '2019-09-07 11:38:28'),
	(96, 'buy', '2019-09-07 11:38:28'),
	(97, 'buy', '2019-09-07 11:38:28'),
	(98, 'buy', '2019-09-07 11:38:29'),
	(99, 'buy', '2019-09-07 11:38:29'),
	(100, 'buy', '2019-09-07 11:38:29'),
	(101, 'buy', '2019-09-07 11:38:29'),
	(102, 'buy', '2019-09-07 11:38:29'),
	(103, 'buy', '2019-09-07 11:38:29'),
	(104, 'buy', '2019-09-07 11:38:29'),
	(105, 'buy', '2019-09-07 11:38:29');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
