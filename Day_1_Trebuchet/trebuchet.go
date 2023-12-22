package main

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

func storeInput(input os.File) [][]rune { 

    scanner := bufio.NewScanner(&input)
    scanner.Split(bufio.ScanLines)
    
    output := [][]rune{}

    for scanner.Scan() {
        var lines = []rune{}
       
        for _,ele := range scanner.Text() {
            if unicode.IsNumber(ele) {
                lines = append(lines, ele)
            }
        }
        output = append(output, lines)
    }

    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return output
}

/*

The newly-improved calibration document consists of 
lines of text; each line originally contained a specific 
calibration value that the Elves now need to recover.
On each line, the calibration value can be found by 
combining the first digit and the last digit (in that order)
to form a single two-digit number.

*/
func calibrateOutput(output [][]rune) int {

    return 0
}


func main() {
    input := "testinput.txt"
    file, err := os.Open(input)

    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    // Start here
    output := storeInput(*file)
    calibrateOutput(output)


    
    
}
