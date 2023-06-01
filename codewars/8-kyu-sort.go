/*
You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.

The returned value must be a string, and have "***" between each of its letters.

You should not remove or add elements from/to the array.

*/

package kata

import (
    "sort"
    "strings"
)

func TwoSort(arr []string) string {
    sort.Strings(arr)
    return strings.Join(strings.Split(arr[0], ""), "***")
}


/*
======================================================
# TEST CASES
======================================================
*/

// Ginkgo BDD Testing Framework <http://onsi.github.io/ginkgo/>
// Gomega Matcher Library <http://onsi.github.io/gomega/>

package kata_test
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  . "codewarrior/kata"
)

var _ = Describe("Tests", func() {
  It("should work for sample tests", func() {
    var s []string

    s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
    Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
    s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
    Expect(TwoSort(s)).To(Equal("a***r***e"))
    s = []string{"lets", "talk", "about", "javascript", "the", "best", "language"}
    Expect(TwoSort(s)).To(Equal("a***b***o***u***t"))
    s = []string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}
    Expect(TwoSort(s)).To(Equal("c***o***d***e"))
    s = []string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}
    Expect(TwoSort(s)).To(Equal("L***e***t***s"))
  })
})
