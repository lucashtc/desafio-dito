USE `dito`;

-- Copiando estrutura para tabela dito.events
DROP TABLE IF EXISTS `events`;
CREATE TABLE IF NOT EXISTS `events` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `EVENT` varchar(15) NOT NULL DEFAULT '',
  `TIMESTAMP` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=142 DEFAULT CHARSET=utf8;

-- Copiando dados para a tabela dito.events: ~61 rows (aproximadamente)
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
	(105, 'buy', '2019-09-07 11:38:29'),
	(106, 'comprou-camisa', '0000-00-00 00:00:00'),
	(107, 'comprou-camisa', '0000-00-00 00:00:00'),
	(108, 'comprou-camisa', '0000-00-00 00:00:00'),
	(109, 'comprou-camisa', '0000-00-00 00:00:00'),
	(110, 'comprou-camisa', '0000-00-00 00:00:00'),
	(111, 'comprou-camisa', '0000-00-00 00:00:00'),
	(112, 'comprou-camisa', '0000-00-00 00:00:00'),
	(113, 'comprou-camisa', '0000-00-00 00:00:00'),
	(114, 'comprou-camisa', '0000-00-00 00:00:00'),
	(115, 'comprou-camisa', '0000-00-00 00:00:00'),
	(116, 'comprou-camisa', '0000-00-00 00:00:00'),
	(117, 'comprou-camisa', '0000-00-00 00:00:00'),
	(118, 'comprou-camisa', '0000-00-00 00:00:00'),
	(119, 'comprou-camisa', '0000-00-00 00:00:00'),
	(120, 'comprou-camisa', '0000-00-00 00:00:00'),
	(121, 'comprou-camisa', '0000-00-00 00:00:00'),
	(122, 'comprou-camisa', '0000-00-00 00:00:00'),
	(123, 'comprou-camisa', '0000-00-00 00:00:00'),
	(124, 'comprou-camisa', '0000-00-00 00:00:00'),
	(125, 'comprou-camisa', '0000-00-00 00:00:00'),
	(126, 'comprou-camisa', '0000-00-00 00:00:00'),
	(127, 'comprou-camisa', '0000-00-00 00:00:00'),
	(128, 'comprou-camisa', '0000-00-00 00:00:00'),
	(129, 'comprou-camisa', '0000-00-00 00:00:00'),
	(130, 'comprou-camisa', '0000-00-00 00:00:00'),
	(131, 'comprou-camisa', '0000-00-00 00:00:00'),
	(132, 'comprou-camisa', '0000-00-00 00:00:00'),
	(133, 'comprou-camisa', '0000-00-00 00:00:00'),
	(134, 'comprou-camisa', '0000-00-00 00:00:00'),
	(135, 'comprou-camisa', '0000-00-00 00:00:00'),
	(136, 'comprou-camisa', '0000-00-00 00:00:00'),
	(137, 'comprou-camisa', '0000-00-00 00:00:00'),
	(138, 'buy', '2019-09-07 16:27:17'),
	(139, 'buy', '2019-09-07 16:27:30'),
	(140, 'com', '0000-00-00 00:00:00'),
	(141, 'comprou-produto', '0000-00-00 00:00:00');
/*!40000 ALTER TABLE `events` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
