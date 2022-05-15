CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(50) NOT NULL,
    userName varchar(50) NOT NULL unique,
    email varchar(80) NOT NULL unique,
    password varchar(30) NOT NULL,
    createdAt timestamp default current_timestamp()
) ENGINE=InnoDB;

CREATE TABLE followers(
    userId int not null, foreign key(userId) references users(id) on delete cascade,

    followerId int not null, foreign key(followerId) references users(id) on delete cascade,

    primary key (userId, followerId)
) ENGINE=InnoDB;