export GOOS=linux
export GOARCH=arm
export GOARM=6
go build ...lpd8806
#sudo -E go install ...lpd8806
go build tests/lpd8806/lpd8806rpi.go
go build tests/spi_lib/main.go