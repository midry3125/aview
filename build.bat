set GOARCH=amd64

set GOOS=windows
go build

set BuildAll=0
if "%1"=="-a" set BuildAll=1
if "%1"=="--all" set BuildAll=1

if %BuildAll%==1 (
set GOOS=linux
go build -o aview-linux

set GOOS=darwin
go build -o aview-mac
)
