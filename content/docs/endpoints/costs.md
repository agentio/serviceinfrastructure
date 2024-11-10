---
weight: 2
title: What does it cost?
---
# Endpoints Costs

A wonderful thing about Service Infrastructure is that not only is it available for the public to use, it's also what Google uses to run its own APIs. Because it's built to run at Google-scale, it's also highly performant and affordable. 

Service Infrastructure is made up of several sets of related APIs, and the two most central ones are Service Management and Service Control. As we've discussed, Service Management is the directory or catalog of all of the APIs managed by Google, and Service Control is the system for controlling and monitoring API usage.

As you might imagine, the number of calls you make to the Service Management APIs is relatively low, and Google does not charge for them (they are free). On the other hand, the number of calls that you make to Service Control can be quite high, usually matching the number of API calls that your users make, but these calls cost very little. The first two million calls per month are free, after that they cost $3/million, and after a billion calls, the price drops to $1.5/million. To make this more tangible, the table below shows a few examples.

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

What about the other APIs? There is no charge for the API Keys API or for the Service Usage API. The Cloud Logging and Cloud Monitoring APIs charge for data stored, but the data written by Service Control is not considered part of that storage. The charge is only for data that users directly store by calling the Cloud Logging and Cloud Monitoring APIs themselves.