// Leap stub file

// Package leap : The package name  expected by the test program.
package leap

const testVersion = 2

// IsLeapYear : 
// PRE: It receive a valid year 
// POST: return a boolean if leap year
func IsLeapYear(year int) bool {
	//If year mod 400 is equal to 0 , it is a leap year
    if (year % 400) == 0 {
        return true
    } 
    //If year mod 100 is equal to 0 , it is not a leap year
    if (year % 100) == 0 {
        return false
    } 
    //IF year mod 4 is equal to zero , it is a leap year
    if (year % 4) == 0 {
        return true    
    } 
    //any other case it is not a leap year
    return false
}
