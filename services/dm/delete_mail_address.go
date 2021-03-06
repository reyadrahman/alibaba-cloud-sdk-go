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

func (client *Client) DeleteMailAddress(request *DeleteMailAddressRequest) (response *DeleteMailAddressResponse, err error) {
	response = CreateDeleteMailAddressResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DeleteMailAddressWithChan(request *DeleteMailAddressRequest) (<-chan *DeleteMailAddressResponse, <-chan error) {
	responseChan := make(chan *DeleteMailAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMailAddress(request)
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

func (client *Client) DeleteMailAddressWithCallback(request *DeleteMailAddressRequest, callback func(response *DeleteMailAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMailAddressResponse
		var err error
		defer close(result)
		response, err = client.DeleteMailAddress(request)
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

type DeleteMailAddressRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MailAddressId        requests.Integer `position:"Query" name:"MailAddressId"`
}

type DeleteMailAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

func CreateDeleteMailAddressRequest() (request *DeleteMailAddressRequest) {
	request = &DeleteMailAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "DeleteMailAddress", "", "")
	return
}

func CreateDeleteMailAddressResponse() (response *DeleteMailAddressResponse) {
	response = &DeleteMailAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
