# Build & Package Scala Spark Applications with Maven and IntelliJ


### Requirements
 - JDK 8
 - Linux / MacOs / WSL (Windows Subsystem for Linux)
 - Spark installed ()

## Step 1
### Generate a Scala project from the maven archetype
`mvn archetype:generate -DarchetypeGroupId=net.alchim31.maven -DarchetypeArtifactId=scala-archetype-simple -DdownloadSources=true`

## Step 2
### Change to a newer Scala version in the pom.xml
`<scala.version>2.12.13</scala.version>`


## Step 3
### Add the latest Spark SQL dependency

```
<dependency>
    <groupId>org.apache.spark</groupId>
    <artifactId>spark-sql_2.12</artifactId>
    <version>3.1.1</version>
</dependency>
```

## Step 4
### Compile and package project with maven
`mvn clean package`

## Step 5
### Execute with spark-submit
`spark-submit --master "local[*]" --class com.oliver.App target/spark-setup-1.0-SNAPSHOT.jar`


