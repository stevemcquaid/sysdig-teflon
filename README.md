# SysDig CRD ideas
  - CRD to define monitors/alerts
  - CRD to define security rule / falco
  - deploy sysdig
    - helm chart
    - operator
    - deploy falco
  - CRD to create dashboard

# TODO
  - [ ] Research sysdig
  - [ ] Research Falco
  - [ ] Read blog posts
    - [ ] Istio stuff?
    - [ ] Prometheus
    - [ ] Custom monitor annotations/definitions
      - [ ] Operator/monitor
  - [ ] Document Sysdig value props & competitive advantages/SWOT?
  - [ ] Compare with grafana & datadog
  - [ ] priv sidecar on deployment with annotation instead of daemon set
  - [ ] what are falco's value prop and competitive advantages?
  - [ ] Where is falco going?
  - [ ] where could it go?
  - [ ] turn this all into a checklist
  - [ ] Falco -f rules.yaml -f rules2.yaml -> kubectl create -f falco-rules-crd.yaml
  - [ ] Falco helm chart
  - [ ] Falco operator might be better for this than a crd, no CRD is better
  - [ ] Pros/Cons of CRD vs operator
  - [ ] Figure out how to reload falco or have falco sigup/load new rulesets
  - [ ] https://github.com/draios/falco/blob/dev/examples/k8s-using-daemonset/k8s-without-rbac/falco-daemonset.yaml
  - [ ] Create helm chart (determine if rbac or not) for deploying falco
  - [ ] look into ksonnet for helm chart generation
  - [ ] config maps for multiple falco rules files
  - [ ] https://sysdig.com/blog/runtime-security-kubernetes-sysdig-falco/
  - [ ] Get falco chart into kubernetes charts repo
  - [ ] Build falco -> alert manager integration
    - [ ] Build falco prometheus metrics endpoint
      - [ ] https://github.com/draios/falco/wiki/Falco-Alerts
        - [ ] syslog
            - [ ] metrics endpoint as a syslog parser?
                - https://github.com/fstab/grok_exporter
                - Need to deal with separate logic in the grok'er, parsing syslog, and format changes
        - [ ] file
            - [ ] shared file/volume, monitor and parse
                - Need to deal with file buffers, rotation, and locks
        - [ ] stdout
            - [ ] logging aggregator parser/alerter
                - Need to deal with separate logic in the log parser, and format changes
        - [ ] spawned program
            - [ ] REST API for metrics server
            - Need to worry about network stability, stable api, and connecting to api
    - [ ] Build falco prometheus collector/alertmanager alerts
  - [ ] Self healing security via falco
    - [ ] reaper
      - reap on infection, hpa to expand cluster?
      - [ ] falco -> prometheus -> alert manager -> reaper webhook
      - [ ] falco -> reaper sidecar to destroy infected pods
      - [x] falco -> webhook to reaper api to delectively destroy infected pods
        - simple, known architecture
        - must have parsing logic and alerting definitions
        - implement metrics endpoint via prom / sysdig monitor to show number of falco detections & pod reapings
        - 
    - [ ] HPA that via custom metrics api
    -
    -

# Option A
You should implement a Kubernetes HPA using metrics coming from Sysdig Monitor. Use Sysdig API to get this metrics and implement a custom metrics server and a configurable autoscaler. Try to bring this as far as your time allows:

Golang implementation
Defensive code
Kubernetes friendly configuration
Dockerize it and document how to deploy/use with Kubernetes
Upload to GitHub and make it contribution friendly (README.md, PRs policy, etc)

Hint: https://sysdig.com/blog/kubernetes-scaler/

