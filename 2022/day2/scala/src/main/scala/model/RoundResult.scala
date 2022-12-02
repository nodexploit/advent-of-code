package model

enum RoundResult(val points: Int) {
  case Defeat extends RoundResult(0)
  case Draw extends RoundResult(3)
  case Win extends RoundResult(6)
}

object RoundResult {
  def apply(line: String): Option[RoundResult] =
    line match {
      case "X" => Some(RoundResult.Defeat)
      case "Y" => Some(RoundResult.Draw)
      case "Z" => Some(RoundResult.Win)
      case _   => None
    }
}
