# Knowing when to use pointers and values

Does adding or removing something from a value of this type need to create a new value or mutate the existing one?
If the answer is create a new value, then use value receivers for your methods.
If the answer is mutate the value, then use pointer receivers.
