module project

go 1.23.2

replace mymodule => ./animalPackage

require (
	github.com/GoesToEleven/puppy v1.3.0
	mymodule v0.0.0-00010101000000-000000000000
)

require github.com/GoesToEleven/dog v0.0.0-20230428023317-90bef1c76cb9 // indirect
