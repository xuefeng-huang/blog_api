#!/bin/bash
# stops execution when error occurs
set -o errexit

mysql -uroot -ppassword -e "create database blog"
mysql -uroot -ppassword -e "create database blog_test"

mysql -uroot -ppassword blog <<eof
create table articles (
    id int auto_increment primary key,
    title varchar(255) not null default '',
    content text not null,
    author varchar(64) not null default ''
);

insert into articles(title, content, author) values ("first", "my first post", "tim"),
("second", "my second post", "xuefeng"), ("third", "third post", "tim");
eof

mysql -uroot -ppassword blog_test <<eof
create table articles (
    id int auto_increment primary key,
    title varchar(255) not null default '',
    content text not null,
    author varchar(64) not null default ''
);
eof