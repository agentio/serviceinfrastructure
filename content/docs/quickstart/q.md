---
weight: 1
title: The q Tool
---
## The q Tool

`q` is a command-line tool for calling the Service Infrastructure APIs that can be used alongside or in place of [gcloud](https://cloud.google.com/sdk/gcloud). In many ways, `q` replicates capabilities that you'll find in `gcloud`, but with a closer correspondence to the underlying APIs. For every method of Service Infrastructure that we discuss, there is a `q` subcommand that allows us to call it directly, and source code is available to show how it works. 

`q` is written in Go, and its code is on GitHub at [github.com/agentio/q](https://github.com/agentio/q). A container that contains both `q` and `gcloud` is available at [ghcr.io/agentio/q:nightly](https://github.com/agentio/q/pkgs/container/q).

To get started with `q`, try the demo in the next section. It provides scripts, support files, and instructions to allow you to quickly set up a managed gRPC service on Google Cloud Run, and you can do everything in the Google Cloud Shell or your own terminal, if you prefer.

For background on how `q` calls the Service Infrastructure APIs, see [How to Call Google APIs](/docs/details/how-to-call-google-apis).

---
#### Continue with [Demo: Deploy a Sample Service](/docs/quickstart/demo).