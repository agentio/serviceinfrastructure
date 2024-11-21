---
weight: 2
title: How to Call Google APIs
---
## How to Call Google APIs

Since the systems that we are using are provided as APIs, working with them will require us to call Google APIs. Here's a quick primer on how to do that.

### Calling Google APIs with gcloud

The [gcloud](https://cloud.google.com/cli) tool allows us to perform many GCP-related tasks, and these often include calling underlying APIs. This can't help us if we want to do something with an API that's not supported by `gcloud`, but that's actually pretty rare.

Getting `gcloud` configured and running is almost always a good way to start (we're assuming you are already a GCP user, if not, [sign up for a free trial here](https://cloud.google.com/free?hl=en)). It gives us something to compare against, it can solve problems in a pinch, and we can even use it to easily get auth tokens to send with our own API calls. 

If you use the [Google Cloud Shell](https://cloud.google.com/shell/docs), you'll find `gcloud` installed and ready. Otherwise, use [Google's instructions for installing gcloud](https://cloud.google.com/sdk/docs/install) and get set up. We use it throughout our discussions, starting with the [quickstart demo](/docs/quickstart/demo).

### Seeing inside gcloud with --log-http

`gcloud` has a nice built-in way to learn about Google APIs. Just add the `--log-http` flag to any `gcloud` command, and it will print all of the API calls that it makes along with their responses. Here's an example:

```
$ gcloud endpoints services describe stores.endpoints.bobadojo.cloud.goog --log-http
=======================
==== request start ====
uri: https://servicemanagement.googleapis.com/v1/services/stores.endpoints.bobadojo.cloud.goog?alt=json
method: GET
== headers start ==
b'accept': b'application/json'
b'accept-encoding': b'gzip, deflate'
b'authorization': --- Token Redacted ---
b'content-length': b'0'
b'user-agent': b'google-cloud-sdk gcloud/498.0.0 command/gcloud.endpoints.services.describe invocation-id/e9f069227bf9476cbca65f6bcd78c1f0 environment/None environment-version/None client-os/LINUX client-os-ver/6.11.0 client-pltf-arch/x86_64 interactive/True from-script/False python/3.11.9 term/xterm-256color (Linux 6.11.0-8-generic)'
b'x-goog-api-client': b'cred-type/u'
== headers end ==
== body start ==

== body end ==
==== request end ====
---- response start ----
status: 200
-- headers start --
Alt-Svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000
Cache-Control: private
Content-Encoding: gzip
Content-Type: application/json; charset=UTF-8
Date: Fri, 15 Nov 2024 04:35:45 GMT
Server: ESF
Transfer-Encoding: chunked
Vary: Origin, X-Origin, Referer
X-Content-Type-Options: nosniff
X-Debug-Tracking-ID: 11591317513918051737;o=0
X-Frame-Options: SAMEORIGIN
X-XSS-Protection: 0
-- headers end --
-- body start --
{
  "serviceName": "stores.endpoints.bobadojo.cloud.goog",
  "producerProjectId": "bobadojo"
}

-- body end --
total round trip time (request+response): 0.373 secs
---- response end ----
----------------------
producerProjectId: bobadojo
serviceName: stores.endpoints.bobadojo.cloud.goog
```

Notice that this hides your auth token. If you want to see that, just run `gcloud config set log_http_redact_token false` and tokens will be displayed.

This also reveals that `gcloud` doesn't use gRPC! But that's not too surprising because `gcloud` is older than gRPC. Digging into the `google-cloud-sdk` release notes, we find this description of its earliest release:
```
## 0.9.0 (2013-04-09)

*   Developer preview release of the Cloud SDK
*   Includes command line tools for:
    *   App Engine
    *   BigQuery
    *   Compute Engine
    *   Cloud Storage
    *   Cloud SQL
*   Includes the new gauth tool for one time common authentication for all tools
```

### Calling Google HTTP/JSON APIs with the APIs Explorer

Many Google APIs are available as HTTP/JSON APIs that we can easily call with the [Google APIs Explorer](https://developers.google.com/apis-explorer), a web-based tool that you'll often find alongside Google API documentation.

As an example, let's call the [Google Cloud Translation API](https://cloud.google.com/translate/docs/reference/rest). We'll focus on the [translateText](https://cloud.google.com/translate/docs/reference/rest/v3/projects/translateText) method, and we can start by using the APIs Explorer on the right hand side of the documentation page. This will automatically get credentials for us, so we'll just need to figure out what we need to put in our request.

Here's the simplest possible example: First, find the id of one of your Google Cloud projects (we'll use "bobadojo") and put it in the parent field with the "projects/" prefix ("projects/PROJECT_ID", where you substitute your project ID for PROJECT_ID). Then put the following in the "request body" section:
```
{
  "contents": [
    "hello"
  ],
  "targetLanguageCode": "es"
}
```
Now be sure the "Google OAuth 2.0" box is checked and press the "EXECUTE" button. The console will help you authenticate and get a credential that is used to call the API. Your input should look like this:

![alt text](/screenshots/translate-explorer.png)

If this is your first time calling the API, you might get an error like this one:
```
{
  "error": {
    "code": 403,
    "message": "Check Error: SERVICE_NOT_ACTIVATED Service 'translate.googleapis.com' is not enabled for consumer 'project_number:622059496897'.",
    "status": "PERMISSION_DENIED"
  }
}
```
To fix this, open another browser window, go to the Google Cloud Console, and find the "APIs & Services" section in the left sidebar. Hover over it and choose the "Enabled APIs & Services" submenu. Make sure that the project that you want to use is selected in the top bar, and then enter "Cloud Translation" in the search box. Select the "Cloud Translation API".

![alt text](/screenshots/translate-selection.png)

On the next screen, look in the "Additional Details" section and be sure that the service name is "translate.googleapis.com". If not, you might have selected the wrong API! But if you've got the right one, click on the blue "ENABLE" button to activate the API.

![alt text](/screenshots/translate-enable.png)

You'll see a progress indicator and then a status view that shows metrics, and you'll probably see there that you haven't made any API calls yet. Keep this window open and go back to the APIs Explorer and call the translation API. You should get a response like this one:
```
{
  "translations": [
    {
      "translatedText": "Hola",
      "detectedLanguageCode": "en"
    }
  ]
}
```
If you don't, wait a minute or two and try again. The service activation might need some more time to propagate.

Now if you wait a few minutes, you can go back to the status view, reload it, and look for metrics for the call you just made... or not! At this time, calls from API Explorer don't seem to be logged. But we'll see something there soon.

### Calling Google HTTP/JSON APIs with curl

Now let's call this API with `curl`. Expand your API Explorer view by clicking on the small dotted square just to the right of "Try this method". You'll get a view that shows a curl command. It should look like this one:

![alt text](/screenshots/translate-curl.png)

That text is meant to be an example, but it has a few problems. First, we don't need the key query parameter, because [the translation v3 API doesn't accept API keys](https://cloud.google.com/translate/docs/authentication#api-keys). Next, we need to fill in a value for `[YOUR_ACCESS_TOKEN]`. Conveniently, we can get this from `gcloud`. We'll just take the output of `gcloud auth print-access-token` and use it in our call. We do that below using shell substitution, and note that we switched to using double quotes around the string where we make the substitution call (this is a shell trick needed to make the substitution work).
```
$ curl "https://translate.googleapis.com/v3/projects/bobadojo:translateText" \
  -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  -d '{"contents":["hello"], "targetLanguageCode":"es"}' \
  -H "Content-Type: application/json"
```

We're almost there, but when we run the above command, we see another problem in the response:
```
{
  "error": {
    "code": 403,
    "message": "Your application is authenticating by using local Application Default Credentials. The translate.googleapis.com API requires a quota project, which is not set by default. To learn how to set your quota project, see https://cloud.google.com/docs/authentication/adc-troubleshooting/user-creds .",
    "status": "PERMISSION_DENIED",
    "details": [
      {
        "@type": "type.googleapis.com/google.rpc.ErrorInfo",
        "reason": "SERVICE_DISABLED",
        "domain": "googleapis.com",
        "metadata": {
          "service": "translate.googleapis.com",
          "consumer": "projects/32555940559"
        }
      }
    ]
  }
}
```

What is this quota project? Read [the documentation](https://cloud.google.com/docs/authentication/adc-troubleshooting/user-creds) if you are interested, but to keep going, just add the `x-goog-user-project` header to your requests like below, and be sure to replace "bobadojo" with your own project id. With that, your `curl` command should succeed.

```
$ curl "https://translate.googleapis.com/v3/projects/bobadojo:translateText" \
  -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  -d '{"contents":["hello"], "targetLanguageCode":"es"}' \
  -H "Content-Type: application/json" \
  -H "x-goog-user-project: bobadojo"
{
  "translations": [
    {
      "translatedText": "Hola",
      "detectedLanguageCode": "en"
    }
  ]
}
```

Whew. Believe it or not, that's the easiest way to call this API, and hopefully having this example will make it easier to `curl` other Google APIs and call them from your own code.

Now you can go back to the API detail view in Cloud Console and see some of the metrics collected for your API calls. Here's a view that is customized to show two graphs. There's also a list of methods at the bottom (with only one entry because we've only called one method).

![alt text](/screenshots/translate-metrics.png)

### Calling Google gRPC APIs with grpcurl

Now let's move to calling this API with gRPC. We'll start by using the [grpcurl](https://github.com/fullstorydev/grpcurl) command line tool.

`grpcurl` needs a description of the request and response messages of this API, and to get that, it reads a file called a descriptor set that we can generate with `protoc`. `protoc` compiles a file describing the API, and we can find that source file in the googleapis repo: [translation_service.proto](https://github.com/googleapis/googleapis/blob/master/google/cloud/translate/v3/translation_service.proto).

Why are these files public? Downstream tools use them, and Google has a long established practice of publishing API descriptions. Before there was OpenAPI or Swagger, Google had the API Discovery Service, which described Google's APIs with a custom JSON format that was created by a team at Google. That was used to drive code generators that Google wrote, and over the years, outside developers wrote their own code generators that created client libraries and other tools for Google APIs.

```
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go: downloading github.com/cespare/xxhash/v2 v2.2.0
go: downloading github.com/cncf/udpa/go v0.0.0-20220112060539-c52dc94e7fbe
go: downloading google.golang.org/genproto/googleapis/api v0.0.0-20231106174013-bbf56f31fb17
go: downloading golang.org/x/oauth2 v0.14.0
go: downloading cloud.google.com/go/compute v1.23.3

$ git clone https://github.com/googleapis/googleapis

$ cd googleapis

$ protoc google/cloud/translate/v3/translation_service.proto \
  --proto_path . \
  --include_imports \
  -o descriptor.pb

grpcurl -protoset descriptor.pb \
    -d @ \
    -H "Authorization: Bearer `gcloud auth print-access-token`" \
    -H "x-goog-user-project: bobadojo" \
    translate.googleapis.com:443 \
    google.cloud.translation.v3.TranslationService/TranslateText <<EOM
{
  "parent": "projects/bobadojo",
  "source_language_code": "en",
  "target_language_code": "es",
  "mime_type": "text/plain",
  "contents": [
    "Hello, World!"
  ]
}
EOM
{
  "translations": [
    {
      "translatedText": "¡Hola Mundo!"
    }
  ]
}
```

### Calling Google gRPC APIs with generated API clients

Google builds and uses tools to generate higher-level client libraries that provide an abstraction layer above the basic generated gRPC clients. Here we'll see how to use the Cloud Translate client for Go.

As a companion to this book, we've written a command line tool that we can use to call each of the APIs that we describe and sometimes perform higher-level API management operations. The CLI is at [github.com/agentio/q](https://github.com/agentio/q), and it includes a simple command handler that calls the same translation API that we just called using gRPC.  The handler is in [cmd/translate/cmd.go](https://github.com/agentio/q/blob/main/cmd/translate/cmd.go).

If you're a Go programmer, this is probably the easiest way to call this API. Note that this is using what Google calls [Application Default Credentials](https://cloud.google.com/docs/authentication/application-default-credentials), which are essentially the same credentials that we've been getting from `gcloud` (there are nuances that you can read about in the Google docs).

For more, see [How Google Makes APIs](/docs/details/how-google-makes-apis).

---
#### Go back to [Going Deeper](/docs/details).