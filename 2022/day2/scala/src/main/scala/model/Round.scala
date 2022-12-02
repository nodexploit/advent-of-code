package model

import RockPaperScissors.Rock
import model.RockPaperScissors

trait Parseable[A]:
  extension (line: Array[String]) def parse: Option[A]

trait Decidable[A]:
  extension (a: A) def decide: Int

sealed trait Round

case class RoundPtI(opponent: RockPaperScissors, selection: RockPaperScissors)
    extends Round

object RoundPtI {

  given Parseable[RoundPtI] with
    extension (line: Array[String])
      def parse =
        if line.isEmpty then None
        else
          for {
            opponentSelection <- RockPaperScissors(line(0))
            selection <- RockPaperScissors(line(1))
          } yield RoundPtI(opponentSelection, selection)

  given Decidable[RoundPtI] with
    extension (r: RoundPtI)
      def decide: Int =
        val RoundPtI(opponent, selection) = r
        val roundOutcome =
          if opponent.defeats(selection) then RoundResult.Defeat
          else if selection.defeats(opponent) then RoundResult.Win
          else RoundResult.Draw

        selection.points + roundOutcome.points
}

case class RoundPtII(opponent: RockPaperScissors, roundResult: RoundResult)
    extends Round

object RoundPtII {

  given Parseable[RoundPtII] with
    extension (line: Array[String])
      def parse =
        if line.isEmpty then None
        else
          for {
            opponentSelection <- RockPaperScissors(line(0))
            result <- RoundResult(line(1))
          } yield RoundPtII(opponentSelection, result)

  given Decidable[RoundPtII] with
    extension (r: RoundPtII)
      def decide: Int =
        val RoundPtII(opponent, result) = r
        val selection = result match {
          case RoundResult.Draw   => opponent
          case RoundResult.Defeat => opponent.defeats
          case RoundResult.Win =>
            Set(
              RockPaperScissors.Rock,
              RockPaperScissors.Paper,
              RockPaperScissors.Scissors
            ).filter(selection =>
              selection != opponent && selection != opponent.defeats
            ).head
        }

        selection.points + result.points
}
