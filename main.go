// Generated by Sideko (sideko.dev)
package sideko_hacker_news

import (
"encoding/json"
	"fmt"
	"io"
    "mime/multipart"
	"net/http"
	"net/url"
    "os"
	"path/filepath"
	"strconv"
	"time"
)

func anyToString(value interface{}) string {
	switch v := value.(type) {
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case string:
		return v
	default:
		return fmt.Sprintf("%v", v)
	}
}

func addFileToFormDataWriter(writer *multipart.Writer, field string, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	part, err := writer.CreateFormFile(field, filepath.Base(file.Name()))
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	fmt.Printf("copied %v", part)

	return nil
}

func addFieldToFormDataWriter(writer *multipart.Writer, field string, value any) error {
	label, err := writer.CreateFormField(field)
	if err != nil {
		return err
	}
	label.Write([]byte(anyToString(value)))
	return nil
}

type RequestError struct {
	StatusCode int
	Method     string
	Url        string
	Data       any
	Request    http.Request
	Response   http.Response
}

func NewRequestError(request http.Request, response http.Response) RequestError {
	body, _ := io.ReadAll(response.Body)
	var unmarshaled any
	err := json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return RequestError{
			StatusCode: response.StatusCode,
			Method:     request.Method,
			Url:        request.URL.String(),
			Data:       string(body),
			Request:    request,
			Response:   response,
		}
	}

	return RequestError{
		StatusCode: response.StatusCode,
		Method:     request.Method,
		Url:        request.URL.String(),
		Data:       unmarshaled,
		Request:    request,
		Response:   response,
	}
}

func (e RequestError) Error() string {
	return fmt.Sprintf("RequestError: received %d from %s %s", e.StatusCode, e.Method, e.Url)
}

func errorForStatus(request http.Request, response http.Response) error {
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return NewRequestError(request, response)
	}
	return nil
}

type Client struct {
	httpClient *http.Client
	baseUrl    string
	username   string
	password   string
    token      string
    apiKey     string
}

// Instantiate a new API client
func NewClient() *Client {
    baseUrl := `https://hacker-news.firebaseio.com/v0`
    httpClient := http.Client{Timeout: time.Duration(3) * time.Second}

    client := Client{httpClient: &httpClient, baseUrl: baseUrl}
	return &client
}

// Updates base url of API client
func (c *Client) SetBaseUrl(url string) {
	c.baseUrl = url
}

// Returns base url of API client
func (c *Client) BaseUrl() string {
	return c.baseUrl
}

// Updates token of API client
func (c *Client) SetToken(token string) {
	c.token = token
}

// Sets request timeout of client
func (c *Client) SetTimeout(timeout time.Duration) {
	httpClient := http.Client{Timeout: timeout}
	c.httpClient = &httpClient
}

// Returns timeout
func (c *Client) Timeout() time.Duration {
	return c.httpClient.Timeout
}

func (c *Client) GetAskStoryIds(request GetAskstoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/askstories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetBestStoryIds(request GetBeststoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/beststories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetItem(request GetItemIdJsonRequest) (Item, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/item/"+anyToString(request.ID)+".json")
    if err != nil {
        return Item{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return Item{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return Item{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Item{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return Item{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Item{}, err
    }

    var castedBody Item
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return Item{}, err
    }
    return castedBody, nil
}
func (c *Client) GetJobStoryIds(request GetJobstoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/jobstories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetMaxItemId(request GetMaxitemJsonRequest) (int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/maxitem.json")
    if err != nil {
        return 0, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return 0, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return 0, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return 0, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    var castedBody int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return 0, err
    }
    return castedBody, nil
}
func (c *Client) GetNewStoryIds(request GetNewstoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/newstories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetShowStoryIds(request GetShowstoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/showstories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetTopStoryIds(request GetTopstoriesJsonRequest) ([]int, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/topstories.json")
    if err != nil {
        return []int{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return []int{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return []int{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return []int{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return []int{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return []int{}, err
    }

    var castedBody []int
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return []int{}, err
    }
    return castedBody, nil
}
func (c *Client) GetUpdates(request GetUpdatesJsonRequest) (GetUpdatesJSONResponse, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/updates.json")
    if err != nil {
        return GetUpdatesJSONResponse{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return GetUpdatesJSONResponse{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return GetUpdatesJSONResponse{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return GetUpdatesJSONResponse{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return GetUpdatesJSONResponse{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return GetUpdatesJSONResponse{}, err
    }

    var castedBody GetUpdatesJSONResponse
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return GetUpdatesJSONResponse{}, err
    }
    return castedBody, nil
}
func (c *Client) GetUser(request GetUserIdJsonRequest) (User, error) {
    rawUrl, err := url.JoinPath(c.baseUrl, "/user/"+anyToString(request.ID)+".json")
    if err != nil {
        return User{}, err
    }

    targetUrl, err := url.Parse(rawUrl)
	if err != nil {
		return User{}, err
	}

    queryParams := targetUrl.Query()
    if request.Print != nil {
        queryParams.Set("print", anyToString(*request.Print))
    }
    targetUrl.RawQuery = queryParams.Encode()


    req, err := http.NewRequest(
        "GET",
        targetUrl.String(),
        nil,
    )
    if err != nil {
        return User{}, err
    }



    resp, err := c.httpClient.Do(req)
    if err != nil {
        return User{}, err
    }
    defer resp.Body.Close()

    statusErr := errorForStatus(*req, *resp)
	if statusErr != nil {
		return User{}, statusErr
	}

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return User{}, err
    }

    var castedBody User
    err = json.Unmarshal(body, &castedBody)
    if err != nil {
        return User{}, err
    }
    return castedBody, nil
}

