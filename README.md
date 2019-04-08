# microservice-course

We will create an API server used to do full text search, via Elasticsearch, in sport
facilities list in Paris deployed with `docker` and `docker-compose`.

## Step 1

Create a Dockerfile to build a simple server with one endpoint `/search` returning
a JSON representation of one sport facility.

Useful documentation:
[Dockerfile reference](https://docs.docker.com/engine/reference/builder/)

## Step 2

Use `docker-compose` to run an Elasticsearch cluster locally and upload data from
`facilities.json` file in to Elasticsearch.

Useful documentation:
[Compose reference](https://docs.docker.com/compose/compose-file/)
[Elasticsearch documentation on Docker](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)

## Step 3

Add your server in your `docker-compose` file.
