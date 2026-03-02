package almanac

import (
	"log"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// Almanac section name prefixes used in the puzzle input.
const SEEDS_PREFIX = "seeds: "
const SEED_TO_SOIL = "seed-to-soil"
const SOIL_TO_FERTILIZER = "soil-to-fertilizer"
const FERTILIZER_TO_WATER = "fertilizer-to-water"
const WATER_TO_LIGHT = "water-to-light"
const LIGHT_TO_TEMPERATURE = "light-to-temperature"
const TEMPERATURE_TO_HUMIDITY = "temperature-to-humidity"
const HUMIDITY_TO_LOCATION = "humidity-to-location"

var namesOfMaps = []string{
	SEED_TO_SOIL,
	SOIL_TO_FERTILIZER,
	FERTILIZER_TO_WATER,
	WATER_TO_LIGHT,
	LIGHT_TO_TEMPERATURE,
	TEMPERATURE_TO_HUMIDITY,
	HUMIDITY_TO_LOCATION,
}

// parseSeedsLine parses the seeds declaration line into seed IDs.
func parseSeedsLine(line string) []uint64 {
	seedsLine := strings.TrimPrefix(line, SEEDS_PREFIX)
	seedData := []uint64{}

	for seed := range strings.SplitSeq(seedsLine, " ") {
		seed = strings.TrimSpace(seed)
		if seed == "" {
			continue
		}
		seedInt, err := strconv.ParseUint(seed, 10, 64)
		if err != nil {
			log.Fatal("Error converting seed to integer:", err)
		}
		seedData = append(seedData, seedInt)
	}

	return seedData
}

// parseAssociatedRangeLine parses one almanac mapping row into an AssociatedRange.
//
// The expected row format is: "destination source range".
func parseAssociatedRangeLine(line string, mapName string) AlmanacMappingRule {
	trimmedLine := strings.TrimSpace(line)
	var parts [3]string
	count := 0

	for part := range strings.SplitSeq(trimmedLine, " ") {
		if part == "" {
			continue
		}
		if count >= 3 {
			count++
			break
		}
		parts[count] = part
		count++
	}

	if count != 3 {
		log.Fatal("Invalid format for " + mapName + ": " + trimmedLine)
	}

	source, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		log.Fatal("Error converting source to integer for " + mapName + ": " + parts[1])
	}
	destination, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		log.Fatal("Error converting destination to integer for " + mapName + ": " + parts[0])
	}
	valuesRange, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		log.Fatal("Error converting range to integer for " + mapName + ": " + parts[2])
	}

	return AlmanacMappingRule{
		Source:      source,
		Destination: destination,
		Range:       valuesRange,
	}
}

// mapNameFromLine returns the matching mapping section name for line.
func mapNameFromLine(line string) (string, bool) {
	for _, mapName := range namesOfMaps {
		if strings.Contains(line, mapName) {
			return mapName, true
		}
	}

	return "", false
}

// parseAssociatedRanges parses all contiguous mapping rows for one map section.
func parseAssociatedRanges(inputData []string, startIndex int, mapName string) []AlmanacMappingRule {
	associatedRanges := []AlmanacMappingRule{}

	for i := startIndex; i < len(inputData) && len(strings.TrimSpace(inputData[i])) > 0; i++ {
		associatedRanges = append(associatedRanges, parseAssociatedRangeLine(inputData[i], mapName))
	}

	return associatedRanges
}

// extractData reads and parses the puzzle input file.
//
// It returns the raw seed data and all almanac mapping ranges grouped by section name.
func extractData(filename string) ([]uint64, map[string][]AlmanacMappingRule) {
	var seedData []uint64
	associatedRangesByName := make(map[string][]AlmanacMappingRule)

	filebuffer, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	inputData := strings.Split(string(filebuffer), "\n")

	for i, line := range inputData {
		if strings.Contains(line, SEEDS_PREFIX) {
			seedData = parseSeedsLine(line)
			log.Println("Extracted seed integers:", seedData)
			continue
		}

		mapName, exists := mapNameFromLine(line)
		if exists {
			associatedRangesByName[mapName] = parseAssociatedRanges(inputData, i+1, mapName)
		}
	}
	return seedData, associatedRangesByName
}

// seedToSoil applies the seed-to-soil mapping.
func seedToSoil(seed uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := seed
	for _, ar := range associatedRangesByName[SEED_TO_SOIL] {
		if ar.In(seed) {
			associatedValue, _ = ar.GetAssociatedValue(seed)
		}
	}
	return associatedValue
}

// soilToFertilizer applies the soil-to-fertilizer mapping.
func soilToFertilizer(soil uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := soil
	for _, ar := range associatedRangesByName[SOIL_TO_FERTILIZER] {
		if ar.In(soil) {
			associatedValue, _ = ar.GetAssociatedValue(soil)
		}
	}
	return associatedValue
}

// fertilizerToWater applies the fertilizer-to-water mapping.
func fertilizerToWater(fertilizer uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := fertilizer
	for _, ar := range associatedRangesByName[FERTILIZER_TO_WATER] {
		if ar.In(fertilizer) {
			associatedValue, _ = ar.GetAssociatedValue(fertilizer)
		}
	}
	return associatedValue
}

// waterToLight applies the water-to-light mapping.
func waterToLight(water uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := water
	for _, ar := range associatedRangesByName[WATER_TO_LIGHT] {
		if ar.In(water) {
			associatedValue, _ = ar.GetAssociatedValue(water)
		}
	}
	return associatedValue
}

// lightToTemperature applies the light-to-temperature mapping.
func lightToTemperature(light uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := light
	for _, ar := range associatedRangesByName[LIGHT_TO_TEMPERATURE] {
		if ar.In(light) {
			associatedValue, _ = ar.GetAssociatedValue(light)
		}
	}
	return associatedValue
}

// temperatureToHumidity applies the temperature-to-humidity mapping.
func temperatureToHumidity(temperature uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := temperature
	for _, ar := range associatedRangesByName[TEMPERATURE_TO_HUMIDITY] {
		if ar.In(temperature) {
			associatedValue, _ = ar.GetAssociatedValue(temperature)
		}
	}
	return associatedValue
}

// humidityToLocation applies the humidity-to-location mapping.
func humidityToLocation(humidity uint64, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	associatedValue := humidity
	for _, ar := range associatedRangesByName[HUMIDITY_TO_LOCATION] {
		if ar.In(humidity) {
			associatedValue, _ = ar.GetAssociatedValue(humidity)
		}
	}
	return associatedValue
}

// seedsToLocation maps each seed ID all the way to a location ID.
func seedsToLocation(seedNumbers []uint64, associatedRangesByName map[string][]AlmanacMappingRule) []uint64 {
	var locations []uint64
	for _, seed := range seedNumbers {
		soil := seedToSoil(seed, associatedRangesByName)
		fertilizer := soilToFertilizer(soil, associatedRangesByName)
		water := fertilizerToWater(fertilizer, associatedRangesByName)
		light := waterToLight(water, associatedRangesByName)
		temperature := lightToTemperature(light, associatedRangesByName)
		humidity := temperatureToHumidity(temperature, associatedRangesByName)
		location := humidityToLocation(humidity, associatedRangesByName)
		locations = append(locations, location)
	}
	return locations
}

// interval represents a half-open range [start, end).
//
// Half-open ranges make interval splitting and merging easier and less error-prone.
type interval struct {
	start uint64
	end   uint64
}

func minUint64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func maxUint64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

// seedRangesToIntervals converts (start, length) ranges to [start, end) intervals.
func seedRangesToIntervals(seedRanges []SeedRange) []interval {
	intervals := make([]interval, 0, len(seedRanges))
	for _, seedRange := range seedRanges {
		intervals = append(intervals, interval{
			start: seedRange.Source,
			end:   seedRange.Source + seedRange.Range,
		})
	}
	return intervals
}

// normalizeIntervals sorts intervals and merges overlapping or adjacent segments.
//
// This limits interval growth between mapping stages and improves runtime.
func normalizeIntervals(intervals []interval) []interval {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i int, j int) bool {
		if intervals[i].start == intervals[j].start {
			return intervals[i].end < intervals[j].end
		}
		return intervals[i].start < intervals[j].start
	})

	merged := make([]interval, 0, len(intervals))
	for _, current := range intervals {
		if current.start >= current.end {
			continue
		}

		if len(merged) == 0 {
			merged = append(merged, current)
			continue
		}

		last := &merged[len(merged)-1]
		if current.start <= last.end {
			last.end = maxUint64(last.end, current.end)
			continue
		}

		merged = append(merged, current)
	}

	return merged
}

// applyMapToIntervals applies one almanac map to all input intervals.
//
// For each source interval, we split it against mapping ranges:
//   - overlapping pieces are translated to destination coordinates,
//   - uncovered pieces keep identity mapping.
func applyMapToIntervals(inputIntervals []interval, associatedRanges []AlmanacMappingRule) []interval {
	if len(inputIntervals) == 0 {
		return nil
	}

	sortedRanges := make([]AlmanacMappingRule, len(associatedRanges))
	copy(sortedRanges, associatedRanges)
	sort.Slice(sortedRanges, func(i int, j int) bool {
		return sortedRanges[i].Source < sortedRanges[j].Source
	})

	output := make([]interval, 0, len(inputIntervals))

	for _, sourceInterval := range inputIntervals {
		cursor := sourceInterval.start

		for _, mapping := range sortedRanges {
			mappingStart := mapping.Source
			mappingEnd := mapping.Source + mapping.Range

			if mappingEnd <= cursor {
				continue
			}
			if mappingStart >= sourceInterval.end {
				break
			}

			if cursor < mappingStart {
				gapEnd := minUint64(mappingStart, sourceInterval.end)
				output = append(output, interval{start: cursor, end: gapEnd})
				cursor = gapEnd
			}

			if cursor >= sourceInterval.end {
				break
			}

			overlapStart := maxUint64(cursor, mappingStart)
			overlapEnd := minUint64(sourceInterval.end, mappingEnd)
			if overlapStart >= overlapEnd {
				continue
			}

			mappedStart := mapping.Destination + (overlapStart - mappingStart)
			mappedEnd := mappedStart + (overlapEnd - overlapStart)
			output = append(output, interval{start: mappedStart, end: mappedEnd})
			cursor = overlapEnd
		}

		if cursor < sourceInterval.end {
			output = append(output, interval{start: cursor, end: sourceInterval.end})
		}
	}

	return normalizeIntervals(output)
}

// seedRangesToMinLocation computes part 2 by propagating intervals across all maps.
//
// This avoids iterating each individual seed and provides much better performance.
func seedRangesToMinLocation(seedRanges []SeedRange, associatedRangesByName map[string][]AlmanacMappingRule) uint64 {
	intervals := normalizeIntervals(seedRangesToIntervals(seedRanges))

	for _, mapName := range namesOfMaps {
		intervals = applyMapToIntervals(intervals, associatedRangesByName[mapName])
	}

	minLocation := uint64(math.MaxUint64)
	for _, locationInterval := range intervals {
		if locationInterval.start < minLocation {
			minLocation = locationInterval.start
		}
	}

	return minLocation
}

// seedRangesFromSeedData converts (start, length) pairs into SeedRange values.
func seedRangesFromSeedData(seedData []uint64) []SeedRange {
	dataSize := len(seedData)
	if dataSize%2 != 0 {
		log.Fatal("Invalid seed data length, must be even")
	}

	var seedRanges []SeedRange
	for i := 0; i < dataSize; i += 2 {
		start := seedData[i]
		rangeSize := seedData[i+1]
		seedRanges = append(seedRanges, SeedRange{Source: start, Range: rangeSize})
	}
	return seedRanges
}

// GetMinimumLocationForPart1 solves part 1 of the puzzle for filename.
func GetMinimumLocationForPart1(filename string) uint64 {
	seeds, associatedRangesByName := extractData(filename)
	locations := seedsToLocation(seeds, associatedRangesByName)
	minLocation := slices.Min(locations)
	return minLocation
}

// GetMinimumLocationForPart2 solves part 2 of the puzzle for filename.
func GetMinimumLocationForPart2(filename string) uint64 {
	seedData, associatedRangesByName := extractData(filename)
	seedRanges := seedRangesFromSeedData(seedData)
	minLocation := seedRangesToMinLocation(seedRanges, associatedRangesByName)
	return minLocation
}
