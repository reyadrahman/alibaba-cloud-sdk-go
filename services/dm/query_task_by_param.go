package dm

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

func (client *Client) QueryTaskByParam(request *QueryTaskByParamRequest) (response *QueryTaskByParamResponse, err error) {
	response = CreateQueryTaskByParamResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) QueryTaskByParamWithChan(request *QueryTaskByParamRequest) (<-chan *QueryTaskByParamResponse, <-chan error) {
	responseChan := make(chan *QueryTaskByParamResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryTaskByParam(request)
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

func (client *Client) QueryTaskByParamWithCallback(request *QueryTaskByParamRequest, callback func(response *QueryTaskByParamResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryTaskByParamResponse
		var err error
		defer close(result)
		response, err = client.QueryTaskByParam(request)
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

type QueryTaskByParamRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	KeyWord              string           `position:"Query" name:"KeyWord"`
	Status               requests.Integer `position:"Query" name:"Status"`
}

type QueryTaskByParamResponse struct {
	*responses.BaseResponse
	RequestId  string                 `json:"RequestId" xml:"RequestId"`
	TotalCount int                    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                    `json:"PageSize" xml:"PageSize"`
	Data       DataInQueryTaskByParam `json:"data" xml:"data"`
}

func CreateQueryTaskByParamRequest() (request *QueryTaskByParamRequest) {
	request = &QueryTaskByParamRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "QueryTaskByParam", "", "")
	return
}

func CreateQueryTaskByParamResponse() (response *QueryTaskByParamResponse) {
	response = &QueryTaskByParamResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
