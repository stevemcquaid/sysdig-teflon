# SysDig Challenge Ideas
  - CRD to define monitors/alerts
  - CRD to define security rule / falco
  - Deploy mechanisms for sysdig / falco
    - helm chart
    - operator
  - CRD to create dashboards/falco rules
  - API to receive falco alerts
    - Delete infected pods
  - Instrumentation/Metrics from falco alerts
  - HPA based upon falco alerts

# TODO
  - [X] Research Sysdig
  - [X] Research Falco
  - [X] Read blog posts
    - [X] Istio stuff
    - [X] Prometheus
    - [X] Falco + K8S
    - [X] Custom prometheus monitor annotations/definitions
  - [X] Figure out how to reload falco or have falco sigup/load new rulesets
  - [ ] Get falco helm chart into kubernetes charts repo
  - [X] Figure out how to expose Falco Alerts to an API - "TEFLON Project"
    - [X] https://github.com/draios/falco/wiki/Falco-Alerts
    - [X] Methods
      - [ ] Standard Output
      - [ ] File
      - [ ] Syslog
      - [X] Spawned program
        - Most reliable, easy to use, and scalable
    - [X] How to parse Falco json in golang
  - [X] Figure out how to expose metrics of these Falco alerts
    - Prometheus Metrics Endpoint - https://blog.alexellis.io/prometheus-monitoring/
    - Sysdig Monitor Ingestion
  - [X] Figure out how to create HPA to scale if under attack
    - https://sysdig.com/blog/kubernetes-scaler/
  - [ ] Start building the different parts of the system
    - 

# Tasks
  - [ ] Ensure Cluster Compatibility
    - [ ] Check if `--horizontal-pod-autoscaler-use-rest-clients` is set on kube-controller-manager
  - [ ] Get target deployment to scale
    - https://github.com/mateobur/kubernetes-scaler-metrics-api
    - This will eventually be the same as the input to the Teflon service
  - [ ] Create a custom-metrics-apiserver - produce a: `MetricValueList`, with `apiVersion`: `custom.metrics.k8s.io/v1beta1`
    - https://github.com/kubernetes-incubator/custom-metrics-apiserver
    - https://github.com/mateobur/kubernetes-scaler-metrics-api
    - `kubectl create -f kubernetes-scaler-metrics-api/scaler/custom-metrics-ns-rbac.yaml`
    - Add in the sysdig API KEY
    - `kubectl create -f kubernetes-scaler-metrics-api/scaler/custom-metric-api-sysdig.yaml -n custom-metrics`
    - Preview the metrics available:
      - `kubectl create clusterrolebinding cli-api-read-access --clusterrole custom-metrics-getter --user system:anonymous`
      - `curl -sSk https://10.96.0.1/apis/custom.metrics.k8s.io/v1beta1/namespaces/default/services/flask/net.http.request.count`
  - [ ] Create Horizontal Pod Autoscaler
      ```bash
    kind: HorizontalPodAutoscaler
    apiVersion: autoscaling/v2beta1
    metadata:
      name: flask-hpa
    spec:
      scaleTargetRef:
        kind: Deployment
        name: flask
      minReplicas: 2
      maxReplicas: 10
      metrics:
      - type: Object
        object:
          target:
            kind: Service
            name: flask
          metricName: net.http.request.count
          targetValue: 100
    ```
    - `kubectl create -f kubernetes-scaler-metrics-api/scaler/horizontalpodautoscaler.yaml`
    - `kubectl describe hpa`



# Archive
  - [ ] Create HPA ingesting the metrics from Teflon
  - [ ] Self healing security via falco
    - [ ] reaper
      - reap on infection, hpa to expand if under attack
      - [X] falco -> webhook to reaper api to selectively destroy infected pods
        - simple, known architecture
        - must have parsing logic and alerting definitions
        - implement metrics endpoint via prom / sysdig monitor to show number of falco detections & pod reapings



# Option A
You should implement a Kubernetes HPA using metrics coming from Sysdig Monitor. Use Sysdig API to get this metrics and implement a custom metrics server and a configurable autoscaler. Try to bring this as far as your time allows:

Golang implementation
Defensive code
Kubernetes friendly configuration
Dockerize it and document how to deploy/use with Kubernetes
Upload to GitHub and make it contribution friendly (README.md, PRs policy, etc)

Hint: https://sysdig.com/blog/kubernetes-scaler/
# Idea
  - Kubernetes HPA using a custom-defined metric coming from sysdig monitor. Implement a custom metrics server and configurable autoscaler.
  - The Custom-defined metric is a function of falco event alerts, ingested via a custom falco.yaml program_output sending to a custom service to parse and expose as a prometheus endpoint
  -
  -
