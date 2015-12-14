package config

// APP_PORT represents web server port
const APP_PORT string = ":8000"

// SWARM_MANAGER host names
const SWARM_MANAGER string = "http://192.168.1.35:8081"

// Discovery services

// ETCD host
const ETCD_HOST string = "http://127.0.0.1:2379"

// Consul host
const CONSUL_HOST string = "http://127.0.0.1:2375"

//Zookeeper host
const ZOOKEEPER_HOST string = "http://127.0.0.1:2375"

const ETCD = 0
const CONSUL = 1
const ZOOKEEPER = 2
const DISCOVERY_SERVICE = ETCD

//Environments
const DEVELOPMENT = 0
const PRODUCTION = 1

const CURRENT_ENV = DEVELOPMENT
