export GOOS=linux
export GOARCH=arm
export GOARM=6
go build ...lpd8806
sudo -E go install ...lpd8806
go build test/lpd8806rpi.go