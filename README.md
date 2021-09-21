# gophercises
Gophercises course
1. Quiz Game
* A CLI quiz program that reads a CSV file of two columns (question and answer; no headings). A timer is also added.
* Commands
    ```go run main.go -f <path-to-file> -t <seconds>```<br>
    The default values for ```-f``` and ```-t``` are ```problem.csv``` and ```30 seconds``` respectively.