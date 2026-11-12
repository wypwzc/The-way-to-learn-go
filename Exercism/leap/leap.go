package leap

func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

/*逻辑最强
package leap

func IsLeapYear(year int) bool {
 	if year % 400 == 0 {
        return true
    }
    if year % 100 == 0 {
        return false
    }
    if year % 4 == 0 {
        return true
    }
    return false
*/
