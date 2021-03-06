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

// DeleteLivePullStreamInfo invokes the cdn.DeleteLivePullStreamInfo API synchronously
// api document: https://help.aliyun.com/api/cdn/deletelivepullstreaminfo.html
func (client *Client) DeleteLivePullStreamInfo(request *DeleteLivePullStreamInfoRequest) (response *DeleteLivePullStreamInfoResponse, err error) {
	response = CreateDeleteLivePullStreamInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteLivePullStreamInfoWithChan invokes the cdn.DeleteLivePullStreamInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletelivepullstreaminfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLivePullStreamInfoWithChan(request *DeleteLivePullStreamInfoRequest) (<-chan *DeleteLivePullStreamInfoResponse, <-chan error) {
	responseChan := make(chan *DeleteLivePullStreamInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteLivePullStreamInfo(request)
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

// DeleteLivePullStreamInfoWithCallback invokes the cdn.DeleteLivePullStreamInfo API asynchronously
// api document: https://help.aliyun.com/api/cdn/deletelivepullstreaminfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteLivePullStreamInfoWithCallback(request *DeleteLivePullStreamInfoRequest, callback func(response *DeleteLivePullStreamInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteLivePullStreamInfoResponse
		var err error
		defer close(result)
		response, err = client.DeleteLivePullStreamInfo(request)
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

// DeleteLivePullStreamInfoRequest is the request struct for api DeleteLivePullStreamInfo
type DeleteLivePullStreamInfoRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	AppName       string           `position:"Query" name:"AppName"`
	StreamName    string           `position:"Query" name:"StreamName"`
}

// DeleteLivePullStreamInfoResponse is the response struct for api DeleteLivePullStreamInfo
type DeleteLivePullStreamInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteLivePullStreamInfoRequest creates a request to invoke DeleteLivePullStreamInfo API
func CreateDeleteLivePullStreamInfoRequest() (request *DeleteLivePullStreamInfoRequest) {
	request = &DeleteLivePullStreamInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DeleteLivePullStreamInfo", "", "")
	return
}

// CreateDeleteLivePullStreamInfoResponse creates a response to parse from DeleteLivePullStreamInfo response
func CreateDeleteLivePullStreamInfoResponse() (response *DeleteLivePullStreamInfoResponse) {
	response = &DeleteLivePullStreamInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
