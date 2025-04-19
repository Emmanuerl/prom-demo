# 📈 Introduction to Prometheus

**Prometheus** is an open-source systems monitoring and alerting toolkit originally developed at SoundCloud. It is now a part of the Cloud Native Computing Foundation (CNCF), and widely used in the DevOps ecosystem.

---

## 🔍 What is Prometheus?

Prometheus is designed to:

- Collect **metrics** from configured targets at given intervals
- Store all data as time series
- Evaluate rule expressions\*\*
- Display results or trigger alerts

A time series is a stream of timestamped values belonging to the same metric and set of labels.

---

## 🔄 How Prometheus Works

```plaintext
+-------------+          +------------------+         +----------------+
| Applications|  ----->  |   Prometheus     | ----->  |  Grafana (UI)  |
| /metrics    |  (Pull)  |  Server          |         +----------------+
+-------------+          | - Scrapes targets|
                         | - Stores data    |
                         | - Exposes UI/API |
                         +------------------+
```

---

- **Multi-dimensional data model** using key/value labels
- **Powerful query language**: PromQL (Prometheus Query Language)
- **Pull model** over HTTP for collecting metrics
- **Built-in time series database**
- **Service discovery** or static target configuration
- **Alerting system** with integration to tools like Alertmanager, Slack, email, etc.
- **Easy integration** with Grafana for dashboards

---

## Metrics AND Labels

Metric - a particular quantity to be measures
Label - qualifiers for a quantity

I want to count the number of devs in this room (metric) who write js and python (labels)

---

## Prom Datatypes - The ones that count

Counter - Quanties that only increase, or reset to zero
Guage - Quantities that increase or decrease

others are Histogram, and Summary

---

# Demo Time 🚀
