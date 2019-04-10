# oTTo ~ The Pluggable Micro Server ~

OttO is a very small and simple microserver that uses plugins to
provide the applications functionality.  For example a couple plugins
include an echo server, a VM lister for digital ocean and a website
walker, just to name a couple.

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
- wally ~ website walker
- dork  ~ digital ocean cloud manager
