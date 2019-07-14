drop table if exists posts;
create table if not exists posts
(
  id      int unsigned not null primary key auto_increment,
  user_id    int not null,
  text    varchar(256) not null
);

create table if not exists users
(
    id      int unsigned not null primary key auto_increment,
    name    varchar(128) not null,
    email    varchar(128) not null,
    password    varchar(256) not null
);