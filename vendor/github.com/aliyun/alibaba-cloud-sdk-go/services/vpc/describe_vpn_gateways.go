package vpc

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

// DescribeVpnGateways invokes the vpc.DescribeVpnGateways API synchronously
// api document: https://help.aliyun.com/api/vpc/describevpngateways.html
func (client *Client) DescribeVpnGateways(request *DescribeVpnGatewaysRequest) (response *DescribeVpnGatewaysResponse, err error) {
	response = CreateDescribeVpnGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpnGatewaysWithChan invokes the vpc.DescribeVpnGateways API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevpngateways.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpnGatewaysWithChan(request *DescribeVpnGatewaysRequest) (<-chan *DescribeVpnGatewaysResponse, <-chan error) {
	responseChan := make(chan *DescribeVpnGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpnGateways(request)
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

// DescribeVpnGatewaysWithCallback invokes the vpc.DescribeVpnGateways API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevpngateways.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpnGatewaysWithCallback(request *DescribeVpnGatewaysRequest, callback func(response *DescribeVpnGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpnGatewaysResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpnGateways(request)
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

// DescribeVpnGatewaysRequest is the request struct for api DescribeVpnGateways
type DescribeVpnGatewaysRequest struct {
	*requests.RpcRequest
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	VpnGatewayId         string           `position:"Query" name:"VpnGatewayId"`
	Status               string           `position:"Query" name:"Status"`
	BusinessStatus       string           `position:"Query" name:"BusinessStatus"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeVpnGatewaysResponse is the response struct for api DescribeVpnGateways
type DescribeVpnGatewaysResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	VpnGateways VpnGateways `json:"VpnGateways" xml:"VpnGateways"`
}

// CreateDescribeVpnGatewaysRequest creates a request to invoke DescribeVpnGateways API
func CreateDescribeVpnGatewaysRequest() (request *DescribeVpnGatewaysRequest) {
	request = &DescribeVpnGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpnGateways", "vpc", "openAPI")
	return
}

// CreateDescribeVpnGatewaysResponse creates a response to parse from DescribeVpnGateways response
func CreateDescribeVpnGatewaysResponse() (response *DescribeVpnGatewaysResponse) {
	response = &DescribeVpnGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
