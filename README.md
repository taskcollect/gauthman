# gauthman
Taskcollect's Google Authentication Manager.

A small microservice with one purpose: take care of token exchange with Google.

## HTTP Spec

`/v1/exchange`

**Query Parameters**
- `code` - the one-time access code from the frontend

**Returns**
```jsonc
{
  "access": "token here", // access token returned by google
  "refresh": "other token here", // refresh token returned by google
  "expires": 1639630704 // when the access token expires
}
```

### IMPORTANT NOTE!
Google will only return a refresh token if this is the user's first time authorizing with the application UNLESS you specify `prompt=consent` in the clientside sign-in code!
