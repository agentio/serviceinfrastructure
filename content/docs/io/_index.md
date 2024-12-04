---
title: IO
type: docs
bookHidden: true
bookToC: true
bookSearchExclude: true

weight: 80
---
# Let's talk about IO

Here we move through a series of perspectives that leads to an architecture for API management for entities that we call "agents".

## API Management

In the [introduction](/docs/intro#api-proxies-and-gateways), we showed this general representation of API management.
```goat
.------------.     .---------------.                  .---------------.     .------------.
|            |     |               |                  |               |     |            |
| API client +---->| forward proxy +----------------->| reverse proxy +---->| API server |
|            |     |               |   API Request    |               |     |            |
.------------.     .---------------.                  .---------------.     .------------.
```

## API Serving

If we are writing an API server, we only need a reverse proxy.
```goat
                  .---------------.     .------------.
                  |               |     |            |
----------------->| reverse proxy +---->| API server |
   API Requests   |               |     |            |
                  .---------------.     .------------.
```

## API Calling

If we are writing an API client, we only need a forward proxy.
```goat
.------------.     .---------------.
|            |     |               |
| API client +---->| forward proxy +----------------->
|            |     |               |   API Requests
.------------.     .---------------.
```

## Introducing IO

But many applications need both. Let's call these applications "agents". Agents accept requests that come through a reverse proxy and make requests to other services using a forward proxy. Let's also suggest that there's a controller for an agent's proxies that we call `IO`. `IO`'s job is to coordinate all communication with the agent so that the agent only needs to do its unique job.

```goat
                  .---------------.     .-----------.     .---------------.
                  |               |     |           |     |               |
----------------->| reverse proxy +---->|   Agent   +---->| forward proxy +----------------->
   API Requests   |               |     |           |     |               |   API Requests  
                  .---------------.     .-----------.     .---------------.
                           ^                                      ^
                  .--------+--------------------------------------+-------.
                  +                          IO                           +
                  .-------------------------------------------------------.

```

## IO with a remote controller

It's likely that `IO` will need to communicate closely with the proxies. If it runs locally with them, it might call out to a centralized service for configuration and control.

```goat
                  .---------------.     .-----------.     .---------------.
                  |               |     |           |     |               |
----------------->| reverse proxy +---->|   Agent   +---->| forward proxy +----------------->
   API Requests   |               |     |           |     |               |   API Requests  
                  .---------------.     .-----------.     .---------------.
                           ^                                      ^
                           |               .-----.                |
                           +-------------->| IO  |<---------------+
                                           .-----.
                                              ^
                                              |
                                              v
                                     .-----------------.
                                     |  Agent Control  |
                                     .-----------------.
```

## Defining an Agent Data Plane

Let's adjust the picture slightly to separate our agent from a "data plane" that consists of the proxies and their local controller and a "control plane" that provides remote configuration and control.

```goat
                                        .-----------.
                                        |           |
                           +----------->|   Agent   +-------------+
                           |            |           |             |
                           |            .-----------.             |
                           |                                      |
                           |                                      v
                  .--------+------.                       .---------------.
                  |               |        .-----.        |               |
----------------->| reverse proxy +------->| IO  |<------>| forward proxy +----------------->
   API Requests   |               |        .-----.        |               |  API Requests  
                  .---------------.           ^           .---------------.
                                              |
                                              v
                                     .-----------------.
                                     |  Agent Control  |
                                     .-----------------.
```

## IO redrawn with Envoy

Now let's recognize that we can use Envoy for both proxies. Now our agent data plane has two layers: the Envoy proxy and its local `IO` controller.

```goat
                                        .-----------.
                                        |           |
                           +----------->|   Agent   +-------------+
                           |            |           |             |
                           |            .-----------.             |
                           |                                      |
                           |                                      v
                  .--------+----------------------------------------------.
                  |                                                       |
----------------->| reverse proxy           Envoy           forward proxy +----------------->
   API Requests   |                                                       |  API Requests  
                  +-------------------------------------------------------+
                                              ^
                                              |
                                              v
                                           +-----+
                                           | IO  |
                                           .--+--+
                                              ^
                                              |
                                              v
                                     .-----------------.
                                     |  Agent Control  |
                                     .-----------------.
```

This is just the sidecar model that is widely used in service meshes. It only looks a bit different here because we're zooming in to focus on agents so that we can pay close attention to their specific needs. This also allows us to manage communications of an agent independently of how any other agent is managed.

## Agent IO Concerns

When we are setting up an IO, what do we care about?

### The IO
- what is its version?
- how do we update it?

### The Agent Controller
- where is it?
- how do we authenticate to it?
- how do we register our identity with it?

### The Proxy
- is it running?
- how do we start and stop it?
- how do we update it?

### The Forward Interface, aka "Calling"
- what is its configuration?
- who is allowed to call it?
- what is it able to call?

### The Reverse Interface, aka "Serving"
- what is its configuration?
- who is allowed to call it?
- what is it able to call?



