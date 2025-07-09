# learn-go-with-tests
Learning go with tests

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor


"pkgsite -open ." --> packages documentation


Commands:
- go run file.go // Run
- go test // Tests
- go test -bench=. // Benchmarking. With -benchmem // info about memory allocation


Code:
.Error() -> Converts an error into a string
t.Fatal() -> Stop running the test completely 


Annotations:

- functions != methods
- 'nil' is the same as 'null' in other languages
- GO copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
