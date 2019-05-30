drop table if exists posts;
create table if not exists posts
(
  id      int unsigned not null primary key auto_increment,
  name    varchar(128) not null,
  text    varchar(256) not null
);
