# ДЗ 3. Репликация

## Roadmap

- Create network for VMs
- Create 2 VMs for master-slave
- Install docker on these VMs
- Create overlay network
- Install MySQL 5.7 InnoDB on them
- Create master-slave replication
- Load data into master database
- Check replicated data on the slave
- Configure app to use slave for reading
- Deploy app to the cloud
- Check operability of the service
- Add monitoring master / slave
- Start YCloud load testing

## Links

https://severalnines.com/blog/introduction-docker-swarm-mode-and-multi-host-networking/
https://docs.docker.com/engine/swarm/swarm-tutorial/create-swarm/

https://www.digitalocean.com/community/tutorials/how-to-set-up-replication-in-mysql

## Commands

docker run --name=mysql1 --mount type=bind,src=/home/admin/mysql/my.cnf,dst=/etc/my.cnf --mount type=bind,src=/home/admin/mysql/datadir,dst=/var/lib/mysql -d --network=host mysql/mysql-server:8.0.25

docker stop mysql1 && docker rm mysql1 && sudo rm -fr datadir && mkdir datadir

ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';

CHANGE REPLICATION SOURCE TO
SOURCE_HOST='10.129.0.9',
SOURCE_USER='replica_user',
SOURCE_PASSWORD='password',
SOURCE_LOG_FILE='mysql-bin.000003',
SOURCE_LOG_POS=1571;


