package boatrace

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const TIME_PREFIX = "Time:"
const DISTANCE_PREFIX = "Distance:"

func extractDataForPart1(filename string) ([]uint64, []uint64) {
	var timeData []uint64
	var distanceData []uint64

	filebuffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	inputData := strings.Split(string(filebuffer), "\n")
	r, _ := regexp.Compile(`\d+`)
	for _, line := range inputData {
		if strings.Contains(line, TIME_PREFIX) {
			timeLine, _ := strings.CutPrefix(line, TIME_PREFIX)
			times := r.FindAllString(timeLine, -1)
			for _, time := range times {
				timeInt, err := strconv.ParseUint(time, 10, 0)
				if err != nil {
					log.Fatal("Error converting time to integer:", err)
				}
				timeData = append(timeData, timeInt)
			}
		}

		if strings.Contains(line, DISTANCE_PREFIX) {
			distanceLine := strings.TrimSpace(strings.TrimPrefix(line, DISTANCE_PREFIX))
			distances := r.FindAllString(distanceLine, -1)
			for _, distance := range distances {
				distanceInt, err := strconv.ParseUint(distance, 10, 0)
				if err != nil {
					log.Fatal("Error converting distance to integer:", err)
				}
				distanceData = append(distanceData, distanceInt)
			}
		}
	}

	return timeData, distanceData
}

func extractDataForPart2(filename string) (uint64, uint64) {
	var timeData uint64
	var distanceData uint64
	var err error

	filebuffer, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	inputData := strings.Split(string(filebuffer), "\n")
	r, _ := regexp.Compile(`\d+`)
	for _, line := range inputData {
		if strings.Contains(line, TIME_PREFIX) {
			timeLine, _ := strings.CutPrefix(line, TIME_PREFIX)
			times := r.FindAllString(timeLine, -1)
			time := strings.Join(times, "")
			timeData, err = strconv.ParseUint(time, 10, 64)
			if err != nil {
				log.Fatal("Error converting time to integer:", err)
			}
		}

		if strings.Contains(line, DISTANCE_PREFIX) {
			distanceLine := strings.TrimSpace(strings.TrimPrefix(line, DISTANCE_PREFIX))
			distances := r.FindAllString(distanceLine, -1)
			distance := strings.Join(distances, "")
			distanceData, err = strconv.ParseUint(distance, 10, 64)
			if err != nil {
				log.Fatal("Error converting distance to integer:", err)
			}
		}
	}

	return timeData, distanceData
}

func NumberOfWaysToWin(duration uint64, bestDistance uint64) uint64 {
	var waysToWin uint64 = 0

	for holdDuration := uint64(0); holdDuration <= duration; holdDuration++ {
		distanceTraveled := (duration - holdDuration) * holdDuration
		if distanceTraveled > bestDistance {
			waysToWin++
		}
	}

	return waysToWin
}

func RecordBreakingsProducts(filename string) uint64 {
	durations, bestDistances := extractDataForPart1(filename)
	log.Println("Extracted time durations:", durations)
	log.Println("Extracted best distances:", bestDistances)
	iterations := len(durations)
	var product uint64 = 1

	for i := range iterations {
		ways := NumberOfWaysToWin(durations[i], bestDistances[i])
		log.Printf("For duration %d and best distance %d, number of ways to win: %d\n", durations[i], bestDistances[i], ways)
		product *= ways
	}

	return product
}

func RecordBreaking(filename string) uint64 {
	duration, bestDistance := extractDataForPart2(filename)
	log.Printf("Extracted time duration: %d\n", duration)
	log.Printf("Extracted best distance: %d\n", bestDistance)
	ways := NumberOfWaysToWin(duration, bestDistance)
	return ways
}
