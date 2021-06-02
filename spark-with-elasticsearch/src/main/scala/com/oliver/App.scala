package com.oliver

import org.apache.spark.sql.SparkSession
import org.elasticsearch.spark.sql._

/** @author ${user.name}
  */
object App {

  def main(args: Array[String]) {

    val spark = SparkSession.builder().getOrCreate()

    import spark.implicits._

    val indexDocuments = Seq(
      AlbumIndex("Led Zeppelin", 1969, "Led Zeppelin"),
      AlbumIndex("Boston", 1976, "Boston"),
      AlbumIndex("Fleetwood Mac", 1979, "Tusk")
    ).toDF


    indexDocuments.saveToEs("demoindex/music")

  }

}

case class AlbumIndex(artist: String, yearOfRelease: Int, albumName: String)
