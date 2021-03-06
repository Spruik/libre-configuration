# Libre Configuration Library
> Libre Manufacturing Data Platform 

The Configuration Library provides a common approach for handling external configuration files in microservices
* Sets a common schema for external json configuration files
* handles the reading and management of configuration data.

## What is the Libre Manufacturing Data Platform?

Libre is an Open-source, Data Science ready continuous improvement and Manufacturing Execution Platform.

The Libre platform allows you to:
* connect to your Manufacturing data sources 
* model your equipment, and the data that your equipment provides
* store machine data in a specialised time-series database with full context from your equipment model
* Define rules for capturing event data from your equipment, and store that event data in the time-series database with context
* Define Materials and recipes that describe the way products are manufactured in your plant
* Access all of your data through an open GraphQL API including the equipment and material models along with data from the time-series database

### Libre Platform Architecture:

The Libre Platform is an Event-Driven Micro-services architecture written in Go.

The Frontend UI application communicates primarily through the GraphQL gateway, and 
receives event notifications directly from the event bus.

Apollo Federation is used to compose a single GraphQL schema from multiple graphql services

Non time-series data is stored in Dgraph which is a GraphQL native database platform.

![Libre Component Architecture](./docs/LibreComponentArchitecture.png)
