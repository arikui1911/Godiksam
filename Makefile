test: dkc/y.go
	go test github.com/arikui1911/Godiksam/t

dkc/y.go: dkc/diksam.go.y
	go tool yacc -o dkc/y.go dkc/diksam.go.y
