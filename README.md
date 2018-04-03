# SysDig Challenge Abstract
  - Destroy & Recreate infected pods
  - Instrument rate of reaping
  - Scale services that are under attack

## TODO
  - [ ] Implement dummy endpoint for "Kubernetes custom metric apiserver spec"
  - [ ] Test http request body falcojson parsing
  - [ ] Test k8s delete pod
  - [ ] Setup code coverage
  - [ ] Setup HPA resource (test metic apiserver spec)
  - [ ] Separate custom metric apiserver to pull from sysdig
  - [ ] E2E automation?

### Components Required
  - [ ] TEFLON
    - [X] API to receive, parse, process falco alerts
    - [X] Delete infected pods
      - [ ] Check for available available replicas avoid downtime
    - [ ] Metrics endpoint to monitor rate of infection/reaping
  - [ ] HPA to scale if under active attack
    - Possible Trigger Metrics (Goal = Have more replicas than hacker can infect):
      - [ ] Number of infections
      - [ ] Number of reapings
      - [ ] % available endpoints vs reaped ones
      - [ ] Time to new infection
      -[ ] Time to recreation
  - [ ] Custom-metrics-apiserver
    - Unsure if this can be same as metrics endpoint
    - Consumed by HPA to determine whether or not to scale

# TODO
  - [X] Figure out how to expose Falco Alerts to an API - "TEFLON Project"
    - [X] How to parse Falco json in golang
  - [X] Figure out how to expose metrics of these Falco alerts
    - Prometheus Metrics Endpoint - https://blog.alexellis.io/prometheus-monitoring/
    - Sysdig Monitor Ingestion
  - [X] Figure out how to create HPA to scale if under attack
    - https://sysdig.com/blog/kubernetes-scaler/

# Tasks to create HPA
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

# Description of system architecture
  - Create HPA ingesting the metrics from Teflon
  - Self healing security via falco
    - reaper
      - reap on infection, hpa to expand if under attack
      - falco -> webhook to reaper api to selectively destroy infected pods
        - simple, known architecture
        - must have parsing logic and alerting definitions
        - implement metrics endpoint via prom / sysdig monitor to show number of falco detections & pod reapings

# Option A
You should implement a Kubernetes HPA using metrics coming from Sysdig Monitor. Use Sysdig API to get this metrics and implement a custom metrics server and a configurable autoscaler. Try to bring this as far as your time allows:
  - Golang implementation
  - Defensive code
  - Kubernetes friendly configuration
  - Dockerize it and document how to deploy/use with Kubernetes
  - Upload to GitHub and make it contribution friendly (README.md, PRs policy, etc)

Hint: https://sysdig.com/blog/kubernetes-scaler/
# Idea
  - Kubernetes HPA using a custom-defined metric coming from sysdig monitor. Implement a custom metrics server and configurable autoscaler.
  - The Custom-defined metric is a function of falco event alerts, ingested via a custom falco.yaml program_output sending to a custom service to parse and expose as a prometheus endpoint
  - Use autosploit to give live demo about hacking rate as function of system load
