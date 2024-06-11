package swaps

import (
	"fmt"
	"reflect"
	"strings"
)

// Objects defines the 3D objects and their 2D shape components
var Objects = map[string]map[string]int{
	"cylinder": {
		"circle": 1,
		"square": 1,
	},
	"cone": {
		"circle":   1,
		"triangle": 1,
	},
	"prism": {
		"triangle": 1,
		"square":   1,
	},
	"sphere": {
		"circle": 2,
	},
	"cube": {
		"square": 2,
	},
	"pyramid": {
		"triangle": 2,
	},
}

type object struct {
	shapes map[string]int
}

func (in *object) deepCopy() *object {
	out := &object{
		shapes: map[string]int{},
	}
	for k, v := range in.shapes {
		out.shapes[k] = v
	}
	return out
}

func deepCopySet(o []object) []object {
	rtn := make([]object, len(o), len(o))
	for k, v := range o {
		rtn[k] = *v.deepCopy()
	}
	return rtn
}

func (a *object) equal(b object) bool {
	for k, v := range a.shapes {
		if v > 0 && a.shapes[k] != b.shapes[k] {
			return false
		}
	}
	for k, v := range b.shapes {
		if v > 0 && a.shapes[k] != b.shapes[k] {
			return false
		}
	}
	return true
}

func deepEqualSet(a, b []object) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].equal(b[i]) {
			return false
		}
	}
	return true
}

type Swap struct {
	componentShape1 string
	posShape1       int

	componentShape2 string
	posShape2       int
}

func (s *Swap) Export() string {
	pos1 := "left"
	if s.posShape1 == 1 {
		pos1 = "middle"
	} else if s.posShape1 == 2 {
		pos1 = "right"
	}

	pos2 := "left"
	if s.posShape2 == 1 {
		pos2 = "middle"
	} else if s.posShape2 == 2 {
		pos2 = "right"
	}

	// Dissecting twice swaps the two composites
	return fmt.Sprintf("Dissect %s on %s and dissect %s on %s\n", s.componentShape1, pos1, s.componentShape2, pos2)
}

// swapShapes repressively goes through all possible solutions.
// If the current and target 3D objects are identical, no swap is needed.
// If the max swaps is reached it returns with no solution found.
// If the current and target 3D objects are not identical, loop through every 3D object combination, and every 2D
// component combination of those objects, swapping them. If it results in matching 3D objects then the solution is found.
// If a solution is found, the function returns with the swap needed to get the matching shapes.
// If the objects and components do not match, recurse with the swapped shapes and the swap count incremented.
// If the recursion returns a solution for the current swap, check if it's lower than the currently found solution or
// is the first solution found. If so, set the swap solution to the current swap plus the returned one, thus
// Adding on the current swap step to the list of swaps.
// Once all combinations are looped through, return the shortest solution or an empty Swap list if a solution is not found
func swapShapes(current, target []object, currSwapCount, maxSwaps int) []Swap {
	if reflect.DeepEqual(current, target) {
		return []Swap{}
	}

	if currSwapCount > maxSwaps {
		return []Swap{}
	}

	var swaps []Swap

	// Loop through every shape combination, and each component combination of that shape
	// If the shapes are the same skip
	for i := 0; i < len(current); i++ {
		for j := i + 1; j < len(current); j++ {
			for k, k2 := range current[i].shapes {
				if k2 <= 0 {
					continue
				}
				for l, l2 := range current[j].shapes {
					if l2 <= 0 {
						continue
					}
					if k == l {
						continue
					}
					newCurrent := deepCopySet(current)
					newCurrent[i].shapes[k]--
					newCurrent[i].shapes[l]++

					newCurrent[j].shapes[k]++
					newCurrent[j].shapes[l]--

					// safe to delete while in a range
					if newCurrent[i].shapes[k] <= 0 {
						delete(newCurrent[i].shapes, k)
					}
					if newCurrent[j].shapes[l] <= 0 {
						delete(newCurrent[j].shapes, l)
					}
					currSwap := Swap{
						componentShape1: k,
						posShape1:       i,

						componentShape2: l,
						posShape2:       j,
					}
					if deepEqualSet(newCurrent, target) {
						return []Swap{currSwap}
					}
					newSwaps := swapShapes(newCurrent, target, currSwapCount+1, maxSwaps)
					if (newSwaps != nil && len(newSwaps) > 0) && (swaps == nil || len(newSwaps) < len(swaps)) {
						swaps = []Swap{currSwap}
						swaps = append(swaps, newSwaps...)
					}
				}
			}
		}
	}

	return swaps
}

func newObject(o string) object {
	rtn := object{
		shapes: map[string]int{},
	}
	for k, v := range Objects[strings.ToLower(o)] {
		rtn.shapes[k] = v
	}
	return rtn
}

func newObjects(o []string) []object {
	var rtn []object
	for _, i := range o {
		rtn = append(rtn, newObject(i))
	}

	return rtn
}

func GetShape(input string) string {
	input = strings.ToLower(input)
	if input == "t" {
		return "cylinder"
	}
	if input == "s" {
		return "cone"
	}
	if input == "c" {
		return "prism"
	}
	return ""
}

func GetSwaps(soloCallout string, currentShapes []string) ([]Swap, error) {
	err := IsValidInput(soloCallout, currentShapes)
	if err != nil {
		return nil, err
	}
	targetShapes := make([]string, len(soloCallout), len(soloCallout))
	for i, i2 := range soloCallout {
		targetShapes[i] = GetShape(string(i2))
	}

	currentObjects := newObjects(currentShapes)
	targetObjects := newObjects(targetShapes)

	swap := swapShapes(currentObjects, targetObjects, 0, 5)
	if len(swap) == 0 {
		return nil, fmt.Errorf("no dissections found")
	}
	return swap, nil
}

func IsValidInput(soloCallout string, currentObjects []string) error {
	// Check input lengths
	if len(soloCallout) != 3 {
		return fmt.Errorf("you must enter 3 solo callout letters")
	}
	if len(currentObjects) != 3 {
		return fmt.Errorf("you must enter three current 3D objects")
	}
	// check if inputs valid
	// check if solo callout is made up of a combination of a single T, S, and C letter
	shapeCount := map[string]int{"T": 0, "S": 0, "C": 0}
	for _, callout := range soloCallout {
		letter := strings.ToUpper(string(callout))
		if _, exists := shapeCount[letter]; !exists {
			return fmt.Errorf("the solo callout letters must be T (triangle), S (square), or C (circle)")
		}
		shapeCount[letter]++
	}
	if shapeCount["T"] != 1 || shapeCount["S"] != 1 || shapeCount["C"] != 1 {
		return fmt.Errorf("the solo callout letters must contain exactly one T, one S, and one C")
	}

	// Validate that the currentObjects exist in the Objects map
	for _, o := range currentObjects {
		if _, exists := Objects[o]; !exists {
			return fmt.Errorf("the 3D object %s is not valid", o)
		}
	}

	// Sum up the number of each shape component in the selected objects
	totalShapes := map[string]int{"circle": 0, "square": 0, "triangle": 0}
	for _, object := range currentObjects {
		for shape, count := range Objects[object] {
			totalShapes[shape] += count
		}
	}
	// Check if there are exactly 2 circles, 2 squares, and 2 triangles
	if totalShapes["circle"] != 2 || totalShapes["square"] != 2 || totalShapes["triangle"] != 2 {
		return fmt.Errorf("the selected objects must add up to exactly 2 circles, 2 squares, and 2 triangles")
	}

	return nil
}
