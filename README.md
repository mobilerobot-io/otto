# oTTo ~ The Pluggable Micro Server ~

OttO is a very small and simple microserver that uses plugins to
provide the applications functionality.  This allows otto to serve up
a wide variety of functionality while allowing the user to only
include the stuff she wants.

Plugins are developed using [Go builtin plugins](http://golang.org/packages/plugin). 

Spinning up servers for PoCs, testing or starting a new project is
a typical thing to do these days [era of the
micro-service](http://wikipedia.org/microservices). 

## Built-In Endpoints

- /routes	~ Dump the routes currently registered
- /plugins	~ Dump the plugins we have registered

The /routes endpoint will dump the routes that we have been registered
by plugins so far...

## Plugins

- echo	~ echo back

- dork  ~ digital ocean cloud manager
- static ~ serve up static files
- store ~ persist info for later
- wally ~ website walker
- wsgi ~ run python Flask scripts


