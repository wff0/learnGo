package testDemo

import (
	"fmt"
	"sort"
	"testing"
)

type person struct {
	Name string
	Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func TestSort(t *testing.T) {
	a := personSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	sort.Sort(a)
	fmt.Println("Sort:", a)

	sort.Stable(a)
	fmt.Println("Stable:", a)

}

type Peak struct {
	Name      string
	Elevation int // in feet
}

func TestSortSlice(t *testing.T) {
	peaks := []Peak{
		{"Aconcagua", 22838},
		{"Denali", 20322},
		{"Kilimanjaro", 19341},
		{"Mount Elbrus", 18510},
		{"Mount Everest", 29029},
		{"Mount Kosciuszko", 7310},
		{"Mount Vinson", 16050},
		{"Puncak Jaya", 16024},
	}

	// does an in-place sort on the peaks slice, with tallest peak first
	sort.Slice(peaks, func(i, j int) bool {
		return peaks[i].Elevation >= peaks[j].Elevation
	})
	fmt.Println(peaks)
}
