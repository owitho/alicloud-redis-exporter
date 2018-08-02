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

// CreateSslVpnClientCert invokes the vpc.CreateSslVpnClientCert API synchronously
// api document: https://help.aliyun.com/api/vpc/createsslvpnclientcert.html
func (client *Client) CreateSslVpnClientCert(request *CreateSslVpnClientCertRequest) (response *CreateSslVpnClientCertResponse, err error) {
	response = CreateCreateSslVpnClientCertResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSslVpnClientCertWithChan invokes the vpc.CreateSslVpnClientCert API asynchronously
// api document: https://help.aliyun.com/api/vpc/createsslvpnclientcert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSslVpnClientCertWithChan(request *CreateSslVpnClientCertRequest) (<-chan *CreateSslVpnClientCertResponse, <-chan error) {
	responseChan := make(chan *CreateSslVpnClientCertResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSslVpnClientCert(request)
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

// CreateSslVpnClientCertWithCallback invokes the vpc.CreateSslVpnClientCert API asynchronously
// api document: https://help.aliyun.com/api/vpc/createsslvpnclientcert.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateSslVpnClientCertWithCallback(request *CreateSslVpnClientCertRequest, callback func(response *CreateSslVpnClientCertResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSslVpnClientCertResponse
		var err error
		defer close(result)
		response, err = client.CreateSslVpnClientCert(request)
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

// CreateSslVpnClientCertRequest is the request struct for api CreateSslVpnClientCert
type CreateSslVpnClientCertRequest struct {
	*requests.RpcRequest
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	SslVpnServerId       string           `position:"Query" name:"SslVpnServerId"`
	Name                 string           `position:"Query" name:"Name"`
}

// CreateSslVpnClientCertResponse is the response struct for api CreateSslVpnClientCert
type CreateSslVpnClientCertResponse struct {
	*responses.BaseResponse
	RequestId          string `json:"RequestId" xml:"RequestId"`
	Name               string `json:"Name" xml:"Name"`
	SslVpnClientCertId string `json:"SslVpnClientCertId" xml:"SslVpnClientCertId"`
}

// CreateCreateSslVpnClientCertRequest creates a request to invoke CreateSslVpnClientCert API
func CreateCreateSslVpnClientCertRequest() (request *CreateSslVpnClientCertRequest) {
	request = &CreateSslVpnClientCertRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateSslVpnClientCert", "vpc", "openAPI")
	return
}

// CreateCreateSslVpnClientCertResponse creates a response to parse from CreateSslVpnClientCert response
func CreateCreateSslVpnClientCertResponse() (response *CreateSslVpnClientCertResponse) {
	response = &CreateSslVpnClientCertResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
