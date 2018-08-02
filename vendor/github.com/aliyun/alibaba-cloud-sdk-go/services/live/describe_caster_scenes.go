package live

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCasterScenes invokes the live.DescribeCasterScenes API synchronously
// api document: https://help.aliyun.com/api/live/describecasterscenes.html
func (client *Client) DescribeCasterScenes(request *DescribeCasterScenesRequest) (response *DescribeCasterScenesResponse, err error) {
	response = CreateDescribeCasterScenesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCasterScenesWithChan invokes the live.DescribeCasterScenes API asynchronously
// api document: https://help.aliyun.com/api/live/describecasterscenes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCasterScenesWithChan(request *DescribeCasterScenesRequest) (<-chan *DescribeCasterScenesResponse, <-chan error) {
	responseChan := make(chan *DescribeCasterScenesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasterScenes(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCasterScenesWithCallback invokes the live.DescribeCasterScenes API asynchronously
// api document: https://help.aliyun.com/api/live/describecasterscenes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeCasterScenesWithCallback(request *DescribeCasterScenesRequest, callback func(response *DescribeCasterScenesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCasterScenesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasterScenes(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCasterScenesRequest is the request struct for api DescribeCasterScenes
type DescribeCasterScenesRequest struct {
	*requests.RpcRequest
	CasterId string           `position:"Query" name:"CasterId"`
	SceneId  string           `position:"Query" name:"SceneId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCasterScenesResponse is the response struct for api DescribeCasterScenes
type DescribeCasterScenesResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Total     int       `json:"Total" xml:"Total"`
	SceneList SceneList `json:"SceneList" xml:"SceneList"`
}

// CreateDescribeCasterScenesRequest creates a request to invoke DescribeCasterScenes API
func CreateDescribeCasterScenesRequest() (request *DescribeCasterScenesRequest) {
	request = &DescribeCasterScenesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasterScenes", "live", "openAPI")
	return
}

// CreateDescribeCasterScenesResponse creates a response to parse from DescribeCasterScenes response
func CreateDescribeCasterScenesResponse() (response *DescribeCasterScenesResponse) {
	response = &DescribeCasterScenesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
