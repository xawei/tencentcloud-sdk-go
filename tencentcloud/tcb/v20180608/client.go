// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180608

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCommonServiceAPIRequest() (request *CommonServiceAPIRequest) {
    request = &CommonServiceAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CommonServiceAPI")
    return
}

func NewCommonServiceAPIResponse() (response *CommonServiceAPIResponse) {
    response = &CommonServiceAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TCB云API统一入口
func (c *Client) CommonServiceAPI(request *CommonServiceAPIRequest) (response *CommonServiceAPIResponse, err error) {
    if request == nil {
        request = NewCommonServiceAPIRequest()
    }
    response = NewCommonServiceAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMysqlInstanceRequest() (request *CreateMysqlInstanceRequest) {
    request = &CreateMysqlInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateMysqlInstance")
    return
}

func NewCreateMysqlInstanceResponse() (response *CreateMysqlInstanceResponse) {
    response = &CreateMysqlInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建mysql实例
func (c *Client) CreateMysqlInstance(request *CreateMysqlInstanceRequest) (response *CreateMysqlInstanceResponse, err error) {
    if request == nil {
        request = NewCreateMysqlInstanceRequest()
    }
    response = NewCreateMysqlInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseACLRequest() (request *DescribeDatabaseACLRequest) {
    request = &DescribeDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDatabaseACL")
    return
}

func NewDescribeDatabaseACLResponse() (response *DescribeDatabaseACLResponse) {
    response = &DescribeDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取数据库权限
func (c *Client) DescribeDatabaseACL(request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvsRequest() (request *DescribeEnvsRequest) {
    request = &DescribeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvs")
    return
}

func NewDescribeEnvsResponse() (response *DescribeEnvsResponse) {
    response = &DescribeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
func (c *Client) DescribeEnvs(request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewIsolateMysqlInstanceRequest() (request *IsolateMysqlInstanceRequest) {
    request = &IsolateMysqlInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "IsolateMysqlInstance")
    return
}

func NewIsolateMysqlInstanceResponse() (response *IsolateMysqlInstanceResponse) {
    response = &IsolateMysqlInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 隔离mysql实例
func (c *Client) IsolateMysqlInstance(request *IsolateMysqlInstanceRequest) (response *IsolateMysqlInstanceResponse, err error) {
    if request == nil {
        request = NewIsolateMysqlInstanceRequest()
    }
    response = NewIsolateMysqlInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseACLRequest() (request *ModifyDatabaseACLRequest) {
    request = &ModifyDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyDatabaseACL")
    return
}

func NewModifyDatabaseACLResponse() (response *ModifyDatabaseACLResponse) {
    response = &ModifyDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改数据库权限
func (c *Client) ModifyDatabaseACL(request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvRequest() (request *ModifyEnvRequest) {
    request = &ModifyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnv")
    return
}

func NewModifyEnvResponse() (response *ModifyEnvResponse) {
    response = &ModifyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新环境信息
func (c *Client) ModifyEnv(request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewOfflineMysqlInstanceRequest() (request *OfflineMysqlInstanceRequest) {
    request = &OfflineMysqlInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "OfflineMysqlInstance")
    return
}

func NewOfflineMysqlInstanceResponse() (response *OfflineMysqlInstanceResponse) {
    response = &OfflineMysqlInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 下线mysql实例
func (c *Client) OfflineMysqlInstance(request *OfflineMysqlInstanceRequest) (response *OfflineMysqlInstanceResponse, err error) {
    if request == nil {
        request = NewOfflineMysqlInstanceRequest()
    }
    response = NewOfflineMysqlInstanceResponse()
    err = c.Send(request, response)
    return
}

func NewUpgradeMysqlInstanceRequest() (request *UpgradeMysqlInstanceRequest) {
    request = &UpgradeMysqlInstanceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "UpgradeMysqlInstance")
    return
}

func NewUpgradeMysqlInstanceResponse() (response *UpgradeMysqlInstanceResponse) {
    response = &UpgradeMysqlInstanceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 升级mysql实例
func (c *Client) UpgradeMysqlInstance(request *UpgradeMysqlInstanceRequest) (response *UpgradeMysqlInstanceResponse, err error) {
    if request == nil {
        request = NewUpgradeMysqlInstanceRequest()
    }
    response = NewUpgradeMysqlInstanceResponse()
    err = c.Send(request, response)
    return
}
