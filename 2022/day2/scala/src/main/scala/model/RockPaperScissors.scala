package model

import model.RoundResult

sealed trait RockPaperScissors(
    val defeats: RockPaperScissors,
    val points: Int
)

object RockPaperScissors {
  case object Rock extends RockPaperScissors(Scissors, 1)
  case object Paper extends RockPaperScissors(Rock, 2)
  case object Scissors extends RockPaperScissors(Paper, 3)

  extension (opponent: RockPaperScissors)
    def defeats(selection: RockPaperScissors) = opponent.defeats == selection

  def apply(input: String): Option[RockPaperScissors] =
    input match {
      case "A" | "X" => Some(Rock)
      case "B" | "Y" => Some(Paper)
      case "C" | "Z" => Some(Scissors)
      case _         => None
    }
}
