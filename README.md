Package Set provides a naïve (finite) set-like object for use with comparable types in Go, along with some common set operations supplied as methods.

A naïve set is a "well defined" collection of objects, in which the ordering of the objects is immaterial. Some common set operations and concepts include union, intersection, complement, power set and equality. These are supplied as methods and functions to this package.



---
Defines the Set object. If setMember is in Set, its value is ```true```, otherwise ```false```. Multiples are not allowed. A short discussion of why can be found in the source code.
```go
type setMember interface{}

type Set struct:
  mem: map[setMember]bool
```
---
Create a new, empty Set.
```go
NewSet() Set
```
---
Append element to s.
```go
s.SetAppend(element)
```
---
Remove element from s. If element is not in s, do nothing.
```go
s.SetRemove(element)
```
---
Return ```true``` if theSet and theOtherSet contain the same elements, else returns ```false```.

```go
SetEquals(theSet, theOtherSet)
```
---

s -> union of s and otherSet.
```go
s.SetUnion(otherSet) Set
```
---

s -> intersection of s and otherSet.
```go
s.SetIntersection(otherSet)
```
---
s -> the set containing the elements of s that are not in universe.

```go
s.SetRelCompl(universe)
```
---
Return a slice containing all subsets of theSet, including the empty set and Set itself.

```go
powSet(theSet) []Set
```
---

Return the string representation of s.
```go
s.SetString(Set) string
```

## License:

Copyright 2018 Jonas Conneryd

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
