
import scala.io.Source
import scala.util.matching.Regex


object Demo {

  def main(args: Array[String]) = {

    val inputArray = readLinesFromFile("2_large.input")
    val regex: Regex = "([a-z]+) ([0-9]+)".r

    var verticalPos = 0;
    var horizontalPos = 0;
    var aim = 0;

    for (operation <- inputArray) {

      operation match {
        case regex(op, amount) => {
          if(op.equals("forward")) {
            horizontalPos += amount.toInt;
            verticalPos += amount.toInt * aim;
          } else if (op.equals("down")) {
            aim += amount.toInt;
          } else if (op.equals("up")) {
            aim -= amount.toInt
          }
      }
      }
    }

    println(verticalPos * horizontalPos)
  }

  def readLinesFromFile(path: String): Array[String] = {
    return Source.fromFile(path).getLines.toArray
  }
}