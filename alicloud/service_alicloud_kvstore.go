package alicloud

import (
	"time"

	"github.com/hashicorp/terraform/helper/resource"

	r_kvstore "github.com/aliyun/alibaba-cloud-sdk-go/services/r-kvstore"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

type KvstoreService struct {
	client *connectivity.AliyunClient
}

func (s *KvstoreService) DescribeKVstoreInstance(id string) (instance *r_kvstore.DBInstanceAttribute, err error) {
	request := r_kvstore.CreateDescribeInstanceAttributeRequest()
	request.RegionId = s.client.RegionId
	request.InstanceId = id
	raw, err := s.client.WithRkvClient(func(rkvClient *r_kvstore.Client) (interface{}, error) {
		return rkvClient.DescribeInstanceAttribute(request)
	})
	if err != nil {
		if IsExceptedError(err, InvalidKVStoreInstanceIdNotFound) {
			return nil, WrapErrorf(Error(GetNotFoundMessage("KVstoreInstance", id)), NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return nil, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*r_kvstore.DescribeInstanceAttributeResponse)
	if len(response.Instances.DBInstanceAttribute) <= 0 {
		return nil, WrapErrorf(Error(GetNotFoundMessage("KVstoreInstance", id)), NotFoundMsg, AlibabaCloudSdkGoERROR)
	}

	return &response.Instances.DBInstanceAttribute[0], nil
}

func (s *KvstoreService) DescribeKVstoreBackupPolicy(id string) (response *r_kvstore.DescribeBackupPolicyResponse, err error) {
	request := r_kvstore.CreateDescribeBackupPolicyRequest()
	request.RegionId = s.client.RegionId
	request.InstanceId = id
	raw, err := s.client.WithRkvClient(func(rkvClient *r_kvstore.Client) (interface{}, error) {
		return rkvClient.DescribeBackupPolicy(request)
	})
	if err != nil {
		if IsExceptedError(err, InvalidKVStoreInstanceIdNotFound) {
			return nil, WrapErrorf(Error(GetNotFoundMessage("KVstoreBackupPolicy", id)), NotFoundMsg, AlibabaCloudSdkGoERROR)
		}
		return nil, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ = raw.(*r_kvstore.DescribeBackupPolicyResponse)
	return
}

func (s *KvstoreService) WaitForKVstoreInstance(id string, status Status, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		object, err := s.DescribeKVstoreInstance(id)
		if err != nil {
			if NotFoundError(err) {
				if status == Deleted {
					return nil
				}
			} else {
				return WrapError(err)
			}
		}
		if object.InstanceStatus == string(status) {
			break
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.InstanceStatus, status, ProviderERROR)
		}
	}
	return nil
}

func (s *KvstoreService) RdsKvstoreInstanceStateRefreshFunc(id string, failStates []string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		object, err := s.DescribeKVstoreInstance(id)
		if err != nil {
			if NotFoundError(err) {
				// Set this to nil as if we didn't find anything.
				return nil, "", nil
			}
			return nil, "", WrapError(err)
		}

		for _, failState := range failStates {
			if object.InstanceStatus == failState {
				return object, object.InstanceStatus, WrapError(Error(FailedToReachTargetStatus, object.InstanceStatus))
			}
		}
		return object, object.InstanceStatus, nil
	}
}

func (s *KvstoreService) WaitForKVstoreInstanceVpcAuthMode(id string, status string, timeout int) error {
	deadline := time.Now().Add(time.Duration(timeout) * time.Second)
	for {
		object, err := s.DescribeKVstoreInstance(id)
		if err != nil && !NotFoundError(err) {
			return err
		}
		if object.VpcAuthMode == string(status) {
			break
		}
		if time.Now().After(deadline) {
			return WrapErrorf(err, WaitTimeoutMsg, id, GetFunc(1), timeout, object.VpcAuthMode, status, ProviderERROR)
		}
	}
	return nil
}

func (s *KvstoreService) DescribeParameters(id string) (ds *r_kvstore.DescribeParametersResponse, err error) {
	request := r_kvstore.CreateDescribeParametersRequest()
	request.RegionId = s.client.RegionId
	request.DBInstanceId = id

	raw, err := s.client.WithRkvClient(func(rkvClient *r_kvstore.Client) (interface{}, error) {
		return rkvClient.DescribeParameters(request)
	})
	if err != nil {
		if IsExceptedErrors(err, []string{InvalidDBInstanceIdNotFound, InvalidDBInstanceNameNotFound}) {
			return nil, WrapErrorf(Error(GetNotFoundMessage("Parameters", id)), NotFoundMsg, ProviderERROR)
		}
		return nil, WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	response, _ := raw.(*r_kvstore.DescribeParametersResponse)
	return response, err
}

func (s *KvstoreService) ModifyInstanceConfig(id string, config string) error {
	request := r_kvstore.CreateModifyInstanceConfigRequest()
	request.RegionId = s.client.RegionId
	request.InstanceId = id
	request.Config = config

	if err := s.WaitForKVstoreInstance(id, Normal, DefaultLongTimeout); err != nil {
		return WrapError(err)
	}
	raw, err := s.client.WithRkvClient(func(rkvClient *r_kvstore.Client) (interface{}, error) {
		return rkvClient.ModifyInstanceConfig(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, id, request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw, request.RpcRequest, request)
	return nil
}
