docker-commander
=====
A GUI to manage Docker

## Some info
* Project structure is defined like [here](https://github.com/golang/go/wiki/GithubCodeLayout):
* We're using [gin-gonic](https://github.com/gin-gonic/gin)
* We'll use [Docker Swarm](https://github.com/docker/swarm) probably and / or Ansible.

## Glossary
* A `Host` is a machine, virtual or not, accessible from the server that is going to host Docker containers.
* The `Server` is the machine hosting Docker Commander
* A `Container` is a Docker container that is being host in a *Host*.
* A `Commander` is a User of the application. Has `Hosts` associated.

## MoSCoW Requirements

### Host Requirements
* As the *Commander*, I *Must* __know__ __if__ a __*Host* is down__ or not so I'll know as fast as possible that one or more containers could be down
  * *Acceptance criteria:*
    * A registered host is turned off. It appears like *down* in the dashboard
    * A registered host is on. It appears like *online* in the dashboard
    * A registered *Host* that appears like *online* is turned off. The app shows *down* after the *refreshing period* has passed
    * A registered *Host* that appears like *down* is turned on. The app shows *online* after the *refreshing period* has passed.


* As the *Commander*, I *Must* __know the Images__ a *Host* is hosting so I have some insight of the host load
  * Acceptance criteria:
    * A *Host* without *Images* appears like empty.
    * A *Host* with 2 *Images* appears correctly.
    * In a *Host* with 2 *Images* we add one more. 3 *Images* must appear after *refreshing period* has passed
    * In a *Host* with 3 *Imags* we delte one. 2 *Images* must appear after *refreshing period* has passed.


* As the *Commander*, I *Should* __know the latency__ I'm getting against a *Host* so I can detect potential connectivity problems
  * Acceptance criteria:
    * In each *refreshing period* cycle, a *latency* for each *Host* must appear in ms must appear in the dashboard.
   * A *Host* is turned off. *Latency* must show *timeout*.


* As the *Commander*, I *Must* __know the running *Containers*__ in a *Host* so I have full knowledge of the *Host's* *Containers*
  * Acceptance criteria:
    * No *Containers* are running. A *Container* is run. It appears after *refreshing period* has passed
    * A *Container* is running. The *Container* is killed. It dissapears from the dashboard after the *refreshing period* has passed.


* As the *Commander*, I *Should* __know__ all __exposed ports__ in a *Host*
  * Acceptance criteria:
    * A *Host* has no exposed ports. No ports appears on the *Container* description.
    * A *Host* has a *Container* running with one exposed port. The mapping appears in the dashboard.
    * A *Host* has a *Container* with 2 exposed ports. Both mapping appears in *Container* description.

### Containers Requirements
* As the *Commander*, I *Must* __know__ the full __status__ and details __of__ a __container__
  * Acceptance criteria:
    * A *More info* query is done against a *Container*. The result must show as much info as possible TODO.


* As the *Commander*, I *Must* __know if__ a selected __*Image* is__ actually __running__
  * Acceptance criteria:
    * An *Image* installed in Docker is "selected". The *Image* has an associated *Container* running. The dashboard must show the associated *Container*.
    * An *Image* installed in Docker is "selected". The *Image* does not have an associated *Container* running. The dashboard must show no associated *Containers*.


* As the *Commander*, I *Should* __know__ the __exposed ports of__ a __*Container*__ and their associated mapping ports in *Host* so I know that available ports in *Host*
  * Acceptance criteria:
    * A *Container* that is running has no exposed *ports*. The result in dashboard shows no ports.
    * A *Container* that is running has one exposed *port*. The result in dashboard shows one mapped port and their associated ports in *Host*.
    * A *Container* that is running has 3 exposed *ports*. The result in dashboard shows 3 mapped ports and their associated ports in *Host*.

### Commander Requirements
TODO


# Architecture

## Gin Gonic Server

### etcd piece

Maintains state of the app and connected Swarm agents in cluster

* Watch 2 events:
  * New Agent in cluster: Emits `new-agent` event in server
  * Agent down in cluster: Emits `kill-agent` event in server

### REST Client

Will be used by server to make REST connections to Swarm Agents and Manager and return JSON info "as is" back to front via web-socket.

### Socket API

An API via socket layer to communicate real-time with front. Because it's bi-directional it is described as follows:


#### Client->Server
  * `cluster`: Returns a json with the entire cluster state. Mainly used when a client has just connected to the app
  * `agent:containers`: Replicates `GET /containers/json`. Needs image IP and optional parameters like in [Docker API](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#list-volumes)
  * `agent:images`: Replicates `GET /images/json`

#### Server->Client
  * `new-agent`: When a new Swarm Agent has joined the cluster
  * `kill-agent`: When an already existing Swarm Agent has left the cluster

## Front (React+Redux)
TODO
