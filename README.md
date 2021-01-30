## NETGAME CHALLENGE

A simple browser based multiplayer game written in Go and Vanila Javascript.


## How it works
This Network game is a lucky number game.

Step 1 Players pick two numbers from 1 to 10 and enter to play with their username.

Game requires a minimum of two players for a game to run. All other's who join after a game starts will be queued for the next game.
Step 2 The computer picks random numbers btween 1 and 10 thirty times. Each time is called a round. 

SCORING
For each round, if the number picked matches those picked by any player, that player is awarded +5 points.
If numbers chosen by the game fall in btween both numbers the players choose they are awarded: +5 - (HigherNumberChosen - LowerNumberChosen)
e.g. player chose 3 & 8 and number for the round is 7 this awards the player 5-(8-3) = 0.
If not a match and not within, say number for the round is 9 the player's score is reduced by one point(-1)
The winner is the player with the highest score at the end of 30 rounds or the one that scores 21 points total first.
In the case of a tie the winner is the one whose name come first this is chosen by ordering their usernames alphabetically.


### Setup:

1. Unzip the content
2. go into the folder netgame-master
3. on mac terminal run ./netgame


### Default setup
By default the game runs on http://localhost:8080

Requires minimum of two player to start.
Works as described in the instructions.

However it can be customized to:
1. Generate fake players in addition to admitting real ones.
2. Set more than two player requirements for games to start.
3. Run on a different port. 
These can be achieved by using flags when starting up the game. eg. 

$ ./netgame 3 8080 4 3

(3  --- 8080 ---   4  ---  3)

(No of games --- Custom Port for game --- Minimum number of players required --- Generate Fake players )

The first number 3 is the number of games to run concurrently.
The 8080 is the browser port where it'll run can be changed to another.
4 is the minimum number of players required for a game default is 2. 
The last 3 just tells it to generate fake players on start. 
It'll generate five fake player and assign random numbers to them at the start of each game.


#### Data Persistence
It is inevitable that accidents will happen.
One way this game can be improved in the area of resillience though not currently a requirement, is by persisting the data and restarting with a daemon to continue where it left off from the data store. A key value store or simple .csv file would make things very light weight.


#### Keep Alive
This feature simply means the system expects that at some point the game could be terminated.
To make sure the system is always up and running, we've adopted two methods of deployment.

1. Deploying it as a **systemd daemon**
2. Deploying it as a **Docker container** that can be deployed in a Kubernetes cluster.
    $ docker run -d --name netgame -p 8080:8080 sirpros/netgame:latest

In this document we are focused on the systemd daemon.





PS: Download zip file from github: https://github.com/Prosp3r/netgame/archive/master.zip
