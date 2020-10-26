# tictactoe module specification

## Abstract

This Cosmos module implements the tictactoe game as a 3x3 matrix. It works by letting users submit an invitation to the network in the form of a transaction. A new game is started when another user sends its own transaction to accept the invitation. Following this, players game submit their moves until a winner emerges or a draw is reached.

## Transaction Commands

The module provides the following transaction commands (CLI format):

1. **create-invitation \<message\>**
2. **delete-invitation \<invitation ID\>**
3. **accept-invitation \<invitation ID\>**
4. **make-move <game ID> \<coordinates\>**

**Note 1:** the coordinates for `make-move` must be *concatenated*. Assuming `jcs` in CLI `fubarapp` submits a move to position (1,2) for game with ID `XPTO`, the command should be:

```
fubarapp tx tictactoe make-move XPTO 12 --from jcs
```

**Note 2:** Once an invitation is accepted, the invitation's ID becomes the game's ID

The REST interface is as follows (assuming `localhost:1317` as the server address):

1. **localhost:1317/tictactoe/invitation (POST and DELETE methods)**
2. **localhost:1317/tictactoe/accept (POST method)**
3. **localhost:1317/tictactoe/move (POST method)**

Just like for CLI format, the move coordinates are also concatenated for the REST interface.

## Query Commands

The query commands are as follows (CLI format):

1. **list-invitation**
2. **list-game**
3. **get-game \<game ID\>**


The output of `fubarapp query tictactoe get-game XPTO` would be something like:

```
Players:

X: jcs
O: scj

Status: Ongoing


Next move: X



   |   | O 
-----------
 O | X | X 
-----------
 X |   | O 


JSON:

{
	"id": "XPTO",
	"state": [
		"AgEA",
		"AAIA",
		"AQIB"
	],
	"players": {
		"1": "scj",
		"2": "jcs"
	},
	"next": 2
}

```

The REST interface for the queries is as follows (all GET methods):

1. **localhost:1317/tictactoe/invitation**
2. **localhost:1317/tictactoe/game**
3. **localhost:1317/tictactoe/game/\<game id\>**