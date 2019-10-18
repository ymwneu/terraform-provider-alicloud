package bssopenapi

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

// CreateResellerUserQuota invokes the bssopenapi.CreateResellerUserQuota API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/createreselleruserquota.html
func (client *Client) CreateResellerUserQuota(request *CreateResellerUserQuotaRequest) (response *CreateResellerUserQuotaResponse, err error) {
	response = CreateCreateResellerUserQuotaResponse()
	err = client.DoAction(request, response)
	return
}

// CreateResellerUserQuotaWithChan invokes the bssopenapi.CreateResellerUserQuota API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/createreselleruserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateResellerUserQuotaWithChan(request *CreateResellerUserQuotaRequest) (<-chan *CreateResellerUserQuotaResponse, <-chan error) {
	responseChan := make(chan *CreateResellerUserQuotaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateResellerUserQuota(request)
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

// CreateResellerUserQuotaWithCallback invokes the bssopenapi.CreateResellerUserQuota API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/createreselleruserquota.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateResellerUserQuotaWithCallback(request *CreateResellerUserQuotaRequest, callback func(response *CreateResellerUserQuotaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateResellerUserQuotaResponse
		var err error
		defer close(result)
		response, err = client.CreateResellerUserQuota(request)
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

// CreateResellerUserQuotaRequest is the request struct for api CreateResellerUserQuota
type CreateResellerUserQuotaRequest struct {
	*requests.RpcRequest
	Amount   string           `position:"Query" name:"Amount"`
	OutBizId string           `position:"Query" name:"OutBizId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
	Currency string           `position:"Query" name:"Currency"`
}

// CreateResellerUserQuotaResponse is the response struct for api CreateResellerUserQuota
type CreateResellerUserQuotaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateCreateResellerUserQuotaRequest creates a request to invoke CreateResellerUserQuota API
func CreateCreateResellerUserQuotaRequest() (request *CreateResellerUserQuotaRequest) {
	request = &CreateResellerUserQuotaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "CreateResellerUserQuota", "bssopenapi", "openAPI")
	return
}

// CreateCreateResellerUserQuotaResponse creates a response to parse from CreateResellerUserQuota response
func CreateCreateResellerUserQuotaResponse() (response *CreateResellerUserQuotaResponse) {
	response = &CreateResellerUserQuotaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
