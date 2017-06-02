package typesort

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//////////////////////////////////////////////
/// aliases
//////////////////////////////////////////////

var println = fmt.Println
var sprintf = fmt.Sprintf
var printf = fmt.Printf

//////////////////////////////////////////////
/// Counter
//////////////////////////////////////////////

type Counter int

func (c Counter) String() string {
	return sprintf("%v times", int(c))
}

func (c *Counter) Up() {
	*c += 1
}

func (c *Counter) Down() {
	*c -= 1
}

//////////////////////////////////////////////
/// sortStr
//////////////////////////////////////////////

type sortStr []rune

func (s sortStr) Len() int {
	return len(s)
}

func (s sortStr) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortStr) String() string {
	return string(s)
}

//////////////////////////////////////////////
/// sortLines
//////////////////////////////////////////////

/*
This is already implemented (probably better) by sort.StringSlice
*/

type sortLines []string

func (l sortLines) Len() int {
	return len(l)
}

func (l sortLines) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l sortLines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l sortLines) String() string {
	return strings.Join(l, "\n")
}

//////////////////////////////////////////////
/// sortInt
//////////////////////////////////////////////

type sortInt int64

func fromBinary(binaryStr string) sortInt {
	i, err := strconv.ParseInt(binaryStr, 2, 64)
	if err == nil {
		return sortInt(i)
	}
	panic(err)
}

func (i sortInt) toBinary() string {
	return strconv.FormatInt(int64(i), 2)
}

func (i *sortInt) sort() {
	bStr := i.toBinary()
	s := sortStr(bStr)
	sort.Sort(s)
	*i = fromBinary(string(s))
}
