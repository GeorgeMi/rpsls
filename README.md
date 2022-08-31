# Rock, Paper, Scissors, Lizard, Spock API
RPSLS-API is a game service that takes input from the user to calculate the winner based on what the computer player chooses. The API can be easily extended by adding new elements in choicesMap. Also the scoreboard history size can be changed from scoreBoard.history var.

---

### Usage ###

#### Build the rpsls-api binary ###

```
go build -o rpsls-api.exe

```

#### Start the api on port :8080 ###

```
./rpsls-api.exe run

```

#### Build the rpsls-api using Docker ###

```
docker build -t rpsls-api .  

```

#### Start the api on port :8080 using Docker###

```
docker run -dp 8080:8080 rpsls-api

```

---

### Endpoints ###

#### Get all the choices that are usable for the UI ####

```
Example valid url: GET /rpsls-api/choices
```

#### Get a randomly generated choice ####

```
Example valid url: GET /rpsls-api/choice
```

#### Play a round against a computer opponent ####

```
Example valid url: POST /rpsls-api/play/
{
  “player”: 0 
}

The choice of the winner uses an implemented algorithm based on the following matrix:
    | 0 | 1 | 2 | 3 | 4 |
| 0 | t | + | - | + | - |
| 1 | - | t | + | - | + |
| 2 | + | - | t | + | - |
| 3 | - | + | - | t | + |
| 4 | + | - | + | - | t |

Where:
	0: paper,
	1: rock,
	2: lizard,
	3: spock,
	4: scissors,
	+: win
	-: lose
	t: tie

```

#### Get scoreboard with the 10 most recent results ####

```
Example valid url: GET /rpsls-api/scoreboard
```

#### Reset scoreboard ####

```
Example valid url: PUT /rpsls-api/scoreboard
```

#### Health #### 

```
Example valid url: GET /rpsls-api/health
```

*For up to date endpoints, view the OpenAPI specs at GET /rpsls-api/docs/*


