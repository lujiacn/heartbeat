# Revel heartbeat module


### How to

1. In app.conf add module

``module.heartbeat=github.com/lujiacn/heartbeat``

2. add in routes
```
module:heartbeat

#GET  /heartbeat                              HeartBeat.Index
GET  /heartbeat

```
3. clone to github.com/lujiacn/heartbeat
