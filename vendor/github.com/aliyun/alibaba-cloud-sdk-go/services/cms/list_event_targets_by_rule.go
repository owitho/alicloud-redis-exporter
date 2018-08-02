package cms

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

// ListEventTargetsByRule invokes the cms.ListEventTargetsByRule API synchronously
// api document: https://help.aliyun.com/api/cms/listeventtargetsbyrule.html
func (client *Client) ListEventTargetsByRule(request *ListEventTargetsByRuleRequest) (response *ListEventTargetsByRuleResponse, err error) {
	response = CreateListEventTargetsByRuleResponse()
	err = client.DoAction(request, response)
	return
}

// ListEventTargetsByRuleWithChan invokes the cms.ListEventTargetsByRule API asynchronously
// api document: https://help.aliyun.com/api/cms/listeventtargetsbyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEventTargetsByRuleWithChan(request *ListEventTargetsByRuleRequest) (<-chan *ListEventTargetsByRuleResponse, <-chan error) {
	responseChan := make(chan *ListEventTargetsByRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEventTargetsByRule(request)
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

// ListEventTargetsByRuleWithCallback invokes the cms.ListEventTargetsByRule API asynchronously
// api document: https://help.aliyun.com/api/cms/listeventtargetsbyrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListEventTargetsByRuleWithCallback(request *ListEventTargetsByRuleRequest, callback func(response *ListEventTargetsByRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEventTargetsByRuleResponse
		var err error
		defer close(result)
		response, err = client.ListEventTargetsByRule(request)
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

// ListEventTargetsByRuleRequest is the request struct for api ListEventTargetsByRule
type ListEventTargetsByRuleRequest struct {
	*requests.RpcRequest
	RuleName string `position:"Query" name:"RuleName"`
}

// ListEventTargetsByRuleResponse is the response struct for api ListEventTargetsByRule
type ListEventTargetsByRuleResponse struct {
	*responses.BaseResponse
	Code              string                                    `json:"Code" xml:"Code"`
	Message           string                                    `json:"Message" xml:"Message"`
	RequestId         string                                    `json:"RequestId" xml:"RequestId"`
	ParameterCount    int                                       `json:"ParameterCount" xml:"ParameterCount"`
	ContactParameters ContactParametersInListEventTargetsByRule `json:"ContactParameters" xml:"ContactParameters"`
	FcParameters      FcParametersInListEventTargetsByRule      `json:"FcParameters" xml:"FcParameters"`
	MnsParameters     MnsParametersInListEventTargetsByRule     `json:"MnsParameters" xml:"MnsParameters"`
}

// CreateListEventTargetsByRuleRequest creates a request to invoke ListEventTargetsByRule API
func CreateListEventTargetsByRuleRequest() (request *ListEventTargetsByRuleRequest) {
	request = &ListEventTargetsByRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "ListEventTargetsByRule", "cms", "openAPI")
	return
}

// CreateListEventTargetsByRuleResponse creates a response to parse from ListEventTargetsByRule response
func CreateListEventTargetsByRuleResponse() (response *ListEventTargetsByRuleResponse) {
	response = &ListEventTargetsByRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
