package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	//	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func changeLines() []string {

	var insertLine string = ""
	var lineAfterNumber, lineAfterName, lineAfterKey, lineAfterSize = false, false, false, false
	var newInserts []string
	Mylines, err := readLines("test.txt")
	if err == nil {
		fmt.Println("Length Of Slice", len(Mylines))
		for _, line := range Mylines {
			line = strings.TrimSpace(line)

			if lineAfterSize {
				lineAfterSize = false
				if _, err := strconv.Atoi(line); err == nil || (line == "") {
					fmt.Println("No Desc: " + line)
					insertLine += "NULL),"
					newInserts = append(newInserts, insertLine)
					insertLine = ""
					insertLine += "(" + line + ",'"
					lineAfterNumber = true
				} else {
					fmt.Println("Desc: " + line)
					insertLine += "'" + line + "'),"
					newInserts = append(newInserts, insertLine)
					insertLine = ""

				}

			} else if _, err := strconv.Atoi(line); err == nil {
				fmt.Println("Number: " + line)
				insertLine = "(" + line + ",'"
				lineAfterNumber = true

			} else if lineAfterNumber {
				fmt.Println("Name: " + line)
				insertLine += "" + line + "',NULL,"
				lineAfterNumber = false
				lineAfterName = true
			} else if lineAfterName {
				fmt.Println("Key: " + line)
				lineAfterName = false
				if line == "" {
					lineAfterName = true
				} else if line == "X" {
					insertLine += "1" + ",'"
					lineAfterKey = true
				} else {
					insertLine += "0" + ",'" + line + "',"
					lineAfterSize = true
				}

			} else if lineAfterKey {
				fmt.Println("Size: " + line)
				insertLine += "" + line + "',"
				lineAfterSize = true
				lineAfterKey = false
			}

		}
	}
	return newInserts
}

func removeWhiteSpace() {
	var newInserts []string

	Mylines, err := readLines("test.txt")
	if err == nil {

		fmt.Println("Length Of Slice", len(Mylines))
		for _, line := range Mylines {
			line = strings.TrimSpace(line)
			newInserts = append(newInserts, line)
		}
	}
	writeLines(newInserts, "test2.txt")

}

func changeToVarchar() {
	var newInserts []string

	MyLines, err := readLines("varchar1.txt")
	if err == nil {

		fmt.Println("Length Of Slice", len(MyLines))
		var valueCheck string

		for num, line := range MyLines {
			fmt.Println("Num: ", num)
			if len(line) > 6 {
				valueCheck = line[0:2]
			}
			if valueCheck == "('" {
				allOnComma := strings.Split(line, ",")
				comma1 := allOnComma[1]
				comma1 = strings.TrimSpace(comma1)
				comma2 := allOnComma[2]
				comma2 = strings.TrimSpace(comma2)
				comma2 = "'" + comma2 + "'"
				comma1 = "'" + comma1 + "'"
				allOnComma[1] = comma1
				allOnComma[2] = comma2
				finalString := strings.Join(allOnComma, ",")
				newInserts = append(newInserts, finalString)
				fmt.Println("FinalString: ", finalString)
				valueCheck = ""

			} else {
				newInserts = append(newInserts, line)
			}

		}
		writeLines(newInserts, "varchar2.txt")

	}

}

func Covariance(x, y []float64) float64 {
	x1 := x
	y1 := y
	avgX := avgFloat(x1)
	avgY := avgFloat(y1)
	var additionNeeded []float64
	for num, _ := range x1 {

		additionNeeded = append(additionNeeded, (x1[num]-avgX)*(y1[num]-avgY))
	}

	covariance := addFloat(additionNeeded) / (float64(len(x1)) - 1)
	fmt.Println("Covariance: ", covariance)
	return covariance
}

func Corrolation() {
	x1 := []float64{2.1, 2.5, 4.0, 3.6}
	y1 := []float64{8, 12, 14, 10}
	sdX := standardDeviationFloat(x1, avgFloat(x1))
	sdY := standardDeviationFloat(y1, avgFloat(y1))

	corrlation := Covariance(x1, y1) / (sdX * sdY)
	fmt.Println("Correlation: ", corrlation)
}

func standardDeviationFloat(slice []float64, avg float64) float64 {
	var squaredNums []float64
	for _, value := range slice {
		squaredNums = append(squaredNums, math.Pow((value-avg), 2))
	}
	summed := addFloat(squaredNums)
	return math.Sqrt((summed / (float64(len(slice) - 1))))

}

func addFloat(slice []float64) float64 {
	var finalSum float64
	for _, value := range slice {
		finalSum += value
	}
	return finalSum

}

func avgFloat(slice []float64) float64 {
	var avg float64
	for _, value := range slice {
		avg += value
	}
	avg = avg / float64(len(slice))
	return avg
}

type State struct {
	Name    string
	SCities []City
	Count   int
}

type City struct {
	Name  string
	Count int
}

func ReadIn() [][]string {
	csvfile, err := os.Open("SC.csv")
	r := csv.NewReader(csvfile)
	//	var sliceOfRows [][]string
	sliceOfRows, err := r.ReadAll()
	//	for {
	//if err == io.EOF{
	//		break
	//	}
	//		if err != nil {
	//			fmt.Println("err: ", err)
	//		}
	//		sliceOfRows = append(sliceOfRows, record)
	//	}
	if err != nil {
		fmt.Println("err: ", err)
	}
	return sliceOfRows
}

func MakeCount() []State {
	var stateList []State
	sliceOfRows := ReadIn()
	fmt.Println("Len: ", len(sliceOfRows))
	//	fmit.Println(sliceOfRows[1421][1])

	for _, value := range sliceOfRows {
		//		delaySecond(1)
		city := value[0]
		state := value[1]
		//	fmt.Printf("State: %v, City: %v\n", state, city)
		var currentState *State

		stateFound := false
		for num, _ := range stateList {
			currentState = &stateList[num]
			//			fmt.Printf("StateName: %v, ValueName: %v\n", stateValue.Name, state)
			if currentState.Name == state && stateFound == false {

				fmt.Println("inside")
				stateFound = true
				cityFound := false
				//			fmt.Println(currentState.SCities)
				var currentCity *City
				for nm, _ := range currentState.SCities {
					currentCity = &currentState.SCities[nm]
					//fmt.Printf("CityName: %v, cityValue: %v\n", currentCity.Name, city)

					if currentCity.Name == city && cityFound == false {
						currentCity.Count = currentCity.Count + 1
						cityFound = true

					}
				}
				if cityFound == false {
					fmt.Println("Added city")
					newCity := City{}
					newCity.Name = city
					newCity.Count += 1
					currentState.SCities = append(currentState.SCities, newCity)

				}
				currentState.Count = currentState.Count + 1
			}
		}
		if stateFound == false {
			fmt.Println("Added State")
			newState := State{}
			newState.Name = state
			newState.Count += 1
			newCity := City{}
			newCity.Name = city
			newCity.Count += 1
			newState.SCities = append(newState.SCities, newCity)
			stateList = append(stateList, newState)
			//		fmt.Println(stateList)
		}

	}
	fmt.Println(stateList)
	return stateList

}
func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}
func WriteState() error {
	stateList := MakeCount()
	file, err := os.Create("StatesCity.csv")
	if err != nil {
		return err
	}
	defer file.Close()
	var total int
	w := bufio.NewWriter(file)
	for _, line := range stateList {
		fmt.Fprintf(w, "%v\t\t%v\t%v\t%v%v\n", line.Name, line.Count, line.Name, line.Name, " cont. ")
		total = total + line.Count
		internalCounter := 0
		for _, value := range line.SCities {
			internalCounter = internalCounter + 1
			fmt.Fprintf(w, "\t%v\t%v\t", value.Name, value.Count)
			if internalCounter == 2 {
				fmt.Fprintf(w, "\n")
				internalCounter = 0
			}

		}
		fmt.Fprintf(w, "\n")

	}
	fmt.Fprintf(w, "Total: %v", total)
	return w.Flush()

}
func main() {
	//	testLine := changeLines()
	//	writeLines(testLine, "test2.txt")
	//removeWhiteSpace()
	//	changeToVarchar()
	//	Covariance()
	// Corrolation()
	WriteState()
	//	fmt.Println(len(stateList))
}
