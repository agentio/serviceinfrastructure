---
weight: 2
title: Better Deployments
---
# Better Deployments

Proxies were originally deployed separately on serverless platforms (GAE, Cloud Run)

Setup was (is) complicated

For example, [Set up Cloud Endpoints gRPC for Cloud Run with ESPv2](https://cloud.google.com/endpoints/docs/grpc/set-up-cloud-run-espv2) includes several steps that overall make it excruciatingly painful:
- getting an endpoints hostname for the proxy by deploying a dummy instance. This appears to only be used to name the service.
- building a new container with docker. This appears to be because Cloud Run didn't allow file mounts
- deploying the proxy to a separate cloud run instance from the backend and granting the proxy access to the backend

Google automated proxy deployment in the API Gateway product

Kubernetes deployments can run the proxies as internal sidecars

Cloud Run (i.e. Knative) also now allows proxies to be run as sidecars

...so we don’t need (or want) API Gateway for Cloud Run