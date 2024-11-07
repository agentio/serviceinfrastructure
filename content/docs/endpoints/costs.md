---
weight: 2
title: What does it cost?
---
# Endpoints Costs

The system that provides Google API management is publicly known as Service Infrastructure, and a wonderful thing about it is that it is available to the public to use. Because it's built to run at Google-scale, it's also highly performant and affordable. 

Service Infrastructure is made up of several sets of related APIs, and the two most central ones are called Service Management and Service Control. Respectively, they provide the management plane and the control plane for the APIs managed with Service Infrastructure. Service Management is the directory or catalog of all of the APIs managed by Google, and Service Control is the system for controlling and monitoring API usage.

As you might imagine, the number of calls you make to the Service Management APIs is relatively few, and these API calls are currently free. Conversely, the number of calls that you make to Service Control can be quite high, usually matching the number of API calls that your users make, and these calls are billed quite inexpensively. The first two million calls per month are free, after that they cost $3/million, but after a billion calls, the price drops to $1.5/million. To make this more tangible, the table below shows a few examples.

| calls per month | cost per month |
| --------------- | ------------ |
| under 2 million | $0 |
| 3 million | $3 |
| 10 million | $24 |
| 100 million | $294 |
| 1 billion | $2,994 |
| 10 billion | $16,494 |

If you're serving 10 billion API calls per month, you're probably doing something so well that even that last number is a bargain. For example, in Feb 2024, Amazon reported [4 billion site visits](https://www.semrush.com/website/amazon.com/overview/). Also, with that much traffic, you also will have other infrastructure costs that will probably make this inconsequential.

This also assumes no caching or batching of requests. In practice, proxies often cache the results of calls that check API keys and batch the calls that log API traffic. This can reduce costs by at least an order of magnitude.

All told, it would be hard to find another API management solution that's a better value than this.

Why is this so inexpensive? In the report mentioned above, Google has the two most-visited websites in the world, totalling nearly 270 billion site visits per month. When you do something that much, it pays to optimize it, and many Google careers have been made on seemingly small optimizations that had enormous impact at scale. That's led to Google's unique way of building APIs, which we discuss in [How Google Makes APIs](/docs/details/how-google-makes-apis).
