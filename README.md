# Collections (Golang)

```go
import "github.com/tasnimzotder/collections"
```

## Collections

- [Stack](stack.go)
- [Queue](queue.go)
- [LinkedList](linkedlist.go)

### [Stack](stack.go)

#### Operations

| Operation | Description |
| --------- | ----------- |
| `Push` | Push element to the stack |
| `Pop` | Pop element from the stack |
| `Peek` | Peek element from the stack |
| `IsEmpty` | Check if the stack is empty |
| `Clear` | Clear the stack |

#### Example

```go
package main


import (
    "fmt"
    "github.com/tasnimzotder/collections"
)


func main() {
    var stack collections.Stack

    stack.Push(1)
    stack.Push(2)
    stack.Push(3)

    fmt.Println(stack.Pop())
    fmt.Println(stack.Pop())
}
```

### [Queue](queue.go)

#### Operations

| Operation | Description |
| --------- | ----------- |
| `Push` | Enqueue element to the queue |
| `Pop` | Dequeue element from the queue |
| `Peek` | Peek element from the queue |
| `IsEmpty` | Check if the queue is empty |
| `Clear` | Clear the queue |

#### Example

```go
package main


import (
    "fmt"
    "github.com/tasnimzotder/collections"
)


func main() {
    var queue collections.Queue

    queue.Push(1)
    queue.Push(2)
    queue.Push(3)

    fmt.Println(queue.Pop())
    fmt.Println(queue.Pop())
}
```

### [LinkedList](linkedlist.go)

#### Operations

| Operation | Description |
| --------- | ----------- |
| `InsertAt` | Insert element at the given index |
| `RemoveFrom` | Remove element at the given index |
| `Get` | Get element at the given index |
| `Clear` | Clear the linked list |

#### Example

```go
package main


import (
    "fmt"
    "github.com/tasnimzotder/collections"
)


func main() {
    var linkedList collections.LinkedList

    linkedList.InsertAt(0, 1)
    linkedList.InsertAt(1, 2)
    linkedList.InsertAt(2, 3)

    fmt.Println(linkedList.Get(0))
    fmt.Println(linkedList.Get(1))
    fmt.Println(linkedList.Get(2))
}
```

## License

This project is licensed under the [MIT License](LICENSE).
