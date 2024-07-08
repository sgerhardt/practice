package time

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	now := time.Now()
	println(now.String())

	println(now.AddDate(0, 0, -1).String())

	// Tricky things with time api
	//1. time.Add is to add a duration, while time.Sub is to substract a time. To substract time use AddDate
	// If the zone abbreviation is unknown, Parse records the time as being in a fabricated location with the given zone abbreviation and a zero offset.
	//2. time package does NOT necessarily _convert_ times when you call .UTC(), it just changes the location
	//3. Time is machine and location dependent - different locale definitions for different operating systems
	println(now.Local().String())

	now.Sub(now)

	now.UTC()

	// convert to UTC
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		print(err)
	}
	println(loc.String())
	parsedTimeUTC, err := time.ParseInLocation(now.String(), now.String(), loc)
	if err != nil {
		print(err)
	}
	println(parsedTimeUTC.String())
}

func TestConvertTimezonesToUTC(t *testing.T) {
	// Note here we can call .UTC() and expect time to tranlsate likely because we can also parse the zones as expected.
	loc, err := time.LoadLocation("America/Chicago")
	require.NoError(t, err)
	println(loc.String())
	now := time.Now()
	// verify CST
	zone, offset := now.Local().Zone()
	assert.Equal(t, "CST", zone)
	// offset is given in seconds
	assert.Equal(t, -21600, offset)

	// Verify expected hours. CST is 6 hours behind UTC. E.g. if it's 10AM in CST, it's 4PM UTC
	assert.Equal(t, -21600/60/60, -6)

	utcTime := now.UTC()
	require.NoError(t, err)
	utcZone, _ := utcTime.Zone()
	assert.Equal(t, "UTC", utcZone)

	assert.Equal(t, now.Add(time.Duration(6)*time.Hour).Hour(), utcTime.Hour())

	// Ambiguous zone info (at least on windows) so offset is 0 and zone is set to UTC despite desire to set to Pacific time (or maybe Phillipines..)
	// This means conversion to UTC won't work as expected
	const longForm = "2006-01-02 15:04:05 MST"
	pacificTime, err1 := time.Parse(longForm, "2016-01-17 23:04:05 PST")
	require.NoError(t, err1)
	now.Year()
	fmt.Printf("PST to UTC: %v\n\n", pacificTime.UTC())
	zone, offset = pacificTime.UTC().Zone()
	assert.Equal(t, 0, offset)
	assert.Equal(t, "UTC", zone)

	assert.Equal(t, "PST", pacificTime.Location().String())
	pacificZone, offset := pacificTime.Zone()
	//NOTE HOW THIS SHOWS AN OFFSET OF 0. THIS IS WRRRROOOOONG SINCE ACTUAL OFFSET IS -8 hours for PT.
	//To fix this, we'd need to Parse with location and a known abbreviation for the zone
	assert.Equal(t, 0, offset)
	assert.Equal(t, "PST", pacificZone)

	loc, err = time.LoadLocation("America/Los_Angeles")
	require.NoError(t, err)
	actualPacificTime, err := time.ParseInLocation(longForm, "2016-01-17 23:04:05 PDT", loc)
	zone, offset = actualPacificTime.Zone()
	assert.Equal(t, "PST", zone)
	assert.Equal(t, -28800, offset)
	pacificToUTC := actualPacificTime.UTC()

	// Pacific time is offset 8 hours
	assert.Equal(t, pacificToUTC.Add(time.Hour*-8).Hour(), actualPacificTime.Hour())
	assert.Equal(t, -28800/60/60, -8)

}

func TestLeapYear(t *testing.T) {
	assert.True(t, isLeapYear(2020))
	assert.True(t, isLeapYear(2000))
	assert.False(t, isLeapYear(2001))
}

func TestDaysInMonth(t *testing.T) {
	value := "Thu, 02/17/22, 10:47PM"
	layout := "Mon, 01/02/06, 03:04PM" // This is tricky... it uses a mnemonic (a device such as a pattern of letters, ideas, or associations that assists in remembering something.)
	//to define layout. An easy way to remember the layout, tho it is ugly is 01/02 03:04:05PM '06 -0700 (note the increments...).
	//Also note that whatever layout you use, the value will need to match the format
	parsedTime, err := time.Parse(layout, value)
	require.NoError(t, err)
	fmt.Println(parsedTime.String())
	day := daysInMonth(parsedTime)
	assert.Equal(t, 28, day)
}
