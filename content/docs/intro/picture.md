---
weight: 3
title: "Get the Picture?"
---
## Get the Picture!

Here's a picture that puts everything in context.

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
                    |                                             |
  +-----------------+---------------------------------------------+-------------------------------+
  |                 |                                             |        Service Infrastructure |
  |                 |                                             v                               |
  |       .---------+----------.                         .-----------------.                      |
  |       |                    |                         |                 |                      |
  |       | Service Management +------------------------>| Service Control |                      |
  |       |                    |                         |                 |                      |
  |       .---------+----------.                         .--------+---+----.                      |
  |                 ^                                         ^   |   |                           |
  |                 |                          +--------------+   |   +---------------+           | Google
  |                 |                          |                  |                   |           |
  |                 |                          |                  v                   v           |
  |                 |                    .-----+-----.       .---------.        .------------.    |
  |                 |                    | API Keys  |       |  Cloud  |        |   Cloud    |    |
  |                 |                    |           |       | Logging |        | Monitoring |    |
  |                 |                    .-----------.       .----+----.        .-----+------.    |
  |                 |                          ^                  |                   |           |
  |                 |                          |                  |                   |           |
  +-----------------+--------------------------+------------------+-------------------+-----------+
                    |                          |                  |                   |
                    | Service                  | User             |                   |
                    | Configuration            | Authorizations   | Logs              | Metrics
                    |                          |                  v                   v 
                 o      
                -+- Service Owner
                / \   

```
If you're reading this, you're probably either a current or a future [Service Owner](/docs/glossary/#service-owner), and you have a group of [Service Users](/docs/glossary/#service-user) writing [API clients](/docs/glossary/#api-client) that call a service that you've implemented with an [API server](/docs/glossary/#api-server). But rather than implementing access controls, rate limiting, monitoring, etc. in your API server, you're delegating that to an [Extensible Service Proxy](/docs/glossary/#extensible-service-proxy) that you will deploy alongside your API server, typically using the same host, VM, or Kubernetes deployment. The proxy is supported by [Service Infrastructure](/docs/glossary/#service-infrastructure), which provides configuration and runtime support for request handling. Your [Service Deployment](/docs/glossary/#service-deployment) can be anywhere within network reach of [Google](/docs/glossary/#google), which runs everything in the Service Infrastructure box. You show up at the bottom of the picture again where you, the Service Owner, drive Service Infrastructure by providing configuration and creating API keys, and you follow the health of your service with logs and metrics.

Did that seem like a lot? This website makes this easy, and in the next section, the [quickstart](/docs/quickstart) will get you up-and-running with a sample service deployment on [Google Cloud Run](https://cloud.google.com/run).

## References
- https://cloud.google.com/service-infrastructure/docs/service-control/reference