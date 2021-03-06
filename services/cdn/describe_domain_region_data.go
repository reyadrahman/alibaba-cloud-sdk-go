package cdn

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

func (client *Client) DescribeDomainRegionData(request *DescribeDomainRegionDataRequest) (response *DescribeDomainRegionDataResponse, err error) {
	response = CreateDescribeDomainRegionDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainRegionDataWithChan(request *DescribeDomainRegionDataRequest) (<-chan *DescribeDomainRegionDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainRegionDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainRegionData(request)
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

func (client *Client) DescribeDomainRegionDataWithCallback(request *DescribeDomainRegionDataRequest, callback func(response *DescribeDomainRegionDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainRegionDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainRegionData(request)
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

type DescribeDomainRegionDataRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	StartTime     string           `position:"Query" name:"StartTime"`
	EndTime       string           `position:"Query" name:"EndTime"`
}

type DescribeDomainRegionDataResponse struct {
	*responses.BaseResponse
	RequestId    string                          `json:"RequestId" xml:"RequestId"`
	DomainName   string                          `json:"DomainName" xml:"DomainName"`
	DataInterval string                          `json:"DataInterval" xml:"DataInterval"`
	StartTime    string                          `json:"StartTime" xml:"StartTime"`
	EndTime      string                          `json:"EndTime" xml:"EndTime"`
	Value        ValueInDescribeDomainRegionData `json:"Value" xml:"Value"`
}

func CreateDescribeDomainRegionDataRequest() (request *DescribeDomainRegionDataRequest) {
	request = &DescribeDomainRegionDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainRegionData", "", "")
	return
}

func CreateDescribeDomainRegionDataResponse() (response *DescribeDomainRegionDataResponse) {
	response = &DescribeDomainRegionDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
