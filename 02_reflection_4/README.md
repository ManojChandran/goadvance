# go type Assertion

  A type assertion doesn't really convert an interface to another data type,
  but it provides access to an interface's concrete value, which is typically
  what you want. The type assertion x.(T) asserts that the concrete value
  stored in x is of type T , and that x is not nil.
