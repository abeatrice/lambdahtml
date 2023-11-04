# GOLANG Lambda Function URL HTML Server

This is an example application written in go that runs a server locally and returns html when invoked on aws lambda with a function url

## Requirements
 - [docker](https://docs.docker.com/)
 - [docker-compose](https://docs.docker.com/compose/)

## Local Development
 - CompileDaemon will recompile and run the app server when .go files are updated.
 - docker-compose will create services: database, app server
```sh
$ make run
```
 - Invoke aws test events in the test_events directory
```sh
$ make invoke EVENT="test_events/get-root.json"
```
