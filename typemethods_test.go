package typemethods

import (
	"sort"
	"testing"
)

//////////////////////////////////////////////
/// Counter
//////////////////////////////////////////////

func TestCounter(t *testing.T) {
	c := Counter(0)
	println(c)
	c.Up()
	println(c)
	c.Down()
	println(c)
}

//////////////////////////////////////////////
/// SortStr
//////////////////////////////////////////////

func TestSortStr(t *testing.T) {
	s := sortStr("This is a string to sort")
	println(s)
	sort.Sort(s)
	println(s)
	println()

	// s1 := "a"
	// s1 += "bc"
	// println(s1)
	// s += "bc"
	// ^error
}

//////////////////////////////////////////////
/// SortLines
//////////////////////////////////////////////

func TestSortLines(t *testing.T) {
	l := sortLines{
		"These",
		"are",
		"some",
		"lines",
		"to",
		"sort",
	}
	println(l)
	println()
	sort.Sort(l)
	println(l)
	println()
}

//////////////////////////////////////////////
/// BinaryConversion
//////////////////////////////////////////////

func TestBinaryConversion(t *testing.T) {
	println(fromBinary("1001"))
	println(fromBinary("1111011"))
	println(sortInt(9).toBinary())
	println(sortInt(123).toBinary())
	println()
}

//////////////////////////////////////////////
/// SortInt
//////////////////////////////////////////////

func sortIt(i int64) {
	si := sortInt(i)
	println(si)
	println(si.toBinary())
	si.sort()
	println(si.toBinary())
	println(si)
	println()
}

func TestSortInt(t *testing.T) {
	intSlice := []int64{7, 13, 14, 8, 9, 123}
	for _, i := range intSlice {
		sortIt(i)
	}
}
