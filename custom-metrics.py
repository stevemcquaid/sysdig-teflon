from flask import Flask, request
import os
from sdcclient import SdcClient

app = Flask(__name__)
sdclient = SdcClient(open("/etc/sysdigtoken/token.txt","r").read().rstrip())
metrics = [{ "id": "net.http.request.count", "aggregations": { "time": "timeAvg", "group": "avg"  }  }]
podname = os.environ['POD_NAME']

def get_request_time():
    podfilter = "kubernetes.pod.name = '%s'" % podname
    metricdata = sdclient.get_data(metrics, -60, 0, 60, podfilter)
    request_time = metricdata[1].get('data')[0].get('d')[0]
    return request_time

@app.route('/')
def handle_page():
    return "Hello from python\n"

@app.route('/metrics')
def handle_metrics():
    return "# TYPE http_request_count gauge\nhttp_request_count %s\n" % get_request_time()

if __name__ == '__main__':
        app.run(app.run(host='0.0.0.0', port=8080))

