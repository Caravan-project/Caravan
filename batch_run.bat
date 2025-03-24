rd /s /q record
del /q log\*
del /q S0_*
del /q S1_*
del /q S2_*
del /q S3_*
del mz.log

start cmd /k go run main.go -S 4 -f 1 -s S0 -n N0 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N1 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N2 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N3 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N4 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N5 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S0 -n N6 -t 20W.csv

start cmd /k go run main.go -S 4 -f 1 -s S1 -n N0 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N1 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N2 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N3 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N4 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N5 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S1 -n N6 -t 20W.csv

start cmd /k go run main.go -S 4 -f 1 -s S2 -n N0 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N1 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N2 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N3 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N4 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N5 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S2 -n N6 -t 20W.csv

start cmd /k go run main.go -S 4 -f 1 -s S3 -n N0 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N1 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N2 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N3 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N4 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N5 -t 20W.csv
start cmd /k go run main.go -S 4 -f 1 -s S3 -n N6 -t 20W.csv


timeout /T 20 /NOBREAK && go run main.go -S 2 -c -t 20W.csv