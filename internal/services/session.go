package services

import "github.com/zkfmapf123/go-llm/config"

type SessionServiceParams struct {
	sessionId int
}

func NewSession(id int) *SessionServiceParams {
	return &SessionServiceParams{
		sessionId: id,
	}
}

func (s *SessionServiceParams) Start() error {
	pg := config.NewPGConn().MustConnect()
	_, err := pg.DB.Exec("INSERT INTO sessions (id) VALUES ($1)", s.sessionId)
	if err != nil {
		return err
	}

	defer pg.Close()
	return nil
}
