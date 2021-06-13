

```golang

type Stack struct {
    items []int
}

func NewStack() *Stack{
    return &Stack{
        items: []int{}
    }
}

func (s *Stack) Push(value int) {
    s.items = append(s.items, value)
}

func (s *Stack)Pop() int {
    if len(s.items) == 0 {
        return -1
    }
    item := s.items[len(s.items) - 1]
    s.items = s.items[0: len(s.items) - 1]
    return item
}

```