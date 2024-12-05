---
title: IO Concerns
type: docs
bookHidden: true
bookToC: true
bookSearchExclude: true

weight: 20
---

# Agent IO Concerns

When we are setting up an IO, what do we care about?

Here are some ideas that could be elements of an IO user interface.

## The IO
- what is its version?
- how do we update it?
- is it healthy? (number of threads or goroutines, memory usage, time of operation, etc)

## The Proxy
- is it running?
- how do we start and stop it?
- how do we update it?
- what is the round trip time from IO to the proxy?

## The Forward Interface, aka "Calling"
- what is its configuration?
- who is allowed to call it?
- what is it able to call?

## The Reverse Interface, aka "Serving"
- what is its configuration?
- who is allowed to call it?
- what is it able to call?

## The Agent Controller
- where is it?
- how do we authenticate to it?
- how do we register our identity with it?