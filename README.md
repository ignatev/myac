# myac

[![Build Status](https://travis-ci.org/ignatev/myac.svg?branch=master)](https://travis-ci.org/ignatev/myac)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ac02cfb27fb04d6a84e0ccbceb232f53)](https://www.codacy.com/project/iskander.ignatev/myac/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ignatev/myac&amp;utm_campaign=Badge_Grade_Dashboard)
[![Go Report Card](https://goreportcard.com/badge/github.com/ignatev/myac)](https://goreportcard.com/report/github.com/ignatev/myac)

Myac is a configuration server inspired by Spring Cloud Config.

* Serve git repo with configs
* Generate URL for configs by folder name

Example git repo with configuration to serve:

```bash
2018/09/11 00:04:33 repository already exists

 ._ _        _.   _
 | | |  \/  (_|  (_
        /
Configuration server

Service running on port 8888
.filesystem-repo
├── .git
├── LICENSE >>> http://localhost:8888/.filesystem-repo
├── service-1
│   ├── generic-service.yml >>> http://localhost:8888/.filesystem-repo/service-1
│   └── service-1-2
│       ├── servcie-1-2-conf >>> http://localhost:8888/.filesystem-repo/service-1/service-1-2
│       └── service-1-3
│           └── service-1-3-conf >>> http://localhost:8888/.filesystem-repo/service-1/service-1-2/service-1-3
└── service-2
    ├── generic-service-dev.yml >>> http://localhost:8888/.filesystem-repo/service-2/generic-service-dev
    └── generic-service.yml >>> http://localhost:8888/.filesystem-repo/service-2/generic-service

```

For this repo structure Myac will generate two URLs:

* <http://myac-server:8888/service-1>
* <http://myac-server:8888/service-2>

## ToDo

* Return JSON instead of plain text
* Create several URLs for folders with more than one config-file (service-1/dev, service-1/prod, etc.)
* Write tests
