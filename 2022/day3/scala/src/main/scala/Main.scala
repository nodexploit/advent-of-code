import scala.io.Source
import scala.collection.mutable.HashSet
import scala.annotation.tailrec
import scala.util.Try

val FilePath = "input.txt"

@main def main: Unit =
  val pt1Priorities = part1Priorities()

  println(s"Sum of priorities Pt. I is: $pt1Priorities")

  val pt2Priorities = part2Priorities()

  println(s"Sum of priorities Pt. II is: $pt2Priorities")

private def part1Priorities(): Int =
  readLines(FilePath).foldLeft(0)((acc, line) =>
    accPriorities(acc, findDuplicate(line))
  )

private def part2Priorities() =
  readLines(FilePath)
    .grouped(3)
    .foldLeft(0)((acc, group) => accPriorities(acc, findDuplicate(group)))

private def splitLine(line: String): (String, String) =
  val halfLength = line.length() / 2
  line.span(char => line.indexOf(char) < halfLength)

private def accPriorities(acc: Int, maybeDuplicate: Option[Char]) =
  val maybePriority = for {
    duplicate <- maybeDuplicate
  } yield priority(duplicate)

  acc + maybePriority.getOrElse(0)

private def readLines(fileName: String) = Source.fromFile(fileName).getLines()

private def findDuplicate(line: String): Option[Char] =
  val (firstHalf, secondHalf) = splitRecursive(line)

  val set = HashSet.from(firstHalf)

  @tailrec def find(elems: String): Option[Char] =
    elems match {
      case elems if elems.isEmpty() => None
      case _ =>
        val element = elems.head
        if set.contains(element) then Some(element)
        else find(elems.tail)
    }

  find(secondHalf)

private def findDuplicate(lines: Seq[String]): Option[Char] =
  if lines.length != 3 then None
  else
    val secondSet = HashSet.from(lines(1))
    val thirdSet = HashSet.from(lines(2))

    @tailrec def find(elems: String): Option[Char] =
      elems match {
        case elems if elems.isEmpty() => None
        case _ =>
          val element = elems.head
          if secondSet.contains(element) && thirdSet.contains(element) then
            Some(element)
          else find(elems.tail)
      }

    find(lines(0))

private def splitRecursive(line: String): (String, String) =
  val halfLength = line.length() / 2
  @tailrec def split(
      counter: Int,
      firstHalf: String,
      line: String
  ): (String, String) =
    if counter == halfLength then (firstHalf, line)
    else split(counter + 1, firstHalf :+ line.head, line.tail)

  split(0, "", line)

private def priority(char: Char): Int =
  // Ascii value offset for a-z (1-26), A-Z (27-52)
  val offset =
    if char.isUpper then 38
    else 96

  val maybePriority = for {
    asciiValue <- Try(char.toInt).toOption
  } yield asciiValue - offset

  maybePriority.getOrElse(0)
