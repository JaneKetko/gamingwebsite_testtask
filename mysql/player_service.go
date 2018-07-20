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
	err := s.DB.QueryRowContext(ctx, fmt.Sprintf("select * from %s where id=%d", s.Name, id)).Scan(&pl.ID, &pl.Name, &pl.Balance)
	if err != nil {
		return nil, fmt.Errorf("cannot get player by ID: %v", err)
	}
	return &pl, nil
}

// AddPlayer inserts new player into table.
func (s PlayerService) AddPlayer(ctx context.Context, name string) (int, error) {
	res, err := s.DB.ExecContext(ctx, fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", s.Name), name)
	if err != nil {
		return 0, fmt.Errorf("could not add player: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("cannot add player: %v", err)
	}
	return int(id), nil
}

// DeletePlayer deletes player by id from table, if possible.
func (s PlayerService) DeletePlayer(ctx context.Context, id int) error {
	_, err := s.DB.ExecContext(ctx, fmt.Sprintf("delete from %s where id=?", s.Name), id)
	if err != nil {
		return fmt.Errorf("cannot delete player: %v", err)
	}
	return nil
}

// UpdatePlayer updates player with player id from table with player.Player, if possible.
func (s PlayerService) UpdatePlayer(ctx context.Context, id int, player player.Player) error {
	_, err := s.DB.ExecContext(ctx, fmt.Sprintf("UPDATE %s SET name=?, balance=? WHERE id=?", s.Name), player.Name, player.Balance, id)
	if err != nil {
		return fmt.Errorf("cannot update player: %v", err)
	}
	return nil
}
