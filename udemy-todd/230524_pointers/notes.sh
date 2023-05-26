
# type: *int, this is a pointer to an int
# a variable with type *int is an address
# * could be used in two manners
# *b is dereferencing a pointer
# if we say x *int it means that x is a pointer to an int
# and we write *x which means that we'll access the value inside
# the pointer, that is, we are dereferencing the pointer

# on the other hand & has only one meaning
# and it is used when we want to take the address of a value
# if a = 10
# then *&a will show 10

# User pointer when the data is large and we don't want to copy 
# everything

# we say mutate a value
# when we use to dereference and store another value in the address

# when we say "pointer receiver" or "pointer value" what do we mean?
# "pointer receiver" refers to a parameter which is a pointer, it could refer to
# the parameter of a object which the "this" or "self"
# "pointer value" refers to an object/value which is a pointer to T

go run main.go









