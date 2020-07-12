# Sample users service #

## Quick Start ##

1. Copy `.env.dist` file to `.env`
    ```bash
    $ cp .env.dist .env
    ```
    Configure the `.env` file as you need or leave it by default
2. Build the project
    ```bash
    $ docker-compose up -d
    $ docker-compose exec app ./users migrate
    $ docker-compose exec app ./users import /data/users_1.jsonl /data/users_2.jsonl
    ```
3. Open [ui](http://127.0.0.1:15400/)
