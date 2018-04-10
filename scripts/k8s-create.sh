#!/bin/bash

set -x

kubectl create -f k8s/deployment.yaml
kubectl create -f k8s/svc.yaml
kubectl create -f k8s/ingress.yaml
