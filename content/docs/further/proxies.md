---
weight: 4
title: Better Proxies
---
# Better Proxies

From the experiences we described in the [proxies](/doc/proxies) section, it appears that the condition of the two ESP proxies is the weakest link in the Service Infrastructure experience. They are buggy, difficult to build, and lightly maintained.

Superficially, we also notice that the proxies are in two separate GitHub organizations. ESP is in [cloudendpoints/esp](https://github.com/cloudendpoints/esp) and ESPv2 is in [googlecloudplatform/esp-v2](https://github.com/googlecloudplatform/esp-v2).

Google publishes builds for these proxies, but what if we need an update sooner? We can use GitHub actions to build them, but there are some significant problems:
- The ESP builds use archaic build environments and tools.
- The ESPv2 builds are fragile and extremely slow.

The ESPv2 builds are slow because ESPv2 is a custom build of Envoy. This might have been a good idea at the time, but since 2017, Envoy has matured significantly (despite the code bloat), adding powerful new APIs for configuration and API management:

- [xDS](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/operations/dynamic_configuration) configures routes, upstream targets, and filters (request processors)
- [ext_authz](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_authz_filter) makes unary gRPC requests to a helper process to authorize requests
- [ext_proc](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/ext_proc_filter) makes streaming gRPC requests to a helper to monitor and control request handling

Envoy continues to provide a useful collection of filters that ESPv2 uses and that can be used without a custom build:
- [jwt_authn](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/jwt_authn_filter)
- [grpc_json_transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- [grpc_web](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_web_filter)

We could also explore building proxies on other tools such as [caddy](https://caddyserver.com/) or even building a proxy from scratch.

However, Envoy has 
[lots of usage](https://www.envoyproxy.io/),
[a security-hardened data plane](https://www.envoyproxy.io/docs/envoy/latest/start/quick-start/securing),
and [many features](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/http_filters) that could give us a significant headstart on the next generation of Service Infrastructure-based proxies.

