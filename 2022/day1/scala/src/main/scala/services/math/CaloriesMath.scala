package services.math

import model.ElvesCalories

object CaloriesMath {

  def maxCalories(elvesCalories: ElvesCalories) = elvesCalories.max

  def topThree(elvesCalories: ElvesCalories) =
    elvesCalories.sortBy(_._2)(Ordering[Int].reverse).slice(0, 3)
}
