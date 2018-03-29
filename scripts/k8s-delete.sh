#!/bin/bash

set -ex

kubectl delete  -f k8s/deployment.yaml
kubectl delete  -f k8s/svc.yaml
