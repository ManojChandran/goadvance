package hashtable

import (
  "github.com/timtadh/data-structures/types"
  "github.com/timtadh/data-structures/errors"
)

type entry struct {
  key Hashable
  value interface{}
  next *entry
}

type Hash struct {
  table []*entry
  size int
}

func NewHashTable (initial_size int) *Hash {
  return &Hash{
    table: make ([]*entry, initial_size),
    size: 0,
  }
}

func ()  {

}
