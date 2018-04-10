#!/bin/bash

set -x

kubectl delete -f k8s/deployment.yaml
kubectl delete -f k8s/svc.yaml
kubectl delete -f k8s/ingress.yaml
