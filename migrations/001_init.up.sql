CREATE TABLE `delegation_list` (
  `list_id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(100) NULL,
  `url` varchar(250) NULL,
  `date_add` datetime NOT NULL,
  `date_update` datetime NOT NULL
);

CREATE TABLE `games` (
  `game_id` varbinary(50) NOT NULL,
  `external_id` int(11) NOT NULL,
  `list_id` int(10) unsigned NOT NULL,
  `home_team` varchar(100) DEFAULT NULL,
  `away_team` varchar(100) DEFAULT NULL,
  `venue` varchar(50) DEFAULT NULL,
  `game_date` datetime DEFAULT NULL,

  KEY `list_id` (`list_id`), CONSTRAINT `games_ibfk_1` FOREIGN KEY (`list_id`) REFERENCES `delegation_list` (`list_id`) ON DELETE CASCADE
)