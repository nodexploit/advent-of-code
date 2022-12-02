package services

import model.Round
import model.Parseable
import model.Decidable
import model.RockPaperScissors

import scala.io.Source

object Parser {

  def parseRounds[Round: Parseable: Decidable](
      fileName: String
  ): Int =
    inputIterator(fileName).foldLeft(0)(reducer)

  private def inputIterator(fileName: String) =
    Source.fromFile(fileName).getLines()

  private def reducer[Round: Parseable: Decidable](
      accumulator: Int,
      line: String
  ) =
    val chars = line.split(' ')
    val maybeRound: Option[Round] = chars.parse

    if maybeRound.isEmpty then accumulator
    else
      val round: Round = maybeRound.get
      accumulator + round.decide
}
