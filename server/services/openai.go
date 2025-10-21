package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func GenerateGoStructFieldsFromPayload(payload map[string]interface{}) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	jsonData, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshalling payload: %w", err)
	}

	prompt := fmt.Sprintf(`Given the following JSON payload, output only the Go struct FIELDS (inside the struct body),
not the type declaration. Use idiomatic Go naming conventions and JSON tags.

The output should be valid Go syntax like this:
{
    FieldName string `+"`json:\"field_name\"`"+`
    AnotherField int `+"`json:\"another_field\"`"+`
}

Do NOT include commas between fields.

JSON:
%s
`, string(jsonData))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0,
		},
	)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	output := strings.TrimSpace(resp.Choices[0].Message.Content)

	re := regexp.MustCompile(`(?s)\{(.*)\}`)
	matches := re.FindStringSubmatch(output)
	if len(matches) > 1 {
		output = strings.TrimSpace(matches[1])
	}

	return output, nil
}
