#!/bin/bash

# this is your part to fill
# minikube addons enable ingress

# IMPORTANT:
# Error handling was not implemented but for real scenario cases such deploy file should take it into account 

eval $(minikube docker-env)

# building docker images 
for app in invoice-app payment-provider
do
  cd $app
  docker build -t $app:challenge .
  cd ..
done

# Create namespaces
# namespace for simple deployment with 'kubectl apply -f' command
kubectl create ns simple
# Deploy invoice-app
echo -n "Deploy invoice-app -> " && kubectl apply -f invoice-app/deployment.yaml -n simple
# Deploy payment-provider
echo -n "Deploy payment-provider -> " && kubectl apply -f payment-provider/deployment.yaml -n simple

# namespace for best practise deployment using helm charts 
kubectl create ns development
# Deploy to another namespace using helm charts:
for app in invoice-app payment-provider; do cd $app/chart && helm upgrade --install $app  . -f envs/dev.yaml --namespace development; cd -; done

minikube service invoice-app-service --url --namespace development &
