rm -rf ./bin
GOOS=windows GOARCH=386 go build -o ./bin/scriptotron_i386.exe
GOOS=windows GOARCH=amd64 go build -o ./bin/scriptotron_amd64.exe
go build -o ./bin/scriptotron