from prometheus_client import Counter, CollectorRegistry,make_asgi_app,Counter,GC_COLLECTOR,PLATFORM_COLLECTOR,PROCESS_COLLECTOR
from fastapi import FastAPI
import uvicorn

registry = CollectorRegistry()
counter = Counter(name="ping_request_count",documentation="No. of requests handled a Ping handler")
[registry.register(collector) for collector in [GC_COLLECTOR,PLATFORM_COLLECTOR,PROCESS_COLLECTOR,counter]]
exporter = make_asgi_app(registry)

app = FastAPI(debug=False)
app.mount("/metrics",exporter)
@app.get("/ping")
def handle_ping():
    counter.inc()
    return "pong"


if __name__ == "__main__":
  uvicorn.run(app,host="127.0.0.1",port=8080)