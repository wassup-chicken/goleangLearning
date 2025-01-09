module project

go 1.23.2

replace mymodule => ./animalPackage

require github.com/GoesToEleven/puppy v1.3.0

require (
	github.com/GoesToEleven/dog v0.0.0-20230428023317-90bef1c76cb9 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8
)
