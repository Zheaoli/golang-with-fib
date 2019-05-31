package engine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Zheaoli/golang-with-fib/config"
	"github.com/Zheaoli/golang-with-fib/utils"
	"io"
	"io/ioutil"
	"os"
)

type InvokeRequest struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type InvokeResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

var MapFunction map[string]func(params interface{}) (interface{}, int)

func InitFunction() {
	MapFunction = map[string]func(params interface{}) (interface{}, int){}
	MapFunction["demo"] = demo
}

func demo(params interface{}) (interface{}, int) {
	fmt.Println(params)
	data, err := ioutil.ReadFile(config.ServerConfig.TemplatePath)
	if err != nil {
		return "", utils.RemoteServerError
	}
	if errWrite := utils.WriteResult(string(data)); errWrite != nil {
		return "", utils.RemoteServerError
	}
	return "", 0
}

func MainEngine() {
	var (
		data InvokeRequest
		resp InvokeResponse
	)

	stdinData, err := parseInValue()
	if err != nil || stdinData == "" {
		return
	}
	if err := json.Unmarshal([]byte(stdinData), &data); err != nil {
		resp.ID = 0
		resp.JsonRPC = "2.0"
		resp.Error = map[string]interface{}{
			"code":    utils.ServerError,
			"message": "params parse error",
		}
		if errOutput := outputResultToJson(resp); errOutput != nil {
			return
		}
		return
	}
	resp.ID = data.ID
	resp.JsonRPC = data.JsonRPC
	if data.Params == nil {

		resp.Error = map[string]interface{}{
			"code":    utils.ServerError,
			"message": "params parse error",
		}
		if errOutput := outputResultToJson(resp); errOutput != nil {
			return
		}
		return
	}
	if value, err := MapFunction[data.Method]; err {
		result, code := value(data.Params)
		if code != 0 {
			resp.Error = map[string]interface{}{
				"code":    code,
				"message": result,
			}
			if errOutput := outputResultToJson(resp); errOutput != nil {
				return
			}
			return
		}
		resp.Result = result
		if errOutput := outputResultToJson(resp); errOutput != nil {
			return
		}
		return
	} else {
		resp.Error = map[string]interface{}{
			"code":    utils.MethodNotFound,
			"message": "method not found",
		}
		if errOutput := outputResultToJson(resp); errOutput != nil {
			return
		}
		return
	}
}

func parseInValue() (string, error) {
	inputReader := bufio.NewReader(os.Stdin)
	data, err := inputReader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return data, nil
}

func outputResultToJson(response InvokeResponse) error {
	data, err := json.Marshal(&response)
	if err != nil {
		return err
	}
	errWrite := utils.WriteResult(string(data))
	if errWrite != nil {
		return errWrite
	}
	return nil
}
