---
title: IO Planning
type: docs
bookHidden: true
bookToC: true
bookSearchExclude: true

weight: 10
---
# Planning Agent IO

Here we move through a series of perspectives that leads to an architecture for API management for entities that we call "agents".

## API Management

In the [introduction](/docs/intro#api-proxies-and-gateways), we showed this general representation of API management. There we said that forward proxies help people (API client developers) make requests and reverse proxies help people (API server developers) serve requests.

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

## Agents

But some applications need both kinds of proxies. Let's call these applications "agents". Agents accept requests from clients and send their own (different) requests to servers. Internally, they may be very complex and powerful, but at their boundaries, agents are just accepting and making requests. Without proxies, it looks like this:

```goat
.-------------.              .-----------.              .-------------.
|             |              |           |              |             |
| API clients +------------->|   Agent   +------------->| API servers |
|             | API Requests |           | API Requests |             |  
.-------------.              .-----------.              .-------------.
```

## Agents and Proxies

Returning to our discussion of proxies, we can simplify agents by using forward and reverse proxies to manage their communications. Using proxies, an agent accepts requests that come through a reverse proxy and makes requests to other services using a forward proxy. This moves the complexity of communication out of the agent so that the it can focus on its unique purpose.

```goat
                  .---------------.     .-----------.     .---------------.
                  |               |     |           |     |               |
----------------->| reverse proxy +---->|   Agent   +---->| forward proxy +----------------->
   API Requests   |               |     |           |     |               |   API Requests  
                  .---------------.     .-----------.     .---------------.
```                  

## Introducing IO

This next picture suggests that there's a controller for an agent's proxies that we call `IO`. `IO`'s job is to configure and control the proxies, which now can be standard components that we use off-the-shelf, like [Envoy](https://envoyproxy.io). Together, IO and the proxies take over and handle all of the agent's communication.

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
```

## IO with a remote controller

It's likely that `IO` will need to communicate closely with the proxies that it controls. If it runs locally with them, it might call out to a remote, centralized service for its own configuration and control.

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
                                     |   IO Control    |
                                     .-----------------.
```

## Defining Agent, IO, and Control Planes

Let's adjust the picture slightly to separate our agent from an "IO plane" that consists of the proxies and their local controller and a "control plane" that provides remote configuration and control.

```goat
                                        .-----------.                            Agent Plane
                                        |           |
                           +----------->|   Agent   +-------------+
                           |            |           |             |
                           |            .-----------.             |
---------------------------+--------------------------------------+--------------------------
                           |                                      |                 IO Plane
                           |                                      v
                  .--------+------.                       .---------------.
                  |               |        .-----.        |               |
----------------->| reverse proxy +------->| IO  |<------>| forward proxy +----------------->
   API Requests   |               |        .-----.        |               |  API Requests  
                  .---------------.           ^           .---------------.
                                              |
----------------------------------------------+----------------------------------------------
                                              |                                Control Plane
                                              v
                                     .-----------------.
                                     |   IO Control    |
                                     .-----------------.
```

## Using Envoy as our Proxy

If we use Envoy as our configurable proxy, it can fill both roles: both forward and reverse proxy. Bringing those together into one component, our `IO` pops out and our data plane has two layers: the Envoy proxy and its local `IO` controller.

```goat
                                        .-----------.                            Agent Plane
                                        |           |
                           +----------->|   Agent   +-------------+
                           |            |           |             |
                           |            .-----------.             |
---------------------------+--------------------------------------+--------------------------
                           |                                      |                 IO Plane 
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
----------------------------------------------+----------------------------------------------
                                              |                                Control Plane  
                                              v
                                     .-----------------.
                                     |   IO Control    |
                                     .-----------------.
```

You might recognize this as the sidecar model that is widely used in service meshes. But here we're zooming in pay close attention to the specific needs of agents. We're going to pay less attention to the "mesh", focusing on managing the communications of an agent independently of how any other agent is managed.

Unlike service meshes that use Envoy, we have a local controller for our Envoy. This is because we'll be proposing such a close integration of Envoy and IO that a remote controller is suboptimal, if not completely impractical. Also, by bringing IO closer to the agent, we bring it into the domain of the agent developer, who will have much to say about its configuration and capabilities.

In this picture, Envoy is potentially an implementation detail of the IO plane. Its control surface may or may not be exposed to the control plane, and if not, we can rely on it as a convenience and have the flexibility to move to other implementations in the future. 
