# <p align="center">Shopping Cart Engine <img src="https://github.githubassets.com/images/icons/emoji/unicode/1f680.png?v8" width=40px height=40px> </p>

## Description
This web service fundamentally uses the hexagonal architecture pattern, which is a framework that is part of my personal portfolio in developing a web service and its various supporting components. This architecture has unique dependencies in its usage, such as the adapter-port method which directly abstracts the working basis of this structure.

## Installation

If want running locally, install first dependencies:  
```shell
make clean && make install
```

run app using Makefile:
```shell
make start
```
or run migration file using : 
```shell
go run main.go db:migrate up
```

**notes: <br />Please create file on (conf/app.yaml) copy from app.example.yaml & create database first on local

Run using docker: 
```shell
docker compose -f deploy/docker-compose.yaml up -d
```
### Dependencies
| Name | Version | Status |
|----------|----------|----------|
| Go | 1.21.1 | &check; |
| Postgres | 1.16.3 | &check; |
| Docker | 26.1.3 | &check; |

### Docs
- [Image registry](https://hub.docker.com/r/philipsjp26/shopping-cart)
- [ERD Database diagram ](https://github.com/philipsjp26/Shopping-Cart/blob/master/assets/diagrams/ERD.png?raw=true)
- [Architectural Diagram](https://github.com/philipsjp26/Shopping-Cart/blob/master/assets/diagrams/diagram.png?raw=true)
- [API Docs ( Postman )](https://github.com/philipsjp26/Shopping-Cart/blob/master/assets/diagrams/CollectionShoppingCart.json)
- [Quick Start API Contract](https://github.com/philipsjp26/Shopping-Cart/blob/master/assets/README.md)

