# Project: 
Lem-in

# Authors: 
* Jacob Pesämaa
* Nafisah Rantasalmi
* Huong Le
* Thy Bui

# Technologies:
Project is built with Golang version 1.19

# Description: 
This project is meant represent a digital version of an ant farm. The program will read from a file (including details about the ants and the colony) given in the arguments. Upon sucessfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

# Basic Features:
* The program reads from a file given in the arguments (the examples). This file includes: the number of the ants, rooms, links, start and end rooms, and comments.
* The program finds the most efficient paths for the ants to reach from start to end rooms. 
* The program will then display the content of the file, and each moves of the ants. 

# Set up and usage: 
1. Clone the repo in VSCode/ your text editor of choice or Terminal.
2. Go run the program: 
```
$ go run . <name-of-file.txt>
```

# Implimentation Details:
1. Check the validity of the file, including:
- Check if the number argument is 2 and if the name of file is valid
- Check if the input file is valid: If the file provides number of ants and if the number of ants is valid (an int and > 0), If there are valid start and end rooms, If the ##start or ##end is EOF, If all the links are valid, If the comments are valid, If the file is empty

2. Create ant hill: 
- Read file
- If the file is not empty, split strings and get data from file
- Format the data taken from the file: Removing all comments; Empty string before, between, and after elements
- Return cleaned data
- Create anthill: Find all rooms and all links

3. Check path from start room to end room:
- Use func CheckPathToEndFromStartRoom to check if there is any path leading from start to end point. If not, the program exits and returns an ERR

4. Create tunnels

5. Find all possible paths:
- FindAllPaths takes in a start room and end room as parameters and returns all paths
- It starts at the start room and appends the next path recursively till the end room

6. Find best paths:
- Append paths that do not have intersections and stores to BestPaths

7. Find the number of possible paths using NumberPossible function - If the maximum of paths is equal to the number of possible paths, release and move the ants, if not, calculate steps and then do the later. 

8. Release the ants by giving each ant a path and move the ants

9. Print results:
- Print out the content of the given file and each moves of the ants