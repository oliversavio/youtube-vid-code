package com.oliver

import org.apache.spark.sql.SparkSession
import org.apache.spark.sql.types.StructType
import org.apache.spark.sql.functions._
import org.apache.spark.sql.streaming.Trigger
import java.util.Properties
import org.apache.spark.sql.SaveMode

object SalStream {
  def main(args: Array[String]) {

    val spark = SparkSession
      .builder()
      .appName("SalStream")
      .getOrCreate()

    val schema = new StructType()
      .add("Id", "string")
      .add("EmployeeName", "string")
      .add("JobTitle", "string")
      .add("BasePay", "double")
      .add("OvertimePay", "double")
      .add("OtherPay", "double")
      .add("Benefits", "string")
      .add("TotalPay", "double")
      .add("TotalPayBenefits", "double")
      .add("Year", "int")
      .add("Notes", "string")
      .add("Agency", "string")
      .add("Status", "string")

    val salaries = spark.read
      .option("header", "true")
      .schema(schema)
      .csv("salary/")

    val meanTotalSalary = salaries
      .agg(count("*") as "count", avg("TotalPay") as "total_pay")

     meanTotalSalary.show()

  }
}

