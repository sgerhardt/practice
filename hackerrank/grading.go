package hackerrank

// HackerLand University has the following grading policy:
//
// Every student receives a grade in the inclusive range from 0 to 100.
// Any grade less than 40 is a failing grade.
// Sam is a professor at the university and likes to round each student's  according to these rules:
//
// If the difference between the grade and the next multiple of 5 is less than 3, round  up to the next multiple of 5.
// If the value of grade is less than 38 grade, no rounding occurs as the result will still be a failing grade.
// Examples
//
// 84 round to 85  (85 - 84 is less than 3)
// 29 do not round (result is less than 40)
// 57 do not round (60 - 57 is 3 or higher)
// Given the initial value of  for each of Sam's  students, write code to automate the rounding process.
//
// # Function Description
//
// Complete the function gradingStudents in the editor below.
//
// gradingStudents has the following parameter(s):
//
// int grades[n]: the grades before rounding
// Returns
//
// int[n]: the grades after rounding as appropriate
// Input Format
//
// The first line contains a single integer, , the number of students.
// Each line  of the  subsequent lines contains a single integer, .
//
// # Constraints
//
// # Sample Input 0
//
// 4
// 73
// 67
// 38
// 33
// Sample Output 0
//
// 75
// 67
// 40
// 33
func gradingStudents(grades []int32) []int32 {
	graded := make([]int32, len(grades))
	for idx, grade := range grades {
		if grade < 38 {
			// do not round
			graded[idx] = grade
			continue
		}
		remainder := grade % 5
		// find next multiple of 5
		addVal := 5 - remainder
		nextMultiple := grade + addVal

		if nextMultiple-grade < 3 {
			// round to next multiple of 5
			graded[idx] = nextMultiple
		} else {
			// do not round
			graded[idx] = grade
		}
	}
	return graded
}
