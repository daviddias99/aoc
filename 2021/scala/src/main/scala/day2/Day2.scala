import scala.util.matching.Regex

object Day2 extends Common {

  def main(args: Array[String]): Unit = {
    this.challenge1();
    this.challenge2();
  }

  def challenge1(): Unit = {
    val inputArray = readLinesFromFile("inputs/2_large.input")
    val regex: Regex = "([a-z]+) ([0-9]+)".r

    var verticalPos = 0;
    var horizontalPos = 0;

    for (operation <- inputArray) {
      operation match {
        case regex(op, amount) => {
          if(op.equals("forward")) {
            horizontalPos += amount.toInt;
          } else if (op.equals("down")) {
            verticalPos += amount.toInt;
          } else if (op.equals("up")) {
            verticalPos -= amount.toInt
          }
        }
      }
    }

    println(verticalPos * horizontalPos)
  }

  def challenge2(): Unit = {
    val inputArray = readLinesFromFile("inputs/2_large.input")
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
}
