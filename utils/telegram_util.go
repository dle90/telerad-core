package utils

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	_error "telerad-core-module/error"

	logger "github.com/BeeTechHub/go-common/logger"
)

func SendMessageTeleSuccess(token string, userId string, serviceCode string, serviceName string) (bool, *_error.SystemError) {
	bodyString := fmt.Sprintf("\\ud83d\\udc4d \\ud83d\\udc4d \\ud83d\\udc4d \\nServiceCode: %s \\nServiceName: %s",
		serviceCode, serviceName)
	return SendMessageTelegram(token, userId, bodyString)
}

func SendMessageTeleFail(token string, userId string, serviceCode string, serviceName string, message string, response string) (bool, *_error.SystemError) {
	bodyString := fmt.Sprintf("\\ud83d\\ude21 \\ud83d\\ude21 \\ud83d\\ude21 \\n- ServiceCode: %s\\n- ServiceName: %s\\n- Message: %s \\n- Response: %s",
		serviceCode, serviceName, message, escapeTelegramMessage(response))
	return SendMessageTelegram(token, userId, bodyString)
}

func SendMessageTelegram(token string, userId string, message string) (bool, *_error.SystemError) {
	telegramUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	messageStandard := escapeTelegramMessage(message)
	bodyString := fmt.Sprintf(`{"chat_id":"%s","text":"%s","disable_notification":true}`, userId, messageStandard)
	var jsonStr = []byte(bodyString)
	logger.Infof("SendMessage to telegram with \n- url: %s \n- body: %s", telegramUrl, jsonStr)
	resp, err := http.Post(telegramUrl, "application/json", bytes.NewReader(jsonStr))
	if err != nil {
		logger.Fatalf("SendMessage to telegram (get) An Error Occured %v", err)
		return false, _error.New(err)
	}
	if resp.StatusCode == http.StatusOK {
		logger.Info("SendMessage to telegram (get) successfully!")
		return true, nil
	} else {
		return false, _error.NewErrorByString(fmt.Sprintf("SendMessage to telegram (get) api:%s httpcode:%v", telegramUrl, resp.Status))
	}
}

// ()_-. are reserved by telegram.
func escapeTelegramMessage(input string) string {
	// return strings.NewReplacer(
	// 	"\"", "\\\"",
	// ).Replace(input)
	//return strings.ReplaceAll(input, `"`, `\"`)
	return strings.ReplaceAll(input, "\"", "\\\"")
}
