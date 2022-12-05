import scala.io.Source
import scala.collection.mutable.Stack
import scala.annotation.tailrec
import scala.concurrent.duration.span

val FilePath = "input.txt"

@main def main: Unit =
  part1()
  part2()

def part1(): Unit =
  val (stacks, instructions) = parseInput(FilePath)
  val movedStacks = execute(stacks, instructions)
  val crates = topCrates(movedStacks)

  println(s"Final top creates are: ${crates}")

def part2(): Unit =
  val (stacks, instructions) = parseInput(FilePath)
  val movedStacks = executeNewCrane(stacks, instructions)
  val crates = topCrates(movedStacks)

  println(s"Final top creates are: ${crates}")

// Each stack takes 3 character separated by space
def groupedInput(line: String): Vector[String] =
  @tailrec def aux(
      counter: Int,
      line: String,
      groups: Vector[String]
  ): Vector[String] =
    if line.isEmpty() then groups
    else
      val groupSize = if counter % 2 == 0 then 3 else 1
      val (group, tail) = (line.take(groupSize), line.drop(groupSize))

      if counter % 2 == 0 then
        aux(counter + 1, tail.mkString, groups :+ group.mkString)
      else aux(counter + 1, tail.mkString, groups)

  aux(0, line, Vector.empty[String])

final case class Instruction(amount: Int, from: Int, to: Int)

type Stacks = Vector[Stack[Char]]

object Stacks {
  def empty: Stacks = Vector.empty[Stack[Char]]
  def init(lineLength: Int): Stacks = Vector.fill(lineLength)(Stack())
}

def readLines(fileName: String) = Source.fromFile(fileName).getLines

def parseInput(fileName: String): (Stacks, Vector[Instruction]) =
  readLines(fileName).foldLeft((Stacks.empty, Vector.empty[Instruction]))(
    (acc, line) =>
      val (stacks, instructions) = acc

      line match {
        case line if line.contains("[") =>
          (parseStacks(line, stacks), instructions)
        case line if line.startsWith("move") =>
          (stacks, parseInstructions(line, instructions))
        case _ => acc
      }
  )

def parseStacks(line: String, stacks: Stacks): Stacks =
  val letters = parseLine(line)
  if stacks.isEmpty then add(letters, Stacks.init(letters.length))
  else add(letters, stacks)

def parseLine(line: String): Vector[Char] =
  def sanitise(elem: String): Char = elem(1)

  groupedInput(line).map(sanitise)

def add(line: Vector[Char], stacks: Stacks): Stacks =
  for ((letter, index) <- line.zipWithIndex)
    if letter != ' ' then stacks(index).push(letter)

  stacks

def parseInstructions(
    line: String,
    instructions: Vector[Instruction]
): Vector[Instruction] =
  val pattern = """^move (\d+) from (\d+) to (\d+)$""".r
  val matches = pattern.findAllIn(line).matchData

  if matches.hasNext then
    val next = matches.next()
    instructions :+ Instruction(
      next.group(1).toInt,
      next.group(2).toInt,
      next.group(3).toInt
    )
  else instructions

def execute(
    stacks: Stacks,
    instructions: Vector[Instruction],
    keepOrder: Boolean = true
): Stacks =
  instructions.foldLeft(stacks.map(_.reverse))((stacks, instructions) =>
    val source = stacks(instructions.from - 1)
    val destination = stacks(instructions.to - 1)
    val crates = (1 to instructions.amount)
      .map(_ => if source.nonEmpty then Some(source.pop()) else None)
      .flatten

    val orderedCrates = if keepOrder then crates else crates.reverse

    orderedCrates.foreach((crate) => destination.push(crate))

    stacks
  )

def executeNewCrane(stacks: Stacks, instructions: Vector[Instruction]): Stacks =
  execute(stacks, instructions, false)

def topCrates(stacks: Stacks): String = stacks.map(_.top).mkString
