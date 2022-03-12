# Product Images

## Uploading

Note: need to use `--data-binary` to ensure file is not converted to text

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @test.png
```
---
* show how to use gzip compression with our services, so that we can compress
files and make them just really quick response from the server with go std lib

* http pkg, a way of uploading file: go is pretty straight forward 
because as part of any handler this http request, and inside http request
we have a stream the request body
* when you send a request to a GO service, GO isn't immediately buffering
all the data that's sent to it, it can give the capability of doing it 
gradually, can read the data as it's kind of sent up to the wire which
means that can do a lot of things like making sure that we don't accept too
much data, we don't stream 10 giga of data when actually the max file upload
is 3 megabytes , and using ioReader  and being able to control the numbers 
of bytes are read 