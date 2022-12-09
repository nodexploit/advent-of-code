import scala.io.Source
val FileName = "test.txt"

@main def main: Unit =
  part1()

def part1(): Unit =
//   DirectoryTree.printTree(example)
//   println("--------")
//   DirectoryTree.printTree(example2)
//   println("--------")
//   DirectoryTree.printTree(example3)

  val tree =
    getLines(FileName).foldLeft(DirectoryTree.init)((directoryTree, line) =>
      val tree = reduceLine(directoryTree, line)
      println(s"------------> $line")
      println("-------------")
      DirectoryTree.printTree(tree)
      tree
    )

  println("----FINAL----")
  DirectoryTree.printTree(tree)

def getLines(fileName: String) = Source.fromFile(fileName).getLines()

def reduceLine(directoryTree: DirectoryTree, line: String): DirectoryTree =
  line match
    case line if line.startsWith("$") => command(directoryTree, line)
    case _                            => directoryTree.addListing(line)

def command(directoryTree: DirectoryTree, line: String) =
  val split = line.split(" ")
  val command = split(1)
  command match
    case "cd" =>
      val dirName = split(2)
      if dirName == ".." then directoryTree.previousDirectory
      else directoryTree.changeDirectory(dirName)
    case _ => directoryTree

val example = Dir("/", Vector(Dir("a", Vector.empty, true), File("b.txt", 148)))

val example2 = DirectoryTree.addChildren(example, File("foo", 1))

val example3 =
  DirectoryTree.addChildren(example.changeDirectory("/"), File("lel", 0))
