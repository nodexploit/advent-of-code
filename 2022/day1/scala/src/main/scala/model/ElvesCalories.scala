package model

type Elve = Int // Elves represented by position in input
type Calories = Int
type ElvesCalories = Vector[(Elve, Calories)]

object ElvesCalories {
  val empty = Vector.empty[(Elve, Calories)]
}
