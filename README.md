# Math Quiz CLI
Math Quiz uses a 'problems.csv' file to provide the math quiz problems.

## Usage

Clone Repo & Navigate to root directory
```shell
git clone <repository-url> && cd <repository-directory>
```

Build Executable
```shell
go build .
```

Run Executable
```shell
./math-quiz-cli
```
Run Executable with flags
```shell
./math-quiz-cli -csv custom.csv -limit=30
```
flags: 
-csv
    A csv file in the format of 'problem,answer'
-limit
    Total time limit for the quiz in seconds

Example
```shell
bash-3.2$ ./math-quiz-cli
Problem #1: 1+0 = 
1
Problem #2: 1+1 = 
3
You scored 1 out of 2.
```

