# Usage examples:

## Games list:
curl http://localhost:8080/games

## Add game:
curl http://localhost:8080/games --include --header "Content-Type: application/json" -d @newgame.json

