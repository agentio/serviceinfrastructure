---
weight: 2
title: ESPv2
---
## ESPv2, The Extensible Service Proxy v2

ESPv2 is the latest proxy developed for [Cloud Endpoints](/cloud-endpoints).

ESPv2 is a custom build of Envoy with source available at [github.com/GoogleCloudPlatform/esp-v2](https://github.com/GoogleCloudPlatform/esp-v2).

## Running ESPv2

Official container images are available at `gcr.io/endpoints-release/endpoints-runtime:2`.

ESPv2 startup options are documented [here](https://cloud.google.com/endpoints/docs/grpc/specify-esp-v2-startup-options).

## Building ESPv2

[github.com/agentio/esp-v2](https://github.com/agentio/esp-v2) is a fork that includes a [Dockerfile](https://github.com/agentio/esp-v2/blob/master/Dockerfile) and configuration to use GitHub actions to build alternate container images that are published at [github.com/agentio/esp-v2/pkgs/container/esp-v2](https://github.com/agentio/esp-v2/pkgs/container/esp-v2). (these builds are currently broken, possibly due to [this issue](https://github.com/envoyproxy/envoy/issues/36650)) 

When they do work, ESPv2 builds take a long time -- builds using GitHub Actions can take several hours to complete. This is because each ESPv2 build includes a complete build of Envoy, which is itself notoriously slow to build (see [Why does Envoy take so long to compile?](https://www.envoyproxy.io/docs/envoy/latest/faq/build/speed)).

## Debugging ESPv2

One interesting thing to note is that in the process of [debugging a problem with ESP](/docs/proxies/esp#debugging-esp), we updated the service config by direct uploads that didn't include the file descriptor set of our API. After doing so, HTTP/JSON transcoding for ESPv2 no longer worked (though it did for ESP). What's the difference? ESP uses the service config to configure transcoding, but ESPv2 uses Envoy's [gRPC transcoding filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter), which requires the [file descriptor set](https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/descriptor.proto#L56) for configuration.

## Differences between ESP and ESPv2

See [Google documentation](https://cloud.google.com/endpoints/docs/openapi/migrate-to-esp-v2) for a list of things to consider when migrating from ESP to ESPv2. Some highlights are noted below:

- The results of [JWT verification](https://cloud.google.com/endpoints/docs/grpc/authenticating-users#receiving_authentication_results_in_your_api) are different.

- gRPC transcoding uses API descriptions in the service config in ESP and the uploaded file descriptor set in ESPv2. 

---
#### Continue with [Cloud Endpoints](/docs/endpoints).