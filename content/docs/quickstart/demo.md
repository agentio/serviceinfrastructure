---
weight: 2
title: Deploy a Sample Service
---
## Demo: Deploy a Sample Service

Here's an easy way to explore the Service Infrastructure APIs. We'll use `q` in the [Google Cloud Shell](https://cloud.google.com/shell/docs) to set up a managed service on [Google Cloud Run](https://cloud.google.com/run). We'll be running the [Boba Dojo Stores](/docs/details/demo#the-boba-dojo-stores-api) API, and you can go read about that now or just jump right in.

This quickstart assumes that you have a Google Cloud account with a project created and access to the Google Cloud Shell. You don't have to start with a fresh project, but we recommend it! It's a good way to keep everything associated with your demo in one easy-to-find place.

## Install and get set up with q

After opening Google Cloud Shell, install `q` with `go install github.com/agentio/q@latest`. This might take a minute or two while Go downloads and builds `q` and its dependencies, but it should complete without errors.

{{< details "Here's sample output of 'go install' in a new Cloud Shell instance." >}}
<pre><code>Welcome to Cloud Shell! Type "help" to get started.
Your Cloud Platform project in this session is set to bobadojo.
Use “gcloud config set project [PROJECT_ID]” to change to a different project.
tim@cloudshell:~ (bobadojo)$ go install github.com/agentio/q@latest
go: downloading github.com/agentio/q v0.0.0-20241116235648-56e1ca548974
go: downloading github.com/spf13/cobra v1.8.1
go: downloading cloud.google.com/go/apikeys v1.1.10
go: downloading cloud.google.com/go/longrunning v0.6.0
go: downloading google.golang.org/protobuf v1.34.2
go: downloading cloud.google.com/go v0.115.1
go: downloading github.com/go-jose/go-jose/v4 v4.0.4
go: downloading golang.org/x/oauth2 v0.22.0
go: downloading cloud.google.com/go/logging v1.11.0
go: downloading google.golang.org/genproto v0.0.0-20240903143218-8af14fe29dc1
go: downloading cloud.google.com/go/monitoring v1.21.0
go: downloading google.golang.org/api v0.196.0
go: downloading google.golang.org/genproto/googleapis/api v0.0.0-20240903143218-8af14fe29dc1
go: downloading cloud.google.com/go/servicecontrol v1.14.1
go: downloading github.com/google/uuid v1.6.0
go: downloading google.golang.org/grpc v1.66.0
go: downloading cloud.google.com/go/iam v1.2.0
go: downloading cloud.google.com/go/servicemanagement v1.9.11
go: downloading cloud.google.com/go/translate v1.12.0
go: downloading cloud.google.com/go/serviceusage v1.8.9
go: downloading github.com/spf13/pflag v1.0.5
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1
go: downloading golang.org/x/crypto v0.26.0
go: downloading cloud.google.com/go/compute/metadata v0.5.0
go: downloading github.com/googleapis/gax-go/v2 v2.13.0
go: downloading cloud.google.com/go/auth v0.9.3
go: downloading cloud.google.com/go/auth/oauth2adapt v0.2.4
go: downloading go.opencensus.io v0.24.0
go: downloading go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0
go: downloading golang.org/x/time v0.6.0
go: downloading golang.org/x/net v0.28.0
go: downloading github.com/google/s2a-go v0.1.8
go: downloading go.opentelemetry.io/otel v1.29.0
go: downloading go.opentelemetry.io/otel/metric v1.29.0
go: downloading go.opentelemetry.io/otel/trace v1.29.0
go: downloading go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0
go: downloading golang.org/x/sys v0.24.0
go: downloading github.com/googleapis/enterprise-certificate-proxy v0.3.3
go: downloading github.com/go-logr/logr v1.4.2
go: downloading github.com/felixge/httpsnoop v1.0.4
go: downloading github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
go: downloading github.com/go-logr/stdr v1.2.2
go: downloading golang.org/x/sync v0.8.0
go: downloading golang.org/x/text v0.17.0
tim@cloudshell:~ (bobadojo)$ 
</code></pre>
{{< /details >}}

When that finishes, run `q` with no arguments.

```prompt
q
```
```output
Manage APIs with Service Infrastructure

Usage:
  q [command]

Available Commands:
  api-keys           Manage API keys with the API Keys API
  compile            Compile a Service Configuration for an API
  completion         Generate the autocompletion script for the specified shell
  demo               Set up a sample managed service
  doctor             Verify necessary dependencies and configuration
  help               Help about any command
  inspect            Read API information from compiled file descriptors
  jwt                Read, verify, and generate JSON Web Tokens
  logging            Write and manage log entries with the Cloud Logging API
  monitoring         Monitor services with the Cloud Monitoring API
  service-control    Control API services with the Service Control API
  service-management Manage service descriptions with the Service Management API
  service-usage      Manage usage of APIs with the Service Usage API
  translate          Translate with the Cloud Translation API

Flags:
  -h, --help   help for q

Use "q [command] --help" for more information about a command.
```

Now run `q doctor` to check your `gcloud` configuration. You should see something like this:

```prompt
q doctor
```
```output
Checking gcloud configuration...
account = YOUR_EMAIL
project = YOUR_PROJECT
Error: run/region required
```

If you need to set a run region, do it with the following command. We use `us-west1` here, but you could use any of the [Cloud Run locations](https://cloud.google.com/run/docs/locations) that Google supports.

```prompt
gcloud config set run/region us-west1
```
```output
Updated property [run/region].
```

Now rerun `q doctor`.

```prompt
q doctor
```
```output
Checking gcloud configuration...
account = YOUR_EMAIL
project = YOUR_PROJECT
run/region = us-west1
Checking application default credentials...
token = ya29.[REDACTED]
Everything looks good!
```

## Use `q` to create a demo instance

Next run `q demo`. This generates a directory of files that you'll use to set up your service. They will be customized to the project and run region that you configured for `gcloud`, so you won't need to edit any of these files.

```prompt
q demo
```
```
To run the demo, see the SETUP.sh script in stores-demo
```

Go inside the generated directory and look around.
```prompt
cd stores-demo; ls
```
```
api_config.yaml  CHECK.sh  descriptor.pb  iam.yaml  service.yaml  SETUP.sh
```

You might recognize some of these files. Here's a description of everything:
- [api_config.yaml](/docs/quickstart/files#api_configyaml) contains a snippet of Service Configuration used to define your demo service.
- [CHECK.sh](/docs/quickstart/files#checksh) can be optionally used to verify access to your Cloud Run service once it's set up.
- [descriptor.pb](/docs/quickstart/files#descriptorpb) is the file descriptor set describing your demo service.
- [iam.yaml](/docs/quickstart/files#iamyaml) contains IAM configuration used to set up your Cloud Run service.
- [service.yaml](/docs/quickstart/files#serviceyaml) describes the demo service that you will set up on Cloud Run.
- [SETUP.sh](/docs/quickstart/files#setupsh) does everything to set up and run your demo service.

Run the `SETUP.sh` script to configure your project and create your instance.
```prompt
sh SETUP.sh
```
```
Enabling the Service Control API.
Enabling the Service Management API.
Enabling the Cloud Run Admin API.
Operation "operations/acf.p2-622059496897-baab248d-0eae-4d5e-9e65-4fe1aa26bae0" finished successfully.
Creating the service from the service descriptor and API config file.
Waiting for async operation operations/serviceConfigs.stores.endpoints.YOUR_PROJECT.cloud.goog:7bbc869f-df3b-40fd-88f7-26bfe4c0c053 to complete...
Operation finished successfully. The following command can describe the Operation details:
 gcloud endpoints operations describe operations/serviceConfigs.stores.endpoints.YOUR_PROJECT.cloud.goog:7bbc869f-df3b-40fd-88f7-26bfe4c0c053

Waiting for async operation operations/rollouts.stores.endpoints.YOUR_PROJECT.cloud.goog:8340c75f-411a-4c5c-a282-1c1312cee543 to complete...
Operation finished successfully. The following command can describe the Operation details:
 gcloud endpoints operations describe operations/rollouts.stores.endpoints.YOUR_PROJECT.cloud.goog:8340c75f-411a-4c5c-a282-1c1312cee543

Service Configuration [2024-10-29r1] uploaded for service [stores.endpoints.YOUR_PROJECT.cloud.goog]

To manage your API, go to: https://console.cloud.google.com/endpoints/api/stores.endpoints.YOUR_PROJECT.cloud.goog/overview?project=YOUR_PROJECT
Creating a service account to run the server container.
- '@type': type.googleapis.com/google.rpc.ResourceInfo
  resourceName: projects/YOUR_PROJECT/serviceAccounts/stores@YOUR_PROJECT.iam.gserviceaccount.com
Giving the service account roles to call the Service Control APIs.
Updated IAM policy for project [YOUR_PROJECT].
bindings:
- members:
  - serviceAccount:stores@YOUR_PROJECT.iam.gserviceaccount.com
  role: roles/cloudtrace.agent
- members:
  - serviceAccount:service-622059496897@containerregistry.iam.gserviceaccount.com
  role: roles/containerregistry.ServiceAgent
- members:
  - serviceAccount:622059496897-compute@developer.gserviceaccount.com
  role: roles/editor
- members:
  - user:YOUR_EMAIL
  role: roles/owner
- members:
  - serviceAccount:service-622059496897@gcp-sa-pubsub.iam.gserviceaccount.com
  role: roles/pubsub.serviceAgent
- members:
  - serviceAccount:service-622059496897@serverless-robot-prod.iam.gserviceaccount.com
  role: roles/run.serviceAgent
- members:
  - serviceAccount:stores@YOUR_PROJECT.iam.gserviceaccount.com
  role: roles/servicemanagement.serviceController
etag: BwYlom1WNEs=
version: 1
Updated IAM policy for project [YOUR_PROJECT].
bindings:
- members:
  - serviceAccount:stores@YOUR_PROJECT.iam.gserviceaccount.com
  role: roles/cloudtrace.agent
- members:
  - serviceAccount:service-622059496897@containerregistry.iam.gserviceaccount.com
  role: roles/containerregistry.ServiceAgent
- members:
  - serviceAccount:622059496897-compute@developer.gserviceaccount.com
  role: roles/editor
- members:
  - user:YOUR_EMAIL
  role: roles/owner
- members:
  - serviceAccount:service-622059496897@gcp-sa-pubsub.iam.gserviceaccount.com
  role: roles/pubsub.serviceAgent
- members:
  - serviceAccount:service-622059496897@serverless-robot-prod.iam.gserviceaccount.com
  role: roles/run.serviceAgent
- members:
  - serviceAccount:stores@YOUR_PROJECT.iam.gserviceaccount.com
  role: roles/servicemanagement.serviceController
etag: BwYlom1zMec=
version: 1
Creating the Cloud Run container.
Applying new configuration to Cloud Run service [stores] in project [YOUR_PROJECT] region [us-west1]
OK Deploying new service... Done.                                                                                                                                                                                                                                
  OK Creating Revision...                                                                                                                                                                                                                                        
  OK Routing traffic...                                                                                                                                                                                                                                          
Done.                                                                                                                                                                                                                                                            
New configuration has been applied to service [stores].
URL: https://YOUR_HOST.us-west1.run.app
Configuring IAM to allow outside access to the container.
Updated IAM policy for service [stores].
bindings:
- members:
  - allUsers
  role: roles/run.invoker
etag: BwYlonNoIzE=
version: 1
```

A lot just happened! Here's a summary:
1. Three runs of `gcloud services enable` turned on three important Google APIs used by the demo.
2. `gcloud endpoints services deploy` uploaded a description of your demo service to Service Management.
3. `gcloud iam service-accounts create` created a service account that is used by your demo service.
4. Two calls to `gcloud projects add-iam-policy-binding` gave necessary permissions to the service account.
5. `gcloud run services replace` used the configuration in `service.yaml` to create your demo service on Cloud Run. We use "replace" to allow this to be called again to upload modified `service.yaml` files with alternate configurations.
6. `gcloud run services set-iam-policy` removed a restriction on your Cloud Run service and opens it to public access. We don't need this restriction because we are protecting the service with a proxy (look inside `service.yaml` for a sneak peek at this).

## Call your demo API with your browser

Now go to the Cloud Run section of the Google Cloud Console. You'll see a list of services, and if you started with a fresh project, there will be just one, like in the screenshot below:

![alt text](/screenshots/cloud-run-services.png)

Click on the name of your service to open a detail view like the one below.

![alt text](/screenshots/cloud-run-service-detail.png)

Near the center, you'll see a highighted URL for your service. Click on that to call your service from your browser. You'll get a response like this:

```
{
    "code": 404,
    "message": "The current request is not defined by this API."
}
```

That's good, because it's true. The root path is not defined by your API. To call your API, add `/v1/stores` to the path.

After you've done that, you'll get JSON like this:
```
{
    "stores": [
        {
            "name": "stores/0",
            "type": "office",
            "title": "Columbus, NM 88029",
            "location": {
                "latitude": 31.8301201,
                "longitude": -107.638199
            },
            "address": {
                "street": "South Main Street",
                "regionCode": "us"
            }
        },
        {
            "name": "stores/1",
            "type": "office",
            "title": "United States Post Office - Gateway Station",
            "location": {
                "latitude": 37.7949409,
                "longitude": -122.399437
            },
            "address": {
                "zipCode": 94111,
                "regionCode": "us"
            }
        },
        ...47 more entries...
        {
            "name": "stores/49",
            "type": "office",
            "title": "Sheboygan Falls Post Office",
            "location": {
                "latitude": 43.7295609,
                "longitude": -87.8121719
            },
            "address": {
                "street": "Maple Street",
                "city": "Sheboygan Falls",
                "state": "WI",
                "zipCode": 53085,
                "regionCode": "us"
            }
        }
    ],
    "nextPageToken": "NTA"
}
```

Now try to get an individual store by changing the path to `/v1/stores/0`.

You'll get an error.
```
{
    "message": "UNAUTHENTICATED: Method doesn't allow unregistered callers (callers without established identity). Please use API Key or other form of API consumer identity to call this API.",
    "code": 401
}
```

This is because this method requires an API key. You can get one with `gcloud`, but note that you'll need to set the `PROJECT` environment variable to the name of your project. Here we set ours to "bobadojo":
```prompt
PROJECT=bobadojo
```

```prompt
gcloud services api-keys create --key-id demo \
  --api-target service=stores.endpoints.$PROJECT.cloud.goog 
```
```output
Operation [operations/akmf.p7-622059496897-0a3a0229-45c0-484d-b770-2f63879d63cb] complete. Result: {
    "@type":"type.googleapis.com/google.api.apikeys.v2.Key",
    "createTime":"2024-10-31T15:26:42.203902Z",
    "etag":"W/\"2j5qS7GWBEUrgsfYdqi0rQ==\"",
    "keyString":"YOUR_KEY",
    "name":"projects/622059496897/locations/global/keys/demo",
    "restrictions":{
        "apiTargets":[
            {
                "service":"stores.endpoints.YOUR_PROJECT.cloud.goog"
            }
        ]
    },
    "uid":"91a05581-a418-4b80-8f5f-18fbc5ce7485",
    "updateTime":"2024-10-31T15:26:42.232775Z"
}
```

In your actual output, "YOUR_KEY" in the response above would be your actual key string.

Add it to your request in the browser by appending it to the path.
```
https://YOUR_HOST.us-west1.run.app/v1/stores/0?api_key=YOUR_KEY
```

## Call your demo API with `curl`

You can also get the key to use at the command line with `gcloud` and `jq`:
```prompt
KEY=$(gcloud services api-keys get-key-string \
  projects/$PROJECT/locations/global/keys/demo --format json \
  | jq .keyString -r)
```

`gcloud` can also give you the hostname for your service:
```prompt
HOST=$(gcloud run services describe stores --format json \
  | jq .status.address.url -r)
```

Now you can easily call your service with `curl`:
```prompt
curl -s $HOST/v1/stores/0 -H "X-Api-Key: $KEY" | jq
```
```
{
  "name": "stores/0",
  "type": "office",
  "title": "Columbus, NM 88029",
  "location": {
    "latitude": 31.8301201,
    "longitude": -107.638199
  },
  "address": {
    "street": "South Main Street",
    "regionCode": "us"
  }
}
```
Note that here we are passing the API key in a header (`X-Api-Key`). We could have also used the `api_key` query parameter:

```prompt
curl -s "$HOST/v1/stores/0?api_key=$KEY" | jq
```
```
{
  "name": "stores/0",
  "type": "office",
  "title": "Columbus, NM 88029",
  "location": {
    "latitude": 31.8301201,
    "longitude": -107.638199
  },
  "address": {
    "street": "South Main Street",
    "regionCode": "us"
  }
}
```

## View your service in the Endpoints console

Now find the Endpoints link in the Cloud Console sidebar (on the left) and select it. You'll find a list of services like the one below.

![alt text](/screenshots/endpoints-services.png)

Select your service by clicking on either the service title or the service name. You'll open a detail screen that looks something like this:

![alt text](/screenshots/endpoints-service-detail.png)

For the screenshot above, we've selected to display only the Requests graph so that the screenshot will include the log of operations at the bottom. The graph and logs above show a lot more traffic than you should expect (for this instance, a cron job is making a batch of API calls every five minutes). But if you select "View logs" on a row corresponding to one of the methods you've called, you'll find logs of your requests in the Logs Explorer view.

![alt text](/screenshots/endpoints-log-getstore.png)

Click on the '>' on the left of one of the log entries to expand it. If you expand all the subsections, you'll find a lot of detail! You'll also notice that if you are looking at the log of a call that used an API key, the key is included in the log! It's blacked out in the screenshot below. You'll want to be careful what you do with these logs (we'll work on this problem later).

![alt text](/screenshots/endpoints-log-getstore-detail.png)

---
#### Continue with [Service Infrastructure](/docs/serviceinfrastructure).