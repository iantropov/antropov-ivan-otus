# ДЗ 4. K8S

## Описание

Выполнить указанные команды
Проверить работу сервиса через curl или через Postman-а (коллекция для локальной версии и в кубере лежит рядом)

```
minikube start --cpus=6 --memory=6g --vm-driver=hyperkit

helm repo add bitnami https://charts.bitnami.com/bitnami
helm install mysql-minikube -f k8s/services/mysql-values.yaml bitnami/mysql

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
helm install nginx -f k8s/services/nginx-ingress.yaml ingress-nginx/ingress-nginx

cd k8s/app
kubectl apply -f configmap.yaml -f secrets.yaml
kubectl apply -f deployemnt.yaml -f service.yaml -f ingress.yaml
kubectl apply -f job.yaml

➜  app git:(microservices-4-k8s) ✗ curl arch.homework/user/1
{"username":"johndoe589","firstName":"John","lastName":"Doe","email":"bestjohn@doe.com","phone":"+71002003040","id":1}
```

## Roadmap

+ Create users app
+ Connect with DB
+ Debug app with docker-compose
+ Install DB from helm
+ Add Configmap
+ Add Job
+ Add Postman

## Commands (for myself)

```
minikube start --cpus=6 --memory=6g --vm-driver=hyperkit

helm repo add mysql-repo https://charts.bitnami.com/bitnami
helm install mysql-minikube -f mysql-values.yaml mysql-repo/mysql
helm uninstall mysql-minikube

MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace default mysql-release -o jsonpath="{.data.mysql-root-password}" | base64 -d)
kubectl run mysql-release-client --rm --tty -i --restart='Never' --image  docker.io/bitnami/mysql:8.0.32-debian-11-r0 --namespace default --command -- bash

kubectl create namespace m && helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/ && helm repo update && helm install nginx ingress-nginx/ingress-nginx --namespace m -f nginx-ingress.yaml

```