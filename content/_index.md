---
title: Introduction
type: docs
bookToC: true
---

![Seattle Infrastructure by Marsh Gardiner, 2018](/images/seattle-infrastructure.jpg)

This website describes a way to manage APIs that is inexpensive, high-performing, and surprisingly easy to use.

It's also great for APIs that are built with [gRPC](https://grpc.io) and [Protocol Buffers](https://protobuf.dev). It's the only API management system that starts with these powerful approaches, and if you are building APIs with them, you'll find it especially familiar.

It emerged from internal API management practices at Google, and as you might hope, it works great on Google Cloud, but it's not limited to that. We'll show you how it can be used with or without Kubernetes on any cloud platform and on systems that you run yourself.

We start with a review of [API Management](/docs/intro) and discuss what's special about [gRPC API Management](/docs/intro/grpc). Then we look at Google's [Service Infrastructure](/docs/serviceinfrastructure) and the [Extensible Service Proxies](/docs/proxies). These are the key parts of our API management solution.

You can get hands-on experience using [q](/docs/quickstart/q), a command-line tool that helps us configure and demonstrate Service Infrastructure-based API management. `q` includes a [demo](/docs/quickstart/demo) that Google Cloud users can use to quickly get a sample managed service running on [Cloud Run](https://cloud.google.com/run). Try it!

Thanks for reading, and please post any feedback to [our issues page](https://github.com/agentio/serviceinfrastructure/issues).

This is an expanded version of a [talk that was presented on October 28](https://timburks.me/2024/10/28/managing-grpc-apis). If you'd like to see it with your team, [get in touch!](https://linkedin.com/in/timburks)

---
#### Let's get started with [Managing APIs](/docs/intro).