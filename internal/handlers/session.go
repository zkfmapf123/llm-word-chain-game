package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/zkfmapf123/go-llm/internal/services"
)

type SessionResponse struct {
	Msg       string
	SessionId int
}

// SessionHandlers godoc
//
//	@summary		Setting Session
//	@description	session 생성
//	@accept			json
//	@success		200	{object}	handlers.SessionResponse
//	@router			/api/session [get]
func SessionHandler(c *fiber.Ctx) error {
	sessionId := int(uuid.New().ID())

	sessionService := services.NewSession(sessionId)
	err := sessionService.Start()
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(SessionResponse{Msg: "session start error", SessionId: 0})
	}

	return c.JSON(SessionResponse{Msg: "session start success", SessionId: sessionId})
}
