# auth-svc

## What is this?
HTTP layer for authenticating user requests to and from orca-svc

## Dependencies?
For now, just redis.

## Why Redis?
Cause I'm too lazy to make a proof of concept with a SQL database.

## What's the schema?
```json
{
  "users": {
    "username-one" : {
      "password": "abcdef",
      "authed_devices": [
        "device-one"
      ]
    },
    "username-two" : {
      "password": "abcdef",
      "authed_devices": [
        "device-two"
      ]
    },
    "username-three" : {
      "password": "abcdef",
      "authed_devices": [
        "device-one",
        "device-two"
      ]
    } 
  }
}
```


## That seems terribly unoptimal and insecure
Yeah it is.