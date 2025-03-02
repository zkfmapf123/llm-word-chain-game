package handlers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/zkfmapf123/go-llm/config"
	"github.com/zkfmapf123/go-llm/internal/services"
	"github.com/zkfmapf123/go-llm/internal/utils"
)

type UserInputRequest struct {
	SesiosnId int    `json:"session_id"`
	Word      string `json:"word"`
}

type UserInputsResponse struct {
	Word      string `json:"word"`
	SesiosnId int    `json:"session_id"`
	YouLost   bool   `json:"you_lost"`
}

// UserInputsHandlers godoc
//
//	@summary		UserInputs
//	@description	userInputs
//	@accept			json
//	@param			request	body		handlers.UserInputRequest	true	"UserInputRequest"
//	@success		200		{object}	handlers.UserInputsResponse
//	@router			/api/play [post]
func UserInputsHandlers(c *fiber.Ctx) error {
	llmClient := config.NewOpenAI()
	req, err := utils.Serialize[UserInputRequest](c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = utils.CheckParallel(req.Word,
		utils.CheckFuncParmas[string]{
			Fn: func(v string) bool {
				return v == ""
			},
			Err: utils.ErrInvalidEmptyWord,
		},
		utils.CheckFuncParmas[string]{
			Fn: func(v string) bool {
				return len(v) < 2
			},
			Err: utils.ErrInvalidWordLen,
		},
	)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	log.Println("--- start ---")

	// [TODO] sesisonId 검사

	// 중복검사
	isOverlab, err := services.NewUserService().CheckOverlab(req.SesiosnId, req.Word)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if isOverlab {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "중복된 단어입니다."})
	}

	log.Println("1. 중복검사... 완료")

	// 유저가 준 단어가 최근단어의 첫단어인가?
	words, err := services.NewUserService().LoadAllWord(req.SesiosnId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if len(words) > 0 && !utils.ComparisonVarsPrefixSuffix(req.Word, words[0]) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s의 마지막글자로 이뤄져야합니다", words[0])})
	}

	log.Println("2. 단어 적합성 완료 / 단어 불러오기... 완료 ", words)

	// 이 단어가 정말 실제하는 단어인가?

	log.Println("3. 실제 존재하는 단어 or 단어 / 명사인지 인지 체크")
	validCheckPrompt := services.NewUserService().WordValidCheckPrompting(req.Word, words)
	validCheck, err := llmClient.SetPrompt(validCheckPrompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	if validCheck == "false" {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": fmt.Sprintf("%s 는 실제로 존재하지 않는 단어입니다.", req.Word)})
	}

	log.Println("3. 실제 존재하는 단어 or 단어 / 명사인지 인지 체크... 완료 ", validCheck)

	// 단어 저장
	err = services.NewUserService().SaveWord(req.SesiosnId, req.Word)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Println("4.단어 저장... 완료 ", req.Word)

	// 프롬프팅
	prompt := services.NewUserService().WordPrompting(req.Word, words)

	log.Println("5.프롬프팅... 완료 ", prompt)

	// 단어 추천
	selectWord, err := llmClient.SetPrompt(prompt)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Println("6.단어 추천... 완료 ", selectWord)

	// LLM이 만든 단어 저장
	err = services.NewUserService().SaveWord(req.SesiosnId, selectWord)
	if err != nil {
		// return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"word": selectWord, "session_id": req.SesiosnId, "you_lost": false})
	}

	log.Println("7.LLM이 만든 단어 저장... 완료 ", selectWord)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"word": selectWord, "session_id": req.SesiosnId, "you_lost": false})
}
