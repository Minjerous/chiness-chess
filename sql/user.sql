CREATE TABLE `users` (
                         `uid` bigint(20) NOT NULL AUTO_INCREMENT,
                         `user_name` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `salt` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `phone` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `pass_word` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `email` varchar(171) COLLATE utf8_croatian_ci NOT NULL,
                         `follow_count` bigint(20) NOT NULL,
                         `follower_count` bigint(20) DEFAULT NULL,
                         `signature` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `avatar` varchar(171) COLLATE utf8_croatian_ci DEFAULT NULL,
                         `coin` bigint(20) DEFAULT NULL,
                         PRIMARY KEY (`uid`,`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_croatian_ci;