package models

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Boardgame struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./boardgame-scores.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func GetGames(count int) ([]Boardgame, error) {

	rows, err := DB.Query("SELECT id, name, description from boardgames LIMIT " + strconv.Itoa(count))

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	games := make([]Boardgame, 0)

	for rows.Next() {
		singleGame := Boardgame{}
		err = rows.Scan(&singleGame.ID, &singleGame.Name, &singleGame.Description)

		if err != nil {
			return nil, err
		}

		games = append(games, singleGame)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return games, err
}

func GetGameById(id string) (Boardgame, error) {

	stmt, err := DB.Prepare("SELECT id, name, description from boardgames WHERE id = ?")

	if err != nil {
		return Boardgame{}, err
	}

	game := Boardgame{}

	sqlErr := stmt.QueryRow(id).Scan(&game.ID, &game.Name, &game.Description)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Boardgame{}, nil
		}
		return Boardgame{}, sqlErr
	}
	return game, nil
}

func AddGame(newGame Boardgame) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO boardgames (name, description) VALUES (?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newGame.Name, newGame.Description)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateGame(ourGame Boardgame, id int) (bool, error) {

	tx, err := DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE boardgames SET name = ?, description = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ourGame.Name, ourGame.Description, id)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteGame(gameId int) (bool, error) {

	tx, err := DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := DB.Prepare("DELETE from boardgames where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(gameId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
