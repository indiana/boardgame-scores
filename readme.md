# Usage examples:

## Games list:
`url http://localhost:8080/game`

## Game data:
`curl http://localhost:8080/game/1`

## Add game:
`curl http://localhost:8080/game --include --header "Content-Type: application/json" -d @newgame.json`

## Update game:
`curl http://localhost:8080/game/1 --include --header "Content-Type: application/json" -d @newgame.json -X PUT`

## Delete game:
`curl http://localhost:8080/game/1 -X DELETE`

## Show options:
`curl http://localhost:8080/game -X OPTIONS`
