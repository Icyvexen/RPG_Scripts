package names_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/Icyvexen/rpgscripts/names"
)

func TestNewBook(t *testing.T) {
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
		{"Not Seeded - cov pages", notSeeded, "A book bound by"},
		{"Not Seeded - cov pov", notSeeded, " binding, as told from a(n) "},
		{"Not Seeded - cov topic", notSeeded, "A book is here with"},
		{"Not Seeded - pages pov", notSeeded, " pages, from a(n) "},
		{"Not Seeded - pages topic", notSeeded, " pages about "},
		{"Not Seeded - pov topic", notSeeded, " perspective for "},
		{"Not Seeded - cov pages pov", notSeeded, "Here is a book in"},
		{"Not Seeded - cov pages topic", notSeeded, " binding and set on "},
		{"Not Seeded - cov pov topic", notSeeded, "This book has "},
		{"Not Seeded - pages pov topic", notSeeded, " perspective, detailing thoughts on "},
		{"Not Seeded - all", notSeeded, " perspective on the subjects of "},
		{"seeded 1", seeded, "Here is a book inCloth binding, pages made of Display Screen, from a(n) Human perspective."},
		{"seeded 2", seeded, "Here is a book inCloth binding, pages made of Display Screen, from a(n) Human perspective."},
		{"seeded 3", seeded, "A book with Etched Glass pages about Drama."},
		{"seeded 4", seeded, "A book bound by Cloth, with Display Screen pages."},
		{"seeded 5", seeded, "A book with Etched Glass pages from a(n) Western perspective, detailing thoughts on Naturalism."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := names.NewBook(tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				if strings.Contains(tt.name, "Not Seeded") {
					for i := 10000; i > 0; i-- {
						if !strings.Contains(got, tt.want) {
							got = names.NewBook()
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
