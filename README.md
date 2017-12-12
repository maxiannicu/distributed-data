# Distributed Web Platform

## Deployment Diagram

Following diagram shows the deployment diagram of platform. Here is described basic flow of system :

- First of all all connections between data nodes are created and mediator is started.
- When a client opens connection to mediator, mediator searches master node in `node group` (node with maximum number of elements) using Multicast.
- After finding master node, it opens a TCP connection to this and requests data. Note that PORT is found in discovery process, what means that mediator is decoupled from Nodes.
- When connection is estabilished, node requires all data from connected nodes. This process is a breadth-first aggregation.
- Once all data is aggregated is returned to mediator.
- Mediator converts into `accept` format specified by client and sends it back.

![Deployment Diagram](/images/deployment_diagram.png)

## Configuration

Configuration is stored in folder `configurations` and it's essential for running platform.

### Mediator

```json
{
  "listenEndPoint": {
    "host" : "127.0.0.1",
    "port" : 31012
  },
  "discoveryEndPoint": {
    "host" : "224.0.0.1",
    "port" : 31013
  },
  "discoveryDuration" : 2000
}
```

### Node

Nodes should stored with naming convention of `node.XXXX.json` where XXXX - is anything you want. This config difines identificator and node's connections. Also it has data inside. There is also an `contentTyp` which specify data format in server. In our case is XML. When a node will ask `node_london` for data, as a response will receive a XML data.

```json
{
  "identificator": "node_london",
  "connections": [
    "node_chisinau",
    "node_tokyo"
  ],
  "contentType": 1,
  "discoveryEndPoint": {
    "host": "224.0.0.1",
    "port": 31013
  },
  "data": [
    {
      "firstName": "Greg",
      "lastName": "Young",
      "age": 17
    },
    {
      "firstName": "Michael",
      "lastName": "Stewart",
      "age": 23
    }
  ]
}
```

## Running

In order to run platform, you have to run following GO main files, with respect to order:

### Node Runner

Run `node_runner.go` with Working Directory to `configurations`

```
[node_runner] 2017/12/12 01:58:47.782242 Loaded 3 node configurations
[node_london] 2017/12/12 01:58:47.782314 Starting node
[node_london] 2017/12/12 01:58:47.782318 Data size = 2
[node_london] 2017/12/12 01:58:47.782320 Starting tcp server
[node_london] 2017/12/12 01:58:47.782406 Tcp server started 127.0.0.1:38057
[node_chisinau] 2017/12/12 01:58:47.782459 Starting node
[node_chisinau] 2017/12/12 01:58:47.782462 Data size = 7
[node_chisinau] 2017/12/12 01:58:47.782464 Starting tcp server
[node_chisinau] 2017/12/12 01:58:47.782481 Tcp server started 127.0.0.1:45335
[node_tokyo] 2017/12/12 01:58:47.782495 Starting node
[node_tokyo] 2017/12/12 01:58:47.782498 Data size = 4
[node_tokyo] 2017/12/12 01:58:47.782499 Starting tcp server
[node_tokyo] 2017/12/12 01:58:47.782512 Tcp server started 127.0.0.1:41511
[node_london] 2017/12/12 01:58:47.782530 Connecting to 127.0.0.1:45335
[node_london] 2017/12/12 01:58:47.782545 Listening for discovery
[node_tokyo] 2017/12/12 01:58:47.782637 Listening for discovery
[node_london] 2017/12/12 01:58:47.782647 Connected 127.0.0.1:45335
[node_london] 2017/12/12 01:58:47.782655 Connecting to 127.0.0.1:41511
[node_london] 2017/12/12 01:58:47.782770 Connected 127.0.0.1:41511
[node_chisinau] 2017/12/12 01:58:47.782778 Connecting to 127.0.0.1:38057
[node_chisinau] 2017/12/12 01:58:47.782844 Listening for discovery
[node_chisinau] 2017/12/12 01:58:47.783082 Connected 127.0.0.1:38057
[node_tokyo] 2017/12/12 01:58:47.783092 Connecting to 127.0.0.1:38057
[node_tokyo] 2017/12/12 01:58:47.783239 Connected 127.0.0.1:38057
```

### Mediator

Run `mediator.go` with Working Directory to `configurations`

```
[mediator] 2017/12/12 01:58:50.646385 Starting TCP server
[mediator] 2017/12/12 01:58:50.646492 TCP server started on 127.0.0.1:31012
[mediator] 2017/12/12 01:58:50.646495 Starting UDP sender
[mediator] 2017/12/12 01:58:50.646517 UDP sender started
[mediator] 2017/12/12 01:58:50.646518 Starting UDP listener
[mediator] 2017/12/12 01:58:50.646566 UDP listener started on 127.0.0.1:37084
```

### Client

Run `mediator.go`

## Client Response

```json
{
  "Data": [
    {
      "FirstName": "Greg",
      "LastName": "Young",
      "Age": 17
    },
    {
      "FirstName": "Nicolae",
      "LastName": "Botnari",
      "Age": 22
    },
    {
      "FirstName": "Maxim",
      "LastName": "Bircu",
      "Age": 22
    },
    {
      "FirstName": "Vladimir",
      "LastName": "Voronin",
      "Age": 22
    },
    {
      "FirstName": "Mihail",
      "LastName": "Botnari",
      "Age": 22
    },
    {
      "FirstName": "Yu",
      "LastName": "Kim",
      "Age": 22
    },
    {
      "FirstName": "Michael",
      "LastName": "Stewart",
      "Age": 23
    },
    {
      "FirstName": "Ji Sun",
      "LastName": "Park",
      "Age": 24
    },
    {
      "FirstName": "Ji",
      "LastName": "Kim",
      "Age": 24
    },
    {
      "FirstName": "Ion",
      "LastName": "Druta",
      "Age": 33
    },
    {
      "FirstName": "Vlad",
      "LastName": "Plahotniuc",
      "Age": 54
    },
    {
      "FirstName": "Son",
      "LastName": "Kim",
      "Age": 54
    },
    {
      "FirstName": "Vladislav",
      "LastName": "Vladislavovici",
      "Age": 89
    }
  ],
  "Size": 13
}
```