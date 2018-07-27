# myac

[![Build Status](https://travis-ci.org/ignatev/myac.svg?branch=master)](https://travis-ci.org/ignatev/myac)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ac02cfb27fb04d6a84e0ccbceb232f53)](https://www.codacy.com/project/iskander.ignatev/myac/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ignatev/myac&amp;utm_campaign=Badge_Grade_Dashboard)

Myac is a configuration server inspired by Spring Cloud Config.

* Serve git repo with configs
* Generate URL for configs by folder name

Example git repo with configuration to serve:

```bash

│
├── service-1
│   └── generic-service.yml
└── service-2
    └── generic-service.yml

```

For this repo structure Myac will generate two URLs:

* <http://myac-server:8888/service-1>
* <http://myac-server:8888/service-2>

## ToDo

* Return JSON instead of plain text
* Create several URLs for folders with more than one config-file (service-1/dev, service-1/prod, etc.)
* Write tests
