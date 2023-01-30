# Homework 2: K8S basics

Как проверить:

1. Применить манифесты
```
kubectl apply -f deployment.yaml -f service.yaml -f ingress.yaml
```
2. Проверить работу
```
➜  2-k8s-basics git:(2-k8s-basics) ✗ curl http://arch.homework/health
{"status": "OK"}%
➜  2-k8s-basics git:(2-k8s-basics) ✗
```
