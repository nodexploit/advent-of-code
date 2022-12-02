import services.Parser
import model.RoundPtI
import model.RoundPtII

val FilePath = "input.txt"

@main def hello: Unit =
  val gameScore = Parser.parseRounds[RoundPtI](FilePath)

  println(s"There was a total score of: ${gameScore}")

  val secondGameScore = Parser.parseRounds[RoundPtII](FilePath)

  println(s"Changing the rules, there was a total score of: ${secondGameScore}")
