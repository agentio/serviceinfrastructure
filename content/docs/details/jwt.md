---
weight: 6
title: Authenticating Users with JWTs
---
## Authenticating Users with JSON Web Tokens

While this is not used by any known Google APIs, the Cloud Endpoints product includes support for user authentication with [JSON Web Tokens](https://en.wikipedia.org/wiki/JSON_Web_Token). This uses standard mechanisms and ESPv2 uses the JWT Envoy filter, making this a fairly straightforward addition to the proxy that is configured with the [AuthProvider](https://github.com/googleapis/googleapis/blob/master/google/api/auth.proto#L107) message of Service Config.

For more details, see 
[Authenticating users](https://cloud.google.com/endpoints/docs/grpc/authenticating-users),
[Authentication between services (gRPC)](https://cloud.google.com/endpoints/docs/grpc/service-account-authentication),
and
[Authentication between services (OpenAPI)](https://cloud.google.com/endpoints/docs/grpc/service-account-authentication)
in the Cloud Endpoints documentation.

---
#### Go back to [Going Deeper](/docs/details).