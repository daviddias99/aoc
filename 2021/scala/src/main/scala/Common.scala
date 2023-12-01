import scala.io.Source

class Common {
  def readLinesFromFile(path: String): Array[String] = {
    return Source.fromFile(path).getLines.toArray
  }
}