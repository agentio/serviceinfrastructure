---
weight: 2
title: Better Deployments
---
# Better Deployments

When the Cloud Endpoints product first appeared, the proxies were deployed in separate containers to work with services running on serverless platforms like Google App Engine and Cloud Run. Setup was (and still is) complicated, especially when users follow the recommended quickstarts.

For example, [Set up Cloud Endpoints gRPC for Cloud Run with ESPv2](https://cloud.google.com/endpoints/docs/grpc/set-up-cloud-run-espv2) includes several steps that overall make it excruciatingly painful:
- Users are told to reserve an endpoints hostname for their proxy by deploying a dummy instance to Cloud Run. This appears to only be used to name the service, and we've seen that this isn't necessary. Services on service infrastructure can be named independently of the name of the server where they run.
- Users are told to use Docker to build a new container containing their ESPv2 proxy. This is said to be necessary to reduce the number of calls that the proxy makes to Service Management, but for anyone starting with Endpoints, this is a non-issue and an ugly barrier to first usage.
- Cloud Run users are told to deploy their proxy to a separate Cloud Run instance from their backend, forcing them to deal with configuration to control the connection from their proxy to their backend server.

To address this, Google automated proxy deployment in the [API Gateway](https://cloud.google.com/api-gateway/docs) product, but took control of their proxy images away from users and continued to deploy proxies independently of their services. Also, the API gateway product is currently only available in [ten GCP regions](https://cloud.google.com/api-gateway/docs/deployment-model).

Meanwhile, the Cloud Run product advanced to allow [sidecar deployments](https://cloud.google.com/blog/products/serverless/cloud-run-now-supports-multi-container-deployments), which we demonstrate for our demo example. With this, it's possible to deploy a service and a proxy side-by-side in a single deployment with a [unified YAML configuration file](/docs/quickstart/files#serviceyaml). With this, the proxies conveniently blend into service deployments and require no explicit management, and their logs and metrics are available, as always, in the Cloud Endpoints console. This is a better developer experience than what API Gateway offers and a strong benefit to using Cloud Run. Also, these Cloud Run deployments are very similar to proxy deployments in Kubernetes, making it easy to move from Cloud Run to more general Kubernetes platforms.

---
#### Continue with [Better Configuration](/docs/further/configuration).