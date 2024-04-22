# Cheat Sheet

## Run the program chossing a version
```
go run . v1
go run . v2
go run . v3
```

## Reading the .prof files
```
go tool pprof ./results/v2cpu.prof
go tool pprof ./results/v3mem.prof
```

## pprof commands
### top
```
top N
```
list top N processes consuming most resources in the current profile (N is optional)

### list
```
list file.name
```
shows consumed resources for each line for the chosen file

### help
```
help
help command
```
"help" shows list of commands and gives basic directions while "help command" gives detailed instructions for given command
### text
```
text
text >file.txt
```
Outputs top entries in text form and writes it to given file if specified

## Comparing profiles

### diff
```
go tool pprof -diff_base=./results/v1mem.prof ./results/v2mem.prof 
```
compares the actual profile v2mem.prof to the previous profile v1mem.prof