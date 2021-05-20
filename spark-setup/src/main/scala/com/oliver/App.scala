package com.oliver

import org.apache.spark.sql.SparkSession

/**
 * @author ${user.name}
 */
object App {
  

  def main(args : Array[String]) {

    val spark = SparkSession.builder()
      .getOrCreate()

    val data = spark.sparkContext.parallelize(Seq("Hello World", "This is some text", "Hello text"))
    val map = data.flatMap(e => e.split(" ")).map(word => (word, 1))
    val counts = map.reduceByKey(_ + _).repartition(1)

    counts.saveAsTextFile("wordcount")
  }

}
