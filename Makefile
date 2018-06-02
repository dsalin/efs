all:
	go build -o efsctl main.go
clean:
	rm -rf ./.efs
