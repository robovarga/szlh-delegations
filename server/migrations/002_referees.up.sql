CREATE TABLE `referees` (
  `referee_id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(250) NOT NULL,
  `date_add` datetime NOT NULL,
  `date_update` datetime NOT NULL
);

CREATE TABLE `game_referees` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `game_id` int NOT NULL,
  `referee_id` int NOT NULL,
  `date_add` datetime NOT NULL,
  `date_update` datetime NOT NULL,
  FOREIGN KEY (`game_id`) REFERENCES `games` (`game_id`),
  FOREIGN KEY (`referee_id`) REFERENCES `referees` (`referee_id`)
);