// PRELIMINARIES:

package Set

import (
	"fmt"
	"strconv"
)

// *****************************************************************************
// STRUCTS AND INTERFACES:

type setMember interface{}

// Lets the members of a set be, in effect, of arbitrary type.

type Set struct {
	mem map[setMember]bool
}

/* The members of a set are a hashmap seeing if a set contains a
certain element. */

/* I do not allow multiplicites. This is a bit of a philosophical standpoint,
but I think the elements of a set are a kind of representation of an object
rather than objects in and of themselves. E.g. in the set {1, 1, 2, 3} there is
no discernible difference between the two ones, and so there is no real
difference between {1, 1, 2, 3} and {1, 2, 3}, since the additional one does not
bring any new information. This standpoint is consistent with the usual
mathematical definition of being a member of a set.*/

// *****************************************************************************
// METHODS AND FUNCTIONS:

// Constructor of the set.
func NewSet() Set {

	var emptySet Set
	emptySet.mem = make(map[setMember]bool)

	return emptySet
}

// -----------------------------------------------------------------------------

// Add an element to the set.
func (s Set) Append(newMem setMember) {

	s.mem[newMem] = true

}

// -----------------------------------------------------------------------------

/* Remove an element from the set. If it is already not in the set, the method
does nothing. */
func (s Set) Remove(removeMem setMember) {

	delete(s.mem, removeMem)

}

// -----------------------------------------------------------------------------

// Set equality: Two sets are equal if they have the same elements.
func Equals(oneSet, otherSet Set) bool {

	if len(oneSet.mem) != len(otherSet.mem) {
		return false
	}
	/* Obviously, if the sets have different numbers of elements, they are not
	equal. */

	for k := range oneSet.mem {
		if otherSet.mem[k] == false {
			return false
		}
	}
	/* If the sets have an equal number of elements and an element in one of the
	sets is not in the other set, they are not equal. */

	return true
}

// -----------------------------------------------------------------------------

/* The union of two sets is the set of members of either (or both) sets. This
function is a method which acts on a set and returns the union of it and
another set specified in the argument of the function. */
func (s Set) Union(otherSet Set) Set {

	unionSet := NewSet()

	for k := range s.mem {
		unionSet.mem[k] = true
	}
	// Adds the elements of the set to the union set.

	for k := range otherSet.mem {
		unionSet.mem[k] = true
	}
	// Adds the elements of the set in the argument to the union set.

	return unionSet
}

// -----------------------------------------------------------------------------

/* The intersection of two sets is the set of members of both sets. This function
is a method which acts on a set and returns the intersection of it and another
set specified in the argument of the function. */
func (s Set) Intersection(otherSet Set) Set {

	interSet := NewSet()

	for k := range s.mem {
		if otherSet.mem[k] == true {
			interSet.mem[k] = true
			// Adds an element to the intersection set if it is a member of both sets.
		}
	}

	return interSet
}

// -----------------------------------------------------------------------------

/* The relative complement of a set relative to another set (call it universe)
is the set of members of set which are not members of universe. This method
returns the relative complement of the set it acts upon relative to the set
specified in the argument.*/
func (s Set) RelCompl(universe Set) Set {

	relSet := NewSet()

	for k := range s.mem {
		if universe.mem[k] == false {
			relSet.mem[k] = true
			// Adds a member to the relative set if it is in s but not in universe.
		}
	}
	return relSet
}

// -----------------------------------------------------------------------------

// Return a slice containing all subsets of mySet.
func PowerSet(mySet Set) []Set {

	powSet := []Set{}
	mySlice := []setMember{}
	mySubSlices := [][]setMember{}

	/* The function first breaks the set down into a slice which is easier to work
	with. The problem then amounts to finding all subslices of this slice. This
	can be done by realizing that, since the power set of a set with N elements
	has 2^N elements, these elements can be represented by binary strings with N
	digits. This is what the subset() function below does. */

	for k := range mySet.mem {
		mySlice = append(mySlice, k)
	}

	mySubSlices = subsets(mySlice)
	// Generate the slice containing all subslices of this slice.

	for i := 0; i < len(mySubSlices); i++ {
		newSubSet := NewSet()
		for k := 0; k < len(mySubSlices[i]); k++ {
			newSubSet.mem[mySubSlices[i][k]] = true
		}
		powSet = append(powSet, newSubSet)
	}

	return powSet
}

/* This function takes a slice as input and returns a slice containing all
subslices of this slice. */
func subsets(mySlice []setMember) [][]setMember {

	indexSlice := []string{}
	subSetSlice := [][]setMember{}
	binaryLen := "%0" + strconv.Itoa(len(mySlice)) + "b"

	for i := 0; i < pow(2, len(mySlice)); i++ {
		indexSlice = append(indexSlice, fmt.Sprintf(binaryLen, i))
	}
	/* Generate all binary numbers from 0 to 2^len(mySlice)-1 and set their number
	   of digits to len(mySlice). */

	for i := 0; i < pow(2, len(mySlice)); i++ {

		subsets := []setMember{}

		for k := 0; k < len(mySlice); k++ {

			if indexSlice[i][k] == byte('1') {
				subsets = append(subsets, mySlice[k])
			}
			/* Every binary number from 0 to 2^(len(mySLice))-1 represents a
				   subset of mySlice: iterate through every binary number, and if the k:th
				   element of the binary number is 1, append the k:th element of mySlice to
				   subsets, which represents a given subset. When all the binary numbers are
				   iterated through, every possible subset of mySlice will have been
				   generated through this process, since there are 2^len(mySlice) unique
				   binary numbers and 2^len(mySlice) subsets of mySlice, counting the empty
			     slice.  */
		}

		subSetSlice = append(subSetSlice, subsets)
		// Append the generated subset to the subSetSlice.
	}

	return subSetSlice
}

/* There is no native integer exponential method in Go. I used this basic
implementation I found on:
http://grokbase.com/t/gg/golang-nuts/139n3edtq3/go-nuts-integer-exponentiation-in-go
*/
func pow(a, exp int) int {
	p := 1
	for exp > 0 {
		if exp&1 != 0 {
			p *= a
		}
		exp >>= 1
		a *= a
	}
	return p
}

// Testing some conditions put on the power set.
func testPowerSet(subSetSlice []Set, theSet Set) {

	allElemsInPowSet := NewSet()

	if len(subSetSlice) != pow(2, len(theSet.mem)) {
		panic("Error in number of elements in PowerSet")
	}
	// Check if the length of the set relative to the power set align.

	for i := 0; i < len(subSetSlice); i++ {
		for k := range subSetSlice[i].mem {
			allElemsInPowSet.Append(k)
			// Collect all elements found in the power set.
			if theSet.mem[k] == false {
				panic("Error in allowed set members in PowerSet")
				// Checks if some elements of the power set are not in the set.
			}
		}
	}

	for i := range allElemsInPowSet.mem {
		if theSet.mem[i] == false {
			panic("Error: elements missing in PowerSet")
		}
	}
	// Check if some elements of the set are not in the power set.
}

// -----------------------------------------------------------------------------

/* Return a string representation of the set. */

func (s *Set) SetString() string {

	mySetSlice := []setMember{}

	for i := range s.mem {
		mySetSlice = append(mySetSlice, i)
	}

	myString := fmt.Sprint(mySetSlice)

	return myString
}

func testSetString(theSetString string, theSet Set) {
	counter := 0
	for i := range theSet.mem {
		for k := 0; k < len(theSetString); k++ {

			if fmt.Sprint(i) == string(theSetString[k]) {
				counter += 1
			}
		}
	}
	if counter != len(theSet.mem) {
		panic("Error in SetString: wrong number of elements")
	}
	// Seeing if every element of theSet can be found as a string in theSetString.
}

// *****************************************************************************
// TESTING:

func main() {

	theSet := NewSet()
	theOtherSet := NewSet()
	emptySet := NewSet()

	// ---------------------------------------------------------------------------
	// TESTING OF EMPTY SET AND ONE-ELEMENT SET OPERATIONS:

	if Equals(theSet, theOtherSet) != true {
		panic("Error handling empty set in Equals")
	}

	testPowerSet(PowerSet(emptySet), emptySet)

	theSet.Append(1)

	testPowerSet(PowerSet(theSet), theSet)

	if Equals(theSet.Intersection(emptySet), emptySet) != true {
		panic("Error handling intersection of set with emptySet")
	}

	if Equals(theSet.Union(emptySet), theSet) != true {
		panic("Error handling union of set with emptySet")
	}

	// Testing union and intersection on emptySet.

	if Equals(theSet, theOtherSet) != false {
		panic("Error handling comparing empty to one-element set in Equals")
	}

	if Equals(theSet, theSet) != true {
		panic("Error handling comparing set to itself in Equals")
	}

	theOtherSet.Append(1)

	if Equals(theSet, theOtherSet) != true {
		panic("Error handling comparing one-element sets in Equals")
	}

	// Testing Equals with one and zero-element sets.

	theSet.Append(1)

	if Equals(theSet, theOtherSet) != true {
		panic("Error in handling same-element appending in Append")
	}

	testPowerSet(PowerSet(theSet), theSet)

	if len(PowerSet(emptySet)) != 1 {
		panic("Error in handling power set of empty set")
	}

	if Equals(theSet.RelCompl(emptySet), theSet) != true {
		panic("Error in handling relative complement of set and empty set in RelCompl")
	}

	if Equals(theSet.RelCompl(theOtherSet), emptySet) != true {
		panic("Error in handling relative complement of one-element sets in RelCompl")
	}

	// Testing RelCompl with one- and zero-element sets.

	if Equals(emptySet.Intersection(emptySet), emptySet) != true {
		panic("Error in handling intersection of empty set with itself in Intersection")
	}

	if Equals(emptySet.Union(emptySet), emptySet) != true {
		panic("Error in handling union of non-empty with itself in Union")
	}

	// Testing of methods on emptySet.

	if Equals(theSet.Intersection(theSet), theSet) != true {
		panic("Error in handling intersection of non-empty set with itself in Intersection")
	}

	if Equals(theSet.Union(theSet), theSet) != true {
		panic("Error in handling union of non-empty set with itself in Union")
	}

	// Testing of methods on non-empty set.

	theOtherSet.Remove(1)

	testPowerSet(PowerSet(theSet), theSet)

	if Equals(theOtherSet, emptySet) != true {
		panic("Error in handling removal of element in one-element set in Remove")
	}

	theOtherSet.Remove("k")

	if Equals(theOtherSet, emptySet) != true {
		panic("Error in handling removal of non-existent element in Remove")
	}

	// Testing Remove.

	// ---------------------------------------------------------------------------
	// TESTING MULTI-ELEMENT SET OPERATIONS:

	theSet.Append(2)
	theSet.Append(3)
	theSet.Append("a")
	theSet.Append("b")

	theOtherSet.Append(1)
	theOtherSet.Append(2)
	theOtherSet.Append(3)
	theOtherSet.Append("a")
	theOtherSet.Append("b")

	// Adding some elements to the sets.

	testPowerSet(PowerSet(theSet), theSet)

	if Equals(theSet, theOtherSet) != true {
		panic("Error in handling multi-element set in Equals")
	}

	theOtherSet.Append("c")

	if Equals(theSet, theOtherSet) != false {
		panic("Error in handling multi-element set in Equals")
	}
	// Testing Equals.

	for i := 0; i < 12; i++ {
		theOtherSet.Append(i)
	}

	testPowerSet(PowerSet(theOtherSet), theOtherSet)
	// Testing PowerSet.

	if Equals(theSet.Intersection(theOtherSet), theSet) != true {
		panic("Error in handling multi-element set intersection")
	}

	if Equals(theSet.Union(theOtherSet), theOtherSet) != true {
		panic("Error in handling multi-element set union")
	}
	/* Testing union and intersection. Since in this state, all elements of
	theSet are members of theOtherSet, these comparisons are valid. */

	if Equals(theSet.RelCompl(theOtherSet), emptySet) != true {
		panic("Error in handling multi-element set relative complement")
	}
	// Case all elements of theSet are elments in theOtherSet.

	oneElementSet := NewSet()

	oneElementSet.Append("d")

	theSet.Append("d")

	if Equals(theSet.RelCompl(theOtherSet), oneElementSet) != true {
		panic("Error in handling multi-element set relative complement")
	}
	// Case not all elements of theSet are elments in theOtherSet.

	testSetString(emptySet.SetString(), emptySet)
	testSetString(theSet.SetString(), theSet)
	// Testing the SetString method.

}
