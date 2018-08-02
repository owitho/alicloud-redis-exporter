package slb

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

// SetLoadBalancerAutoReleaseTime invokes the slb.SetLoadBalancerAutoReleaseTime API synchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerautoreleasetime.html
func (client *Client) SetLoadBalancerAutoReleaseTime(request *SetLoadBalancerAutoReleaseTimeRequest) (response *SetLoadBalancerAutoReleaseTimeResponse, err error) {
	response = CreateSetLoadBalancerAutoReleaseTimeResponse()
	err = client.DoAction(request, response)
	return
}

// SetLoadBalancerAutoReleaseTimeWithChan invokes the slb.SetLoadBalancerAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLoadBalancerAutoReleaseTimeWithChan(request *SetLoadBalancerAutoReleaseTimeRequest) (<-chan *SetLoadBalancerAutoReleaseTimeResponse, <-chan error) {
	responseChan := make(chan *SetLoadBalancerAutoReleaseTimeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLoadBalancerAutoReleaseTime(request)
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

// SetLoadBalancerAutoReleaseTimeWithCallback invokes the slb.SetLoadBalancerAutoReleaseTime API asynchronously
// api document: https://help.aliyun.com/api/slb/setloadbalancerautoreleasetime.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLoadBalancerAutoReleaseTimeWithCallback(request *SetLoadBalancerAutoReleaseTimeRequest, callback func(response *SetLoadBalancerAutoReleaseTimeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLoadBalancerAutoReleaseTimeResponse
		var err error
		defer close(result)
		response, err = client.SetLoadBalancerAutoReleaseTime(request)
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

// SetLoadBalancerAutoReleaseTimeRequest is the request struct for api SetLoadBalancerAutoReleaseTime
type SetLoadBalancerAutoReleaseTimeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	Tags                 string           `position:"Query" name:"Tags"`
	LoadBalancerId       string           `position:"Query" name:"LoadBalancerId"`
	AutoReleaseTime      requests.Integer `position:"Query" name:"AutoReleaseTime"`
}

// SetLoadBalancerAutoReleaseTimeResponse is the response struct for api SetLoadBalancerAutoReleaseTime
type SetLoadBalancerAutoReleaseTimeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSetLoadBalancerAutoReleaseTimeRequest creates a request to invoke SetLoadBalancerAutoReleaseTime API
func CreateSetLoadBalancerAutoReleaseTimeRequest() (request *SetLoadBalancerAutoReleaseTimeRequest) {
	request = &SetLoadBalancerAutoReleaseTimeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "SetLoadBalancerAutoReleaseTime", "slb", "openAPI")
	return
}

// CreateSetLoadBalancerAutoReleaseTimeResponse creates a response to parse from SetLoadBalancerAutoReleaseTime response
func CreateSetLoadBalancerAutoReleaseTimeResponse() (response *SetLoadBalancerAutoReleaseTimeResponse) {
	response = &SetLoadBalancerAutoReleaseTimeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
