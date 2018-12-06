# otto ~ The Problem Solver ~ 

## TODO: Define Problem Otto Solves

- [Problem We Are Solving ](doc/problem-we-are-solving.md)

## TODO: What is OttO

## Workflow 

> status can be used at any time to see where we are at.

Modern _DevOps_ workflow.  This is the proposed workflow.  This
_workflow_ is really a _cycle_, a _chaotic cycle_, in other words,
reality dictates that my tidy definition below will not always be
orderly or complete.

Truth is, we will jump from milestone to milestone leaving a task here
and there incomplete, for one reason or another.  Is this necessary
and unavoidable?  Maybe, maybe not.  

> Always expect the unexpected (never assume something else "should
> not have happend").

### Launch ~ Provision and Configuration 

1. M0 - Plan			~ Gather Requirement and Create Plan (GrandPlan)
1. M1 - Provision		~ Acquire Resources and Ready with Creds
1. M2 - Configure		~ Bring all resources to required state
3. M4 - Log				~ Organized, intellegent history of app
3. M5 - Tickets			~ Bugs and Enhancements: prioritize and move forward
3. M3 - Integrate		~ Roll changes from dev into production
2. M6 - Optimization	~ Customers, History and Data
5. M7 - HAL 2010		~ A decade late, but we are Here!

### Workflow Requirements: Expectations and Outcomes

Using a traditional unix style build process and _Makefile_, our
application lifecycle can be created, updated, observed and destroyed
with the following commands (which is also an API we'll try to stick
with.) 

The following example will, set the blueprint, inventory and
providers.

```
% make providers = do, vagrant, gcp, azure
% make inventory = /srv/inventory/{{provider}} 
% make plans = lb-hap-ngx
%
% make images		=> produces proviimages
% make provision
% make configure
% make status
% make update
% make destroy
```

## Table of Contenets

This repo contains the following files and directories.

- README.md	~ you are reading this file (tell you what's what)
- Makefile	~ runs all build, config, and infra updates
- Vagrantf..~ provision vagrant / virtual box test env 
- config	~ ansible = configuration management
- doc		~ documentation (site for hugo) 
- etc		~ stuff that i have not tended to
- img		~ packer = duplicate provider images
- prov		~ terriform = provision 
- src		~ sources for otto the helper


### M0 Images 

Packer will be used to create standard or _Golden Images_ for things
like nginx servers, elk stack, databases, etc.

Packer will ensure our providers: DO, GCP, Vagrant, Docker, CloudStack
all have the same version of nginx server (with all same software
configruation, etc.)

> This saves a tremendous amount of configuration time!

This is a short list of Images that we will benefit from having
pre-created. 

- nginx server	~ web server 
- haproxy		~ load balancing
- otto			~ build server and monitor
- clowd			~ box of clowds


### 1. Provision Infrastructure ~ Terraform

Terraform will a given application (network infra), determine the
resources that do exist, and do not exist.  Terraform will 

= create the resources that do not exit
- assasinate tainted resources and build new resource (server) from
  scratch. 

- terraform is finished when all resources
  - have been created
  - up and running

Terraform will create the desired clowd infra structure. It will also
check inventory and correct problems when they are detected.

- multi clowd
- create configured infra structure
- ensure configured infra structure is healthy, complete and correct
- change management, all additions and deletions of infra will be
  handled by the terraform.
  
> Terraform brings a site to life from a configuration file(s).  It
> ensures that site is correct when run.
  
### 2. Config Infrastructure &  ~ Ansible 

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

