const client = require("prom-client");
const express = require("express");

const registry = new client.Registry();
client.collectDefaultMetrics({ register: registry });
const counter = new client.Counter({
  name: "ping_request_count",
  help: "No of requests handled a Ping handler",
});

function handlePing(req, res) {
  counter.inc();
  res.send("pong");
}

function main() {
  const app = express();
  app.all("/ping", handlePing);
  app.get("/metrics", async (req, res) => {
    const metrics = await registry.metrics();
    res.set("Content-Type", "text/plain");
    res.send(metrics);
  });

  const port = 8080;
  app.listen(port, () => console.log("node js server running on port", port));
}

main();
