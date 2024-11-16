---
weight: 4
title: The API Keys API
---
## The API Keys API

The API Keys API creates and manages API keys that consumers use to call APIs.

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

### CreateKey

Let's use `q` to create an API key using the `CreateKey` method. We'll create a second key like the one that we created in the [quickstart](/demo/quickstart/demo). Recall that there we used this `gcloud` command:
```
$ KEY=$(gcloud services api-keys get-key-string projects/YOUR_PROJECT/locations/global/keys/demo --format json | jq .keyString -r)
```

Now doing the same thing with `q`:
```
q api-keys create-key --parent projects/bobadojo/locations/global --keyid demo2 --service stores.endpoints.bobadojo.cloud.goog | jq
{
  "name": "operations/akmf.p7-1046800315646-a778a84e-6829-4426-b26c-7bd71525cc8c"
}
```

```
KEY=$(q api-keys get-key-string projects/bobadojo/locations/global/keys/demo2 | jq .keyString -r)
```

```
$ curl -s https://stores-1046800315646.us-west1.run.app/v1/stores/0 -H "X-Api-Key: $KEY" | jq
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

### ListKeys

```
$ q api-keys list-keys bobadojo | jq
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

```
$ q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
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

```
$ q api-keys get-key-string projects/bobadojo/locations/global/keys/demo2 | jq
{
  "keyString": "REDACTED"
}

```

### UpdateKey

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


```
$ q api-keys update-key demo2.jq 
{"name":"operations/akmf.p10-1046800315646-274fe790-2c2f-409f-8b58-e2494a489259"}
```

```
$ q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
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

```
$ q api-keys delete-key projects/bobadojo/locations/global/keys/demo2 
{"name":"operations/akmf.p12-1046800315646-3934e2bb-5b66-4e68-b640-753b5fbf60a1"}
```

```
$ q api-keys get-key projects/bobadojo/locations/global/keys/demo2 | jq
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

```
$ curl -s https://stores-1046800315646.us-west1.run.app/v1/stores/0 -H "X-Api-Key: $KEY" | jq
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

It still works! That's because our proxy is caching the key.

```
$ curl -s https://stores-1046800315646.us-west1.run.app/v1/stores/0 -H "X-Api-Key: $KEY" | jq
{
  "code": 400,
  "message": "INVALID_ARGUMENT: API key expired. Please renew the API key."
}
```

### UndeleteKey

```
$ q api-keys undelete-key projects/bobadojo/locations/global/keys/demo2 
{"name":"operations/akmf.p13-1046800315646-d95498f7-4a62-400d-b12e-2e31291e910b"}

$ curl -s https://stores-1046800315646.us-west1.run.app/v1/stores/0 -H "X-Api-Key: $KEY" | jq
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

### LookupKey

```
$ q api-keys lookup-key $KEY | jq
{
  "parent": "projects/1046800315646/locations/global",
  "name": "projects/1046800315646/locations/global/keys/demo2"
}

```

## The Operations service

This is the same service that we discussed for the Service Management API, so we won't discuss it further here.

## Usage Notes

### What happens if someone uses an invalid key?

We can test this by creating an API key with another GCP account and using that key to call our API.

```
$ curl 'https://stores-1046800315646.us-west1.run.app/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI'  -i
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

When they search for the service, they won't find it. It is only visible to the projects have been granted access to it.
![alt text](/screenshots/service-not-found.png)


Delete the key and you'll get this response:
```
$ !curl
curl 'https://stores-1046800315646.us-west1.run.app/v1/stores/0?api_key=AIzaSyBPgB7_IGKATETdWcrYvolr4-LuECEL6uI'  -i
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