---
weight: 3
title: "Google Cloud Endpoints: serving your API to the world (Google Cloud Next '17)"
bookHidden: true
---

# Transcript: Google Cloud Endpoints: serving your API to the world (Google Cloud Next '17)

{{< youtube bR9hEyZ9774 >}}

[MUSIC PLAYING] 

FRANCESC CAMPOY: Hello, everybody. I'm Francesc Campoy, and today I'm going to be talking about Cloud Endpoints. The whole point of the talk is to give you an idea of what Google Cloud Endpoints is able to do today and to also give you a-- let you imagine what we will be adding later on. A little bit about myself, I'm a Developer Advocate at Google.
So my job is to advocate for developers and make them happy. I've been working with the Go Team for quite a while.
And I care mostly about developer experience. That is basically my whole goal. To have developers that are happy
means that their experience should be good. And that's why I'm so excited about Google Cloud Endpoints. Because it makes developer experience
for many things actually nice, which is not something that-- it's not always the word that you
can use when you're talking about Auth2 and stuff like that. So the agenda for today is going to be-- we're
going to start talking first about what is actually an API, and specifically, a REST API, which is part of what
Cloud Endpoints can do for you. Then we're going to talk about, how do you document an API?
An API is many things. But it's specifically something that helps you use someone else's services.
You need to tell them what they can do with your service. How do you do that? We're going to be talking about that.
And then we're going to talk about Cloud Endpoints. We're going to talk about what you can do with it. What is the developer experience?
But also about how it's actually architectured behind it so you can understand exactly how it works and understand
how you could extend it. And finally, we're going to do what the talk is about,
serving your API to the world. And the world will be myself, in this case, but you know, it's something.
So let's start with, what is an API?
I'm sure that you all know what an API is. You've all used at least once an API.
And API's an Application Programming Interface. And basically, the whole point is
that it allows other programs use your programs.
An API is not for people to use it. It's for machines. That's pretty much it. So you could have a web page.
We have a webpage serving some data from a data store.
And you can either use the webpage-- you're a human being and you go in there click on buttons.
But if you're a program, you don't want to go there clicking buttons, because then you need to scrape HTML. And if you've ever done that, it's not super fun
and also not always allowed. So very often we allow this other way
of using the program, same program, just a different way of using it. Cool, so what is a REST API?
I'm not going to try to remember what REST stands for, because there's no point and it really
doesn't mean that much, at least in my opinion. But REST is about these three things.
We have resources. And those resources are your things, the things that you care about on your business logic.
So for instance, you could think about accounts, products--
I was going to say sales, but that is not the word. Whatever. You can find many different things, products, accounts,
or whatever. That is the main example. Then you have the methods, which is what you can do with those.
So you could create a new customer. You could delete it. You could update a customer.
You could do many of the things that you can do. And the whole point is that in a REST API those actions that you can do are actually
tied to well-understood HTTP methods. Right? So you have POST. When you do POST, normally you create a new one.
When you do PUT, you give a key, you update it. GET is to get stuff, as it says, et cetera.
And finally we have the representation. And representation, in the case of rest the APIs, is normally JSON.
So for example, we can have here our /product that represents
our resource, so all the products that we have. If we do a POST and send this JSON here, what we're doing
is we're creating a new resource with that POST method,
pretty simple. Now, this is REST, right? This is not REST.
This is gRPC. But with gRPC you can actually do things that are very, very, very similar to REST.
You can create REST, not RESTful, but RESTish-- I decided to call them-- RESTish APIs that basically allow you to do the same things
with the same style but with a different way of naming the resources.
Normally, in this case, it could be product service. You can create them calling it, create method.
And the representation, rather than being JSON, is protocol buffers. How many of you have used protocol buffers before?
OK. So just a quick description of protocol buffers, protocol buffers is the way we do everything
at Google, basically. Many of the teams at Google, we say that what we do is we decode protocol buffers, and then
we code them again somewhere else. That's pretty much what you do most of the time. They're basically like JSON, but when you encode them
they're encoded as binary format. Then coding and decoding is way faster and also more performance when you want to send it.
At the same time, they are better for versioning. So if you have never used protocol buffers
I think it's a good thing to do. But today we're going to be talking about REST APIs-- so JSON.
Cool. So let's talk about the demo that we're going to be using today.
We're going to be using one of the simplest demos, because the whole point of the demo is not to have a very complicated, very
complicated domain. It's to show you all the things that Cloud Endpoints can do for you without having to add, basically, any code.
OK. So let me see-- there you go.
I had to restart my computer right before coming on stage, which is never a good idea.
OK. So I have my Container Engine here.
I have my container cluster. You can see my container cluster has three nodes running there.
And if I go here and I do kubectl, or kube, C-T-L proxy,
I can access this 8001, UI.
And we should be able to see that I have a service right here.
And that service is running my service. I'm running it on Container Engine. You could be running on App Engine.
You could be running anywhere. You could be running on on-premise. That's totally fine. Cloud Endpoints works on any platform
that you might want to use. So I can get that URL.
Open postman. Those are gophers because I write Go. [CHUCKLES]
Postman, there you go. So I can now send a request there.
OK. So if I'm able to do a copy/paste correctly.
There you go, Body. OK. So if I send--
what is going on? Raw. There you go. OK. So I can send some JSON and say, my message is hello.
And then I can send it. And I receive it back. That's it. That's everything this demo does.
You send a message, and it receives this back, right? Right now there's no Cloud Endpoints going on. It's just some Containers running on Kubernetes, GKE,
Google Container Engine. That's pretty much it. Now what we're going to do, of course, is to actually add Cloud Endpoints to this
and see everything we can do. If I remember correctly-- there you
go-- if you write something that doesn't correspond, it's not good JSON, it will fail. And also, if you try to send something
to a URL that doesn't exist, you will get a 404. If you get a 404 that is kind of an ugly one,
it's just some text. So if you're serving an API, this is probably not what you want to do. But we'll see how Endpoints actually fixes that directly
without you having to do anything.
Demoing with two computers at the same time is not easy. [CHUCKLES]
OK. So we have that demo, right? And the next question is, well, I want to let people use it,
to use it, right? How can they use it? What is the message they should send?
What is the message they should receive? What are the paths that are accepted? What are the methods on those paths that are accepted?
Right? And you could do many things. You could just tell them. You can go and talk to every single one of the developers
and tell them, oh, yeah, it's /echo. And it needs to be POST. It needs to send JSON that has message--
but message needs to be lowercase, by the way, just in case you forget.
It's something that you could do, but I don't recommend that, especially if you're serving this to the world.
So what do we do instead? Well, we could show the code.
And in this case, I could say, right, I'm telling you. I've showed you the code.
So you can see this code. If you know how to write some Go code, you'll be like, OK,
so there it is right there. If you send a POST method to echo,
then the echoHandler function is cool. So then I could follow up and try to understand what we're trying to decode, what are we
trying to encode? Not perfect, but this is a little bit better than just telling people.
Or you could go contract first. And by contract first what I'm saying is, rather than writing the code and telling people
about it, what I'm going to do is I'm actually going to agree with people. This is my API.
This is how it's going to be used. And from there, I'm going to start writing my code.
And now this has a bunch of benefits. One of the benefits it has is that it improves collaboration
across teams or across projects. If I clarify my API clearly, afterwards I
could change completely the way I'm implementing behind. I could move from App Engine to Compute Engine.
I could do whatever I want. If I don't break that API, I should not be impacting any of the teams that are using my service.
So that is pretty important. But also, there is many other ones. And what are you looking for when
you're talking about the way you document APIs? Right? Like, we're saying, if you specify very well what you can do with an API, then you're good.
You get all these advantages. Well, the first thing you need to do is you need to be very accurate. You need to specifically say all the things
that you need to do to use my service, which means that also you need to be very expressive.
Because imagine that if the language that you're trying to use to express all the things about your API
doesn't allow you to say, oh, by the way, the JSON object should look like this.
And finally, it needs to be readable. And when I say readable, there's actually two readables here. There's readable by humans and readable by machines.
Right? I want something that I, as a person, can go, and read, and understand correctly, and have a very good
idea of how to use this API. But also, as a software engineer-- which is a kind of person maybe--
what I want to do is I want to parse that description of the API so I can build stuff on top of that.
So free text, is it accurate? No. Free text feels accurate but is way too artsy.
You can write many things in many different ways that might be ambiguous, especially if you're not a native English speaker, like me.
There's things that you could write that actually means something slightly different depending on the context. So free text, not the best.
But expressiveness, it's amazing. You can express everything in the world with free text, so, of course.
And readability, it is awesome for humans, awful for machines. We're working on that with machine learning.
But I would not recommend using machine learning to parse recommendation for APIs
to create all the tools. That-- it's cool, as a demo. But I don't think it's a very good idea.
So let's go down the other side. It's incredibly accurate. It says exactly what you want to say.
There's no ambiguity. It is very expressive, because you're actually writing the whole domain.
Very often the source code does-- expresses way more than just your API. It expresses exactly all the process.
But readability-wise, I'd say that even though for machines it's pretty easy to read-- I mean, it depends on the prominent language,
but most of the prominent languages are done to be parsed. Breathability for human beings is not that good.
Because first, you need to talk to someone that programs and in your programming language.
I program in Go. If you're a programmer that has never written in Go and only knows-- I don't know, let's say, Haskell.
From Haskell to Go, there's a big difference. If you need to understand Go to be able to use my API,
that's not a good thing. And then finally we have design languages.
And those design languages, there's many of them. But the design language is-- basically the idea is that you create a new language in between.
Right? That language, the whole idea is that it's accurate and expressive enough to be
able to describe your API without ambiguity but not completely.
You're not going to be able to express the whole thing that your API might do in all the corner cases with design language.
It's not a programming language. But on the other hand, it's more readable. It is more readable, in general, for human beings.
And it's definitely more readable for machines. That is the whole point. You're building something that is very possible.
So it feels like we should go with the one that has the better tradeoffs. So I'd go with the last one.
Which one? Well, you could go with RAML. You could go with WSDL--
I think it's pronounced. You can go with WADL, or WADL, depending on where you're from.
Or you can go with OPEN API Initiative also previously known as Swagger.
How many of you have used Swagger? Cool. OK. So that is the one we're using.
If you use Swagger, the whole point is that you're able to do many cool things. You specify mainly for components.
So when you're doing a Swagger API, you have information about the API itself.
It's basically documentation. You have paths. So what are the paths in the URL that you're serving?
What are the operations in them? What are the methods for every single one of the paths? And then for input and output, you're
basically defining those messages. Those messages are what we call definitions. And finally, there are security definitions.
So basically, what you're going to have is for information, in this case, we have description, "A Google Cloud
Endpoints API example," a title, a version, a host-- which is, where should I go to find this API?
And then you're going to have also, what are you receiving? What are you sending? So in this case, application JSON.
We're receiving JSON, sending back JSON. And we're serving everything on HTTPS. All of this doesn't say anything about what
you can do with the API. This is just like the documentation for humans to read and a little bit of where to find it.
I went too fast with that one.
Then we have the paths. And in paths, for instance, in this case you could say, well, if you go to /echo and you send a request,
which is a POST method, then you have this operation. OK? That operation, what you're doing is saying,
well, there's the description. This is the operationID. And we're going to see what this is useful for.
OperationID is not compulsory, but it's very useful when you're generating code. And you'll again say what it's producing, application JSON.
And then you have responses and parameters. And this is very interesting, because basically what you're saying is-- what do I get from the request?
So what are the parameters that I'm receiving? And what do I send back?
In our case, the responses we're going to say, well, there's two cases. One is 200, also known as the status code, OK.
And in which case, we're sending the echo back. Our [AUDIO OUT] says. It just sends the message back.
So that's what it does. And it sends something with schema, echoMessage.
And that is actually a reference to a definition that we're going to see in a minute. So basically you're saying, what are you going to send?
But in other cases, default, we're going to be sending an error message. And the error message is also well-defined.
So we're saying it's an error message, which is also in definitions.
What do we receive? Well, that you're going to have in parameters. And we're going to say that we're receiving a message.
So the schema is again the echoMessage. But what is interesting is we are actually saying
these will be in the body. Right? You could have many different places. You could have this as a parameter in the URL.
You could have it as part of the path. There is many places where that information could be. And Swagger, Open API specification
allows you to specify exactly where this is coming from.
What does the definition look like? Well, error message. We're saying, for instance, this is a type: object.
So it has a bunch of fields. And the fields are both required, messaging code,
and messages of type: string. Code is an integer. And you can even say it's from 100 to 600,
because that's what HTTP codes are for. Cool.
So what do I get from this? Actually, a bunch of stuff. You get really cool things.
So you can create documentation from this that is very readable for humans. You can create code.
And you can even create developer portals, which is very useful. If you've ever used any Google API, for me,
one of the coolest things we have is the Google-- the API playground, the API Explorer. Where you can go, search your APIs,
go send the message directly without having to install anything. So you understand the API before you start writing any code.
So let's demo that.
OK. So what I'm going to do first is make sure that my server is running locally.
Not that one but this one here.
OK. There it is. I should get something complaining. There you go.
OK. So my server is now running. And what I can do is I can visit editor--
wrong window-- editor, Swagger editor.
Wrong window. I have too many windows. There you go, this one. And with this one, what I'm doing
is basically I just copied that Swagger definition onto this editor. This editor is something that is very simple.
It's online. You don't need to install anything. Now the cool thing about this is that we can see on the right side this documentation, where
you can see if you send POST to echo, the description, you have all the parameters, you can understand quite easily how to use it.
The more documentation you add in the specification file, you get a documentation that is easier to use for your developers.
And then on top of that, you can try an operation. You can say exactly what you're going to be
sending, what you're receiving. It is able to tell you-- actually,
the echoMessage object has a field, which is required, message, and you then send it.
So it's able to do validation for yourself. It's able to do this validation even in a proxy
that you don't need to write yourself, which is pretty cool. So you can write, hello, Google Cloud Next, send it there.
And we get a response with a message coming back. And we can see what it looks exactly.
And if you think about this, if you're offering these tools to the developers that
are trying to use your API, this is way more useful than amazing documentation.
In my opinion, documentation is great to read it, but when I want to use your API, I want this.
I want to be able to find what operation is doing what and what are the exact fields that I need to send, try it, and then write
my code that will do the same.
There's an important thing if you try to do this at home. There is the fact that there is what is called the Cross-Origin
HTTP request. I don't know if you heard about this, CORS, C-O-R-S. The whole point is security.
If you try to do this onto a server that is not running locally, you start to need to take care of these things.
There's documentation about that. I feel like every single person that has tried to use this tool
has discovered the options method. It's a method in HTTP that nobody
uses except for this tool. So you need to support it.
Cool. OK. So this brings a lot of value, right?
But again, this is still not Cloud Endpoints. What is Cloud Endpoints? Well, Cloud Endpoints is something
that has been built to-- given just this specification--
give you more value, give you a bunch of things that you need to do. If you want to expose your API to the world,
there's things like monitoring authentication that you should be doing. But you do not need to write those.
So that is exactly what Cloud Endpoints does. Let's see that.
Whoops. OK.
So if you've used Cloud-- has anyone used Cloud Endpoints previously, like one year ago?
OK. So these may be familiar to you. Cloud Endpoints is a framework that runs on App Engine standard environment,
either for Java or Python, and that's it. No.
Not anymore. That's the whole point of Cloud Endpoints is getting that idea that people loved and make
it accessible to everywhere-- no matter if you're running Java, Python, or Go, or PHP,
or whatever language you want. And if you're running it on APP Engine, or in a Container, or somewhere in a machine that is hiding in your closet.
We do not care about that. You are able to run it wherever.
So what it is is an API gateway, and it's an API gateway
that is distributed. We're going to see that later on the architecture diagrams. But the cool thing is that this is very integrated, right?
So it is integrated with deployments for App Engine. If you're using App Engine flexible environment,
it's a great experience. It is simply-- you are just-- one little field in your app in the demo.
You redeploy and that's it. You're done, which is amazing. But for Container Engine it is-- which is what we're going to be doing-- it's very easy.
And if you want to do anything else, it is doable. It might be a little bit harder to implement, but it is definitely doable.
And it also integrates with Stackdriver logging and Stackdriver Trace. So whenever you receive a request,
you're going to be able to see that in your logs and also in traces. So you're going to be able to track, to measure latency,
and things like that, without having to carry it in your code. So the key features are authentication.
So you're able to authenticate users and control who has access to what.
All of this, just by adding things, security definitions to your open API specification.
You're going to get also logging and monitoring. And this is really fast.
It is incredibly fast. And it scales really well. And we're going to see why it's-- I'm not a marketing person, so I'm
going to explain why this is fast and scales very well. You can run it anywhere.
You can run the Endpoints thing, that API gateway, you can run it on App Engine.
So if you run on App Engine, you can run on App Engine flexible for whatever you want to do, or App Engine standard with the frameworks for Java and Python.
You can run it on Container Engine. You can run it on Compute Engine. Or really, you can run it anywhere where you are able to run a Container.
If it runs on a Container, you can run it. So how does this actually work?
Too much partying yesterday. My voice is destroyed.
So first what you do is you deploy an API.
What you do is you code with GCloud. You send the API specification, that openapi.yaml file that we saw before, and you send it to the Google Service
Management. What you're saying here is you're declaring something new. And you're going to get an API name and a version.
Then when you run that gateway, the gateway is going to communicate with the Google Service Management.
And it's going to say, hey, I want to serve this API name and this version. Give me everything I need.
Basically it's going to get all the information from the API, plus authentication, and all the stuff.
This is just the deployment part. That's it. So that's why you can run that Container Engine anywhere,
as long as you can-- Container Engine, sorry, that container for Endpoints. As long as you're able to use the Google Service Management
API, which is simply a REST API, you're able to do this. And then when you get a request, no matter
what client is sending it, rather than going directly to your code, it goes first to that container.
The container checks with the Google Service Control if this is allowed or not.
If it is allowed, it sends it to your actual code.
That, again, could be running anywhere. And finally it logs everything to Stackdriver for logs,
and monitoring, and traces. That's it. This is the whole thing.
Now, the important thing is that-- I don't know if you're able to see it. I'm not really able to see it--
but there's a little blue box running both the Extensible Service Container and my code.
That box is a node. And I'm calling it a node because it could be anything. It could be a path on kubernetes,
but it could also be an instance, an actual node, a virtual machine, or it could be anything else. The whole point is that this is running together, right?
And the point of running together is that since you're exactly on the same machine, there's no network hubs to be added.
So this is very fast. So it is faster and scales with your app.
What do I mean by it scales with your app? You have one gateway per container, or per unit
of whatever you're trying to use as your back-end.
So you could put a load balancer there, and the request will go to the load balancer.
And rather than going directly from the load balancer to your code, you're sending from the load balancer
to the Extensible Service Proxy. And that one will call the functions, the codes,
as we said before, and then send it to your back-end. And this you can add more and more.
As you can see here, the Extensible Service Proxy SV is never the bottleneck, because it keeps on scaling with you.
So you can imagine how this scales very well with kubernetes. You just create a pod with two containers.
One container is your back-end code. The other container is the ESP as a sidecar we call it.
So let's do that.
OK. So the important thing here is my Go program will not change at all.
It's going to be exactly the same code. So I then need to change it.
I'm going to kill the server. What I'm going to do is, I'm going to go to my zero, one.
I have my openapi.yaml. Openapi.yaml is exactly what I showed you before,
expressing all the methods for every path, the definitions, et cetera.
So you have the echoMessage, the errorMessage, and that's it. There's no security here yet. So what I'm going to do is, gcloud service
managament deploy openapi.yaml.
Wow, my computer is not the fastest. There you go. So now what we're doing here is, as I said before, we're
actually talking to-- GCloud is talking-- oh, permission denied. Oh, I'm in the wrong account.
gcloud auth list, endpoint rock. That one, gcloud config set project.
Let me make sure that is the good one, cloud--
cloud-next-endpoints. Cloud-next-endpoints. OK.
Try it again. OK, cool. So now what this is doing is sending that openapi.yaml
to Google Service Manager. And it's going to deploy that information. It's going to analyze it.
So you're going to see a warning there saying, hey, the POST in echo is not protected at all.
Anyone can do whatever they want with this. Are you sure? For now, yes, I am pretty sure that is what I want to do.
But the whole point is that now it is deployed. We're ready to continue. There's two pieces of information
that we care about in here. We get this version here, the API version.
And also we get the echo api.endpoints.cloud next endpoints.cloud.goog. That is the your API name, your API ID, which is unique to you.
You can see it's unique because-- cloud.goog is not unique, but this is the project name.
Cloud-next-endpoints is your project name. So no one else other than you has access to this.
Cool. So now what I'm going to do is I'm going to edit
my container-engine.yaml. So container-engine.yaml, what we have is here.
We can see there's two containers.
There's the container down there. It's the echo container, which is the important one, it is our business code.
And it's simply a container in Google Container Registry that I purchased previously, because the Wi-Fi is not
the greatest. But it's already there. And then we have the container on top, which is ESP.
And on ESP you can see that we have a couple parameters. One of the parameters is what port this is serving,
8081, which means that a load balancer will be sending traffic to that port.
It has the traffic where it should be sending these back. So where is your back-end, which is localhost 8080, which is where my echo server is running.
And then we have the name and the version. So I'm going to change the version. I'm going to update that.
I'm going to set the new one. And now everything I need to do is redeploy my server.
So to do that, simply I do kubectl
apply -f container-engine. So now what I did is I updated this version.
So now my new servers are running here. We can go check that out.
Oh, that is already running. So we can go directly to localhost8001/ui.
We should be able to see that there's a deployment. But if we go to paths, there's one single path
running with two containers. That has now those two containers that we decided.
If we want to add a little bit more, we can say container-engine.yaml.
The replica should be, let's say, three. I reapply.
And quite fast we should be able to see that now there's three paths. Every single one of those paths has one container
running my back-end, one container running the Endpoint proxy.
Cool. So let's try this again.
Let's make sure that I'm still running on the same service. This the external IP that I'm going
to be using, which I don't think changed. Let me see.
No it did not change. So now I can send a message, and you see that it works.
It is amazing. We did not do anything. But if you do things that are not correct, you're going to see that this is actually a little bit different
now. So if you send a path that is not accepted, this path, this request, never got to my back-end.
Instead it's the ESP container saying
that that was not allowed because the method does not exist.
Same here. If you send something that doesn't work, it doesn't work either.
Now, another benefit that I forgot to mention before is the fact that-- OK, so now you have Endpoints, you have Swagger.
You're able to have all this documentation. But there's actually more than that. You're now also able to--
go away. No. It doesn't want to go away. There you go.
You can't create-- generate code. And there's a bunch of options that you can generate.
The cool thing about this is that, yes, I can now use your REST API. Yes, there's now some limitations
that have been added. There's logging and all that stuff running. But there's also the fact that if I want to generate, let's
say, PHP. Why not? I can come here, open my PHP code, and the library, model.
It created the echoMessage. And I'm going to show much code just to give you an idea that this was all to generate directly.
So now you have the capacity of generating customer libraries for all the languages that Swagger supports,
which is actually many of them. There's actually not that many languages that people really use in production that are not supported by Swagger by now.
OK. So now if we go back to my console here--
that one is-- we have our Endpoints.
So many of them now. I remember two years ago there was nothing there.
And we have Cloud Next Endpoints-- this one here. So that one was added now because I deployed this, right?
So if I send requests-- let's send a bunch of requests with this.
A request that works maybe. So this is my load testing.
[LAUGHTER] There you go. It's like Candy Crush but with HTTP.
OK. So now you can see that I got to an amazing 0.05 requests per second, which is pretty sad.
But you can see how the monitoring went up directly, right? I did not have to do anything.
Now I have monitoring. I am able to see what's going on with that measure. I'm able to start to create alerts.
I'm able to say-- normally, I have around one request per second. If I don't have any requests per second for a little bit,
something's wrong. You're able to alert your developers about these kind of things, which is pretty amazing because it's simply Stackdriver.
You can also go here and see a little bit more information about the logs.
What are the requests you actually received? And you can see there's all this information.
There's not much. But you can actually see that for every single log what is the status code that was returned
and who rejected it, basically.
OK. So as we said before, this is amazing. Now everyone can use my service.
But that's maybe not what I want to do. I don't want everyone, everyone in the world, to be able to use my service.
For an echo server I'm pretty confident we're fine. But if you're doing something that involves actually
requests that are expensive, for instance, you might want to secure this. So how do we add authentication?
We have monitoring, tracing, how do we add authentication? Let's go back there.
Let's talk about authentication now. There's four methods that are supported by Cloud Endpoints.
There's Firebase, Auth0, there's API keys, and there's JSON Web Tokens--
also known as JWT, I've heard, which is crazy.
So what I'm going to be demoing today is API keys only, because I think it's pretty interesting.
If you're a developer and you have an API that you're exposing, think about how hard would it
be to implement this safely. To implement something that anyone in the world
can say, create a new API key, just for themselves, and start using it, and make sure that no one else has
access to it, no one else-- and that you're able to successfully identify who is using what.
So this is actually incredibly simple on the Swagger point of view. In Swagger, everything you're going to say
is-- you're going to create that security definition that says, API key. And you say, the type is API key,
and it's the parameter that's going to be in the query. So basically, question mark key equals that API key.
That's what you're specifying there. And then you're saying that the POST method on echo
is protected by API key. That's it. That's everything you need to do.
So let's do that.
So here I should have my openapi.yaml.
Oops, 02. Openapi.yaml. Here on openapi.yaml it just added these lines
at the end, security definitions saying there's an API key that is in the query.
And then I said that I'm requiring API keys for POST on echo.
So I'm going to do gcloud service-management deploy openapi.yaml. And this is what it's going to do.
It's going to redeploy it again, same thing. And now does that mean that I'm actually squashing the previous version and all of a sudden I
just migrated everyone to the new version? No. That'd be pretty awful.
Because it would mean that you're not able to do any tests, basically. What is going to happen is it's going to give us a new version.
And that version, no one is using it yet. We're going to need to then update our kubernetes cluster, or App Engine,
or whatever you're using to use that new version. And then there you can do rolling updates or whatever
you'd normally do to make sure that your new version doesn't break prod.
So we've got our new version right here.
So I'm going to change my container yaml, go to v--
I forgot this one here.
Whoa. Vim is fun sometimes.
Cool. That looks good. OK. So I just changed that. And then kubectl apply.
Now this changed the service. Now the new ESP container is actually using the new version.
So if we go back here and try to send a message
that should not work. That's it. Now it's-- we have API security, API keys that is exposed here.
Method doesn't allow unregistered callers. Please use API key or the form of API consumer identity
to call this API. Who wrote that? Not me. Yeah. This is super simple to use.
And it's now secured. My Go code has not changed yet. So how do we make this work?
I want to call that API. So the way it's going to work is I need to go to here.
And I'm going to create an API manager.
I'm going to go to credentials. We're going to create a new API key.
We're going to copy here. Go back and say, key equals that.
I'm going to send it. And now it works. And this is pretty cool, because it means that anyone can go and create a new API key
and start using it and be identified. So make sure--
I'm not following the script at all anymore. And for the demos I want to make sure I don't forget anything that I want to tell you.
Cool. So, yeah. This now works. Our next step is, OK, so API keys are cool,
but are they actually safe? There's a very good document written on this, on documentation that we published not that long ago.
And you can read it. But the interesting thing about it is that they are not,
basically. An API key has a big problem, which is you sending it with every single request.
So you're sending an API key over HTTP in a URL. That's going to get logged somewhere.
Everyone's going to find it. You're out of luck. You could also put it in your HTML if you want to.
And now everyone can have it too. So it is not super safe. But when you're doing authentication
across different projects, like a client that has an API key. This is super simple and very simple to use.
You don't need to go overkill with OAuth 2 if you actually trust what you're doing. If you have the binary running somewhere,
it can use an API key quite safely. But for other cases, what do we do?
Well, so as we said, there's the API key that you can use with the API producer,
the API producer being my echo server. But also, if you have a user that wants to use this,
very often what you're going to have is the end user is going to use some kind of authentication. It's going to obtain an authentication token.
And they're going to send that authentication token to my back-end. The ESP, the Endpoints--
I never remember what it stands for, the S. I think it's extensible, which doesn't make any sense--
the Endpoints Extensible Proxy will take care of that authentication for you,
too, to verify that this is actually sending some authentication. So for that we can use Firebase or Auth0.
Also you can use Jason Web Tokens. Jason Web Tokens doesn't require extra steps, actually, your back-end can generate it.
Firebase and Auth, too, I think they're very nice. Because what they allow you to do is to have third party authentication, which
means that you don't need to care about what authentication you allow. It's very easy to say, I want people
to authenticate with Google, but also with Facebook, and Twitter, and Yahoo, and every other ones.
Your code does not really change with this. What you're doing is you're changing the code on the client to obtain that token.
That token will be from a different origin. But your code in the back-end will be pretty much the same, which is pretty nice.
So how do you do it? In Firebase, you just create a Firebase security definition. That's it.
With Auth0, same thing-- you just
create an Auth0 security definition. And then you redeploy.
And if you want to know more about authentication, which is actually a very interesting topic, specifically how to authenticate different services,
like service-to-service, there's a session tomorrow from 2:40.
That's what I think, but that might be changed. I've heard that schedules are switching a little bit. So pay attention to that.
It might be a little bit later maybe. And we're going to have two of the people from the Endpoints team talking about what are the best practices to do
authentication across services. So if you're interested in that, go check it out, Service-to-service Authentication
in Cloud Endpoints. OK. Last part, we said that we're going
to share the API to the world. We have not shared that yet. The only person that is allowed to use it is anyone that has
access to my project, my Cloud Endpoints Next project,
that project ID on my personal-- well, not personal account-- on my Go account, which is not good enough.
I want to expose that API to anyone. How do I do that? Well, there's a couple of ways of doing it.
One of the ways-- I'll get to the message in a minute. I thought I had a slide there.
One of the ways is by email. I can say, this person with this email is allowed to use it.
So now what you can do is any project where that email is authorized to use it, they are able to enable that API.
So that is a way of doing it, by email. You can also do it by Google Mail Groups, so Google Groups.
You can create a Google Group and say, any address that is in that Google Group, any project that
is part of that address that is part of the Google Group will have access to this API. So you can do--
you can start doing this quite simply, sharing with only the people that you care about.
So how do we do that?
So what I'm going to be doing is I'm going to have two different accounts. I have my @golang.org account, and I also
have my Gmail account. And what I'm going to do is I'm going to share that API with myself, because I'm very generous.
And then I'm going to enable it and call it with an API key that is generated from my personal account.
OK. So first step, we're going to go here to our Endpoints
right there. This is the one. We're going to click on share API.
And I'm going to just-- OK, I'm going to remove it. I'm going to share it with myself.
So I'm going to share it with campoy83@gmail.com. Now you know how old I am.
And now I have shared that API. What's next? I'm going to log in with my personal account, which
if I'm not completely mistaken is here. OK. That is my personal account.
You can see it's a completely different picture. That's the proof. And now what I can see is I can go to API manager.
It's one more API, nothing crazy here. You enable the same way you enabled a Compute
API, you enable any other API. I'm going to click on enable API.
And then you have that Private APIs right there, which I'm very thankful it's working.
And here you can see echo-api.endpoints-- oh, this one here-- echo api.endpoints.cloud next
endpoints.cloud.goog, which is the name of the API that we had. You click there.
You're going to see documentation a little bit. You can enable it. And when you enable it, it says, would you
like to create credentials? Which you should, because otherwise you will not be able to call it.
So now I can create-- let's say, other API.
What? OK. Let me do it again. Create an API key.
I'm going to get that one. And now if I'm here, let's try--
with some API key that is not allowed, it does not work. With API key that is allowed that I created--
I copy/pasted wrongly I think.
Copy.
That is the same. No. I did something wrong.
OK. What I'm going to do is I'm going to fake it. So I'm going to go here.
I'm going to go to my credentials. Credentials, my API key.
I'm going to put it there. And you can imagine that that is actually a different one, and that this actually worked.
Boom. Done. Yeah. OK. So the whole idea-- I'm sure I missed some steps in the demo--
but the whole idea is that the API key creation now is not controlled by me.
It's controlled by Google. But I do decide who has access to what. And also, when those calls are done, now they are monitored.
And you're actually logging who is accessing what. And if you want to remove that access you can easily do it by revoking the--
I'm sharing that API with someone. So all of a sudden you have all of this functionality.
And my code is still the same, which that makes me happy.
So what else can you do with Cloud Endpoints? There's actually lots of really cool stuff.
If you have never used gRPC-- I love gRPC. I think it's way better than the REST.
But it's the little problem that gRPC's amazing. It's way better than REST. But people love REST.
So you need to expose REST APIs, right? There's a really cool thing called the gRPC gateway.
And gRPC gateway is a project that is pretty popular now. What it does is you give it a gRPC specification.
And from there it generates a REST API, which means that you can have both at the same time
without writing the documentation or anything twice. This is pretty amazing. And this is actually something very similar to what
we do internally at Google. If you see Google APIs very often come in two flavors-- gRPC and REST.
What they do is pretty much the same. We generate the gRPC. We write gRPC. The gRPC is the fast one.
It's the amazing one. And REST is the one that is there for convenience and very often is fast enough.
But if you're doing things like Bigtable, Spanner, things like that, the fact that you're using JSON
and you're encoding, decoding that JSON, it's expensive enough that you might consider going to gRPC directly.
And then that logo-- Cloud Functions. We're working on that. There's a lot of people asking for this, Cloud Functions.
I think it was announced today that they're now beta. I'm almost sure. And if not, I just announced something for you.
I'm almost sure it was announced. But it's now beta, so you can go play with it.
And Cloud Endpoints for Cloud Functions will be there at some point. So we're working on it. And I'm very excited about this, because the experience
of having serverless plus not having to write any code for authentication, logging, or monitoring, is going to be amazing.
And with that I'm going to finish because there's much, much more to learn.
The cool thing is that we're just starting. Cloud Endpoints is a very new project.
And we have basically shown what our idea is, what is the core technology of this.
But we are far from done. We are going to be adding more and more features. So if you have features in mind, let us know.
We're always looking for-- as I said at the beginning, I'm a Developer Advocate. My goal is to make developers happy.
So if you have an idea on how I could make you happy. Let me know. Thank you.

[APPLAUSE]

[MUSIC PLAYING]
