package mts

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

// ReportFacerecogJobResult invokes the mts.ReportFacerecogJobResult API synchronously
// api document: https://help.aliyun.com/api/mts/reportfacerecogjobresult.html
func (client *Client) ReportFacerecogJobResult(request *ReportFacerecogJobResultRequest) (response *ReportFacerecogJobResultResponse, err error) {
	response = CreateReportFacerecogJobResultResponse()
	err = client.DoAction(request, response)
	return
}

// ReportFacerecogJobResultWithChan invokes the mts.ReportFacerecogJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportfacerecogjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportFacerecogJobResultWithChan(request *ReportFacerecogJobResultRequest) (<-chan *ReportFacerecogJobResultResponse, <-chan error) {
	responseChan := make(chan *ReportFacerecogJobResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReportFacerecogJobResult(request)
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

// ReportFacerecogJobResultWithCallback invokes the mts.ReportFacerecogJobResult API asynchronously
// api document: https://help.aliyun.com/api/mts/reportfacerecogjobresult.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReportFacerecogJobResultWithCallback(request *ReportFacerecogJobResultRequest, callback func(response *ReportFacerecogJobResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReportFacerecogJobResultResponse
		var err error
		defer close(result)
		response, err = client.ReportFacerecogJobResult(request)
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

// ReportFacerecogJobResultRequest is the request struct for api ReportFacerecogJobResult
type ReportFacerecogJobResultRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
	Facerecog            string           `position:"Query" name:"Facerecog"`
	Details              string           `position:"Query" name:"Details"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

// ReportFacerecogJobResultResponse is the response struct for api ReportFacerecogJobResult
type ReportFacerecogJobResultResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	JobId     string `json:"JobId" xml:"JobId"`
}

// CreateReportFacerecogJobResultRequest creates a request to invoke ReportFacerecogJobResult API
func CreateReportFacerecogJobResultRequest() (request *ReportFacerecogJobResultRequest) {
	request = &ReportFacerecogJobResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "ReportFacerecogJobResult", "mts", "openAPI")
	return
}

// CreateReportFacerecogJobResultResponse creates a response to parse from ReportFacerecogJobResult response
func CreateReportFacerecogJobResultResponse() (response *ReportFacerecogJobResultResponse) {
	response = &ReportFacerecogJobResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
