# Distributed Tracing with OpenTelemetry (Go + JavaScript on AWS Lambda)

This project demonstrates end-to-end distributed tracing using OpenTelemetry across multiple AWS Lambda services written in Go and JavaScript.

It shows how traces propagate across services, how context is passed via HTTP headers, and how telemetry is exported to a local Jaeger backend.

---

## Architecture Overview

Go Lambda (parent span)  
→ HTTP request with trace context (`traceparent`)  
→ JavaScript Lambda (child span)  
→ OpenTelemetry Collector  
→ Jaeger (trace visualization)

Local Jaeger is exposed to AWS using Ngrok.

---

## Services

### 1. Go Lambda (Parent Service)

- Creates root span
- Injects trace context into outgoing HTTP request
- Calls downstream service

Built using Go

---

### 2. JavaScript Lambda (Downstream Service)

- Extracts trace context from incoming request
- Creates child span linked to parent trace
- Sends telemetry via OpenTelemetry exporter

Built using JavaScript

---

### 3. Observability Stack

- Jaeger for trace visualization
- Ngrok to expose local Jaeger endpoint to AWS
- OpenTelemetry Collector for receiving OTLP traces

---

## Infrastructure

Provisioned using AWS CDK:

- Two AWS Lambda functions
- Function URLs (no API Gateway used)
- IAM roles for logging and tracing permissions

---

## Key Concepts Demonstrated

- Distributed tracing with OpenTelemetry
- Context propagation using `traceparent`
- Parent-child span relationships
- Cross-language tracing (Go ↔ JavaScript)
- OTLP export pipeline
- Cloud-to-local observability bridging

---

## Running the Project

1. Deploy infrastructure using CDK
2. Start OpenTelemetry Collector locally
3. Start Jaeger backend
4. Expose Jaeger using Ngrok
5. Invoke Go Lambda endpoint
6. View traces in Jaeger UI

---

## Result

Single distributed trace across:
Go Lambda → JavaScript Lambda → Jaeger

---

## Future Improvements

- Add more services to extend trace depth
- Enable auto-instrumentation
- Integrate AWS X-Ray
- Add log correlation with traces

---
