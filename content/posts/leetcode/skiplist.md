

```golang

type Element struct {
    score float64
    value interface{}
    forward []*Element
}

type SkipList struct {
    header *Element
    len int
    level int
}

```