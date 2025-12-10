# heaputil

Tiny helper library around Go's `container/heap` that gives you a simple, typed `int` heap which can behave as either a **min-heap** or a **max-heap**.

No more copy-pasting the same `IntHeap` boilerplate into every LeetCode solution, coding exercise, or production service.

## Features

- `int`-only heap, focused and fast
- Min-heap and max-heap via separate constructors
- Fully compatible with `container/heap`
- Pointer receivers only (no mixed receiver semantics)
- Tiny API, easy to remember

## Quick Start

### Min-heap

```go
package main

import (
	"fmt"

	"github.com/baggito/heaputil"
)

func main() {
	// Create a min-heap with initial values
	h := heaputil.NewMinIntHeap(4, 5, 8, 2)

	fmt.Println("len:", h.Len())  // 4
	fmt.Println("min:", h.Peek()) // 2

	h.PushInt(1)
	fmt.Println("min after push:", h.Peek()) // 1

	for h.Len() > 0 {
		fmt.Println("pop:", h.PopInt())
	}
}
```

**Example output:**

```
len: 4
min: 2
min after push: 1
pop: 1
pop: 2
pop: 4
pop: 5
pop: 8
```

### Max-heap

```go
package main

import (
	"fmt"

	"github.com/baggito/heaputil"
)

func main() {
	// Create a max-heap with initial values
	h := heaputil.NewMaxIntHeap(4, 5, 8, 2)

	fmt.Println("len:", h.Len())  // 4
	fmt.Println("max:", h.Peek()) // 8

	h.PushInt(10)
	fmt.Println("max after push:", h.Peek()) // 10

	for h.Len() > 0 {
		fmt.Println("pop:", h.PopInt())
	}
}
```

### Example: K-th Largest Element in a Stream

Classic use case: maintain the k-th largest value in a stream of test scores.

```go
package kthlargest

import "github.com/baggito/heaputil"

// KthLargest keeps track of the k-th largest value using a min-heap of size k.
type KthLargest struct {
	k int
	h *heaputil.IntHeap // min-heap
}

func NewKthLargest(k int, nums []int) *KthLargest {
	h := heaputil.NewMinIntHeap()
	kl := &KthLargest{
		k: k,
		h: h,
	}
	for _, n := range nums {
		kl.Add(n)
	}
	return kl
}

func (kl *KthLargest) Add(val int) int {
	if kl.h.Len() < kl.k {
		kl.h.PushInt(val)
	} else if val > kl.h.Peek() {
		kl.h.PopInt()
		kl.h.PushInt(val)
	}
	return kl.h.Peek()
}
```

**Usage:**

```go
kl := NewKthLargest(3, []int{4, 5, 8, 2})

fmt.Println(kl.Add(3))  // 4
fmt.Println(kl.Add(5))  // 5
fmt.Println(kl.Add(10)) // 5
fmt.Println(kl.Add(9))  // 8
fmt.Println(kl.Add(4))  // 8
```

## API

### Constructors

```go
// NewMinIntHeap constructs a min-heap with optional initial values.
func NewMinIntHeap(nums ...int) *IntHeap

// NewMaxIntHeap constructs a max-heap with optional initial values.
func NewMaxIntHeap(nums ...int) *IntHeap
```

Both:

- Copy the provided values into the internal slice
- Call `heap.Init` internally
- Return a ready-to-use `*IntHeap`

### Core Methods

```go
type IntHeap struct {
    // internal fields are not exported
}

func (h *IntHeap) Len() int
func (h *IntHeap) PushInt(x int)
func (h *IntHeap) PopInt() int
func (h *IntHeap) Peek() int
func (h *IntHeap) Data() []int
```

| Method | Description |
|--------|-------------|
| `Len()` | Number of elements in the heap |
| `PushInt(x)` | Push an element, keeping heap invariant |
| `PopInt()` | Pop and return the root (min or max) |
| `Peek()` | Look at the root without removing it (caller must ensure `Len() > 0`) |
| `Data()` | Underlying slice (read-only use recommended) |

### Interop with `container/heap`

`IntHeap` implements `heap.Interface`, so you can still use the standard functions:

```go
import (
	"container/heap"
	"github.com/baggito/heaputil"
)

h := heaputil.NewMinIntHeap(3, 1, 2)
heap.Push(h, 0)          // heap.Push takes interface{}, still works
x := heap.Pop(h).(int)   // 0
```

## Design Notes

- Single type `IntHeap` for both min-heap and max-heap behavior
- Behavior is driven by an internal comparator:
  - min-heap: `a < b`
  - max-heap: `a > b`
- All methods use pointer receivers (`*IntHeap`) to avoid the usual value/pointer confusion
- Focused on `int` for simplicity and minimal API surface

## Installation

```bash
go get github.com/baggito/heaputil
```

Then:

```go
import "github.com/baggito/heaputil"
```