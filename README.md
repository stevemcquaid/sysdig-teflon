# Teflon - Partner to Sysdig Falco
Teflon is a kubernetes module which integrates with Sysdig-Falco to take programmatic action when malicious activity is detected. The idea is that even if an intruder were to exploit an application, Falco would detect the event and in turn Teflon would be notified.  If specific criteria are met, Teflon then destroys the infected pod (while kubernetes handles re-starting the killed pod). In this strategy, by limiting the amount of time the exploit exists on our infrastructure, we are afforded more time to fix the underlying issue and limit the likelihood of the attacker's ability to cause harm.

By using this approach we are actually able to treat malicious exploitation as a very specific type of load on the system. I believe this is compelling because it allows cluster administrators to focus less on the one-off application vulnerability and instead focus on the bigger picture of improving response times and the triage/mitigation process as a whole. As long as new code is being written, software vulnerabilities will continue to be created. So by treating the problem as load on a system, attention can be paid to the level of active threats on the system and engineering resources can be better appropriated.  

## Teflon Goals:
  1) Smoothly integrate with Sysdig Falco
  2) Expose metrics so that malicious exploitation can be treated as a function of load
  3) Destroy & recreate infected pods while maintaining service uptime

# Solution
  - Kubernetes HPA using a custom-defined metric coming from sysdig monitor. Implement a custom metrics server and configurable autoscaler.
  - The Custom-defined metrics are created when recieving the event from falco alerts. (They are custom falco.yaml program_output sending a request to the teflon service.) Teflon parses the requests, takes the necessary action, and expose the metrics as a prometheus-compatible endpoint.
  - Utilize autosploit to simulate a real-world attack on a vulnerable docker image to show live demo about hacking rate as function of system load. Scale out the number of autosploit containers to simulate multiple attackers.

### Components Required
  - TEFLON
    - API to receive, parse, process falco alerts
    - Delete infected pods
      - Check for available available replicas to avoid downtime
    - Metrics endpoint to monitor rate of infection/reaping
  - HPA to scale if under active attack
    - Possible Trigger Metrics (Goal = Have more pod replicas than hackers can infect)
      - Number of infections
      - Number of reapings
      - % available endpoints vs reaped ones
      - Time to new infection
      - Time to recreation
  - Custom-metrics-apiserver
    - Consumed by HPA to determine whether or not to scale

