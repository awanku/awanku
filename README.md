Awanku Stack
============

![core-api](https://github.com/awanku/awanku/workflows/core-api/badge.svg?branch=master) ![console-webui](https://github.com/awanku/awanku/workflows/console-webui/badge.svg?branch=master) ![landing-webui](https://github.com/awanku/awanku/workflows/landing-webui/badge.svg?branch=master)

## Setting your machine for development

1. Install Docker https://docs.docker.com/get-docker/

1. OpenVPN Client

## Stack

1. Landing page, code is in [landing](landing) folder

2. Console page, code is in [console](console) folder

3. Backend code, in [backend](backend) folder

All stacks have auto reload enabled, any changes will cause the app to reload, including installing new dependencies (if any).

## Environments

### Production

Domain: `awanku.id`

### Staging

Domain: `staging.awanku.xyz`

This is where frontend development is done. Your frontend code will run in your local machine but backend code will run on staging servers. **You need VPN to connect to the environment**.

Running the environment: `make staging-up`

### Development

Domain: `dev.awanku.xyz`

All domains in this environment point to `127.0.0.1`.
