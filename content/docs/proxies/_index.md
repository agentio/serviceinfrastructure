---
weight: 4
bookFlatSection: true
title: "The Proxies"
---
# The Extensible Service Proxies

The Extensible Service Proxies run alongside API servers and are configured and controlled by [Service Infrastructure](/docs/serviceinfrastructure). Reviewing the simplified picture below, we see that these proxies sit alongside API server implementations and take traffic from API clients and check and log it using Service Infrastructure. In this way, the Extensible Service Proxy handles the general-purpose API management needs of the server, allowing the server implementation to focus on its own unique features. The proxy isn't fully self-contained; instead it calls out to Service Infrastructure for configuration, access control, and logging. This makes it lightweight and convenient to deploy alongside an API server.

```goat

                                                   .------------------------------------------.
                                                   |                                          |
  .-------------.                                  |    .--------------.     .------------.   |
  |             |                                  |    |  Extensible  |     |            |   |  o  Service
  | API clients +----------------------------------+--->|   Service    +---->| API server |   | -+-  Owner
  |             |           API Requests           |    |    Proxy     |     |            |   | / \
  .-------------.                                  |    .---------+----.     .------------.   |
     o   o   o                                     |         ^    |                           |
    -+- -+- -+-                                    |         |    |        Service Deployment |
    / \ / \ / \                                    .---------+----+---------------------------.
   Service Users                                             |    |
                                                             |    |
                                                             |    | Request Handling:
                              Proxy Configuration            |    | - Check
                    +----------------------------------------+    | - Allocate Quota
                    |                                             | - Report
                    |                                             |
                    |                                             v
  +-----------------+---------------------------------------------+-------------------------------+
  |                                                                        Service Infrastructure |
  |                                                                                               | Google
  |                                                                                               |
  +-----------------+--------------------------+------------------+-------------------+-----------+
                    ^                          ^                  |                   |
                    | Service                  | User             |                   |
                    | Configuration            | Authorizations   | Logs              | Metrics
                    |                          |                  v                   v 
                 o      
                -+- Service Owner
                / \   

```

Often the ESP and API server are built into their own containers that are deployed together in a single Kubernetes or Cloud Run deployment. This isn't the only way to use an ESP; it could also be run elsewhere and handle requests for multiple backends. But because it is lightweight, and because we generally like to minimize the network delay between the ESP and the API server, we prefer this [sidecar deployment](/docs/glossary#sidecar-deployment).

There are two Google-supported versions of these proxies:
1. [ESP](/docs/proxies/esp), a custom proxy based on nginx.
2. [ESPv2](/docs/proxies/espv2), a custom proxy based on [Envoy](/docs/details/envoy).
## Troubleshooting

See "Troubleshooting Cloud Endpoints configuration deployment" [grpc](https://cloud.google.com/endpoints/docs/grpc/troubleshoot-config-deployment) and [openapi](https://cloud.google.com/endpoints/docs/openapi/troubleshoot-config-deployment) for help from Google for troubleshooting proxy deployments.

## Running the proxies in Docker

For testing and experimentation, it's useful to run the proxies locally. Unfortunately, their implementations are a bit too complex to easily run them directly, but we can use Docker to run them in their containers.

Let's try it for our demo API.

First we'll run our `stores-server` container locally with Docker:
```prompt
docker run -p 8080:8080 ghcr.io/bobadojo/stores-server
```
```
2024/11/19 03:04:20 listening on port 8080
```

### ESPv2

Now, leaving the `stores-server` container running, we run one of the proxies (ESP v2), also in Docker. We'll do it from a script:
```
#!/bin/sh

docker run \
	-v /home/tim/Desktop/bobadojo/stores-server/kubernetes/key.json:/key.json \
	-p 8081:8081 \
	-p 19000:19000 \
	gcr.io/endpoints-release/endpoints-runtime:2 \
	--listener_port=8081 \
	--admin_port=19000 \
	--backend=grpc://192.168.86.200:8080 \
	--service=stores.endpoints.bobadojo.cloud.goog \
	--rollout_strategy=managed \
	--non_gcp \
	--service_account_key=/key.json 
```

Run this script to run the ESP v2 proxy. Note that it contains some things that you'll need to customize.
- You should have a service account key in `key.json`.
- You should fix the path to your local version `key.json`.
- You should set the backend address to the address of the machine running your `stores-server` container.
- You should set the --service option to use the name of your own service.

With this, you should be able to access your API on port 8081 of your local machine. Also, be sure to check out port 19000, the [Envoy administration interface](https://www.envoyproxy.io/docs/envoy/latest/operations/admin).

When you're finished, kill this container with CTRL-C.

### ESP

We can also run the original proxy (ESP) in Docker with slightly different command-line options.
```
#!/bin/sh

docker run \
	-v /home/tim/Desktop/bobadojo/stores-server/kubernetes/key.json:/key.json \
	-p 8081:8081 \
	-p 19000:19000 \
	gcr.io/endpoints-release/endpoints-runtime:1 \
	--http_port=8081 \
	--backend=grpc://192.168.86.200:8080 \
	--service=stores.endpoints.bobadojo.cloud.goog \
	--rollout_strategy=managed \
	--service_account_key=/key.json 
```

But when you run it, something interesting and surprising might happen. The proxy might exit.

```prompt
sh local-espv1.sh
```
``` 
INFO:Constructing an access token with scope https://www.googleapis.com/auth/service.management.readonly
INFO:Service account email: stores@bobadojo.iam.gserviceaccount.com
INFO:Refreshing access_token
INFO:Fetching the service config ID from the rollouts service
INFO:Fetching the service configuration from the service management service
nginx: [warn] Using trusted CA certificates file: /etc/nginx/trusted-ca-certificates.crt
2024/11/19 22:00:02[error]1#1: Cannot load ESP configuration protocol buffer.
nginx: [emerg] Failed to load service configuration files in /etc/nginx/endpoints/nginx.conf:112
nginx: [emerg] There were errors with Endpoints api service configuration.
```

Uh, what? It appears that ESP thinks there's something wrong with our service configuration, but this same configuration just worked with ESP v2. We'll come back to this when we discuss ESP in detail.

---
#### Continue with [ESP](/docs/proxies/esp).