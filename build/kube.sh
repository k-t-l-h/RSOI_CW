#!/bin/bash

docker-compose up
#kompose convert #превращает докер-файл в список сервисов
#kubectl create deployment auth --image=docker.io/ktlh/auth
#kubectl expose deployment auth --type=LoadBalancer --port=8080
#kubectl apply -f auth.yaml
#kubectl get deployments hello-node
#kubectl describe deployments hello-node
#kubectl apply -f ingress.yaml
#minikube service hello-node