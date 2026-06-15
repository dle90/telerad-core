package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	_error "telerad-core-module/error"

	logger "github.com/BeeTechHub/go-common/logger"
)

func RunHealthCheckRequestGet(url string) (bool, *_error.SystemError) {
	resp, err := http.Get(url)
	if os.IsTimeout(err) {
		logger.Warnf("heal-check (get timeout error: %v\n", err)
		return false, _error.New(err)
	}
	if err != nil {
		logger.Warnf("heal-check (get) An Error Occured %v\n", err)
		return false, _error.New(err)
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		errorString := "Không parse được body của response!!!"
		defer resp.Body.Close()
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			logger.Warnf("heal-check (get) An Error Occured when parse body on response %v\n", err)
		} else {
			errorString = string(b)
		}
		return false, _error.NewErrorByString(fmt.Sprintf("Health-check (get) api:%s httpcode:%v response:%s", url, resp.Status, errorString))
	}
}

func RunHealthCheckRequestPost(url string) (bool, *_error.SystemError) {
	//Leverage Go's HTTP Post function to make request
	resp, err := http.Post(url, "application/json", nil)
	//Handle Error
	if err != nil {
		logger.Fatalf("heal-check (post) An Error Occured %v", err)
		return false, _error.New(err)
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	} else {
		return false, _error.NewErrorByString(fmt.Sprintf("Health-check (post) api:%s httpcode:%v", url, resp.Status))
	}
}
