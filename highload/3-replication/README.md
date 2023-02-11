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
+ Improve application to use variables from ENV (in addition to .env file)
+ Deploy app to the cloud
+ Configure app to use slave for reading
+ Check operability of the service
+ Add monitoring master / slave
+ Use JMeter load testing
+ Try Async replication
+ Try Semi-sync replication

## Configuration

otus-vm-1 - master
otus-vm-2 - replica (async) (semi-sync + GTID)
otus-vm-3 - replica (semi-sync + GTID)
otus-social-network-3 - serverless container with the service otus

## Links

https://severalnines.com/blog/introduction-docker-swarm-mode-and-multi-host-networking/
https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/

https://www.digitalocean.com/community/tutorials/how-to-set-up-replication-in-mysql

https://www.redhat.com/sysadmin/gtid-replication-mysql-servers

https://dev.mysql.com/doc/refman/8.0/en/replication-formats.html

https://dev.mysql.com/doc/refman/8.0/en/replication-semisync-installation.html

https://www.percona.com/blog/how-to-createrestore-a-slave-using-gtid-replication-in-mysql-5-6/

## Commands

### Containers

```
docker stop mysql1 && docker rm mysql1 && sudo rm -fr mysql/datadir && mkdir mysql/datadir

docker stop mysql1 && docker rm mysql1

docker run --name=mysql1 --mount type=bind,src=/home/admin/mysql/my.cnf,dst=/etc/my.cnf --mount type=bind,src=/home/admin/mysql/datadir,dst=/var/lib/mysql -d --network=host mysql/mysql-server:8.0.25

docker logs mysql1 -f

docker exec -it mysql1 /bin/bash

docker restart mysql1

ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';
```

### Replication (async)

```
CREATE USER 'replica_user'@'10.129.0.22' IDENTIFIED WITH mysql_native_password BY 'password';
GRANT REPLICATION SLAVE ON *.* TO 'replica_user'@'10.129.0.22';
FLUSH PRIVILEGES;

FLUSH TABLES WITH READ LOCK;
SHOW MASTER STATUS;
UNLOCK TABLES;

CHANGE REPLICATION SOURCE TO
SOURCE_HOST='10.129.0.22',
SOURCE_USER='replica_user',
SOURCE_PASSWORD='password',
SOURCE_LOG_FILE='mysql-bin.000003',
SOURCE_LOG_POS=123377;

START REPLICA;

STOP REPLICA;

SHOW REPLICA STATUS\G

CREATE DATABASE otus;

CREATE TABLE example_table (
example_column varchar(30)
);

INSERT INTO example_table VALUES
('This is the first row'),
('This is the second row'),
('This is the third row');

USE otus;
SHOW TABLES;
SELECT * FROM example_table;
```

### Replication (semi-sync + GTID)

```
SET @@GLOBAL.read_only = ON;

gtid_mode=ON
enforce-gtid-consistency=ON


gtid_mode=ON
enforce-gtid-consistency=ON
log-replica-updates=ON
skip-replica-start=ON

SET @@GLOBAL.read_only = OFF;

INSTALL PLUGIN rpl_semi_sync_master SONAME 'semisync_master.so';
INSTALL PLUGIN rpl_semi_sync_slave SONAME 'semisync_slave.so';

SELECT PLUGIN_NAME, PLUGIN_STATUS
       FROM INFORMATION_SCHEMA.PLUGINS
       WHERE PLUGIN_NAME LIKE '%semi%';

SET GLOBAL rpl_semi_sync_master_enabled = 1;
SET GLOBAL rpl_semi_sync_slave_enabled = 1;

CHANGE REPLICATION SOURCE TO
SOURCE_HOST='10.129.0.22',
SOURCE_USER='replica_user_2',
SOURCE_PASSWORD='password',
SOURCE_AUTO_POSITION=1;

STOP SLAVE IO_THREAD;
START SLAVE IO_THREAD;
```

### Diagnostics

```
SELECT host, user FROM mysql.user;

show variables like 'character%';

show variables like 'binlog%';
```

### Load data

```
mysql -u root --local_infile=1 -p
set global local_infile=true;
LOAD DATA LOCAL INFILE './people.csv' INTO TABLE users CHARACTER SET utf8 FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' (@name,age,city) SET id = UUID(), first_name = SUBSTRING_INDEX(@name, ' ', 1), second_name = SUBSTRING_INDEX(@name, ' ', -1);

CREATE TABLE users (id VARCHAR(36) NOT NULL, first_name VARCHAR(128) NOT NULL, second_name VARCHAR(128) NOT NULL, age INT NOT NULL, password VARCHAR(128), biography VARCHAR(255), city VARCHAR(64), PRIMARY KEY (`id`));

create index users_name_idx on users(second_name,first_name) using BTREE;
```

### Remote connection

```
sudo ufw allow 3306

CREATE USER 'service_user'@'%' IDENTIFIED WITH mysql_native_password BY 'password';
GRANT ALL ON *.* TO 'service_user'@'%';

docker login --username oauth --password ******** cr.yandex

docker run -e PORT=8080 -e DB_ADDR=158.160.13.236:3306 -e DB_NAME=otus -e DB_USER=service_user -e DB_PASS=password --network=host cr.yandex/crprmec70gabitr67n5f/social-network-3
```
