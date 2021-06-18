# Writing a Spark DataFrame to an Elasticsearch Index


## Installing Elasticsearch with Docker

### Documentation
```
https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html
```

### Pull in the Docker image
```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.12.1

```

### Start a single node cluster
```
docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.12.1

```

## Maven Assembly plugin
```
https://maven.apache.org/plugins/maven-assembly-plugin/
```

## Executing our code with spark-submit
```
spark-submit --class com.oliver.App --master "local[*]" --conf spark.es.nodes=localhost --conf spark.es.port=9200 --conf spark.es.nodes.wan.only=true target/spark-with-elasticsearch-1.0-SNAPSHOT-jar-with-dependencies.jar
```
