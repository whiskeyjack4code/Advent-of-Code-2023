package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/*

  We have a multiline file where each line is a string and contains chars and numbers.
  We have to parse the line and extract the numbers, converting them to their integer form and concatenating
  them together to form a whole number.
  We need to have a running total for the calibration amount, and add the wholenumber to it after each
  loop of the line.

My Algorithm for this problem:

    1. Open the file for reading and check for errors
    2. Loop over each line of the file and in turn, loop over each character of the i'th line.
    3. Attempt to convert the char to it's integer form, if it errors out, it's not a true integer so continue the loop
      to the next character.
    4. One the first real integer digit is found, store it in a variable and break to the outer loop because we only
      need the first and last digit, a break will allow control this.
    5. Loop backwards over the same line and repeat steps 3 to 4.
    6. At this point we should have the first and last digit of the line, if any.
    7. We now convert them to a string and concatenate them
    8. Once the numbers are concatenated as a string, we covnert it back to its integer value giving
      us the whole number.
    9. Add this whole number to the running total.
    10. Print out the running total.

*/

func main() {

	fileName := "input.txt"
	file, err := os.Open(fileName)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	calibrationTotal := 0

	var firstDigit int = 0
	var lastDigit int = 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, v := range line {
			firstDigit = 0 // reset the variable to prevent carryover on a line with no integers
			num, err := strconv.Atoi(string(v))

			if err != nil {
				continue // We skip non-integer values
			}
			firstDigit = num
			break
		}

		for i := len(line) - 1; i >= 0; i-- {
			lastDigit = 0 // reset the variable to prevent carryover on a line with no integers
			v := line[i]
			num, err := strconv.Atoi(string(v))

			if err != nil {
				continue // We skip non-integer values
			} else {
				lastDigit = num
				break
			}
		}

		// We need to combine firstNum with lastNum to make the wholeNum and convert it back to an int
		// to add it to the running calibration total
		firstAsString := strconv.Itoa(firstDigit)
		lastAsString := strconv.Itoa(lastDigit)
		wholeNum, err := strconv.Atoi((firstAsString + lastAsString))

		// Check if conversion failed
		if err != nil {
			fmt.Println(err)
		}
		calibrationTotal += wholeNum
	}
	fmt.Println(calibrationTotal) // Print the answer
}
