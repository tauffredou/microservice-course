# microservice-course

We will create an API server used to do full text search, via Elasticsearch, in sport
facilities list in Paris deployed with `docker` and `docker-compose`.

## Step 1

Create a Dockerfile to build a simple server with one endpoint `/search?q=Mouchotte` returning
a JSON representation of one sport facility.

```
{
  "type": "Salle multisports",
  "nature": "Intérieur",
  "city": "Paris",
  "commissioning_year": "1985-1994",
  "name": "Centre Sportif Du Commandant Mouchotte",
  "address": "Rue Du Commandant Mouchotte",
  "nb_facilities": 3,
  "facility_id": 141190,
  "zip_code": 75014
}
```

Useful documentation:
[Dockerfile reference](https://docs.docker.com/engine/reference/builder/)

## Step 2

Use `docker-compose` to run an Elasticsearch cluster locally and upload data from
`facilities.json` file in to Elasticsearch.

```
curl -XPOST -H "Content-Type: application/x-ndjson" "http://localhost:9200/facilities/default?pretty" -d "@facilities.json"
```

Useful documentation:
[Compose reference](https://docs.docker.com/compose/compose-file/)
[Elasticsearch documentation on Docker](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)

## Step 3

Use `main.go` file to launch an implemented version of the server.
Add this server to your `docker-compose` file.
Fix the last line of the HTTP handler to loop over all search results and print them
to the `ResponseWriter`.
