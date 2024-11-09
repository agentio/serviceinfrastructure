---
weight: 1
title: Managing APIs
bookFlatSection: true
type: docs
---
# Managing APIs

**API management** is a thriving industry with many vendors and definitions. Here we take the perspective of a software developer and say that _API management addresses the common concerns that arise when developers make and use networked APIs._

For people making APIs (API producers), these concerns usually include:
* controlling who uses an API
* controlling how much an API is used
* monitoring the usage and performance of an API

People using APIs (API consumers) have similar concerns from the other side:
* getting access to an API, usually using credentials
* using an API without exceeding limits set by the producer
* monitoring the availability and performance of APIs that they are using

These are _common concerns_ because nearly all API producers and consumers have them. Developers often start by addressing these concerns in their service and client implementations, but soon they begin moving their solutions into libraries and frameworks, and eventually they turn to solutions outside of their code. Typically these external solutions run "out-of-process" and are part of the operating system or environment where the service or client runs.
