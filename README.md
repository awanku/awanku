Awanku Stack
============

![core-api](https://github.com/awanku/awanku/workflows/core-api/badge.svg?branch=master) ![console-webui](https://github.com/awanku/awanku/workflows/console-webui/badge.svg?branch=master) ![landing-webui](https://github.com/awanku/awanku/workflows/landing-webui/badge.svg?branch=master)

## Setting up development environment

1. Install Docker https://docs.docker.com/get-docker/

1. Clone this repo, then run:

    ```
    make run-build
    ```

1. Nex you need to run database migration, in another tab run:


    ```
    make db-migrate
    ```

    **TIP**: you can use `make db-nuke` then `make db-migrate` to revert your database into clean state.


1. Your development environment will be ready at:

    ```
    Landing page: http://awanku.xyz

    Console page: http://console.awanku.xyz

    API: http://api.awanku.xyz
    ```

## Stack

1. Landing page, code is in [landing](landing) folder

2. Console page, code is in [console](console) folder

3. Backend code, in [backend](backend) folder

All stacks have auto reload enabled, any changes will cause the app to reload, including installing new dependencies (if any).
