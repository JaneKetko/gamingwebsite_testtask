package mysql

import (
	"context"
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
)

// PlayerService is service for working with table in mysql DB.
type PlayerService struct {
	DB   *Session
	Name string
}

// PlayerByID returns player by id, if exist.
func (s PlayerService) PlayerByID(ctx context.Context, id int) (*player.Player, error) {
	var pl player.Player
	err := s.DB.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id=?", s.Name), id).Scan(&pl.ID, &pl.Name, &pl.Balance)
	if err != nil {
		return nil, fmt.Errorf("cannot get player by ID: %v", err)
	}
	return &pl, nil
}

// AddPlayer inserts new player into table.
func (s PlayerService) AddPlayer(ctx context.Context, name string) (int, error) {
	res, err := s.DB.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", s.Name), name)
	if err != nil {
		return 0, fmt.Errorf("cannot add player: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("cannot add player: %v", err)
	}
	return int(id), nil
}

// DeletePlayer deletes player by id from table, if possible.
// Return error if cannot find player.
func (s PlayerService) DeletePlayer(ctx context.Context, id int) error {
	res, err := s.DB.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE id=?", s.Name), id)
	if err != nil {
		return fmt.Errorf("cannot delete player: %v", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("something wrong while deleted: %v", err)
	}
	if n == 0 {
		return fmt.Errorf("cannot delete player: player not exists")
	}
	return nil
}

// UpdatePlayer updates player with player id from table with player.Player, if possible.
// Return error if cannot find player.
func (s PlayerService) UpdatePlayer(ctx context.Context, id int, player player.Player) error {
	res, err := s.DB.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET name=?, balance=? WHERE id=?", s.Name), player.Name, player.Balance, id)
	if err != nil {
		return fmt.Errorf("cannot update player: %v", err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("something wrong while updated: %v", err)
	}
	if n == 0 {
		return fmt.Errorf("cannot update player: player not exists")
	}
	return nil
}
func (s PlayerService) deleteAllPlayers(ctx context.Context) error {
	_, err := s.DB.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s ", s.Name))
	if err != nil {
		return fmt.Errorf("cannot delete all players: %v", err)
	}
	return nil
}

func (s PlayerService) listAllPlayers(ctx context.Context) ([]*player.Player, error) {
	rows, err := s.DB.QueryContext(ctx, fmt.Sprintf("SELECT * FROM %s ", s.Name))
	if err != nil {
		return nil, fmt.Errorf("cannot get all players: %v", err)
	}

	players := make([]*player.Player, 0)
	for rows.Next() {
		var pl player.Player
		err = rows.Scan(&pl.ID, &pl.Name, &pl.Balance)
		if err != nil {
			return nil, fmt.Errorf("cannot get player: %v", err)
		}
		players = append(players, &pl)
	}
	return players, nil
}
