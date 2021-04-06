# snakegame

Instructions to run the snakeGame:
1) Run the following command by providing height and width of the snake board as integers (10<height<40, 10<width<60):  
    go run cmd/main.go height width  
    Eg: For running the snake game with board of height 40 and width 60  
    go run cmd/main.go 40 60  
2) Press ESC key to quit the game in between.
    
    
Time taken to complete the assignment: approx 25 hours.

Implementation design decisions:
1) First problem was taking the input and displaying the output properly without creating a mess. Since arrow keys is widely used, to get input for arrow keys, I used the https://github.com/nsf/termbox-go library. Not only it solved the input problem, it also solved the problem for displaying the output of snake board. Since termbox-go opens a seperate mini terminal and refreshes it after every instruction, every user move updated the mini terminal giving it the feel of smooth real time snake game instead of a terminal where snakeboards are stacked one after another after every move.
2) Second major problem was how to move the snake and update the position of each body part of snake. The best solution was to maintain a seperate 2-D slice that contains the move on a particular cell the snake's head took(or user took) when it was on that position. With this thing in mind, it was easier to update the position of the snake and how many cells it was occupying. Since each time head moved to a new cell, it would leave its current cell and the following body parts will follow the same path. So instead of updating position of each part by clearing it's current position and updating the new position, only the snake's head and snake's tail needed to be updated. Instead of marking the cell as empty when snake's head move, keep it marked as occupied, so it would be like the body part following the head took it's place. Only mark the cell as unoccupied when snake's tail has left that cell. And since we have been updating every move of head when it passes a cell, the same move would be used by tail to move to the next cell.
