---
weight: 2
title: Better Deployments
---
# Better Deployments

Proxies were originally deployed separately on serverless platforms (GAE, Cloud Run)

Setup was (is) complicated

Google automated proxy deployment in the API Gateway product

Kubernetes deployments can run the proxies as internal sidecars

Cloud Run (i.e. Knative) also now allows proxies to be run as sidecars

...so we don’t need (or want) API Gateway for Cloud Run