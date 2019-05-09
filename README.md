# oTTo ~ The Pluggable Micro Server ~

OttO is a very small and simple microserver that facilitates machine
to machine (and human) communication, it is simple and scalable, uses
plugins to provide functionality.

This allows otto to serve up a wide variety of functionality allowing
the developer to focus on the application, not worry too much about
the corresponding infrastructure.

Plugins are developed using [Go builtin plugins](http://golang.org/packages/plugin). 

Spinning up servers for PoCs, testing or starting a new project is
a typical thing to do these days [era of the
micro-service](http://wikipedia.org/microservices). 

## Supported Protocols

- Static Websites
- Single Page Apps
- REST Server
- Websockets
- MQTT Client
- USB/Serial for micro-controllers

## Built-In Endpoints

- /routes	~ Dump the routes currently registered
- /plugins	~ Dump the plugins we have registered
- /mqtt		~ MQTT information and stats
- /serial   ~ serial port information

The /routes endpoint will dump the routes that we have been registered
by plugins so far...

## Plugins

- echo	~ echo back

- dork  ~ digital ocean cloud manager
- static ~ serve up static files
- store ~ persist info for later
- wally ~ website walker
- wsgi ~ run python Flask scripts


