simple-http-client
==================

<2MB Docker image to use when you just want to run an http request.

### Usage
```
# Usage: METHOD URL [BODY]
$ simple-http-client POST https://example.com/
1970/01/01 00:00:00 POST https://example.com/ 200
```

### Setting timeout
Default timeout is 5 seconds.
```
$ HTTP_TIMEOUT=15 simple-http-client POST https://example.com/
1970/01/01 00:00:00 POST https://example.com/ 200
```