---
weight: 1
title: Get the q Tool
---
## The q Tool

`q` is a command-line tool for calling the Service Infrastructure APIs. It can be used alongside or in place of [gcloud](https://cloud.google.com/sdk/gcloud). `q` is written in Go, and its code provides examples of Go usage of the Service Infrastructure APIs.

To get started with `q`, try the [demo](/docs/quickstart/demo). It provides scripts, support files, and instructions to allow you to quickly set up a managed gRPC service on Google Cloud Run, and you can do everything in the Google Cloud Shell.

Full source code for `q` is on GitHub at [github.com/agentio/q](https://github.com/agentio/q). A container that contains both `q` and `gcloud` is available at [ghcr.io/agentio/q:nightly](https://github.com/agentio/q/pkgs/container/q).

In many ways, `q` replicates capabilities that you'll find in `gcloud`, but with a closer correspondence to the underlying APIs. For every method of Service Infrastructure that we discuss, there is a `q` subcommand that allows us to call it directly, and Go source code is available to show how it works. 

For background on how `q` calls the Service Infrastructure APIs, see [How to Call Google APIs](/docs/details/how-to-call-google-apis).