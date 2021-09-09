#!/bin/bash

kubectl create deployment auth --image=docker.io/ktlh/auth
kubectl expose deployment auth --type=LoadBalancer --port=8080
kubectl apply -f auth.yaml
kubectl get deployments hello-node
kubectl describe deployments hello-node
minikube service hello-node