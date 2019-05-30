#!/bin/bash -e

echo waiting db for 30 sec
sleep 1

echo db init start
mysql -u root -h db -e 'create database if not exists twitter character set utf8mb4 collate utf8mb4_bin'
mysql -u root -h db twitter < init-tables.sql

echo db init finish
