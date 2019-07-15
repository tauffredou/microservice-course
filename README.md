# microservice-course

We will create an API server used to do full text search, via Elasticsearch, in sport
facilities list in Paris deployed with `docker` and `docker-compose`.

## Step 1

Create a Dockerfile to deploy the Golang application in `main.go`. It is recommended to use the official `golang:1.12` image.

When you start the application with the following command:
```
docker run -it --rm <your image name>
```

The application should panic with the following error:
```
panic: health check timeout: Head http://127.0.0.1:9200: dial tcp 127.0.0.1:9200: connect: connection refused: no Elasticsearch node available

goroutine 1 [running]:
main.main()
	/usr/src/main.go:76 +0xdc2
```

At this point, try to identify why the application is returning an error and what should we do avoid it.

Useful documentation:
 - [Dockerfile reference](https://docs.docker.com/engine/reference/builder/)

## Step 2

Use `docker-compose` to run an Elasticsearch cluster besides the application.

**NOTE**: the Elasticsearch cluster can't start without the env var `dicovery.type=single-node`.

At this point, the application should start and load data into the Elasticsearch cluster if necessary. You can reset the Elasticsearch index with the following command:
```
curl -XDELETE http://<hostname of the elasticsearch container>:9200/_all
```

After starting (it could take multiple seconds the first time), the service should be available and you can test it with the following command:
```
curl http://<hostname of the application container>:<port exposed>/search?q=Paris
```

Useful documentation:
 - [Compose reference](https://docs.docker.com/compose/compose-file/)
 - [Elasticsearch documentation on Docker](https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html)

Petclinic features
https://docs.google.com/document/d/1V28rAQzsKc9_5XF5kljyKiOLF0OjbbeBwZY3qXyIwNw
