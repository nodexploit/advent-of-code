import scala.io.Source
import scala.util.Try

val FilePath = "input.txt"

@main def hello: Unit =
  val nFullyContains = readLines(FilePath).count((line) => fullyContains(line))

  println(s"Number of fully contains: $nFullyContains")

  val nOverlaps = readLines(FilePath).count((line) => overlaps(line))

  println(s"Number of fully contains: $nOverlaps")

final case class Range(start: Int, end: Int)

object Range {
  def apply(range: String): Option[Range] =
    val numbers = range.split('-')

    if numbers.length != 2 then None
    else
      for {
        start <- Try(numbers(0).toInt).toOption
        end <- Try(numbers(1).toInt).toOption
      } yield Range(start, end)

  extension (range: Range)
    def contains(other: Range): Boolean =
      range.start <= other.start && range.end >= other.end

  extension (range: Range)
    def overlaps(other: Range): Boolean =
      range.start <= other.end && range.end >= other.start
}

final case class Assignment(firstElf: Range, secondElf: Range)

object Assignment {
  def apply(assignment: String): Option[Assignment] =
    val ranges = assignment.split(',')

    if ranges.length != 2 then None
    else
      for {
        firstElf <- Range(ranges(0))
        secondElf <- Range(ranges(1))
      } yield Assignment(firstElf, secondElf)
}

def readLines(fileName: String) = Source.fromFile(fileName).getLines()

def fullyContains(line: String): Boolean =
  (for {
    Assignment(firstElf, secondElf) <- Assignment(line)
  } yield firstElf.contains(secondElf) || secondElf.contains(firstElf))
    .getOrElse(false)

def overlaps(line: String): Boolean =
  (for {
    Assignment(firstElf, secondElf) <- Assignment(line)
  } yield firstElf.overlaps(secondElf) || secondElf.overlaps(firstElf))
    .getOrElse(false)
