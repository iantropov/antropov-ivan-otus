# ДЗ 5. Работа с шардированием

## Описание

Доработал приложение с прошлых ДЗ с целью поддержки друзей и постов
Реализовал функционал ленты постов друзей
Реализовал кэширование ленты постов друзей

Приложение разворачивается из docker-compose

## Links

Swagger - https://editor.swagger.io/

Go-migrate - https://github.com/golang-migrate/migrate

## Roadmap

+ Migration for MySQL from Golang
+ Extract current user from JWT claims
+ 'Friends' feature
+ 'Posts' feature
+ 'Feed' feature
+ Redis for 'Feed' feature
+ Save feed in Redis
+ Invalidation of cache

## Commands

```
mysql -uuser -p -P6032 -h 127.0.0.1 --prompt='Admin> '
mysql -uuser -p -P6033 -h 127.0.0.1 --prompt='Client> '

insert into mysql_query_rules(rule_id,active,match_pattern,destination_hostgroup,apply) values(4, 1, "e1b77b7e-49d7-48ee-a51e-edb1c6aafb06", 1, 1);
LOAD MYSQL QUERY RULES TO RUNTIME;

select * from stats_mysql_processlist;

select * from stats_mysql_prepared_statements_info;

```