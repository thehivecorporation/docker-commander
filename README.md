docker-commander
=====
A GUI to manage Docker

## Some info
* Project structure is defined like here: https://github.com/golang/go/wiki/GithubCodeLayout
* We're using gin-gonic: https://github.com/gin-gonic/gin

## Glossary
* A ***Host*** is a machine, virtual or not, accessible from the server that is going to host Docker containers.
* The ***Server*** is the machine hosting Docker Commander
* A ***Container*** is a Docker container that is being host in a *Host*.

## MoSCoW Requirements

### Host Requirements
* As the User, I *Must* know if a *Host* is down or not so I'll know as fast as possible that one or more containers could be down
  * *Acceptance criteria:*
    * A registered host is turned off. It appears like *down* in the dashboard
    * A registered host is on. It appears like *online* in the dashboard
* As the User, I *Must* know how many containers a *Host* is hosting so I have some insight of the host load
  * Acceptance criteria:
    * TODO
* As the User, I *Should* know the latency I'm getting against a *Host* so I can detect potential connectivity problems
  * Acceptance criteria:
    * TODO
* As the User, I *Must* know the names of the *Containers* that a *Host* has so I know all *Containers* currently installed
  * Acceptance criteria:
    * TODO
* As the User, I *Must* know the names of the running *Containers* so I have full knowledge of the *Host* *Containers*
  * Acceptance criteria:
    * TODO
* As the User, i *Should* know all exposed ports in a *Host*
  * Acceptance criteria:
    * TODO

### Containers Requirements
* As the User, I *Must* know the full status and details of a container
  * Acceptance criteria:
    * TODO
* As the User, I *Must* know if a selected *Container* is actually running
  * Acceptance criteria:
    * TODO
* As the User, I *Should* know the exposed ports of a *Container*
  * Acceptance criteria:
    * TODO
