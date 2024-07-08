package time

//
//import (
//	"errors"
//	"fmt"
//	"os/exec"
//	"runtime"
//	"time"
//)
//
//// To execute Go code, please declare a func main() in a package "main"
//
//package main
//
//import( "fmt"
//"runtime"
//"os/exec"
//"time"
//"errors"
//)
//
//func main() {
//	println(runtime.GOOS)
//	println(runtime.GOARCH)
//	cmd := exec.Command("sh", "-c", "lsb_release -a")
//	stdoutStderr, err := cmd.CombinedOutput()
//	fmt.Printf("%s\n", stdoutStderr)
//	if err != nil {
//		return
//	}
//
//
//	now := time.Now()
//	println(now.String())
//	// subtract a day
//	println(now.AddDate(0, 0, -1).String())
//
//	// verify zone and offset
//	zone, offset := now.Local().Zone()
//	fmt.Printf("Zone is: %+v\n", zone)
//	fmt.Printf("Offset is: %+v\n", offset)
//
//	// Test conversion
//	const longForm = "2006-01-02 15:04:05 PDT"
//	loc, err := time.LoadLocation("America/Los_Angeles")
//	if err != nil{
//		println(err)
//		return
//	}
//	actualPacificTime, err := time.ParseInLocation(longForm, "2016-01-17 23:04:05 PDT", loc)
//	zone, offset = actualPacificTime.Zone()
//	fmt.Printf("Zone is: %+v\n", zone)
//	fmt.Printf("Offset is: %+v\n", offset)
//
//	utcTime, err := convertTimeToUTC(actualPacificTime)
//	if err != nil{
//		println(err)
//		return
//	}
//	fmt.Printf("pacific time converted to UTC time: %+v\n", utcTime)
//
//}
//
//
//func convertTimeToUTC(location string, time Time)(string, error){
//	if location == ""{
//		return "", errors.New("location not provided")
//	}
//
//
//	return "", nil
//}
