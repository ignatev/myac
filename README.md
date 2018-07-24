# myac
[![Build Status](https://travis-ci.org/ignatev/myac.svg?branch=master)](https://travis-ci.org/ignatev/myac)

Myac is a configuration server inspired by Spring Cloud Config.
  - Serve git repo with configs
  - Generate URL for configs by folder name

Example git repo with configuration to serve:

```
│
├── service-1
│   └── generic-service.yml
└── service-2
    └── generic-service.yml

```
For this repo structure Myac will generate two URLs:
  - http://myac-server:8888/service-1
  - http://myac-server:8888/service-2

#### ToDo
  - Return JSON instead of plain text
  - Create several URLs for folders with more than one config-file (service-1/dev, sevice-1/prod, etc.)
  - Write tests
