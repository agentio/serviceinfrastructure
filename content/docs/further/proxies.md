---
weight: 4
title: Better Proxies
---
# Better Proxies

The proxies are in separate GitHub orgs… what’s the story?

The proxies require security updates… who can we expect to do this?

Google publishes builds, but what if we need an update sooner?

We might build our proxies with GitHub Actions

- ESPv1 - archaic build environment and unclear demand
- ESPv2 - huge build times

Why are ESPv2 builds so slow?...why are Envoy builds so slow?...

Why are we rebuilding Envoy?

New-ish Envoy APIs allow us to manage APIs with standard Envoy plus a helper

[xDS](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration) configures routes, upstream targets, and filters (request processors)

[ext_authz](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter) makes unary gRPC requests to a helper process to authorize requests

[ext_proc](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_proc_filter) makes streaming gRPC requests to a helper to monitor and control request handling

Existing Envoy filters can be easily plugged in
- [jwt_authn](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/jwt_authn_filter)
- [grpc_json_transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- [grpc_web](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter)

Envoy has 
[lots of usage](https://www.envoyproxy.io/),
[a security-hardened data plane](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/securing),
and [lots of features](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/http_filters).

