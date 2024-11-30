---
weight: 3
title: "Cloud Endpoints with Dan Ciruli and Sepehr Ebrahimzadeh: GCPPodcast 44"
bookHidden: true
---

# Transcript: Cloud Endpoints with Dan Ciruli and Sepehr Ebrahimzadeh: GCPPodcast 44

{{< youtube jYh0e1i_16M >}}

== Transcript from the Google Cloud Podcast

https://www.youtube.com/watch?v=jYh0e1i_16M

[MUSIC PLAYING] 

FRANCESC CAMPOY: Hello, and welcome to episode number
44 of the weekly Google Cloud Platform Podcast. I am Francesc Campoy, and I'm here with my colleague, Mark Mandel.
Hey, Mark. 

MARK MANDEL: Hey, Francesc. How are you doing today? 

FRANCESC CAMPOY: Very happy to have you back.
How was Strange Loop? 

MARK MANDEL: Strange Loop was fantastic. That is definitely one of my favorite conferences of the year. 

FRANCESC CAMPOY: I'm very jealous of the t-shirt
you're wearing right now. I think we should tweet a picture of that. 

MARK MANDEL: Yeah. 

FRANCESC CAMPOY: Because that's pretty amazing. 

MARK MANDEL: They have a great Stranger Things-themed Strange
Loop t-shirt. 

FRANCESC CAMPOY: Yeah. 

MARK MANDEL: It's very well done. 

FRANCESC CAMPOY: It is very cool. I kind of want one now. But anyway, so today we're going to be talking
with two of our coworkers, Dan Ciruli and Sep Ebrahimzadeh,
both working in Google Cloud Endpoints. 

MARK MANDEL: Yeah, we have a really interesting conversation about what Endpoints is, what it does, why you'd want to use it,
how you can use it, not just within Google Cloud but also across Cloud. It's really, really cool. 

FRANCESC CAMPOY: Yeah, it is very interesting.
Especially, if you are serving any kind of API. I think that it is worth for you listening to this interview,
you're going to learn some stuff and maybe make your life easier. 

MARK MANDEL: Yeah, I think I actually want to integrate it into some demos that I have now.

FRANCESC CAMPOY: I'm going to integrate with some demos that I have. I'm already working on it. And it is actually incredibly nice, especially
not having to deal with [INAUDIBLE]. 

MARK MANDEL: Yes. 

FRANCESC CAMPOY: Because it is not my thing.
Then after that, we'll have the question of the week. And today we do not have the Cool Things of the Week.
We have the Questions of the Week. 

MARK MANDEL: Yeah. 

FRANCESC CAMPOY: We're going to have two questions of the week.
One is related to a topic that we covered before. And of course, some of the week, how to access
metadata from a pension. The second one is more about, OK,
so I have a bunch of different front ends to my back end. How do I manage this? Things to make it reusable and as simple as possible.
But before that, we're going to have the Cool Thing of the Week, which is about a German philosopher.

MARK MANDEL: Kafka. 

FRANCESC CAMPOY: I don't know if he's German now, actually. 

MARK MANDEL: I have no idea. 

FRANCESC CAMPOY: Actually, I think so.
[LAUGH] 

FRANCESC CAMPOY: Let me check because I'm almost sure. Philosopher.

MARK MANDEL: Francesc is looking on the air. So if you're not familiar with Apache Kafka while Francesc is looking things up,
it's a messaging system that is very popular for a lot of people. Open source.
We have Pub/Sub as a mini messaging system here on Google Cloud Platform. But a lot of people do use Kafka in and outside of Google Cloud
platform. 

FRANCESC CAMPOY: And he wrote in German, but he was actually from the Czech Republic,
from the Austro-Hungarian Empire. 

MARK MANDEL: Important information for everyone. 

FRANCESC CAMPOY: Very, very important information.
I like important information. Basically, what we are announcing, which I think is really cool, is the fact
that if you are a Kafka user, you're going to be able to integrate very well with a bunch
of our technologies. So one of the ones that I'm the most excited about is the fact that you can integrate with Cloud Pub/Sub.
Both ways. So you can send from Pub/Sub to Kafka and back. But then since you have that, there's a bunch of other things
that also work. One of them is, you can use it as a source for BigQuery.
So you can stream into BigQuery using Pub/Sub. You can also do it now with Kafka.
And finally, there's also the possibility of doing the same thing streaming to Apache Beam
or Dataflow. Actually, Dataflow and Apache Beam are pretty much the same thing. 

MARK MANDEL: Yeah, no, I really like this
because if you want to put Pub/Sub in your Workflow, you can. So you can go Kafka to Pub/Sub and then Pub/Sub to a number of different things
like BigQuery or Dataflow. But if you don't want to have another messaging system within your system, maybe for complexity's sake,
wanting to make your life easier. You can go straight from Kafka to Dataflow, or straight from Kafka to BigQuery, which I think that's really, really cool.

FRANCESC CAMPOY: Yeah, it is very interesting, especially for those that are already using Kafka and were considering using things
like BigQuery for big data analysis and understanding better their business. I think it's a great opportunity.

MARK MANDEL: Absolutely, I really, really like it. 

FRANCESC CAMPOY: Yeah, and we'll have links to both the blog post and also all the code, which
is on GitHub. So you can go check it out. 

MARK MANDEL: Absolutely. Cool. Well why don't we go talk to Dan and Sep about Endpoints.

FRANCESC CAMPOY: Sounds good. Let's do that. So I'm very excited to welcome Dan Ciruli and Sep
Ebrahimzadeh, Product Manager and Software Engineer to the Google Club Platform Podcast.
Hello. How are you doing? 

DAN CIRULI: Hello. Doing great. 

SEP EBRAHIMZADEH: I'm doing very well. Thank you. 

FRANCESC CAMPOY: Hello.
So we're here today to talk a little bit about Cloud Endpoints. But before that, I would like to know a little bit more
about who you are. So could you tell the audience who are you and what do you do at Google?
So we can start with you, Dan. 

DAN CIRULI: My name is Dan Ciruli. I'm a Product Manager on Google Cloud Endpoints and Google's
API infrastructure. I've been here for about three years, and my team not only builds Endpoints but builds all the infrastructure that
serves all of Google's APIs. Virtually all of Google's APIs. 

FRANCESC CAMPOY: Nice.
What about you, Sep? 

SEP EBRAHIMZADEH: Great, so I'm Sep Ebrahimzadeh and I'm a Tech Lead Software Engineer.
I've been at Google for about 3 and 1/2 years, and I work on Endpoints and the API infrastructure.
I previously worked on Computer Engine and other Cloud products.
I'm very excited to be here chatting with you guys. 

MARK MANDEL: OK, cool. So Cloud Endpoints recently got re-released, or released again,
or enhanced. It's different. It's new. Tell us what Cloud Endpoints is now. What does it do?

DAN CIRULI: Endpoints is-- Cloud Endpoints is Google's API management product. And you're right.
There's a previous incarnation. I'm sure we'll get into it a little bit. It used to be a feature of App Engine. And it's now a product on its own.
And Endpoints lets you manage your APIs. And by manage, I mean you want to get information from them,
know who's using them, how much they're using them. You want to be able to control, authorize, things like that.
Google Cloud Endpoints is what lets customers do that on Google Cloud.

FRANCESC CAMPOY: So basically, it's going to allow me to provide REST APIs?

DAN CIRULI: Well, there's a lot of ways for you to provide REST APIs. Right. 

FRANCESC CAMPOY: Yeah 

DAN CIRULI: I don't know what your favorite language is.

FRANCESC CAMPOY: Python. [LAUGHTER] 

DAN CIRULI: Let's say it was Python. You might want to use Flask or Django to make a REST API.
You might like a language where people don't use frameworks so much to make APIs. And that's OK.
What we want to do with Endpoints is control who has access to that API and give you really good monitoring information about
who's using that API and how. Is it healthy? 

FRANCESC CAMPOY: OK, so it is not much as about serving an API as much as giving
like a layer of protection and authentication, and stuff like that. On top of the REST API.

DAN CIRULI: Right, and we do have a couple of API frameworks. 

FRANCESC CAMPOY: OK. 

DAN CIRULI: One in Java, one in Python.
Those started its App Engine features. They've now been open sourced and can run anywhere. Those actually are lightweight frameworks
that can help you serve an API. But most people already have a favorite framework
that they're using. 

FRANCESC CAMPOY: Yeah. 

DAN CIRULI: We already have a language of choice. We know that we want to be able to work with anybody, regardless of language and framework.

MARK MANDEL: So if I got this right. So if I'm building a REST Endpoint and I want common features, say like [? logging ?]
or authentication, or some other things along those sort of lines. That's the sort of stuff that I'm looking at.
So what then are the-- if you had to give four or five bullet points for that tooling that I wanted with Endpoints,
is it [? logging? ?] You mentioned authorization. What are those points? 

DAN CIRULI: So we did a bunch of surveys
before we got going with this to find out what our customers needed. And people use APIs in a big variety of situations.
Sometimes it's to power their mobile apps. Sometimes it's between micro-services. Sometimes it's for their partners to use it.
But there is a certain core set of features that everybody seemed to need. First of all is authentication.
Right. Whether it's a service that's making a call or it's a user that's holding your mobile app or logged into your website, you need
to know on behalf of whom that API call is being made. Second is logging and monitoring.
Ops are very important. Status is very important. And you want to be able to know is the API working?
Is it working well and what are latencies? Scalability always comes in really huge.
Even customers who aren't running at large scale, it actually comes in number one on their long list of requirements.
They want to know that, if it takes off, if I become the next Pokemon Go, if I even need 1,000 requests
a second, can this do it? And then the last thing is, they want to know who's making the call? Which developer is making the call?
Every API call actually represents two entities, right? There's the developer who wrote that code. You want to know that.
And you want to know the user who's holding that mobile device. 

FRANCESC CAMPOY: Interesting.
So could you talk a little bit more about that. So you talked about authentication and who is doing it.
How do you explain that? What are the ways I can do that?
Right. So some of the ways that Google's own APIs provide authentication to their consumers
are with Google's API keys and OAuth tokens and service-to-service auth.
And all of those ways are actually available for Endpoints service producers as well. So if you have your own API and want
to put your own mobile app in front of it, you can use the same authentication mechanisms
that you would use if that mobile app was calling one of Google's APIs. 

FRANCESC CAMPOY: Nice.

MARK MANDEL: Well, you mentioned previously, Dan, that there was a previous version of Endpoints that works with App Engine.
So are there any limitations technologically in terms of how I can do my rest endpoints, or is that just like a free-for-all?
Can I put it anywhere? 

DAN CIRULI: It's a free-for-all. You can put it anywhere. As you may or may not know, we are members
and we're founding members of the Open API Initiative. And the Open API Initiative is a bunch
of people, organizations from around the internet-- SmartBear, who owns the Swagger Project, Apogee, Microsoft, IBM, Capital One.
A bunch of different providers came together and said, we should all agree on a specification format for an API.
And there's a bunch of reasons for that, and maybe I'll get into some of those. But this is the spec, that used to be called the Swagger spec,
is now called the Open API spec is how you tell us about your APIs. So I don't care how you develop it.
I don't care what language you use. I don't care what framework you use. As long as you can give endpoints, an Open API
specification for that API, then Endpoints can manage the API for you.

MARK MANDEL: Cool. I'm going to redirect to Sep, because you're being really quiet over there.
But there are frameworks that you said have come out. There's a Java framework and a Python framework.
What does that give me, above and beyond the tooling, that's sort of like the infrastructure tooling that I have available?

SEPEHR EBRAHIMZADEH: Yeah. I want to turn your question around a little bit. I think what you're trying to ask
is, why should I use Endpoints? These solutions and the problems that we talked about
are not new, and there are many other products in this space that provide logging and monitoring
and authentication. Why should I do Endpoints? And we talked about the customer who
suddenly has a lot of success with their app and has a lot of users. Well, it turns out that need for logging and monitoring
and authentication happens very quickly. As soon as you get a couple hundred, a few thousand users,
it turns out that your users will care a lot about whether your app is up and running or not or whether your APIs are up and running or not.
And at that point, it's a little too late to go spend a couple of weeks learning about a logging
service and a monitoring service and then a separate authentication framework.
What Endpoints does is it packages all of those API management features into one solution,
and it tries to reduce the amount of work that the developer has to do in order to integrate with each one of these features one
by one, by aggregating all of them and say, you, focus on your core competency,
pick any framework you like, pick any language you like, and build your back end APIs.
What we will do is we will provide you all of these features, and all you
have to do in order to take advantage of them is provide an Open API specification for your API,
right? So what is it that you get from the frameworks specifically? The frameworks allow you to annotate your code
and generate that API specification from it. And then you can choose to use the API management features
or just use that Open API specification to generate clients in different languages that
can call your API or take advantage of other features, like [INAUDIBLE] UIs that let you explore your API.
Does that answer your question? 

MARK MANDEL: Yeah, I think so. It sounds like basically the frameworks make it easier
to come up with that Swagger or Open API specs rather than trying to write it by hand, for example. Would that be correct?

SEPEHR EBRAHIMZADEH: Yes. 

FRANCESC CAMPOY: Cool. So the way you explained it, I think it's very interesting, because it really sounds like there's basically a bundle of features that
are opinionated, and it makes your life easier as a developer. It really sounds like App Engine to me.
And App Engine also comes with some constraints, right? What are the constraints, other than the fact
that you need to use an Open API or a Swagger spec for Endpoints? Is there anything else that we should know?

SEPEHR EBRAHIMZADEH: I will redirect that to Dan, I think. 

DAN CIRULI: Yeah, no, I don't think so.
In some ways our frameworks are opinionated. But I would say that the use of the Open API spec
is, in a sense, to avoid "opinionation," if that's the word. [LAUGHTER]
We want you to be able to choose your back end. And by, the way, Cloud Endpoints-- I'm surprised we got this far in the podcast before I mentioned it,
bad product manager-- runs on App Engine. But you can also use it from Compute Engine.
You can use it from a Container Engine. You can use it from wherever. The proxy that performs the API management features
is designed to be containerized and deployed anywhere. So one of the big requests we had from existing people
who were using our frameworks was they wanted to move to App Engine flexible environments they wanted to move to Container Engine.
Now they can do that. 

FRANCESC CAMPOY: Very nice. So could it work also on App Engine with other frameworks
other than Java and Python? 

DAN CIRULI: That's a fantastic question, and the answer is not yet, but we're working on that.
We need some work from the App Engine team. We need some work from our networking team as well. But we would love you to be able to,
yeah, pick your own favorite framework and run that on App Engine. That's going to take a coordinated effort,
but we'd love for that to happen. 

MARK MANDEL: Is that a restriction that's specific to App Engine Standard or App Engine Standard and Flexible?

DAN CIRULI: Flexible-- you can do anything in Flexible. 

MARK MANDEL: So Flexible's fine. 

FRANCESC CAMPOY: So it is just like, for instance, if you have an App Engine app running
Go, on Standard environment, for now, it could not work. But you could use Flexible, and it could work.

DAN CIRULI: That's absolutely right. In fact, that's why they call it Flexible because it's, indeed-- 

FRANCESC CAMPOY: It's more flexible. That is actually [INAUDIBLE] naming there.

DAN CIRULI: Among other reasons, that's a little known fact. 

FRANCESC CAMPOY: Great naming right there. 

MARK MANDEL: So we've talked a lot about REST APIs here.
But as many people who listened to the podcast know, we're huge fans of GRPC. Can I use this with GRPC, Sep?

SEPEHR EBRAHIMZADEH: Yes, you can, actually. Our GRPC support is coming very shortly.
It is currently in the alpha stages. And our customers can sign up to use that.
And there are advantages to using GRPC, and there are advantages to using JSON REST.
So you'll have to see which one is the right fit for your application. 

DAN CIRULI: I was actually really pleasantly surprised
at how quickly we got that question from our alpha participants and how readily they adopted it.
It's sometimes a little bit hard to gauge the success of an open source project, GRPC in particular. We don't know how many people are using GRPC.
It's more than I had thought. And, yeah, we want that same feature set, because whether your API's talking REST or not,
you still want logging and monitoring. You still need authentication. You still need API keys. So you want that same dashboard regardless of what
that protocol is, right? 

MARK MANDEL: That's really cool. Now, I'm actually quite curious. Since GRPC is kind of HTP2 binary spec, like Protos coming
in and out, do you get any introspection into like the data going in and out or anything like that? Do you get any cool tools that way?

DAN CIRULI: You don't right now. In fact, in either, whether you're using REST JSON or GRPC,
we're not introspecting your payload. So we're not looking at that payload in either case. But we can tell you what percentage of your calls
are errors per method, what your latencies are like at median, again, per method, which client is making the call.
And maybe you realize, hey, our errors went up, but it's all in a particular method, and it's all from only the iOS client.
So it lets you tell very quickly, OK, this is probably a client problem with the new release we did. 

FRANCESC CAMPOY: So how do you manage
exposing a Swagger or an Open API spec when what you are using is GRPC?
How does that work? GRPC has a design language. 

DAN CIRULI: Do you want to take that, Sep?

SEPEHR EBRAHIMZADEH: Sure. So our GRPC customers inside and outside of Google
are currently using Proto to define their API specification.
We're currently working on integration with Open API for GRPC services, using cyber extensions,
for example, and other options. But internally, what we do is we convert the API specifications
to the same common format. And the APIs we use for that conversion
and exposing the API specification have also been published. So if you're curious, you can take a look at them,
and these would be under our Service Management API set.

DAN CIRULI: Yeah, and we can also make a-- we'll put in the Resources section a link to the GitHub project.
So we have this thing called the Google API Compiler that can take an Open API specification or a Proto
and a Yaml, that together describe how to manage a GRPC API. They compile them this thing we call a normalized service
config, and that's actually the format that all our downstream tools take. So one thing we would like to see, as members of the Open API
Initiative we are working with, could there be a single canonical representation of an API,
regardless of whether it's HTP JSON and HTP 1.1 or Proto and HTP 2.
So as we work with that committee, the Open API technical committee, we'd like to see if there could
be a single representation. But for now, it's either an Open API spec or a Proto and a Yaml. 

MARK MANDEL: That's really cool.
Now I'm just thinking about how we were talking about the GRPC gateway that'll switch a GRPC call into an HTP call.
So then you can put Endpoints on top of both, and that's pretty powerful. 

DAN CIRULI: Right, and then certainly, we're doing that
internally at Google. A lot of our APIs now are serving both the GRPC circuits as well as a JSON surface.
And we'll have that, for sure. 

FRANCESC CAMPOY: Yeah, we had Brandon [? Philips ?] last week
talking about how to use GRPC [INAUDIBLE], and they also mentioned that they were using the GRPC gateway, which was very nice to hear.

DAN CIRULI: Obviously I know that, because I listen every week. 

FRANCESC CAMPOY: As you should. [LAUGHTER] 

SEPEHR EBRAHIMZADEH: Yeah, the nice thing, as you mentioned,
is that you write your back end, and it typically handles either GRPC or REST.
But once you use Endpoints and have that extensible server proxy in front of your back end, you
can transport back and forth. And so your clients are free to use whichever suits their needs better.

DAN CIRULI: And we do want to work with the standards committee on making sure that it's a canonical transformation
that you are describing completely with either one of those documents what your API surface is.
FRANCESC: Cool. I had a question regarding one of my least favorite topics, which is OAuth 2.
What other kind of authentications do you have? You talked about one of my favorite ones, which is API keys, because they're super easy to use.
Do you have any other options? 

SEPEHR EBRAHIMZADEH: Yeah, so we do support Gitkit, or Firebase
Auth, as it's called now. We support other third-party authentication providers,
like Auth 0. Am I missing any others, Dan? 

DAN CIRULI: No, I think those are the important ones.
And both those, Firebase Auth and Auth 0, are used by mobile developers, especially
who want their users to log in, say, with Facebook or with username and password or your Twitter
login or your GitHub login. It doesn't matter. They have integrations with those. So your end user, maybe they just log in with Facebook,
both those solutions will generate a token, a JSON web token that is passed across.
Endpoint can validate it and say, yes, this really was generated by your client. It's an auth token.
It's valid, and here's your user ID out of it. 

MARK MANDEL: So you've got options there. So if you've got something like mobile clients
and you want that particular user to be logging in, and then also if you're doing maybe like service to service, you've got those options, too. 

DAN CIRULI: Right.
And that's one of the interesting things we saw that many of our customers are actually doing two types of auth.
They've got an external-facing service that is validating a user-generated auth token,
say the user logged in with Facebook. But then that's, in turn, calling micro services. And so they're using Endpoints on those micro services
as well, and those are taking service to service. So they're validating an auth token that was created in one of their other services.

FRANCESC CAMPOY: Yeah. At least from the point of view of the client calling the server, having things like Firebase Auth
is so much better, just because you don't need to handle the OAuth 2 dance by yourself,
and that's a huge thing. I don't know a single developer that is happy to write that part of the code, ever.

DAN CIRULI: Right. No one wants to do that. That's right. 

MARK MANDEL: So both of you have actually mentioned this thing called an extensible server proxy,
or something along those sort of lines. Does one of you want to tell us what that is, and why is it extensible, and why is it a proxy?
Like, what does it do?

SEPEHR EBRAHIMZADEH: (LAUGHING) Sure, I can tackle that. So as I mentioned, we want the developers
to be free to choose any language, any framework that they want, and build and host their APIs anywhere.
So what we're providing is this extensible server proxy, which is an Nginx-based web server that sits
in front of their applications. It's extensible because currently it provides features like authentication,
monitoring, and logging. In the future, we will be adding other features that our API management customers are looking for.
And it's also extensible in the sense that, as you know, Nginx is an open source project, and we
are working towards open sourcing the proxy as well in order to have the open source
community be able to take it and add other features that they want.

DAN CIRULI: One interesting thing about the way Endpoints is architected-- it's different than virtually everybody, certainly every other cloud provider's solution-- is
that our proxy is designed to be containerized. It's not what we'd think of as a fat-shared proxy, where
every customer's API calls are going through the same one and then being routed to your back end. Rather this is a thin container that gets deployed
alongside with your back end. Internally at Google, we made a switch from a shared proxy
to what we call a distributed or server local proxy about two years ago.
We did it for performance reasons. It reduces by an order of magnitude the amount of latency introduced at the proxy.
It reduces a network hop. And it's much more scalable. So for all those reasons--

SEPEHR EBRAHIMZADEH: Yeah, the other nice thing about it is the decoupling. If you have multiple APIs or multiple micro services talking
to each other, you can update the proxy sitting in front of one of them without affecting the others.

FRANCESC CAMPOY: That's a very good point. So just make sure I'm understanding this, that means that rather than having that exposing cloud
endpoint IP, where people go always there and then be handed there maybe a load balancer and then
be handed a bunch of machines, what you're saying is that you could still have basically the same architecture, where you're exposing
an IP from your load balancer. And then when it gets to every single machine, those machines are running a proxy.

DAN CIRULI: That's right. Think of it as just another container running on those machines. And if you're using, say, GKE or App Engine Flex,
we make it trivial. With App Engine Flex, that container just pops up. 

FRANCESC CAMPOY: Yeah, it's just there. 

DAN CIRULI: We see that there's an API as part of your app,
and so we spin it up. We wire traffic through it. And so you really don't have to do anything.
If you're using GKE, you just say, throw this container here. And as your app scales and, say, you
go from serving 10 requests a serving 1,000 requests a second and you scale up that back end, each one that comes up,
a server local proxy comes up with it. 

FRANCESC CAMPOY: Nice.

SEPEHR EBRAHIMZADEH: Right. I just want to clarify, though, that even though we're big fans of containers and use them widely within Google
and we're big fans of the Docker Project, you don't actually have to use containers. They make your life a lot easier.
But if you want to just deploy your code on a raw GCEVM or a virtual machine somewhere else,
you can definitely do that, too. 

DAN CIRULI: I like that-- they make your life easier. I mean, if you like pain, sure, we got pain. [LAUGHTER]

MARK MANDEL: So Sep, you said something kind of really interesting, but you kind of glanced past it, which is you said the proxy's open source?

SEPEHR EBRAHIMZADEH: Nginx Project, which the proxy is based on, is open source. So we're working on open sourcing this.
But there are a lot of challenges open sourcing a very large project.

DAN CIRULI: And I would say that the challenges were there some teams we're aligning with internally at Google.
This is a very interesting area-- as well as outside. One of the reasons that projects like GRPC and Kubernetes
have done very well is that we actually got partners in there and working with the code before it went to open source, and it makes
sure you're in a certain state. And so we're kind of in that process right now. There's a lot of people very interested in this project,
and we want to make sure that we don't have to make too big of changes after going [INAUDIBLE] with it.
So we hope to do something soon. 

MARK MANDEL: All right, sweet.

SEPEHR EBRAHIMZADEH: Right. And as Dan mentioned, our tooling and the implementation
of the APIs that I mentioned are basically open source right now. It's very important to us to not lock our customers in, to say,
you have to use Google Cloud or you have to use Endpoints. We would like these components to be able to stand alone and for the APIs that
expose the infrastructure, which is really doing the heavy lifting. It's the same infrastructure that's running Google's APIs.
You could leverage it to host your APIs and take advantage of the same API management features.
So that's what we are really putting out there, and we would like people to be able to look at the proxy code,
look at the tooling code that converts and analyzes the Open API specification and also to be able to extend it, maybe
contribute features. 

FRANCESC CAMPOY: So does this mean that since we can run Endpoints basically anywhere,
I could run the back end for Google Cloud Endpoints outside of the Google Cloud?

DAN CIRULI: Yes.

SEPEHR EBRAHIMZADEH: That is correct. We are not providing documentation and support
for doing so yet, but it is work in progress. 

FRANCESC CAMPOY: Nice. I think that can be very interesting, especially
for people considering multi-cloud, the fact that you can, with Stackdriver, you can have monitoring multi-cloud.
And now with Cloud Endpoints, you can actually have also that in multi-cloud. That is very cool. 

MARK MANDEL: Very cool indeed.
So we're slowly running out of time, not too bad yet. I'll ask the standard question that we always
ask-- is there anything we haven't talked about yet that you want to mention or plug or talk about?
Dan, why don't you go first. 

DAN CIRULI: Well, there is one thing. And I think, because your listeners are really
going to wonder why we didn't mention it-- 

FRANCESC CAMPOY: I actually heard the name once. [LAUGHTER] 

DAN CIRULI: Yeah, I did say it once.
We mentioned it. We did go on a little shopping spree last week and bought a company called Apogee.
They are one of the leaders in the API management space. And of course, the reason we haven't spent much time on it is we told you going in because the deal has been announced
but not yet closed. Regulators on this continent and others prevent us from talking about the roadmap.
So those of you who are listening and hoping to get the inside scoop here, I can't say anything other
than this should make it clear that APIs are a big deal at Google.
And we are really excited about our product line. We're excited about the acquisition
and excited about the possibilities. 

MARK MANDEL: So basically what I'm hearing is you're going to come back again once everything is closed.

DAN CIRULI: 45 days from now. Put it on the calendar. 

MARK MANDEL: Yeah, we can have a longer discussion. 

FRANCESC CAMPOY: Awesome. That's a deal. 

MARK MANDEL: Sep, anything from you?

SEPEHR EBRAHIMZADEH: I would just like to send a message to myself a few years ago when I was building a couple mobile apps that had a few thousand
users. And I suddenly overnight had to start
Googling for monitoring services and logging services and how do I do authentication.
And say that if you are building a mobile app, if you are hosting APIs, definitely give Google Cloud
Endpoints a try. I think you'll be pleasantly surprised about how easy it is to set it up, and that you set it up once,
and you get all the logging, monitoring, authentication features, with more features to be added in the future.
So, yeah, give it a try. 

DAN CIRULI: And you get the benefit of the open API toolchain as well, right?
So that's documentation generation. It's explorers to play with your API in a web browser.
It's client library generation. So there's a whole host of stuff that the open source community is producing that we're helping to produce now
that you get as well. 

FRANCESC CAMPOY: Yeah all the cool things that you've ever heard about Swagger are there. That is very nice indeed.

MARK MANDEL: Excellent. Well, thank you to both of you for coming on. Really appreciate you taking the time.
It's been a really interesting conversation. 

FRANCESC CAMPOY: Yeah., and see you again in 45 days, apparently?

SEPEHR EBRAHIMZADEH: Of course. 

DAN CIRULI: We'll see you then.

MARK MANDEL: We'll book in you right now. That was a delightful conversation with both Dan and Sep.
I really enjoyed learning about Endpoints. I actually hadn't had a chance to look at them too much yet. So it's really nice to come here and do the podcast
and hear about stuff that's new, interesting on GCP that I may not know about. 

FRANCESC CAMPOY: Yeah, we get to learn a lot.
I had been involved with a Google Cloud Endpoints project a little bit from the beginning.
But still, I actually did not know that running in multi-cloud was an option and stuff like that.
I think it's very interesting. 

DAN CIRULI: Yeah, I super hope they end up open sourcing some of that stuff. I think that would be really, really awesome.

FRANCESC CAMPOY: I'm sure they will. 

MARK MANDEL: Cool. So why don't we get stuck into our Questions of the Week.
We have a couple of them. I'm probably going to butcher people's names, but I will try my very best.
We had two questions. The first one's from [? Rashandovorani? ?] 

FRANCESC CAMPOY: I would say so-- [? Roshandarani. ?]

MARK MANDEL: Yeah, OK, cool. Hi, Rashan. He has a question about project level custom data, or metadata,
being able to access it from App Engine. he doesn't seem to be able to do it.
And he was like, hey, Franscesc, can you help me, because this doesn't seem to work for me? 

FRANCESC CAMPOY: So yeah, we answered a question regarding
where to store information like variables, environment variables and stuff like that, a couple episodes ago.
And we mentioned that one of the places was the metadata service that has the benefit that you can actually access it from both Compute Engine
but also a Container Engine and even App Engine. And so to use it from a App Engine,
though, the important thing is that rather than targeting the metadata.Google.internal URL, that is indeed internal
and it is not accessible from App Engine unfortunately, what you need to do is you need to use the Compute API.
So the Compute API is a REST API, and you can use it from anywhere, as long as you have the authentication.
And the default authentication for App Engine provides that, as long as you ask for the compute scope.
And basically what you're going to ask is you going to ask for the information about the project. It is not a metadata thing.
It is actually inside of the project resource. You're going to have a bunch of things. And of them is all the metadata variables
attached to that project, not only to the instance. So it is something that you can definitely do.
It may be a little bit confusing at the beginning if you don't know where to look at it. So I think it was a question that was worth answering here.

MARK MANDEL: So basically, rather than trying to access it like you're running on Compute Engine, access it like you're an external service coming in.

FRANCESC CAMPOY: Yeah, basically, like if it was a mobile app or something like that, that's how you could access it exactly the same way.
The good thing is you don't need to really care about authentication because the default authentication from App Engine, it is good enough to access that.

MARK MANDEL: Yeah. That makes sense. Cool. 

FRANCESC CAMPOY: And then, yeah, we had another question from Leonard Benoit or Bennett.
I don't know. 

MARK MANDEL: Sounds good. I like it. 

FRANCESC CAMPOY: Everybody that sends emails has-- 

MARK MANDEL: An interesting name.

FRANCESC CAMPOY: Names that are hard to pronounce, which is awesome. Mine is not easy either, so, yeah.
So Leonard asks, OK, so we love Google Cloud Platform. That is good. [INAUDIBLE].
And they have an application, which they're running the back end on Google Cloud Platform. But then they have front ends which
are an Android app, an iOS app, and then our browser web interface. And the problem is like, OK, we don't
want to write the code for those applications over and over for every single time we have a little change have to update all our clients.
So how can we make this better to reuse more code, basically?
And I think that you may know about this. 

MARK MANDEL: So we might have talked about it a little bit today, which I think is really cool.
So, yeah, if you want to look at Cloud Endpoints, I think that's a really good way to go, particularly
because it conforms to the Swagger or the API spec. 

FRANCESC CAMPOY: Yeah, the Open API spec, yeah.

MARK MANDEL: And what's cool about that is then you can generate clients for it, and that's where I think you're going to get a lot of benefit
from things. 

FRANCESC CAMPOY: Yeah, that is actually basically what Google Cloud Endpoints used to do for App Engine for Java
and Python. That was exactly what it did, right? Like, it could generate some specification,
which was not Open API spec. It was a different one. But from there, you can autogenerate JavaScript code, Objective C code, and Java code.
With Swagger or Open API Spec, you can do the same thing, but with many other languages, which
is pretty amazing. 

MARK MANDEL: Yeah. I'd be interested-- depending on what it is you're doing, you could also go down the GRPC route.
But you may have some troubles there on the web front. 

FRANCESC CAMPOY: Yeah, and depending on what you're doing-- actually, no.
I was going to say you could use Go because Go can run both on Android and iOs, but not web.

MARK MANDEL: Yeah. 

FRANCESC CAMPOY: Yeah, so yeah, probably not a good idea. 

MARK MANDEL: There is a Go to JavaScript transcompiler.

FRANCESC CAMPOY: I know-- [? Go for ?] [? TS, ?] and it is pretty amazing. But still, I think there may be a little bit oo many moving parts.

MARK MANDEL: (LAUGHING) You're trying to go too far? 

FRANCESC CAMPOY: Yeah, a little bit too many moving pieces. But, yeah, basically, once you have one open API
spec or a GRPC spec, as you're saying, basically you're going to get the possibility of autogenerating the client libraries you get called will
be specific for each platform. And that will make your life much easier. 

MARK MANDEL: Yeah, and definitely
listen to the beginning part of this podcast if you haven't already. 

FRANCESC CAMPOY: Oh, yeah, absolutely. 

MARK MANDEL: Cool.
Well, thank you very much, Francesc, for joining me once again. Are you up to any interesting things lately?

FRANCESC CAMPOY: Not really. (LAUGHING) No, I'm still planning my trip to Brazil. That is my next big trip.
I'm also going to be-- at the beginning of October I will be in New York. And I'm organizing a couple of Go events,
and I might be able to run a Google Cloud Platform specific event.
But we're still planning on that. 

MARK MANDEL: Ooh, nice. I'm going to be wandering around. Google's having an Google Play Indie
Festival, so for indie game developers, on this Saturday after this podcast will come out. So I'll be wandering about that just chatting to people
and saying hello. 

FRANCESC CAMPOY: That sounds fun. Where is that? 

MARK MANDEL: That's here in San Francisco. 

FRANCESC CAMPOY: Cool. 

MARK MANDEL: Yeah, I think there's a wait list now.
So if you're already registered, yay. And if you're not, you might want to get on the wait list. And then in a couple of weeks--

FRANCESC CAMPOY: We'll have a llink in the show notes to that. 

MARK MANDEL: Absolutely. And a couple weeks after that I'll be at another gaming conference called Siege,
where I'll be presenting. 

FRANCESC CAMPOY: Awesome. Sounds like a fun week. 

MARK MANDEL: Absolutely. 

FRANCESC CAMPOY: Well, yeah, thank you again,
and talk to you all next week. 

MARK MANDEL: See you then.
