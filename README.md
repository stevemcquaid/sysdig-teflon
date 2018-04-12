# SysDig Challenge Abstract
  - Destroy & Recreate infected pods
  - Instrument rate of reaping
  - Scale services that are under attack

# Solution
  - Kubernetes HPA using a custom-defined metric coming from sysdig monitor. Implement a custom metrics server and configurable autoscaler.
  - The Custom-defined metrics are created when recieving the event from falco alerts. (They are custom falco.yaml program_output sending a request to the teflon service.) Teflon parses the requests, takes the necessary action, and expose the metrics as a prometheus-compatible endpoint.
  - If time allows, use autosploit and a vulnerable docker image to show live demo about hacking rate as function of system load. Scale out the number of autosploit containers to simulate multiple attacks.

### Components Required
  -  TEFLON
    - API to receive, parse, process falco alerts
    - Delete infected pods
      - Check for available available replicas to avoid downtime
    -  Metrics endpoint to monitor rate of infection/reaping
  -  HPA to scale if under active attack
    - Possible Trigger Metrics (Goal = Have more pod replicas than hackers can infect):
      -  Number of infections
      -  Number of reapings
      -  % available endpoints vs reaped ones
      -  Time to new infection
      -  Time to recreation
  -  Custom-metrics-apiserver
    - Consumed by HPA to determine whether or not to scale
    - 

