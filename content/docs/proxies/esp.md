---
weight: 1
title: ESP
bookToC: false
---
## ESP

ESP, the "Extensible Service Proxy", is the original proxy developed for [Cloud Endpoints](/cloud-endpoints).

ESP is a custom build of nginx with source available at [github.com/cloudendpoints/esp](https://github.com/cloudendpoints/esp).

### Running ESP

Official container images are available at `gcr.io/endpoints-release/endpoints-runtime:1`.

ESP startup options are documented [here](https://cloud.google.com/endpoints/docs/grpc/specify-proxy-startup-options).

### Building ESP

[github.com/agentio/esp](https://github.com/agentio/esp) is a fork that includes a [Dockerfile](https://github.com/agentio/esp/blob/master/Dockerfile) and configuration to use GitHub actions to build alternate container images that are published at [github.com/agentio/esp/pkgs/container/esp](https://github.com/agentio/esp/pkgs/container/esp).

Note that ESP is lightly maintained, with just [a few recent commits](https://github.com/cloudendpoints/esp/commits/master/). We can see that a [recent commit](https://github.com/cloudendpoints/esp/commit/73ae1115ae1a47f6d74269e0219497f44a278f6c#diff-288b593fcb4342d79d4a324f301eb1ea8e4b84d4fed01288c3facd07fee84abdR134) addressed the problem that we noticed of API keys appearing in the service logs. [This later commit](https://github.com/cloudendpoints/esp/commit/4b377082d3f50ad1be94d1b27b41976064c12ed9) made it the default. We still saw the API keys because we were using ESPv2.

You'll also see that the Dockerfile builds ESP on Ubuntu 16.04. This follows the [build instructions in the main repo](https://github.com/cloudendpoints/esp/blob/master/doc/build-esp-on-ubuntu-16-04.md). It also uses a very old version of Bazel (0.21.0) (Bazel is now [past 9.0.0](https://bazel.build/release/rolling)). This makes building and debugging ESP challenging.

### Debugging ESP

Possibly the easiest way to modify ESP and debug it is to modify the source locally and use `docker build` to build a local container image. Noting the error message that we saw when we had problems running ESP with our demo API, we find that it occurs once in the source code [here](https://github.com/cloudendpoints/esp/blob/1a01fe7e15152e4fd35dbb56f6b3ca829fce229b/src/api_manager/config.cc#L431):

```
bool Config::LoadService(ApiManagerEnvInterface *env,
                         const std::string &service_config) {
  if (!service_config.empty()) {
    if (!ReadConfigFromString(service_config, &service_)) {
      env->LogError("Cannot load ESP configuration protocol buffer.");
      return false;
    }
```

To verify that we're getting the correct service config, we print the service config as part of our error reporting, adding one line to the above code:
```

bool Config::LoadService(ApiManagerEnvInterface *env,
                         const std::string &service_config) {
  if (!service_config.empty()) {
    if (!ReadConfigFromString(service_config, &service_)) {
      env->LogError(service_config);
      env->LogError("Cannot load ESP configuration protocol buffer.");
      return false;
    }
```

After building a container locally with `docker build` and modifying our script to use this local container, we see that the proxy is getting the service config that we expect:
```
$ sh LOCAL-espv1.sh 
INFO:Constructing an access token with scope https://www.googleapis.com/auth/service.management.readonly
INFO:Service account email: stores@bobadojo.iam.gserviceaccount.com
INFO:Refreshing access_token
INFO:Fetching the service config ID from the rollouts service
INFO:Fetching the service configuration from the service management service
nginx: [warn] Using trusted CA certificates file: /etc/nginx/trusted-ca-certificates.crt
2024/11/19 22:41:19[error]1#1: {
  "apis": [
    {
      "methods": [
        {
          "name": "ListStores",
          "options": [
            {
              "name": "google.api.http",
              "value": {
                "@type": "type.googleapis.com/google.api.HttpRule",
                "get": "/v1/stores"
              }
            }
          ],
          "requestTypeUrl": "type.googleapis.com/bobadojo.stores.v1.ListStoresRequest",
          "responseTypeUrl": "type.googleapis.com/bobadojo.stores.v1.ListStoresResponse"
        },
        ... details deleted ...
      ],
      "name": "bobadojo.stores.v1.Stores",
      "sourceContext": {
        "fileName": "bobadojo/stores/v1/stores.proto"
      },
      "syntax": "SYNTAX_PROTO3",
      "version": "v1"
    }
  ],
  "authentication": {},
  "backend": {
    "rules": [
      {
        "selector": "bobadojo.stores.v1.Stores.ListStores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.FindStores"
      },
      {
        "selector": "bobadojo.stores.v1.Stores.GetStore"
      }
    ]
  },
  "configVersion": 3,
  "control": {
    "environment": "servicecontrol.googleapis.com"
  },
  ... more details deleted...
}
2024/11/19 22:41:19[error]1#1: Cannot load ESP configuration protocol buffer.
nginx: [emerg] Failed to load service configuration files in /etc/nginx/endpoints/nginx.conf:112
nginx: [emerg] There were errors with Endpoints api service configuration. 
```

The `ReadConfigFromString` function appears to be failing to parse this service config, which we see is JSON, and we can guess that there's something in this service config that is causing it to fail.

We can get the service config with `q service-management get-service-config`, and we can experiment by deleting selective parts of it and uploading them with `q service-management create-service-config`. Each time we upload a new config, we need to roll it out (`q service-management create-service-rollout` will do that). After some trial-and-error and bisection, we find that the failed parsing is due to the presence of two options:

```
$ diff original.config working.config 
125,138d123
<       "options": [
<         {
<           "name": "google.api.resource",
<           "value": {
<             "@type": "type.googleapis.com/google.api.ResourceDescriptor",
<             "type": "stores.bobadojo.io/Store",
<             "pattern": [
<               "stores/{store}"
<             ],
<             "plural": "stores",
<             "singular": "store"
<           }
<         }
<       ],
359,365d343
<               }
<             },
<             {
<               "name": "google.api.resource_reference",
<               "value": {
<                 "@type": "type.googleapis.com/google.api.ResourceReference",
<                 "type": "stores.bobadojo.com/Store"
```

These options are in the `types` section of the service config and aren't used by either proxy. They were introduced by one of Google's API design standards, [AIP-122](https://google.aip.dev/122), long after the creation of ESP and typically don't appear in APIs managed by Endpoints.

Arguably, they should be omitted from the service config by service compilation, but the proto parser in ESP should also be less brittle.

If we decide to work seriously with ESP, we will eventually want to address this. One way to do so would be just remove these options from our original source protos, but in case someone wants to fix this in ESP, we've filed [this issue on the ESP repo](https://github.com/cloudendpoints/esp/issues/876).

---
#### Continue with [ESPv2](/docs/proxies/espv2).