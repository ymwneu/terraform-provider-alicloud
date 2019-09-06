package alikafka

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

// GetAllInstanceIdList invokes the alikafka.GetAllInstanceIdList API synchronously
// api document: https://help.aliyun.com/api/alikafka/getallinstanceidlist.html
func (client *Client) GetAllInstanceIdList(request *GetAllInstanceIdListRequest) (response *GetAllInstanceIdListResponse, err error) {
	response = CreateGetAllInstanceIdListResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllInstanceIdListWithChan invokes the alikafka.GetAllInstanceIdList API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getallinstanceidlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllInstanceIdListWithChan(request *GetAllInstanceIdListRequest) (<-chan *GetAllInstanceIdListResponse, <-chan error) {
	responseChan := make(chan *GetAllInstanceIdListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllInstanceIdList(request)
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

// GetAllInstanceIdListWithCallback invokes the alikafka.GetAllInstanceIdList API asynchronously
// api document: https://help.aliyun.com/api/alikafka/getallinstanceidlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetAllInstanceIdListWithCallback(request *GetAllInstanceIdListRequest, callback func(response *GetAllInstanceIdListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllInstanceIdListResponse
		var err error
		defer close(result)
		response, err = client.GetAllInstanceIdList(request)
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

// GetAllInstanceIdListRequest is the request struct for api GetAllInstanceIdList
type GetAllInstanceIdListRequest struct {
	*requests.RpcRequest
}

// GetAllInstanceIdListResponse is the response struct for api GetAllInstanceIdList
type GetAllInstanceIdListResponse struct {
	*responses.BaseResponse
	Success     bool                   `json:"Success" xml:"Success"`
	RequestId   string                 `json:"RequestId" xml:"RequestId"`
	Code        int                    `json:"Code" xml:"Code"`
	Message     string                 `json:"Message" xml:"Message"`
	InstanceIds map[string]interface{} `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateGetAllInstanceIdListRequest creates a request to invoke GetAllInstanceIdList API
func CreateGetAllInstanceIdListRequest() (request *GetAllInstanceIdListRequest) {
	request = &GetAllInstanceIdListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2018-10-15", "GetAllInstanceIdList", "alikafka", "openAPI")
	return
}

// CreateGetAllInstanceIdListResponse creates a response to parse from GetAllInstanceIdList response
func CreateGetAllInstanceIdListResponse() (response *GetAllInstanceIdListResponse) {
	response = &GetAllInstanceIdListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
