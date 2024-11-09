---
weight: 3
title: "Get the Picture?"
---
# Get the Picture

Here's a picture that puts everything in context:

```goat

                                                   .------------------------------------------.
                                                   |                                          |
  .-------------.                                  |    .--------------.     .------------.   |
  |             |                                  |    |  Extensible  |     |            |   |  o  Service
  | API clients +----------------------------------+--->|   Service    +---->| API server |   | -+-  Owner
  |             |           API Requests           |    |    Proxy     |     |            |   | / \
  .-------------.                                  |    .---------+----.     .------------.   |
     o   o   o                                     |         ^    |                           |
    -+- -+- -+-                                    |         |    |         Service Deployment|
    / \ / \ / \                                    .---------+----+---------------------------.
   Service Users                                             |    |
                                                             |    |
                                                             |    | Proxy Operation:
                              Proxy Configuration            |    | - Check
                    +----------------------------------------+    | - Allocate Quota,                
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
  |                 |                    |  API Keys |       |  Cloud  |        |   Cloud    |    |
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
If you're reading this, you're probably either a current or a future **Service Owner**, and you have a group of **Service Users** writing **API clients** that call a service that you've implemented with an **API server**. But rather than implement access controls, rate limiting, monitoring, etc. in your **API server**, you're delegating that to a **Extensible Service Proxy** that you will deploy alongside your **API server**, typically using the same host, VM, or Kubernetes deployment. The proxy is supported by **Service Infrastructure**, which provides configuration and runtime support for request handling. **Google** runs everything in the **Service Infrastructure** box, but you show up at the bottom of the picture again where you, the **Service Owner**, drive **Service Infrastructure** by providing configuration and creating API keys, and you follow the health of your service with logs and metrics.

Did that seem like a lot? This website makes this easy, and in the next section, the [quickstart](/docs/quickstart) will get you up-and-running with a sample service deployment on [Google Cloud Run](https://cloud.google.com/run).