---
weight: 3
title: How Google Makes APIs
---
## How Google Makes APIs

Looking at things that can be observed externally, we can learn a lot about how Google makes APIs and about large-scale API design in general. 

As I wrote in [I got a golden ticket: what I learned about APIs in my first year at Google](https://medium.com/apis-and-digital-transformation/i-got-a-golden-ticket-what-i-learned-about-apis-in-my-first-year-at-google-556e1f02f9ab), Google makes a lot of APIs and Google APIs get called a lot. These two dimensions of scale, in development and operations, have had a big influence on API practice at Google and led to things that might not be obvious to someone using an online tutorial to code their first API (which is a great first step!)

## Scale

### How many APIs does Google have?

To get an idea of how many APIs Google produces, we could start by just asking Google directly by calling the [API Discovery Service](https://developers.google.com/discovery/v1/reference/apis/list)

```
$ curl -s https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[].name' -r | wc -l
463
```

This is a list of API versions. If we pare the list down to unique APIs, we get a smaller list, but still it's large!
```
$ curl -s https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[].name' -r | sort | uniq | wc -l
286
```

If you're curious, here's a quick way to list all of the API-version pairs.
```
$ curl https://discovery.googleapis.com/discovery/v1/apis | jq '.items.[] | "\(.name) \(.version)"' -r
```

### How many APIs are being managed with Service Infrastructure?

Outside of Google, there's no way to know how many APIs are managed with Service Infrastructure (since the majority of them are private), but we can look at the number of public APIs that are listed. We can get this by listing the managed services with no project specified. Here we'll get them, filter out their names, put them in a file, and count them:
```
$ q service-management list-services "" | jq .[].serviceName -r > SERVICES

$ wc -l SERVICES
6308 SERVICES
```

If we count the number of APIs with each domain name, we get the following:
```
3737 cloud.goog
2010 cloudpartnerservices.goog
476 googleapis.com
32 appspot.com
4 mongodb.com
2 shpconn.com
2 run.app
2 netapp.com
2 floreysoft.net
1 zype.com
1 zdatainc.com
1 zazmic.io
1 xmatters.com
1 wowza.com
1 vogsy.io
1 ucraft-gcloud.com
1 trifacta.com
1 tackle.io
1 superquery.io
1 resiot.io
1 redisenterprise.com
1 qubole.com
1 plugon.us
1 perimeterx.com
1 omnixlabs.com
1 neuton.ai
1 neo4j.io
1 miclient.in
1 mavatar.com
1 managedkube.com
1 kafkaesque.io
1 itopia.com
1 inventurist.com
1 informaticacloud.com
1 hycu.com
1 humanapi.co
1 headspin.io
1 harness.io
1 gwapps.com
1 gpcloudtest.com
1 gliacloud.com
1 fivetran.com
1 elastic.co
1 datastax.com
1 confluent.cloud
1 buzzboard.com
1 blueantoinette.com
1 bkper.com
1 bizstats.ai
1 appranix.com
```

## Philosophy

### Make APIs managable by controlling their design

Google APIs are governed by a strict review and approval process. Reviewers are trained in expectations and best practices and are granted "API readability", which allows them to approve new APIs and API design changes.

This degree of governance is hard to enforce in any company, and particularly a large one. At Google it was possible because Google APIs are published through a central infrastructure, giving that central team the opportunity to establish and enforce requirements on APIs.

### Consistent APIs allow simpler, more powerful tools

Instead of buying a powerful API management system from a vendor (and since these systems were in their infancy), Google engineers built a series of API management systems that grew in capacity while pursuing simplicity. Over the years, hundreds of engineers were involved, and the system described here is a credit to their creativity and collaboration.

It is fair to say that the consistency of Google APIs had as important a part in the development of Google’s API management strategy as any of the code that was written.

## Methodology

A methodology is a set of tools and best practices that an organization uses to address a particular class of problems. Here we focus on the problem of developing and operating APIs.

### Protocol Buffers

The most obvious feature of Google's API methodology is Protocol Buffers. Emerging and maturing over many years, we can think of Protocol Buffers as three distinct and complimentary things:

- A message serialization mechanism. Protocol Buffer encoding puts messages in a concise binary format that is relatively fast to serialize and deserialize, especially compared to its predecessor (XML) and it's most common current alternative (JSON).
- A language for describing APIs. Protocol Buffer descriptions of APIs are written in a dedicated language that is independent of the programming languages used to implement API servers and clients. This makes Protocol Buffers programming language-neutral, which has surely helped stabilize the description language, and the relative simplicity of this language has allowed Google and others to build many Protocol Buffer-based analysis and development tools.
- A methodology for implementing and using APIs. At least one of the tools that Google developed (`protoc`) is used by nearly every Protocol Buffer-based API design, and tools exist in many languages that generate support code for API servers and clients in those languages. Often those generators are written in the same programming languages that they generate, which allows language experts to easily improve the quality of the generated code. Many other tools exist for Protocol Buffer-based APIs, including documentation generators and API management systems, the subject of this website.

### A Protos Repo

It's well known that Google keeps all of its source code in a single repository. This means that all of its API descriptions are also in the "monorepo". This makes it easy for teams to find and learn about APIs, and it has created some patterns that other Protocol Buffer users would be wise to follow.

One aspect of having proto files in a monorepo is having them in a well-defined directory hierarchy. The [googleapis](/docs/details/googleapis) repo partially mirrors that, and is effectively a monorepo of Google API descriptions.

We can make some important observations:
- Proto files are best defined in a directory hierarchy because they are rarely used in isolation. Proto files often include other proto files, and the compiler needs to know where to find them.
- API version names are the last segment of the proto path.
- Third-party (non-Google) APIs should be organized in similar hierarchies that sit alongside the `google` directory. Don't put anything in `google`, or in other directories that are outside the scope of your organization. For example, we put the Boba Dojo APIs under `bobadojo` in  [github.com/bobadojo/apis](https://github.com/bobadojo/apis). This allows us to put standard commonly-included protos alongside our API in a local copy of the `google` directory (the common protos in the `google` directory very rarely change).

### Generated Code

Protobuf and gRPC support code is almost always generated (actually, we've never seen a counterexample).

See [github.com/bobadojo/go](https://github.com/bobadojo/go) for an example repo containing code generated from a repo of protos.

### gRPC

An RPC framework for HTTP/2 (and beyond) that works great with protocol buffers and Google APIs.

Why does Google prefer gRPC?
- performance
- client-side influence (retry)

### Transcoding

Transcoding is a standard mechanism for making gRPC APIs available using HTTP/JSON.

https://cloud.google.com/endpoints/docs/grpc/transcoding

https://google.aip.dev/127

### Client libraries

Google provides client libraries for most Google APIs.

(Aside: Google doesn't do this for all APIs, but people are trying!)

#### Discovery-based Client Libraries

Google API Discovery Service

#### gRPC Client Libraries

http://github.com/googleapis/googleapis

### The gcloud command-line interface

gcloud can call Google APIs and often provides higher-level capabilities.

### AIPs and api-linter

Much of the above requires consistency! For this Google teams spent years internally developing standards that are now shared publicly as AIPs.

https://cloud.google.com/apis/design

https://aip.dev

https://linter.aip.dev/

To see the value of AIPs, consider pagination, as described in [AIP-158](https://google.aip.dev/158).

For more on "the Google school" of API design, see [API Design Patterns](https://www.manning.com/books/api-design-patterns) by JJ Geewax.

---
#### Go back to [Going Deeper](/docs/details).