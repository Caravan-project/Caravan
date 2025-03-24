rd /s /q record
del /q log\*
del /q S0_*
del /q S1_*
del mz.log

start cmd /k go run main.go -S 2 -f 1 -s S0 -n N0 -t 20W.csv
start cmd /k go run main.go -S 2 -f 1 -s S0 -n N1 -t 20W.csv
start cmd /k go run main.go -S 2 -f 1 -s S1 -n N0 -t 20W.csv
start cmd /k go run main.go -S 2 -f 1 -s S1 -n N1 -t 20W.csv

timeout /T 20 /NOBREAK && go run main.go -S 2 -c -t 20W.csv