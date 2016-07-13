package testutil

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

type ConveyRequirement struct {
	Name      string
	Predicate func() bool
}

/*
	Require that the tests are not running with the "short" flag enabled.
*/
var RequiresLongRun = ConveyRequirement{"run long tests", func() bool { return !testing.Short() }}

func RequiresTestLabel(label string) ConveyRequirement {
	return ConveyRequirement{
		Name: fmt.Sprintf("have test label %q", label),
		Predicate: func() bool {
			labels := strings.Split(os.Getenv("DANGGIT_TEST"), ",")
			for _, l := range labels {
				if l == "ALL" || l == label {
					return true
				}
			}
			return false
		},
	}
}

/*
	Decorates a GoConvey test to check a set of `ConveyRequirement`s,
	returning a dummy test func that skips (with an explanation!) if any
	of the requirements are unsatisfied; if all is well, it yields
	the real test function unchanged.  Provide the `...ConveyRequirement`s
	first, followed by the `func()` (like the argument order in `Convey`).
*/
func Requires(items ...interface{}) func(c convey.C) {
	// parse args
	// not the most robust parsing.  just panics if there's weird stuff
	var requirements []ConveyRequirement
	for _, it := range items {
		if req, ok := it.(ConveyRequirement); ok {
			requirements = append(requirements, req)
		} else {
			break
		}
	}
	action := items[len(items)-1]
	// examine requirements
	var widest int
	for _, req := range requirements {
		if len(req.Name) > widest {
			widest = len(req.Name)
		}
	}
	// check requirements
	var requirementsListing bytes.Buffer
	var names []string
	allSat := true
	for _, req := range requirements {
		sat := req.Predicate()
		allSat = allSat && sat
		names = append(names, req.Name)
		fmt.Fprintf(&requirementsListing, "requirement %*q: %v\n", widest+2, req.Name, sat)
	}
	// act
	if allSat {
		return func(c convey.C) {
			// attempted: inserting another convey that makes a single 'true=true' assertion so we see the prereqs and a green check mark.
			// doesn't work: doing so causes a leaf node, in which everything is run :/ even if skipped, the remaining `So` that aren't
			// in another block get attached to it, which makes verrry odd reading, and causes an extra repetition of anything
			// that isn't in another convey block.
			//	convey.SkipConvey(title, func() { convey.So(true, convey.ShouldBeTrue) })
			switch action := action.(type) {
			case func():
				action()
			case func(c convey.C):
				action(c)
			}
		}
	} else {
		title := "Prereqs: " + strings.Join(names, ", ")
		return func(c convey.C) {
			convey.Convey(title, nil)
			c.Println()
			c.Print(requirementsListing.String())
		}
	}
}
