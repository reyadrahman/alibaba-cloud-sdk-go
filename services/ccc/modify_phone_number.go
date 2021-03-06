package ccc

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

func (client *Client) ModifyPhoneNumber(request *ModifyPhoneNumberRequest) (response *ModifyPhoneNumberResponse, err error) {
	response = CreateModifyPhoneNumberResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ModifyPhoneNumberWithChan(request *ModifyPhoneNumberRequest) (<-chan *ModifyPhoneNumberResponse, <-chan error) {
	responseChan := make(chan *ModifyPhoneNumberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPhoneNumber(request)
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

func (client *Client) ModifyPhoneNumberWithCallback(request *ModifyPhoneNumberRequest, callback func(response *ModifyPhoneNumberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPhoneNumberResponse
		var err error
		defer close(result)
		response, err = client.ModifyPhoneNumber(request)
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

type ModifyPhoneNumberRequest struct {
	*requests.RpcRequest
	InstanceId    string `position:"Query" name:"InstanceId"`
	PhoneNumberId string `position:"Query" name:"PhoneNumberId"`
	Usage         string `position:"Query" name:"Usage"`
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
}

type ModifyPhoneNumberResponse struct {
	*responses.BaseResponse
	RequestId      string      `json:"RequestId" xml:"RequestId"`
	Success        bool        `json:"Success" xml:"Success"`
	Code           string      `json:"Code" xml:"Code"`
	Message        string      `json:"Message" xml:"Message"`
	HttpStatusCode int         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PhoneNumber    PhoneNumber `json:"PhoneNumber" xml:"PhoneNumber"`
}

func CreateModifyPhoneNumberRequest() (request *ModifyPhoneNumberRequest) {
	request = &ModifyPhoneNumberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifyPhoneNumber", "", "")
	return
}

func CreateModifyPhoneNumberResponse() (response *ModifyPhoneNumberResponse) {
	response = &ModifyPhoneNumberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
