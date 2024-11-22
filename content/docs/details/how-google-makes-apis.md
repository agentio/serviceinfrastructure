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
[ 5 domains with two APIs each]
[ 41 domains with one API each]
```

## Philosophy

### Make APIs managable by controlling their design

Google APIs are governed by a strict review and approval process. Reviewers are trained in expectations and best practices and are granted "API readability", which allows them to approve new APIs and API design changes.

This degree of governance is hard to enforce in any company, and particularly a large one. At Google it was possible because Google APIs are published through a central infrastructure, giving that central team the opportunity to establish and enforce requirements on APIs.

### Consistent APIs allow simpler, more powerful tools

Instead of buying a powerful API management system from a vendor (and since these systems were in their infancy), Google engineers built a series of API management systems that grew in capacity while pursuing simplicity. Over the years, hundreds of engineers were involved, and the system described here is a credit to their creativity and collaboration.

It is fair to say that the consistency of Google APIs had as important a part in the development of Google’s API management strategy as any of the code that was written.

## Methodology

A _methodology_ is a set of tools and best practices that an organization uses to address a particular class of problems. Here we focus on the problem of developing and operating APIs.

### Protocol Buffers

The most obvious feature of Google's API methodology is [Protocol Buffers](https://protobuf.dev/). Emerging and maturing over many years, we can think of Protocol Buffers as three distinct and complimentary things:

- A message serialization mechanism. Protocol Buffer encoding puts messages in a concise binary format that is relatively fast to serialize and deserialize, especially compared to its predecessor (XML) and it's most common current alternative (JSON).
- A language for describing APIs. Protocol Buffer descriptions of APIs are written in a dedicated language that is independent of the programming languages used to implement API servers and clients. This makes Protocol Buffers programming language-neutral, which has surely helped stabilize the description language, and the relative simplicity of this language has allowed Google and others to build many Protocol Buffer-based analysis and development tools.
- A methodology for implementing and using APIs. At least one of the tools that Google developed (`protoc`) is used by nearly every Protocol Buffer-based API design, and tools exist in many languages that generate support code for API servers and clients in those languages. Often those generators are written in the same programming languages that they generate, which allows language experts to easily improve the quality of the generated code. Many other tools exist for Protocol Buffer-based APIs, including documentation generators and API management systems, the subject of this website.

### gRPC

gRPC was [developed and first shared in 2015](https://grpc.io/blog/principles/), but it is based on practices that developed over many years inside Google that include an RPC framework called "Stubby" that sends Protocol Buffer-encoded messages over HTTP. One possibly non-obvious aspect of gRPC is that it is more than just a messaging protocol or a set of conventions for using HTTP/2. gRPC represents an explicit desire to put opinionated code in clients. This can be a bit of a shock to application developers ("what is this code you want me to put into my app?") but as we see with [gRPC service configuration](https://grpc.io/docs/guides/service-config/), implementing retry policies in clients is an important tool for improving reliability. HTTP/2 and Protocol Buffer encoding also both improve performance for clients and servers at the cost of adding complexity; mitigating that complexity with client side code is also "helping"... so gRPC is arguably helping both sides of the experience (servers and clients). And when this client-side code is reliable and easy-to-integrate (two nontrivial challenges), gRPC and Protocol Buffers can really simplify API consumption.

### Service Configuration

As we've discussed, [Service Configuration](/docs/serviceinfrastructure/serviceconfig) is the root of API description at Google. Every published Google API has a service config that describes its interface (in terms of protos) and specifies its operation. API documentation is also included in service config, which we can see in several examples in our public [compendium of Google service configs](https://serviceinfra.dev/docs/details/services/). It's ironic but unsurprising that Service Config is itself described with a [proto](https://github.com/googleapis/googleapis/blob/master/google/api/service.proto#L47); in a way, it's the [Inception](https://en.wikipedia.org/wiki/Inception) of API description at Google.

### A Protos Repo

It's well known that Google keeps all of its source code in a single repository. This means that all of its API descriptions are also in the "monorepo". This makes it easy for teams to find and learn about APIs, and it has created some patterns that other Protocol Buffer users would be wise to follow.

One aspect of having proto files in a monorepo is having them in a well-defined directory hierarchy. The [googleapis](/docs/details/googleapis) repo partially mirrors that, and is effectively a monorepo of Google API descriptions.

We can make some important observations:
- Proto files are best defined in a directory hierarchy because they are rarely used in isolation. Proto files often include other proto files, and the compiler needs to know where to find them.
- API version names are the last segment of the proto path.
- Third-party (non-Google) APIs should be organized in similar hierarchies that sit alongside the `google` directory. Don't put anything in `google`, or in other directories that are outside the scope of your organization. For example, we put the Boba Dojo APIs under `bobadojo` in  [github.com/bobadojo/apis](https://github.com/bobadojo/apis). This allows us to put \commonly-included standard protos alongside our API in a local copy of the `google` directory (the common protos in the `google` directory very rarely change).

### Generated Code

Protobuf and gRPC support code is almost always generated (actually, we've never seen a counterexample). This is *much* easier to do with a monorepo, because the organizational consistency that comes with that makes it much easier to write tools, and when the tools that you are using are written for a monorepo, it can be miserable to use them without one.

Products like the Buf Schema Registry make it easier to generate code without an explicit monorepo, but just having a "monorepo" of protos goes a long way. See [github.com/bobadojo/apis](https://github.com/bobadojo/apis) for an example protos repo and [github.com/bobadojo/go](https://github.com/bobadojo/go) for an example of a downstream repo containing generated code. Here we see examples of several layers of generated code:
- `rpc` code generated with `protoc-gen-go`. This is the data structures and serialization for Protocol Buffer messages.
- `grpc` code generated with `protoc-gen-go-grpc`. This is the service layer that makes gRPC requests. 
- `gapic` code generated with `protoc-gen-go_gapic`. This is a usability layer that wraps gRPC calls in idiomatic language-specific client libraries.
- a CLI generated with `protoc-gen-go_cli`. This is an API-specific CLI that builds on the Go `gapic`.

### Transcoding

Transcoding is a standard mechanism for making gRPC APIs available using HTTP/JSON. It predates gRPC and is "authoritatively" documented by the comments in [http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto), but there are other useful references:
- Google's [AIP-127](https://google.aip.dev/127)
- The Cloud Endpoints [Transcoding HTTP/JSON to gRPC](https://cloud.google.com/endpoints/docs/grpc/transcoding) documentation
- Envoy's [gRPC-JSON transcoder](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_json_transcoder_filter)
- Microsoft's [gRPC JSON transcoding in ASP.NET Core](https://learn.microsoft.com/en-us/aspnet/core/grpc/json-transcoding?view=aspnetcore-9.0)
- The independently-implemented [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).

### Client Libraries

To make it easier for people to use Google APIs (and to get a foothold for adding some client-side code), Google publishes API client libraries, and as we've noted, many of these are automatically-generated.

#### Discovery-based Client Libraries

The first generation of Google's generated client libraries were based on the API Discovery Service. These used JSON over HTTP (which was often provided by early transcoding support). Many of these are still available and continue to work. Google still maintains [a page documenting them](https://developers.google.com/api-client-library), and if you looked at it, you would think that these were still fully-supported and preferred by Google. These client libraries are open source and generated by open source tools, which generally are all available on GitHub in the These are also open source and available on GitHub in the [googleapis](https://github.com/googleapis) organization.

#### gRPC Client Libraries

But since the advent of gRPC, Google has preferred for developers to use gRPC to call Google APIs, and over time built a family of client libraries and code generators that produce them. These are also open source and available on GitHub in the [googleapis](https://github.com/googleapis) organization.

### The gcloud command-line interface

gcloud can call Google APIs and often provides higher-level capabilities... and that [--log-http](https://serviceinfra.dev/docs/details/how-to-call-google-apis/#seeing-inside-gcloud-with---log-http) flag is really something!

### AIPs and api-linter

Much of the above requires consistency! For this Google teams spent years internally developing standards that were first published as the Google [API design guide](https://cloud.google.com/apis/design) are now shared publicly as [API Improvement Proposals](https://aip.dev).

Although they are often waived for backward compatibility reasons, AIPs are checked for all new API designs at Google using the [API Linter](https://linter.aip.dev/), which is available for public use. As a demonstration, we run it in [bobadojo/apis/Makefile](https://github.com/bobadojo/apis/blob/main/Makefile).

To see the value of AIPs, consider pagination, as described in [AIP-158](https://google.aip.dev/158). By standardizing on an API interface for paginated APIs, AIP-158 allows client developers to build language-idiomatic interfaces that iterate over collections in all of the AIP-compliant APIs.

Long-running operations ([AIP-151](https://google.aip.dev/151)) is another API pattern that makes client libraries better, allowing them to use language-standard constructs like futures to work with asynchronous operations.

For more on "the Google School" of API design, see [API Design Patterns](https://www.manning.com/books/api-design-patterns) by JJ Geewax.

{{<figure 
   height="30%" width="30%" 
   src="/geewax.png" 
   link="https://www.manning.com/books/api-design-patterns"
>}}

---
#### Go back to [Going Deeper](/docs/details).