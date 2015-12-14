package config

// APP_PORT represents web server port
const APP_PORT string = ":8000"

// SWARM_MANAGER host names
const SWARM_MANAGER string = "http://192.168.1.35:8081"

// ETCD host
const ETCD string = "http://192.168.1.35:2379"

//Environments
const DEVELOPMENT = 0
const PRODUCTION = 1

const CURRENT_ENV = DEVELOPMENT
