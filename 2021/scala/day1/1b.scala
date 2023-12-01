
import scala.io.Source

object Demo {

  def main(args: Array[String]) = {
    val inputArray = readLinesFromFile("1_large.input").map(_.toInt)
    var result: Int = 0
    val slidingWindowSize: Int = 3
    var currentSum: Int = 0

    for (a <- 0 to inputArray.length - 1) {
      if(a >= slidingWindowSize) {
        val newSum = currentSum - inputArray(a - slidingWindowSize) + inputArray(a)
        if(newSum > currentSum) {
          result += 1
        }
      }
      currentSum += inputArray(a)
    }

    println(result)
  }

  def readLinesFromFile(path: String): Array[String] = {
    return Source.fromFile(path).getLines.toArray
  }
}