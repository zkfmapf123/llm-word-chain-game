package services

import (
	"fmt"

	"github.com/zkfmapf123/go-llm/config"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// 단어 중복 체크
func (u *UserService) CheckOverlab(sessionId int, word string) (bool, error) {
	pg := config.NewPGConn().MustConnect()

	var exists bool
	err := pg.DB.QueryRow("select exists(select 1 from words where session_id=$1 and word=$2)", sessionId, word).Scan(&exists)
	if err != nil {
		return true, err
	}

	defer pg.Close()
	return exists, err
}

// 단어 저장
func (u *UserService) SaveWord(sessionId int, word string) error {
	pg := config.NewPGConn().MustConnect()

	_, err := pg.DB.Exec("INSERT INTO words (session_id, word) VALUES ($1, $2)", sessionId, word)
	if err != nil {
		return err
	}

	defer pg.Close()
	return nil
}

func (u *UserService) LoadAllWord(sessionId int) ([]string, error) {
	pg := config.NewPGConn().MustConnect()

	var words []string
	rows, err := pg.DB.Query("select word from words where session_id=$1 order by created_at desc", sessionId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var word string
		rows.Scan(&word)
		words = append(words, word)
	}

	defer pg.Close()
	return words, nil
}

// prompting...
func (u *UserService) WordPrompting(word string, words []string) string {
	wr := []rune(word)

	prompt := "너는 한국어 끝말잇기 전문가야. 사용자가 입력한 단어를 기반으로 끝말잇기 규칙을 준수하며 새로운 단어를 추천해야 해.\n"
	prompt += fmt.Sprintln("다음 조건을 만족하는 단어를 추천해줘:")
	prompt += fmt.Sprintf("- 반드시 '%s'로 시작하는 단어를 추천할 것\n", string(wr[len(wr)-1]))
	prompt += "- 아래 단어들은 절대 사용하면 안 됨:\n"
	for _, w := range words {
		prompt += fmt.Sprintf("  - %s\n", w)
	}
	prompt += "- 최대한 긴 단어를 추천할 것 (한자어 및 고유명사 포함 가능)\n"
	prompt += "- 끝말잇기에서 유리한 단어를 추천해야 함 (예: 상대가 답하기 어려운 단어)\n"
	prompt += "- 단어 하나만 출력할 것. 그 외의 설명은 필요 없음.\n"

	return prompt
}

// prompting...
func (u *UserService) WordValidCheckPrompting(word string, words []string) string {
	prompt := "너는 한국어 끝말잇기 전문가야. 사용자가 입력한 단어를 기반으로 끝말잇기 규칙을 준수하며 새로운 단어를 추천해야 해.\n"
	prompt += fmt.Sprintf("%s 이 단어가 실제로 존재하는 단어인지 확인해줘.\n", word)
	prompt += fmt.Sprintln("답변은 아래처럼 답변해줘")
	prompt += fmt.Sprintln("1. 존재하는 단어라면 true, 아니라면 false")
	return prompt
}
