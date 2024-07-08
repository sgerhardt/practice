package arrays_and_string_manipulation

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

func isMul(n int) bool {
	return (n+1)%5 == 0
}

func getMult(n int) int {
	if n/1000 > 0 {
		return 1000
	} else if n/100 > 0 {
		return 100
	} else if n/10 > 0 {
		return 10
	}

	return 1
}

func subProblem(num int) {

}

//	string subProb(int &num) {
//	    map<int,string> val;
//	    val[1] = "I";
//	    val[5] = "V";
//	    val[10] = "X";
//	    val[50] = "L";
//	    val[100] = "C";
//	    val[500] = "D";
//	    val[1000] = "M";
//	    int multiple = getMult(num);
//	    string ans = "";
//	    int value = num / multiple;
//	    if(value == 4) ans += (val[multiple] + val[5 * multiple]), value -=4;
//	    else if(value == 9) ans += (val[multiple] + val[10 * multiple]), value -=9;
//	    else if(value >= 5) {
//	        ans += val[5 * multiple];
//	        value -= 5;
//	    }
//	    for(int i = 0; i < value; i ++) ans += val[multiple];
//	    num %= multiple;
//	    return ans;
//	}
func intToRomanGeneric(num int) string {

	return ""
}
