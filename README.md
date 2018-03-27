# Chunked Transfer Demo Server
This is a web server demo that serves a static HTTP Live Streaming (HLS) manifest and video segments. The video segments are returned using HTTP/1.1 chunked transfer encoding.

## Building
```
docker build . -t muxinc/chunked-transfer-demo
```

## Running
```
docker run -p 8080:8080 muxinc/chunked-transfer-demo:latest
```

## Accessing Media
### HLS Manifest
```
ffplay http://localhost:8080/manifest.m3u8
```

### Individual Segments
There are segments numbered 0 through 5 which can be accessed as follows:
```
ffplay http://localhost:8080/ts/0.ts
```
