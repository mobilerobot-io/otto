# otto

OttO handles a number of details regarding getting the clowdops
monitoring site up and running.

## Contents 

The contents of this directory are as follows:

- README.md: you are reading this file (tell you what's what)
- Makefile ~ start infra, check status and destroy
- Vagrantfile ~ start a local Virtualbox net with otto and nginx
- infra ~ provision & config the clowdops hosting site
- sites ~ sites that will be hosted by this hosting entity
- src ~ source code for otto

## Status ~ Checking Things Out

This repository is a loaded gun, it is very effective at bringing up
"sites", put in the wrong (ignorant) hands, it can easily do damage to
a production deployment, running test deployments and leaving things
running, cranking up un-necessary bills.

This software is very powerful, use it carefully and you will do
things (in a moment) that are amazing!

> make status to see what is going on ...

Make status will let you know what has and has not been provisioned,
the health of this site or app.

### Provision, Configure, Test, Deploy

```bash
$ make prov		# provision sites (terraform or vagrant)
$ make config	# run configuration management (ansible)
$ make status   # quick health check (tf & ans, vag)
$ make destroy  # stop and terminate all resources
```

## Inventory, Domains and Sites

Basic CRUD operations are provided for each of the _stores_ Inventory,
Domains and Sites.

```
- get		/inv
- get		/inv/{item}
- post		/inv/{item}
- post		/inv/{item} body=json
- delete	/inv/{item}
```

Additional calls available:

### DNS Modifications

```
- get /dom/ns/{domain}  => get nameservers for domain
- post /dom/ns/{domain}?ns='ns1,ns2'
- delete /dom/ns/{domain}

- get /dom/dns/{domain} => return host records
- set /dom/dns/{domain}?rec=foo
- delete /dom/dns/{domain}
```

## Walker
```
- get	/site/walk/					- get list of recent site walks
- get	/site/walk/{site}/			- return a list of walkids and the last walk
- get	/site/walk/{site}/{walkid}	- return the walk the the specific id
- put	/site/walk/{site}?params="."  - schedule walk params
- delete /site/walk/{site}			- git rid of the walkers
```  

