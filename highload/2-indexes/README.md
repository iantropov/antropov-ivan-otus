# Homework 2: Indexes

Заготовка социальной сети

Чтобы запустить:

1. Завести .env файл, например, `cp sample.env .env`
2. Поднять сервисы через `docker-compose up`
3. Проверить работу через Postman (импортировав коллекцию `social_network.postman_collection.json`)

### Links

https://go.dev/doc/tutorial/database-access

### Roadmap (for myself)

+ Реализовать поиск по пользователям
+ Попробовать поискать с explain
+ Вынести хранение данных из контейнера
+ Написать запрос импорта пользователей
+ Импортировать пользователей
+ Поискать пользователей без индекса с explain
+ Установить jmeter
+ Протестировать приложение по заданию
+ Добавить индексы
+ Поискать пользователей с индексом с explan
+ Протестировать приложение по заданию
- Сформировать отчёты

## Scripts

mysql -uuser -ppassword social-network

LOAD DATA INFILE './people.csv'
INTO TABLE users
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS
(@name,age,city)
SET id = UUID()
SET first_name=SUBSTRING_INDEX(@name, ' ', 1)
SET last_name=SUBSTRING_INDEX(@name, ' ', -1);

UUID()
SUBSTRING_INDEX(@name, ' ', 1)
SUBSTRING_INDEX(@name, ' ', -11)

LOAD DATA LOCAL INFILE './people.csv' INTO TABLE users FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 ROWS (@name,age,city) SET id = UUID(), first_name = SUBSTRING_INDEX(@name, ' ', 1), second_name = SUBSTRING_INDEX(@name, ' ', -1);

explain select * from users where first_name like 'Абр%' and second_name like 'Ег%' order by id;

select count(*) from users where first_name like 'Абр%' and second_name like 'Ег%' order by id;

create index users_first_name_idx on users(first_name) using BTREE;
drop index users_first_name_idx on users;
create index users_second_name_idx on users(second_name) using BTREE;

create index users_name_idx on users(first_name,second_name) using BTREE;