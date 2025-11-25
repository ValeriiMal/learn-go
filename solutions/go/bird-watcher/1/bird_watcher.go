package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
	for _, v := range birdsPerDay {
		count += v
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	count := 0
	weekSlice := birdsPerDay[(week-1)*7 : week*7]
	for _, v := range weekSlice {
		count += v
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	fixedSlice := birdsPerDay
	for i := range birdsPerDay {
		dayNumber := i + 1
		if dayNumber%2 == 1 {
			fixedSlice[i] = fixedSlice[i] + 1
		}
	}
	return fixedSlice
}
