package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Config 结构用于保存从 config.json 中读取的配置值
type Config struct {
	Token              string `json:"token"`
	CertificateFile    string `json:"certificateFile"`
	CertificateKeyFile string `json:"certificateKeyFile"`
	Password           string `json:"password"`
	TaskInterval       int    `json:"taskInterval"`
}

func main() {
	// 从本地的 config.json 文件中读取配置信息
	config, err := readConfig("config.json")
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	// 在程序启动时执行一次时间任务
	UpdateSLL(config)

	// 设置定时任务，每隔指定天数执行一次时间任务
	ticker := time.NewTicker(time.Hour * 24 * time.Duration(config.TaskInterval))
	defer ticker.Stop()

	for range ticker.C {
		UpdateSLL(config)
	}
}

func UpdateSLL(config *Config) {
	// 读取证书文件内容
	certificate, err := readFromFile(config.CertificateFile)
	if err != nil {
		fmt.Println("Error reading certificate file:", err)
		return
	}

	// 读取证书密钥文件内容
	certificateKey, err := readFromFile(config.CertificateKeyFile)
	if err != nil {
		fmt.Println("Error reading certificate key file:", err)
		return
	}

	// 构建请求体
	requestBody := map[string]string{
		"certificate":    certificate,
		"certificateKey": certificateKey,
		"password":       config.Password,
	}

	// 将请求体转换为 JSON 格式
	payload, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// 发送 POST 请求
	resp, err := sendPOSTRequest("https://api.cachefly.com/api/2.5/certificates", config.Token, payload)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取服务器返回的数据
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// 根据状态码和返回数据处理结果
	switch resp.StatusCode {
	case http.StatusOK:
		// 服务器返回200 OK，解析返回数据并打印
		var data ServerResponse
		err = json.Unmarshal(respBody, &data)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}
		fmt.Println("Response Status:", resp.Status)
		fmt.Println("Response Data:", data)
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden:
		// 服务器返回400, 401, 403，解析错误信息并打印
		var errorData ServerResponse
		err = json.Unmarshal(respBody, &errorData)
		if err != nil {
			fmt.Println("Error decoding error response:", err)
			return
		}
		fmt.Println("Error Status:", resp.Status)
		fmt.Println("Error Message:", errorData.Message)
	default:
		// 处理其他状态码
		fmt.Println("Unexpected Status:", resp.Status)
		fmt.Println("Response Data:", string(respBody))
	}
}

// 从指定的文件中读取配置信息，并返回 Config 结构
func readConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据到 Config 结构
	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// 从指定的文件中读取内容并返回字符串
func readFromFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 读取文件内容
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// 发送 POST 请求
func sendPOSTRequest(url, token string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	// 设置请求头部
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 结构体用于保存服务器返回的数据
type ServerResponse struct {
	Message string `json:"message"`
	// 其他需要处理的字段根据服务器返回的数据格式添加
}
