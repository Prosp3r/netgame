##NETGAME CHALLENGE

A simple browser based multiplayer game written in Go and Vanila Javascript.



Usage:
1. Unzip the content
2. go into the folder netgame-master
3. on mac terminal run ./netgame

By default the game runs on http://localhost:8080
Works as described in the instructions.

However it can be customized to
1. Generate fake players
2. Set more than two player requirements for games to start.
3. Run on a different port by using flags. eg. (./netgame 3 8080 4 3)
The first number 3 is the number of games to run concurrently.
The 8080 is the browser port where it'll run can be changed to another.
4 is the minimum number of players required for a game default is 2the last 3 just tells it to generate fake players on start. It'll generate five fake player and assign random numbers to them at the start of each game.

#### Data Persistence & Keep alive
It is inevitable that accidents will happen.
One way this game can be improved in the area of resillience though not currently a requirement, is by persisting the data and restarting with a daemon to continue where it left off from the data store. A key value store or simple .csv file would make things very light weight.








PS: Download zip file from github: https://github.com/Prosp3r/netgame/archive/master.zip
