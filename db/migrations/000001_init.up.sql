create table users
(
    id       serial primary key,
    username varchar(255) not null,
    name     varchar(255) not null,
    password varchar(255) not null
);