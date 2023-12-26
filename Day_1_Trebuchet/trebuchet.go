package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)


func outputPrinter(output [][]string) {
    for i, innerSlice := range output {
        fmt.Printf("Slice %d: ", i + 1)
        for _, number := range innerSlice {
            fmt.Printf("%s ", number)
        }
        fmt.Println()
    }
}


func findFirstNum(input string, m map[string]string) string {
    word := "" 
    for i := 0; i < len(input); i++ {
       
        if unicode.IsNumber(rune(input[i])) {
            return string(input[i])
        } 
        
        word += string(input[i])

        if testDigitPrefix(m, word) {
            _, ok := m[word]
            if !ok {
                continue 
            }
            break
        }
        word = string(input[i])
    } 
    return m[word]
}

func findLastNum(input string, m map[string]string) string {
    word := "" 
    for i := len(input) - 1; i >= 0; i-- {         
        if unicode.IsNumber(rune(input[i])) {
            return string(input[i])
        } 
        
        word = string(input[i]) + word
        if testDigitSuffix(m, word) {
            _, ok := m[word]
            if !ok {
                continue 
            }
            break
        }
        word = word[:len(word)-1] 
    } 
    return m[word]
}

func storeInput(input *os.File) [][]string { 

    scanner := bufio.NewScanner(input)
    scanner.Split(bufio.ScanLines)
    
    var numberMap = map[string]string {
        "one": "1",
        "two": "2", 
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6", 
        "seven": "7",
        "eight": "8", 
        "nine": "9", 
    }


    output := [][]string{}

    for scanner.Scan() {
        var lines = []string{}
        lines = append(lines, findFirstNum(scanner.Text(), numberMap))
        lines = append(lines,findLastNum(scanner.Text(), numberMap))
        output = append(output, lines)
    }
 
    outputPrinter(output)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return output
}

func testDigitPrefix(m map[string]string, s string) bool {
    for str,_ := range m {
        if strings.HasPrefix(str, s) {
            return true
        }
    }
    return false
}

func testDigitSuffix(m map[string]string, s string) bool {
    for str,_ := range m {
        if strings.HasSuffix(str, s) {
            return true
        }
    }
    return false
}

func calibrateOutput(output [][]string) int {
    sum := 0
    for _, chrs := range output {

        fstr := chrs[0]
        sstr := chrs[len(chrs) - 1]
        combo, err := strconv.Atoi(fstr + sstr)

        if err != nil {
            log.Fatal(err)
        }
               
        sum += combo

    }

    return sum
}


func main() {
    args := os.Args[1]

    file, err := os.Open(args)
    if err != nil {
        fmt.Printf("Invalid file descriptor: please check spelling\n")
        log.Fatal(err)
    }

    defer file.Close()

    // Start here --- yes this is ugly lol
    fmt.Println("The value the Nisse Need:", calibrateOutput(storeInput(file)))
 
}
