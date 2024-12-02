---
weight: 4
title: The API Keys API
---
## The API Keys API

The API Keys API creates and manages API keys that consumers use to call APIs. Anyone can use these API keys to call an API. There is no need that they be a Google customer or even have an identity known to Google. If you're a service owner, the API Keys API is an easy way for you to give individually-trackable access to your users.

The API Keys API is defined in the [googleapis](/docs/details/googleapis) repo in [apikeys_v2.yaml](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys_v2.yaml). It includes two services.

| Service | Purpose |
| ------- | ------- |
| [ApiKeys](#the-apikeys-service) | Methods managing the API keys associated with projects |
| [Operations](#the-operations-service) | A mix-in that handles long-running operations |

## The ApiKeys service

The [ApiKeys](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys.proto#L37) service is defined in [apikeys.proto](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/apikeys.proto) and provides support for managing and using API keys.

The full names of these methods begin with `google.api.apikeys.v2.ApiKeys.`

| Method | Description |
| ------ | ----------- |
| [CreateKey](#createkey) | Creates a new API key |
| [ListKeys](#listkeys) | Lists the API keys owned by a project |
| [GetKey](#getkey) | Gets the metadata for an API key |
| [GetKeyString](#getkeystring) | Get the key string for an API key |
| [UpdateKey](#updatekey) | Patches the modifiable fields of an API key |
| [DeleteKey](#deletekey) | Deletes an API key |
| [UndeleteKey](#undeletekey) | Undeletes an API key which was deleted within 30 days |
| [LookupKey](#lookupkey) | Find the parent project and resource name of the API key that matches the key string in the request |

We'll work through these methods, calling each with `q` for the example project that we set up in the [quickstart](/docs/quickstart).

### **Keys**

Unsurprisingly, the main resource of this service is the [Key](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/resources.proto#L32), described in the proto as "the representation of a key managed by the API Keys API." 

Here's a JSON representation of a key:
```
{
  "name": "projects/1046800315646/locations/global/keys/demo",
  "uid": "cafe4d44-ff21-4273-a429-4dd0426ee8cd",
  "createTime": "2024-11-04T19:21:35.811160Z",
  "updateTime": "2024-11-16T16:57:49.195879Z",
  "restrictions": {
    "apiTargets": [
      {
        "service": "stores.endpoints.bobadojo.cloud.goog"
      }
    ]
  },
  "etag": "W/\"/ytdCWujsYbVbVftxBeBsg==\""
}
```

Some of its fields are worth describing:

- `name` is the name of the key. It includes the project id to make it globally unique.

- `uid` is a system-generated unique id. If no key id was specified when the key was created, this would also be the id of the key.

- `display_name` is an optional human-readable name of the key. This can be modified after the key is created.

- `key_string` is the key value. It is only returned by the `GetKeyString` method.

- `create_time` and `update_time` are timestamps representing the time of creation and last update.

- `delete_time` is only provided if a key has been deleted and is set to the time that they key was deleted. Note that this means that deleted keys are kept.

- `restrictions` specify controls on key usage. We use this to restrict our key to only be usable in calls to our demo API. The kinds of restrictions that we can specify are listed in the description of the [Restrictions](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/resources.proto#L93) message. They include HTTP referrers, calling IP addresses, Android and iOS app ids, and specific API targets. For API targets, keys can be restricted for use [with specific methods only](https://github.com/googleapis/googleapis/blob/master/google/api/apikeys/v2/resources.proto#L174).
 
- `etag` is a checksum that we can use to verify any calls we make to update a key. If we send an etag that doesn't match the current value, our update request will be rejected.

### CreateKey

Let's use `q` to create an API key using the `CreateKey` method. We'll create a second key like the one that we created in the [quickstart](/demo/quickstart/demo). Recall that there we used this `gcloud` command:
```prompt
KEY=$(gcloud services api-keys get-key-string \
  projects/YOUR_PROJECT/locations/global/keys/demo --format json | jq .keyString -r)
```

Now let's do the same thing with `q`:
```prompt
q api-keys create-key --parent projects/bobadojo/locations/global \
    --keyid demo2 \
    --service stores.endpoints.bobadojo.cloud.goog | jq
```

```
{
  "name": "operations/akmf.p7-1046800315646-a778a84e-6829-4426-b26c-7bd71525cc8c"
}
```

This returned an id of a longrunning operation that, when complete, results in a new API key. It's quick, so we won't wait for it here.

### ListKeys

We can list the keys in a project with `ListKeys`, and now when we do, we'll see the new key that we created along with the one we made in the [quickstart](/docs/quickstart/demo).

```prompt
q api-keys list-keys bobadojo | jq
```

```
{
  "keys": [
    {
      "name": "projects/1046800315646/locations/global/keys/demo2",
      "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
      "createTime": "2024-11-16T17:02:44.459793Z",
      "updateTime": "2024-11-16T17:02:44.484273Z",
      "restrictions": {
        "apiTargets": [
          {
            "service": "stores.endpoints.bobadojo.cloud.goog"
          }
        ]
      },
      "etag": "W/\"I57Ym+/tTkfOlWmuJrPVnQ==\""
    },
    {
      "name": "projects/1046800315646/locations/global/keys/demo",
      "uid": "cafe4d44-ff21-4273-a429-4dd0426ee8cd",
      "createTime": "2024-11-04T19:21:35.811160Z",
      "updateTime": "2024-11-16T16:57:49.195879Z",
      "restrictions": {
        "apiTargets": [
          {
            "service": "stores.endpoints.bobadojo.cloud.goog"
          }
        ]
      },
      "etag": "W/\"/ytdCWujsYbVbVftxBeBsg==\""
    }
  ]
}
```

### GetKey

We can also use `GetKey` to directly get a key by specifying its name. Note that we can use either the string or numeric versions of the project id in this name (both work). Also note that the key string is not returned.

```prompt
q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
```

```
{
  "name": "projects/1046800315646/locations/global/keys/demo2",
  "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
  "createTime": "2024-11-16T17:02:44.459793Z",
  "updateTime": "2024-11-16T17:02:44.484273Z",
  "restrictions": {
    "apiTargets": [
      {
        "service": "stores.endpoints.bobadojo.cloud.goog"
      }
    ]
  },
  "etag": "W/\"I57Ym+/tTkfOlWmuJrPVnQ==\""
}

```

### GetKeyString

We can get the key string with `GetKeyString`. This is a separate method to reduce opportunities for accidental exposure of keys.

```prompt
q api-keys get-key-string projects/bobadojo/locations/global/keys/demo2 | jq
```

```
{
  "keyString": "REDACTED"
}
```

That's not the simplest representation for us to use, but we can easily extract the `keyString` field with `jq`:

```
KEY=$(q api-keys get-key-string projects/bobadojo/locations/global/keys/demo2 | jq .keyString -r)
```

Now we can call our API with this new key. We'll also use the `HOST` variable that we set in the [quickstart](/demo/quickstart/demo).

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

### UpdateKey

`UpdateKey` lets us modify properties of a key. We can use it to modify the display name and the restrictions on a key. Let's use `q` to add a display name to our new key. First let's get a JSON representation of the key:
```prompt
q api-keys get-key projects/bobadojo/locations/global/keys/demo2 > demo2.json
```

Now edit `demo2.json` to add the `displayName` line below:
```
{
  "name": "projects/1046800315646/locations/global/keys/demo2",
  "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
  "displayName": "Demo key",
  "createTime": "2024-11-16T17:02:44.459793Z",
  "updateTime": "2024-11-16T17:02:44.484273Z",
  "restrictions": {
    "apiTargets": [
      {
        "service": "stores.endpoints.bobadojo.cloud.goog"
      }
    ]
  },
  "etag": "W/\"I57Ym+/tTkfOlWmuJrPVnQ==\""
}
```

Finally, use `q` to update the key.
```prompt
q api-keys update-key demo2.json | jq 
```

```
{"name":"operations/akmf.p10-1046800315646-274fe790-2c2f-409f-8b58-e2494a489259"}
```

This returns a quick-to-complete operation that updates the key. We can verify the change with `q`:
```prompt
q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
```

```
{
  "name": "projects/1046800315646/locations/global/keys/demo2",
  "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
  "displayName": "Demo key",
  "createTime": "2024-11-16T17:02:44.459793Z",
  "updateTime": "2024-11-16T17:06:31.841817Z",
  "restrictions": {
    "apiTargets": [
      {
        "service": "stores.endpoints.bobadojo.cloud.goog"
      }
    ]
  },
  "etag": "W/\"Y1Y7yLy/aqK4oVGAeedn3w==\""
}
```

### DeleteKey

`DeleteKey` lets us delete keys. Let's try it.
```prompt
q api-keys delete-key projects/bobadojo/locations/global/keys/demo2 
```

```
{"name":"operations/akmf.p12-1046800315646-3934e2bb-5b66-4e68-b640-753b5fbf60a1"}
```

Now when we get our key, we find that it still exists, but now it has a value for `deleteTime` that shows that it has been deleted.
```prompt
q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
```

```
{
  "name": "projects/1046800315646/locations/global/keys/demo2",
  "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
  "displayName": "Demo key",
  "createTime": "2024-11-16T17:02:44.459793Z",
  "updateTime": "2024-11-16T17:07:58.456224Z",
  "deleteTime": "2024-11-16T17:07:58.415258Z",
  "restrictions": {
    "apiTargets": [
      {
        "service": "stores.endpoints.bobadojo.cloud.goog"
      }
    ]
  },
  "etag": "W/\"Y1Y7yLy/aqK4oVGAeedn3w==\""
}
```

We can try to use our deleted key, and we'll find that for a while, it still works:
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

It still works because our proxy is caching the key. Wait a while (seriously, a few minutes... go make a fresh pot of coffee) and then try it again.
```prompt
curl -s $HOST/v1/stores/0 -H "X-Api-Key: $KEY" | jq
```

```
{
  "code": 400,
  "message": "INVALID_ARGUMENT: API key expired. Please renew the API key."
}
```

### UndeleteKey

Deleted keys are retained so that we can use `UndeleteKey` if we decide to restore them. Let's try it:

```prompt
q api-keys undelete-key projects/bobadojo/locations/global/keys/demo2 
```

```
{"name":"operations/akmf.p13-1046800315646-d95498f7-4a62-400d-b12e-2e31291e910b"}
```

This operation completes quickly, and we can quickly see its effects by using the undeleted key:
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
\
```

### LookupKey

The last method of the API is `LookupKey`, and with it, we can see how service implementations might verify an API key string.

```prompt
q api-keys lookup-key $KEY | jq
```

```
{
  "parent": "projects/1046800315646/locations/global",
  "name": "projects/1046800315646/locations/global/keys/demo2"
}
```

Here we've passed our key string and received the name of the key that corresponds to it. If we need to know details about the key, we can use `GetKey` to retrieve them.

## The Operations service

This is the same service that we discussed for the Service Management API, so we won't discuss it in detail here. But just as we did for Service Management, if we want to check the status of operations, we can call `GetOperation`, but this time, the target of our API call should be the API Keys service.

Here is a quick example that uses `curl` to check on the status of the `UndeleteKey` request that we made earlier:
```prompt
curl https://apikeys.googleapis.com/v1/operations/akmf.p13-1046800315646-d95498f7-4a62-400d-b12e-2e31291e910b -H "Authorization: Bearer $(gcloud auth print-access-token)"
```

```
{
  "name": "operations/akmf.p13-1046800315646-d95498f7-4a62-400d-b12e-2e31291e910b",
  "done": true,
  "response": {
    "@type": "type.googleapis.com/google.api.apikeys.v2.Key",
    "name": "projects/1046800315646/locations/global/keys/demo2",
    "displayName": "Demo key",
    "keyString": "AIzaSyBeIXSjqgFW7FJndJYF56sCG62U2kwTQYE",
    "createTime": "2024-11-16T17:02:44.459793Z",
    "uid": "2bd9e341-1e78-42a0-ab13-2c667632f83f",
    "updateTime": "2024-11-16T17:22:20.052262Z",
    "restrictions": {
      "apiTargets": [
        {
          "service": "stores.endpoints.bobadojo.cloud.goog"
        }
      ]
    },
    "etag": "W/\"Y1Y7yLy/aqK4oVGAeedn3w==\""
  }
}
```

As another point of interest, if we try calling `ListOperations`, we can see that it is not implemented for the API Keys API.
```prompt
curl https://apikeys.googleapis.com/v1/operations -H "Authorization: Bearer $(gcloud auth print-access-token)"
```

```
{
  "error": {
    "code": 404,
    "message": "Method not found.",
    "status": "NOT_FOUND"
  }
}
```

## Usage Notes

### What happens if someone uses an invalid key?

We know that keys are associated with projects and that keys can only be used for APIs that are enabled for the associated project and that match any restrictions on those keys.

What happens when someone tries to use a key associated with some other project to call our API? We can explore this by creating an API key with another Google Cloud account and using that key to call our API.

```prompt
curl "$HOST/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI"  -i
```

```
HTTP/2 403 
content-type: application/json
x-envoy-decorator-operation: ingress GetStore
x-cloud-trace-context: a8d45728e34f1d205dff18bf9a922ff4
date: Fri, 08 Nov 2024 15:49:12 GMT
server: Google Frontend
content-length: 117
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000

{"message":"PERMISSION_DENIED: API stores.endpoints.bobadojo.cloud.goog is not enabled for the project.","code":403}
```

The message is accurate and helpful -- and when this user searches for this service in the Cloud Console, they won't find it. It is only visible to the projects have been granted access to it.
![alt text](/screenshots/service-not-found.png)

If that key is deleted and used again, the caller will get this response:
```prompt
!curl
```

```
curl "$HOST/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI"  -i
HTTP/2 400 
content-type: application/json
x-envoy-decorator-operation: ingress GetStore
x-cloud-trace-context: 0efb7df218e906bf48255cfd072af922;o=1
date: Fri, 08 Nov 2024 15:58:13 GMT
server: Google Frontend
content-length: 91
alt-svc: h3=":443"; ma=2592000,h3-29=":443"; ma=2592000

{"code":400,"message":"INVALID_ARGUMENT: API key not found. Please pass a valid API key."}
```

## Summarizing

We've seen that the API Keys API makes it easy to create and manage API keys. We can use it to create keys for our API users, and because Service Control logs those keys, we can observe their usage with the Cloud Logging API, which we examine next.

Anyone can use these API keys to call your API. There is no need that they be a Google customer or even have an identity known to Google.

However, Google limits the API Keys API to allow [at most 300 keys per project](https://cloud.google.com/docs/authentication/api-keys#limits). If you need more, we can suggest a few ways to adapt to this:

1. Create new projects as needed using the [CreateProject](https://github.com/googleapis/googleapis/blob/master/google/cloud/resourcemanager/v3/projects.proto#L91) method in the [Cloud Resource Manager API](https://cloud.google.com/resource-manager/reference/rest). Then you could use these projects as pools of API keys. (Note that this link is to REST documentation for the Cloud Resource Manager API. It is also available as a gRPC service, but that is currently undocumented.)

2. Expanding on the above approach, you could create your projects in [tenancy units](https://cloud.google.com/service-infrastructure/docs/glossary#tenancy) that are dedicated to individual customers and managed with the [Service Consumer Management API](https://cloud.google.com/service-infrastructure/docs/service-consumer-management/getting-started) (The Service Consumer API appears to only be available as a REST service.)

3. If your API users are also Google Cloud users, you can grant them access to your service using the [SetIamPolicy](/docs/serviceinfrastructure/servicemanagement#setiampolicy) method of the Service Management API. Then they can use the [Service Usage API](/docs/serviceinfrastructure/serviceusage) to enable your service in their own Google Cloud Project and create their own API keys with the API Keys API.

---
#### Continue with [the Cloud Logging API](/docs/serviceinfrastructure/cloudlogging).