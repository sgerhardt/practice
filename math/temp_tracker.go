package math

import "math"

//Write a class TempTracker with these methods:
//
//insert()—records a new temperature
//getMax()—returns the highest temp we've seen so far
//getMin()—returns the lowest temp we've seen so far
//getMean()—returns the mean ↴ of all temps we've seen so far
//getMode()—returns a mode ↴ of all temps we've seen so far
//Optimize for space and time. Favor speeding up the getter methods getMax(), getMin(), getMean(), and getMode() over speeding up the insert() method.
//
//getMean() should return a double, but the rest of the getter methods can return integers. Temperatures will all be inserted as integers. We'll record our temperatures in Fahrenheit, so we can assume they'll all be in the range 0..1100..110.
//
//If there is more than one mode, return any of the modes.

type TempTracker struct {
	maxTemp int
	minTemp int
	mode    int
	mean    float64

	// to quickly calc mean
	tempTotal        int
	tempOccurrences  []int
	numberOfReadings int
}

func initTempTracker() *TempTracker {
	return &TempTracker{
		maxTemp:          math.MinInt,
		minTemp:          math.MaxInt,
		mode:             0,
		mean:             0,
		tempTotal:        0,
		tempOccurrences:  make([]int, 110),
		numberOfReadings: 0,
	}
}
func (t *TempTracker) insert(temp int) {
	t.numberOfReadings++
	t.tempTotal += temp
	t.tempOccurrences[temp]++

	t.mean = float64(t.tempTotal) / float64(t.numberOfReadings)

	if temp > t.maxTemp {
		t.maxTemp = temp
	}

	if temp < t.minTemp {
		t.minTemp = temp
	}

}

func (t *TempTracker) getMean() float64 {
	return t.mean
}

func (t *TempTracker) getMax() int {
	return t.maxTemp
}

func (t *TempTracker) getMin() int {
	return t.minTemp
}

func (t *TempTracker) getMode() int {
	return t.mode
}

func (t *TempTracker) getModes() []int {
	return t.tempOccurrences
}
