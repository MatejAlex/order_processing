#!/usr/bin/env zsh
minikube start
kubectl config set-context --current --namespace=order-processor
