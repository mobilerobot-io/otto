# otto

OttO handles a number of details regarding getting the clowdops
monitoring site up and running.

## Contents 

The contents of this directory are as follows:

- README.md		~ you are reading this file (tell you what's what)
- Makefile		~ runs all build, config, and infra updates
- Vagrantfile	~ provision vagrant / virtual box test env 
- clowd			~ image index, provision, cfgmgmt, monitor, ci/cd
- images		~ images created by packer for clowd
- sites			~ sites being monitored by clowd
- src			~ source for otto

## Workflow 

### 0. Golden Images ~ Packer

Basically creating applicances of our _"Golden Images"_.  Here are a
list of appliances we would use.

> Use Packer to create "Golden images"

- nginx server	~ web server 
- haproxy		~ load balancing
- otto			~ build server and monitor
- clowd			~ box of clowds

> Packer creates identicle images to run on multiple providers. (DO,
> GCP, Vagrant, AWS, etc)

### 1. Provision Clowd ~ Terraform

Terraform will create the desired clowd infra structure. It will also
check inventory and correct problems when they are detected.

- create configured infra structure
- ensure configured infra structure is healthy, complete and correct
- change management, all additions and deletions of infra will be
  handled by the terraform.
  
> Terraform brings a site to life from a configuration file(s).  It
> ensures that site is correct when run.
  
### 2. Config Clowd ~ Ansible 

Configuration management with Ansible. Based on the inventory we give
ansible (or it snarfs up from a program we give it), ansible will
ensure all servers (and groups of servers) are configured and
operating correctly.

People like to point out that Ansible is **idempotent** (more
acurrately, it is, if the modules it uses are *idempotent*, its
modules should strive to be, idempontent)

> Ansible, when run will scan existing configuration state of
> network.  It will correct differences in network configuration and
> observed state.

Changes to configurations of any server, application, etc are handled
by ansible.

### 3. Bootstrap

1. Determine when complete site has been brought up, ensure that all of
our webservers are operational as well as our load balancer
(haproxy). 

2. Ensure www servers have all sites, nginx servers are running and
   accessable. 
   
3. loadbalancer is doing its job

### 4. CI/CD ~ Software Changes Automatically Deployed

Software changes (merges / commits to master) invoke an event that
causes a pipeline to be invoked that will pull the latest image,
validate the changes then deploy to server.

### 5. Monitor & Logging ~ Hunting Bugs with ELK

Install ELK stack and start benefiting from logging.

```bash
$ export CLOWD_PROVIDERS="do, gcp, vagrant"
$ make images
$ make provision
$ make config
$ make test
$ make status 
$ make destroy
```

The "lifecycle" of every _site_ or _application_ comes in stages (or
transitions from one state to another at times).  These are the
"lifecycle stages" that clowd goes through and what it does.

Clowd very much takes the "site" or "application" POV when the
configuration files are created.  That means that infrastructure is
aquired as needed.


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

