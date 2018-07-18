package mysql

import (
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
)

type PlayerService struct {
	DB   *Session
	Name string
}

func (s PlayerService) PlayerByID(id int) (*player.Player, error) {
	var pl player.Player
	err := s.DB.QueryRow(fmt.Sprintf("select * from %s where id=%d", s.Name, id)).Scan(&pl.ID, &pl.Name, &pl.Balance)
	if err != nil {
		return nil, fmt.Errorf("cannot get player by ID: %v", err)
	}
	return &pl, nil

}

func (s PlayerService) AddPlayer(name string) (int, error) {
	res, err := s.DB.Exec(fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", s.Name), name)
	if err != nil {
		return 0, fmt.Errorf("could not add player: %v", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("cannot add Player: %v", err)
	}
	return int(id), nil
}

func (s PlayerService) DeletePlayer(id int) error {
	_, err := s.DB.Exec(fmt.Sprintf("delete from %s where id=?", s.Name), id)
	if err != nil {
		return fmt.Errorf("cannot delete Player: %v", err)
	}
	return nil
}

func (s PlayerService) UpdatePlayer(id int, player player.Player) error {
	_, err := s.DB.Exec(fmt.Sprintf("UPDATE %s SET name=?, balance=? WHERE id=?", s.Name), player.Name, player.Balance, id)
	if err != nil {
		return fmt.Errorf("cannot update Player: %v", err)
	}
	return nil
}
