import model.ElvesCalories
import services.parsing.CaloriesParser
import services.math.CaloriesMath.{topThree, maxCalories}

import scala.io.Source
import scala.util.Try

val FilePath = "input.txt"

@main def hello: Unit =
  val elvesCalories = CaloriesParser.parse(FilePath)

  println(s"Max calories carried by the Elves: ${maxCalories(elvesCalories)}")

  val top = topThree(elvesCalories)
  println(s"Top three elves carrying calories: ${top}")

  println(s"Total calories among top three elves: ${top.map(_._2).sum}")
