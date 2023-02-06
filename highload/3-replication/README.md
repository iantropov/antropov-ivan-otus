# ДЗ 3. Репликация

## Roadmap

+ Create network for VMs
+ Create 2 VMs for master-slave
+ Install docker on these VMs
+ Create overlay network
+ Install MySQL 5.7 InnoDB on them
+ Create master-slave replication
+ Load data into master database
+ Check replicated data on the slave

- Improve application to use variables from ENV (in addition to .env file)
- Deploy app to the cloud
- Configure app to use slave for reading
- Check operability of the service
- Add monitoring master / slave
- Start YCloud load testing

## Links

https://severalnines.com/blog/introduction-docker-swarm-mode-and-multi-host-networking/
https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/

https://www.digitalocean.com/community/tutorials/how-to-set-up-replication-in-mysql

## Commands

docker stop mysql1 && docker rm mysql1 && sudo rm -fr mysql/datadir && mkdir mysql/datadir

docker run --name=mysql1 --mount type=bind,src=/home/admin/mysql/my.cnf,dst=/etc/my.cnf --mount type=bind,src=/home/admin/mysql/datadir,dst=/var/lib/mysql -d --network=host mysql/mysql-server:8.0.25

docker logs mysql1 -f

docker restart mysql1

ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';

sudo ufw allow 3306

CREATE USER 'replica_user'@'10.129.0.29' IDENTIFIED WITH mysql_native_password BY 'password';
GRANT REPLICATION SLAVE ON *.* TO 'replica_user'@'10.129.0.29';
FLUSH PRIVILEGES;

FLUSH TABLES WITH READ LOCK;
SHOW MASTER STATUS;
UNLOCK TABLES;

CREATE DATABASE otus;

mysql -u root --local_infile=1 -p
set global local_infile=true;
LOAD DATA LOCAL INFILE './people.csv' INTO TABLE users CHARACTER SET utf8 FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 ROWS (@name,age,city) SET id = UUID(), first_name = SUBSTRING_INDEX(@name, ' ', 1), second_name = SUBSTRING_INDEX(@name, ' ', -1);

SELECT host, user FROM mysql.user;

CHANGE REPLICATION SOURCE TO
SOURCE_HOST='10.129.0.9',
SOURCE_USER='replica_user',
SOURCE_PASSWORD='password',
SOURCE_LOG_FILE='mysql-bin.000003',
SOURCE_LOG_POS=187015073;

START REPLICA;

SHOW REPLICA STATUS\G;

CREATE TABLE example_table (
example_column varchar(30)
);

INSERT INTO example_table VALUES
('This is the first row'),
('This is the second row'),
('This is the third row');

show variables like 'character%';

USE otus;
SHOW TABLES;
SELECT * FROM example_table;


CREATE TABLE users (id VARCHAR(36) NOT NULL, first_name VARCHAR(128) NOT NULL, second_name VARCHAR(128) NOT NULL, age INT NOT NULL, password VARCHAR(128), biography VARCHAR(255), city VARCHAR(64), PRIMARY KEY (`id`));