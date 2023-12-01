object Day1 extends Common {

  def main(args: Array[String]): Unit = {
    this.challenge1();
    this.challenge2();
  }

  def challenge1(): Unit = {
    val inputArray = this.readLinesFromFile("inputs/1_small.input").map(_.toInt)

    var result = 0

    for (a <- 1 to inputArray.length - 1) {
      if (inputArray(a) - inputArray(a - 1) > 0) {
        result = result + 1
      }
    }

    println(result)
  }

  def challenge2(): Unit = {
    val inputArray = readLinesFromFile("inputs/1_large.input").map(_.toInt)
    var result: Int = 0
    val slidingWindowSize: Int = 3
    var currentSum: Int = 0

    for (a <- 0 to inputArray.length - 1) {
      if (a >= slidingWindowSize) {
        val newSum =
          currentSum - inputArray(a - slidingWindowSize) + inputArray(a)
        if (newSum > currentSum) {
          result += 1
        }
      }
      currentSum += inputArray(a)
    }

    println(result)
  }
}
