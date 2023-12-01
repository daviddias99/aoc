
import scala.io.Source

object Demo {

  def main(args: Array[String]) = {
    val inputArray = readLinesFromFile("1_small.input").map(_.toInt)
    
    var result = 0

    for (a <- 1 to inputArray.length - 1) {
      if(inputArray(a) - inputArray(a - 1) > 0) {
        result = result + 1
      }
    }

    println(result)
  }

  def readLinesFromFile(path: String): Array[String] = {
    return Source.fromFile(path).getLines.toArray
  }
}