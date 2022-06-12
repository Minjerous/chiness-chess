CREATE TABLE `rooms` (
                         `id` bigint(20) NOT NULL AUTO_INCREMENT,
                         `name` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `user_one` bigint(20) DEFAULT NULL,
                         `user_two` bigint(20) DEFAULT NULL,
                         `pass_word` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `room_id` bigint(20) DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;