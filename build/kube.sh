#!/bin/bash

kubectl create deployment balanced auth --image=docker.io/ktlh/auth
kubectl expose deployment balanced auth --type=LoadBalancer --port=8080
minikube service auth