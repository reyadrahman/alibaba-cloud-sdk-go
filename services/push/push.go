package push

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

func (client *Client) Push(request *PushRequest) (response *PushResponse, err error) {
	response = CreatePushResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) PushWithChan(request *PushRequest) (<-chan *PushResponse, <-chan error) {
	responseChan := make(chan *PushResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.Push(request)
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

func (client *Client) PushWithCallback(request *PushRequest, callback func(response *PushResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushResponse
		var err error
		defer close(result)
		response, err = client.Push(request)
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

type PushRequest struct {
	*requests.RpcRequest
	AppKey                         requests.Integer `position:"Query" name:"AppKey"`
	Target                         string           `position:"Query" name:"Target"`
	TargetValue                    string           `position:"Query" name:"TargetValue"`
	PushType                       string           `position:"Query" name:"PushType"`
	DeviceType                     string           `position:"Query" name:"DeviceType"`
	Title                          string           `position:"Query" name:"Title"`
	Body                           string           `position:"Query" name:"Body"`
	SendSpeed                      requests.Integer `position:"Query" name:"SendSpeed"`
	JobKey                         string           `position:"Query" name:"JobKey"`
	PushTime                       string           `position:"Query" name:"PushTime"`
	ExpireTime                     string           `position:"Query" name:"ExpireTime"`
	StoreOffline                   requests.Boolean `position:"Query" name:"StoreOffline"`
	BatchNumber                    string           `position:"Query" name:"BatchNumber"`
	AndroidNotifyType              string           `position:"Query" name:"AndroidNotifyType"`
	AndroidOpenType                string           `position:"Query" name:"AndroidOpenType"`
	AndroidActivity                string           `position:"Query" name:"AndroidActivity"`
	AndroidOpenUrl                 string           `position:"Query" name:"AndroidOpenUrl"`
	AndroidXiaoMiActivity          string           `position:"Query" name:"AndroidXiaoMiActivity"`
	AndroidXiaoMiNotifyTitle       string           `position:"Query" name:"AndroidXiaoMiNotifyTitle"`
	AndroidXiaoMiNotifyBody        string           `position:"Query" name:"AndroidXiaoMiNotifyBody"`
	AndroidPopupActivity           string           `position:"Query" name:"AndroidPopupActivity"`
	AndroidPopupTitle              string           `position:"Query" name:"AndroidPopupTitle"`
	AndroidPopupBody               string           `position:"Query" name:"AndroidPopupBody"`
	AndroidMusic                   string           `position:"Query" name:"AndroidMusic"`
	AndroidNotificationBarType     requests.Integer `position:"Query" name:"AndroidNotificationBarType"`
	AndroidNotificationBarPriority requests.Integer `position:"Query" name:"AndroidNotificationBarPriority"`
	AndroidNotificationChannel     string           `position:"Query" name:"AndroidNotificationChannel"`
	AndroidExtParameters           string           `position:"Query" name:"AndroidExtParameters"`
	AndroidRemind                  requests.Boolean `position:"Query" name:"AndroidRemind"`
	IOSApnsEnv                     string           `position:"Query" name:"iOSApnsEnv"`
	IOSRemind                      requests.Boolean `position:"Query" name:"iOSRemind"`
	IOSRemindBody                  string           `position:"Query" name:"iOSRemindBody"`
	IOSMusic                       string           `position:"Query" name:"iOSMusic"`
	IOSBadge                       requests.Integer `position:"Query" name:"iOSBadge"`
	IOSBadgeAutoIncrement          requests.Boolean `position:"Query" name:"iOSBadgeAutoIncrement"`
	IOSSilentNotification          requests.Boolean `position:"Query" name:"iOSSilentNotification"`
	IOSSubtitle                    string           `position:"Query" name:"iOSSubtitle"`
	IOSNotificationCategory        string           `position:"Query" name:"iOSNotificationCategory"`
	IOSMutableContent              requests.Boolean `position:"Query" name:"iOSMutableContent"`
	IOSExtParameters               string           `position:"Query" name:"iOSExtParameters"`
	SmsTemplateName                string           `position:"Query" name:"SmsTemplateName"`
	SmsSignName                    string           `position:"Query" name:"SmsSignName"`
	SmsParams                      string           `position:"Query" name:"SmsParams"`
	SmsDelaySecs                   requests.Integer `position:"Query" name:"SmsDelaySecs"`
	SmsSendPolicy                  requests.Integer `position:"Query" name:"SmsSendPolicy"`
}

type PushResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	MessageId string `json:"MessageId" xml:"MessageId"`
}

func CreatePushRequest() (request *PushRequest) {
	request = &PushRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Push", "2016-08-01", "Push", "", "")
	return
}

func CreatePushResponse() (response *PushResponse) {
	response = &PushResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
