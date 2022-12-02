package services.parsing

import model.ElvesCalories

import scala.io.Source
import scala.util.Try

object CaloriesParser {

  def parse(fileName: String): ElvesCalories =
    inputIterator(fileName).foldLeft(ElvesCalories.empty)(parseLine)

  private def inputIterator(fileName: String) =
    Source.fromFile(fileName).getLines()

  private def parseLine(elvesCalories: ElvesCalories, caloriesInput: String) =
    if caloriesInput.trim().isEmpty() then
      elvesCalories :+ (elvesCalories.length, 0)
    else
      val currentCalories = Try(caloriesInput.toInt).getOrElse(0)
      elvesCalories match {
        case Vector() => ElvesCalories.empty :+ (0, currentCalories)
        case xs =>
          val (elve, calories) = xs.last
          xs.init :+ (elve, calories + currentCalories)
      }
}
