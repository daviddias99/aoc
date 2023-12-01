import scala.util.matching.Regex
import scala.util.control.Breaks._
import scala.collection.mutable.ListBuffer

object Day3 extends Common {

  def main(args: Array[String]): Unit = {
    // this.challenge1();
    this.challenge2();
  }

  def challenge1(): Unit = {
    val inputArray = readLinesFromFile("inputs/3_large.input")
    val countsArr: Array[Int] = new Array(inputArray(0).length)

    for (value <- inputArray) {
      for (i <- 0 to value.length - 1) {
        countsArr(i) += value(i).toString.toInt
      }
    }

    var resultGamma = ""
    var resultEpsilon = ""

    for (i <- 0 to inputArray(0).length - 1) {
      if (countsArr(i) > inputArray.length / 2) {
        resultGamma = resultGamma.concat("1")
        resultEpsilon = resultEpsilon.concat("0")
      } else {
        resultGamma = resultGamma.concat("0")
        resultEpsilon = resultEpsilon.concat("1")
      }
    }

    var gamma = Integer.parseInt(resultGamma, 2);
    var epsilon = Integer.parseInt(resultEpsilon, 2);
    println(gamma * epsilon)
  }

  def getCountsArr(input: Array[String], pos: Int): Int = {
    var res: Int = 0;

    for (value <- input) {
      res += value(pos).toString.toInt
    }

    return res;
  }

  def getCountsArr(input: ListBuffer[String], pos: Int): Int = {
    var res: Int = 0;

    for (value <- input) {
      res += value(pos).toString.toInt
    }

    return res;
  }

  def challenge2(): Unit = {
    val inputArray: Array[String] = readLinesFromFile("inputs/3_large.input")
    val countsArr: Array[Int] = new Array(inputArray(0).length)

    var i = 0;

    var nextArr: ListBuffer[String] = new ListBuffer[String]();
    var currArray: ListBuffer[String] = new ListBuffer[String]();

    for (v <- inputArray) {
      currArray += v;
    }

    while (i < inputArray(0).length && currArray.length > 1) {
      var count = getCountsArr(currArray, i)
      nextArr = new ListBuffer[String]();
      for (j <- 0 to currArray.length - 1) {
        if (count >= currArray.length * 1.0 / 2 ) {
          if (currArray(j)(i) == '1') {
            nextArr += currArray(j);
          }
        } else {
          if (currArray(j)(i) == '0') {
            nextArr += currArray(j);
          }
        }
      }

      currArray = nextArr;
      i += 1;
    }

    var o2 =  Integer.parseInt(nextArr(0),2)

    i = 0;

    nextArr = new ListBuffer[String]();
    currArray = new ListBuffer[String]();

    for (v <- inputArray) {
      currArray += v;
    }

    while (i < inputArray(0).length && currArray.length > 1) {
      var count = getCountsArr(currArray, i)
      nextArr = new ListBuffer[String]();
      for (j <- 0 to currArray.length - 1) {
        if (count < currArray.length * 1.0 / 2 ) {
          if (currArray(j)(i) == '1') {
            nextArr += currArray(j);
          }
        } else {
          if (currArray(j)(i) == '0') {
            nextArr += currArray(j);
          }
        }
      }

      currArray = nextArr;
      i += 1;
    }
    var co2 = Integer.parseInt(nextArr(0), 2) 

    println(co2 * o2)
  }
}
