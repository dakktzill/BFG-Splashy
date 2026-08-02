

rmdir /S /Q "..\builds"

mkdir "../builds"
mkdir "../builds/linux"
mkdir "../builds/windows"

go build -o ../builds/linux/server ./cmd/server
go build -o ../builds/linux/setup ./cmd/setup

go build -o ../builds/windows/server.exe ./cmd/server/
go build -o ../builds/windows/setup.exe ./cmd/setup/

xcopy "..\db" "..\builds\linux\db" /E /I
xcopy "..\dlc" "..\builds\linux\db" /E /I
copy "..\README.md" "..\builds\linux\README.md"
copy "..\LICENSE" "..\builds\linux\LICENSE"
copy "..\constants.json" "..\builds\linux\constants.json"

xcopy "..\dlc" "..\builds\windows\db" /E /I
copy "..\README.md" "..\builds\windows\README.md"
copy "..\LICENSE" "..\builds\windows\LICENSE"
copy "..\constants.json" "..\builds\windows\constants.json"

powershell -Command "Compress-Archive -Path '..\builds\linux\' -DestinationPath '..\builds\linux.zip'"
powershell -Command "Compress-Archive -Path '..\builds\windows\' -DestinationPath '..\builds\windows.zip'"

echo Build Complete!

pause