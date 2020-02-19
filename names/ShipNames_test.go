package names_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/Icyvexen/rpgscripts/names"
)

func TestNewShipName(t *testing.T) {
	type args struct {
		params []int64
	}

	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	seeded := args{params: []int64{seed}}
	notSeeded := args{}

	tests := []struct {
		name string
		args args
		want string
	}{
		//NOTE: to validate the tests more quickly, check if 1 word from each section is present
		{"Not Seeded - d s1", notSeeded, "Northern"},
		{"Not Seeded - p s1", notSeeded, "Hangman's"},
		{"Not Seeded - s1", notSeeded, "Maiden"},
		{"Not Seeded - d s2", notSeeded, "Glory"},
		{"seeded", seeded, "Obedient Sky's Gypsy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := names.NewShipName(tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				if strings.Contains(tt.name, "Not Seeded") {
					for i := 10000; i > 0; i-- {
						if !strings.Contains(got, tt.want) {
							got = names.NewShipName()
							time.Sleep(time.Microsecond) //give the RNG call time to use nanosecond seed
							fmt.Println(got)
						} else {
							fmt.Println("CONDITION FOUND FOR", tt.name, ". NEXT!")
							break
						}
					}
					if !strings.Contains(got, tt.want) {
						fmt.Println("Failed to find", tt.name, "**********************************************")
						t.Errorf("Was unable to generate desired item in 10000 iterations. This may indicate a faulty test condition.")
					}
				} else {
					t.Errorf("NewShipName() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
