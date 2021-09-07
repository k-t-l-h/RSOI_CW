#!/bin/bash

kubectl create deployment auth --image=docker.io/ktlh/auth
kubectl expose deployment auth --type=LoadBalancer --port=8080
minikube service auth