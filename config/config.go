package config

// APP_PORT represents web server port
const APP_PORT string = ":8000"

// SWARM_MANAGER host names
var SWARM_MANAGER_DEVELOPMENT string = "http://192.168.1.35:8081"
var SWARM_MANAGER_PRODUCTION string = "http://192.168.1.35:8081"

// Discovery services

const ETCD = 0
const CONSUL = 1
const ZOOKEEPER = 2
const DISCOVERY_SERVICE = ETCD

// ETCD host
const ETCD_HOST_DEVELOPMENT string = "http://192.168.1.35:2379"
const ETCD_HOST_PRODUCTION string = "http://127.0.0.1:2379"

// Consul host
const CONSUL_HOST_DEVELOPMENT string = "http://127.0.0.1:2375"
const CONSUL_HOST_PRODUCTION string = "http://127.0.0.1:2375"

//Zookeeper host
const ZOOKEEPER_HOST_DEVELOPMENT string = "http://127.0.0.1:2375"
const ZOOKEEPER_HOST_PRODUCTION string = "http://127.0.0.1:2375"

//Environments
const MOCK = 0
const DEVELOPMENT = 1
const PRODUCTION = 2

const CURRENT_ENV = MOCK

// Front -> Back -> Front communication strategy
const WEBSOCKET = 0
const AJAX = 1
const COMMUNICATIONS_STRATEGY = WEBSOCKET

// Front -> Back actions
const CONNECTION_ACTION_CLUSTER = "cluster"
