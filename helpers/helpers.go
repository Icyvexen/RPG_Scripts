//Package helpers - helper functions for use inside this module
package helpers

import (
	"fmt"
	"strings"
	"testing"
)

//TCStrings - stores Name (string), Ar(guments, []int64), and Want (string)
type TCStrings struct {
	Name string
	Ar   []int64
	Want string
}

//StringCombiner - take every incoming string, return 1 string
func StringCombiner(words ...string) string {
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
	}
	return sb.String()
}

/*LoopForTestStrings - takes the testing.T, the func to test, and the list of tests;
for each test run 10k iterations if not seeded, else just 1.
*/
func LoopForTestStrings(t *testing.T, tf func(...int64) string, tests []TCStrings) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if strings.Contains(tt.Name, "Not Seeded") {
				i := 10000
				for i > 0 {
					got := tf(tt.Ar...)
					if strings.Contains(got, tt.Want) {
						fmt.Println("SUCCESS +++ ", tt.Name)
						break
					}
					i--
				}
				if i == 0 {
					t.Errorf("%v --- Could not generate \"%v\" after 10000 iterations. Maybe the test is broken?", tt.Name, tt.Want)
				}
			} else {
				got := tf(tt.Ar...)
				if !strings.Contains(got, tt.Want) {
					t.Errorf("%v --- Generated \"%v\", want \"%v\" .", tt.Name, got, tt.Want)
				}
			}
		})
	}
}
