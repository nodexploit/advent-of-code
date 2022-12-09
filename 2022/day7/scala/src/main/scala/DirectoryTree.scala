import scala.collection.View.Empty
import scala.annotation.tailrec

sealed trait DirectoryTree
final case class Dir(
    name: String,
    children: Vector[DirectoryTree] = Vector.empty[DirectoryTree],
    currentDirectory: Boolean = false
) extends DirectoryTree
final case class File(name: String, size: Int) extends DirectoryTree

given Showable[DirectoryTree] with
  def show(t: DirectoryTree): String =
    t match
      case d: Dir =>
        s"- ${d.name} (dir${if d.currentDirectory then " X" else ""})"
      case f: File => s"- ${f.name} (file, size=${f.size})"

trait Showable[A]:
  def show(a: A): String

  extension (a: A) def showIt = show(a)

object DirectoryTree {

  def init: DirectoryTree = Dir("/", Vector.empty[DirectoryTree], false)

  //  TODO: Tail recursion?
  def printTree(tree: DirectoryTree)(using s: Showable[DirectoryTree]): Unit =
    def aux(tree: DirectoryTree, depth: Int): Unit =
      print(" ".repeat(depth))
      tree match
        case f: File => println(f.showIt)
        case d @ Dir(name, children, _) =>
          println(d.showIt)
          children.foreach(aux(_, depth + 1))

    aux(tree, 0)

  def addChildren(tree: DirectoryTree, child: DirectoryTree): DirectoryTree =
    tree match
      case Dir(name, children, currentDirectory) =>
        if currentDirectory then Dir(name, children :+ child, currentDirectory)
        else if children.nonEmpty then
          Dir(
            name,
            children.map(c => addChildren(c, child)),
            currentDirectory
          )
        else tree
      case _ => tree

  extension (t: DirectoryTree)
    def changeDirectory(dirName: String): DirectoryTree =
      t match
        case Dir(name, children, currentDirectory) =>
          if children.nonEmpty then
            Dir(
              name,
              children.map(c => c.changeDirectory(dirName)),
              name == dirName
            )
          else Dir(name, children, name == dirName)
        case _ => t

    def previousDirectory: DirectoryTree =
      t match
        case d @ Dir(name, children, currentDirectory) =>
          if children.nonEmpty then
            val activeChildren = children.exists((child) =>
              child match
                case Dir(name, children, currentDirectory) => currentDirectory
                case _                                     => false
            )
            println(s"$activeChildren -------")

            if activeChildren then
              Dir(
                name,
                children.map(c => c.deactivateDirectory),
                true
              )
            else
              Dir(
                name,
                children.map(c => c.previousDirectory),
                currentDirectory
              )
          else d
        case _ => t

    def deactivateDirectory: DirectoryTree =
      t match
        case Dir(name, children, currentDirectory) => Dir(name, children, false)
        case _                                     => t

    def addListing(line: String): DirectoryTree =
      val split = line.split(" ")
      split(0) match
        case "dir" => addChildren(t, Dir(split(1)))
        case _ =>
          val size = split(0)
          val name = split(1)
          addChildren(t, File(name, size.toInt))
}
