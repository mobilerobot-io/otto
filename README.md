# otto

This is OttO, otto is responsible for managing and monitoring a
_network of hosts_ that make up a _site_ or _application_.

Otto manages the following scenarios

1. Inventory of online resources
1. Provision resources for applications with Terraform
2. Configuration Management with Ansible
3. Site / Application monitoring with ELK stack, influxdb and
   prometheus 
4. Dashboard to insights

## Sections

1. Inventory
   1. Physical Inventory
   1. Online Inventory

2. Domains
   1. Domains registration
   2. renewals & expiration monitoring
   3. DNS records

3. Sites 
   1. provision
   2. configuration management
   3. monitoring for health and performance
   4. ci/cd

## WorkFlow ~ Provision the Application

Terraform workflow level of abstraction, not resource level of
abstraction. 

## Infrastructure As Code

1. 100% Automated, every run you get exact same results
2. Version Controlled - test, track, rollback
3. Annotated history of changes
4. Repeatable ~ spin up old versions, new versions
5. Replicatable ~ as many copies of your infrastructure as you can
   afford
6. Scalable and elastic ~ grows exactly how you do

## Parallelize and Democratize Development

This model allows different people and organizations to take ownership
of various parts of a bigger system.  Loosley connected interfaces
provide for a tremendous amount of flexibility amoung different parts
of an organization, or across organizations even.

### Larger Application into Microservices

Each microservice can be responsible for provisioning its own
infrastructure and customizing accordingly, yet still stick to project
patterns and conventions for transferability and reusability.
