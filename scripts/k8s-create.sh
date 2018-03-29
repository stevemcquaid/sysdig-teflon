#!/bin/bash

set -ex

kubectl create -f k8s/deployment.yaml
kubectl create -f k8s/svc.yaml
