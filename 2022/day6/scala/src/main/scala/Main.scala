import scala.io.Source
import scala.annotation.tailrec
import scala.collection.immutable.Set
import String.allDifferent

val FileName = "input.txt"

@main def main: Unit =
  part1()

  part2()

def part1(): Unit =
  findMarkerInLines(startOfPacketMarker)

def part2(): Unit =
  findMarkerInLines(startOfMessageMarker)

def findMarkerInLines(markerFinder: (line: String) => Option[Marker]) =
  getLines(FileName)
    .map(markerFinder)
    .zipWithIndex
    .foreach((maybeMarker, lineN) =>
      println(s"In line $lineN marker was found at position $maybeMarker")
    )

def getLines(fileName: String) = Source.fromFile(fileName).getLines()

type Marker = Int

def startOfPacketMarker(line: String) = findMarker(line, 4)

def startOfMessageMarker(line: String) = findMarker(line, 14)

def findMarker(line: String, packetSize: Int): Option[Marker] =
  @tailrec
  def aux(line: String, index: Int): Option[Marker] =
    val outerLimit = index + packetSize
    if outerLimit < line.length then
      val packet = line.slice(index, outerLimit)
      if packet.allDifferent then Some(outerLimit)
      else aux(line, index + 1)
    else None

  aux(line, 0)

object String {
  extension (s: String) def allDifferent: Boolean = s.length == Set.from(s).size
}
