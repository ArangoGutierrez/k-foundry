# \SlurmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbV0040DeleteAccount**](SlurmdbAPI.md#SlurmdbV0040DeleteAccount) | **Delete** /slurmdb/v0.0.40/account/{account_name} | Delete account
[**SlurmdbV0040DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0040DeleteAssociation) | **Delete** /slurmdb/v0.0.40/association | Delete association
[**SlurmdbV0040DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0040DeleteAssociations) | **Delete** /slurmdb/v0.0.40/associations | Delete associations
[**SlurmdbV0040DeleteCluster**](SlurmdbAPI.md#SlurmdbV0040DeleteCluster) | **Delete** /slurmdb/v0.0.40/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0040DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0040DeleteSingleQos) | **Delete** /slurmdb/v0.0.40/qos/{qos} | Delete QOS
[**SlurmdbV0040DeleteUser**](SlurmdbAPI.md#SlurmdbV0040DeleteUser) | **Delete** /slurmdb/v0.0.40/user/{name} | Delete user
[**SlurmdbV0040DeleteWckey**](SlurmdbAPI.md#SlurmdbV0040DeleteWckey) | **Delete** /slurmdb/v0.0.40/wckey/{id} | Delete wckey
[**SlurmdbV0040GetAccount**](SlurmdbAPI.md#SlurmdbV0040GetAccount) | **Get** /slurmdb/v0.0.40/account/{account_name} | Get account info
[**SlurmdbV0040GetAccounts**](SlurmdbAPI.md#SlurmdbV0040GetAccounts) | **Get** /slurmdb/v0.0.40/accounts | Get account list
[**SlurmdbV0040GetAssociation**](SlurmdbAPI.md#SlurmdbV0040GetAssociation) | **Get** /slurmdb/v0.0.40/association | Get association info
[**SlurmdbV0040GetAssociations**](SlurmdbAPI.md#SlurmdbV0040GetAssociations) | **Get** /slurmdb/v0.0.40/associations | Get association list
[**SlurmdbV0040GetCluster**](SlurmdbAPI.md#SlurmdbV0040GetCluster) | **Get** /slurmdb/v0.0.40/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0040GetClusters**](SlurmdbAPI.md#SlurmdbV0040GetClusters) | **Get** /slurmdb/v0.0.40/clusters | Get cluster list
[**SlurmdbV0040GetConfig**](SlurmdbAPI.md#SlurmdbV0040GetConfig) | **Get** /slurmdb/v0.0.40/config | Dump all configuration information
[**SlurmdbV0040GetDiag**](SlurmdbAPI.md#SlurmdbV0040GetDiag) | **Get** /slurmdb/v0.0.40/diag | Get slurmdb diagnostics
[**SlurmdbV0040GetInstance**](SlurmdbAPI.md#SlurmdbV0040GetInstance) | **Get** /slurmdb/v0.0.40/instance | Get instance info
[**SlurmdbV0040GetInstances**](SlurmdbAPI.md#SlurmdbV0040GetInstances) | **Get** /slurmdb/v0.0.40/instances | Get instance list
[**SlurmdbV0040GetJob**](SlurmdbAPI.md#SlurmdbV0040GetJob) | **Get** /slurmdb/v0.0.40/job/{job_id} | Get job info
[**SlurmdbV0040GetJobs**](SlurmdbAPI.md#SlurmdbV0040GetJobs) | **Get** /slurmdb/v0.0.40/jobs | Get job list
[**SlurmdbV0040GetQos**](SlurmdbAPI.md#SlurmdbV0040GetQos) | **Get** /slurmdb/v0.0.40/qos | Get QOS list
[**SlurmdbV0040GetSingleQos**](SlurmdbAPI.md#SlurmdbV0040GetSingleQos) | **Get** /slurmdb/v0.0.40/qos/{qos} | Get QOS info
[**SlurmdbV0040GetTres**](SlurmdbAPI.md#SlurmdbV0040GetTres) | **Get** /slurmdb/v0.0.40/tres | Get TRES info
[**SlurmdbV0040GetUser**](SlurmdbAPI.md#SlurmdbV0040GetUser) | **Get** /slurmdb/v0.0.40/user/{name} | Get user info
[**SlurmdbV0040GetUsers**](SlurmdbAPI.md#SlurmdbV0040GetUsers) | **Get** /slurmdb/v0.0.40/users | Get user list
[**SlurmdbV0040GetWckey**](SlurmdbAPI.md#SlurmdbV0040GetWckey) | **Get** /slurmdb/v0.0.40/wckey/{id} | Get wckey info
[**SlurmdbV0040GetWckeys**](SlurmdbAPI.md#SlurmdbV0040GetWckeys) | **Get** /slurmdb/v0.0.40/wckeys | Get wckey list
[**SlurmdbV0040PostAccounts**](SlurmdbAPI.md#SlurmdbV0040PostAccounts) | **Post** /slurmdb/v0.0.40/accounts | Update accounts
[**SlurmdbV0040PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0040PostAccountsAssociation) | **Post** /slurmdb/v0.0.40/accounts_association | Add accounts with conditional association
[**SlurmdbV0040PostAssociations**](SlurmdbAPI.md#SlurmdbV0040PostAssociations) | **Post** /slurmdb/v0.0.40/associations | Set associations info
[**SlurmdbV0040PostClusters**](SlurmdbAPI.md#SlurmdbV0040PostClusters) | **Post** /slurmdb/v0.0.40/clusters | update clusters
[**SlurmdbV0040PostConfig**](SlurmdbAPI.md#SlurmdbV0040PostConfig) | **Post** /slurmdb/v0.0.40/config | Load all configuration information
[**SlurmdbV0040PostQos**](SlurmdbAPI.md#SlurmdbV0040PostQos) | **Post** /slurmdb/v0.0.40/qos | Set QOS info
[**SlurmdbV0040PostTres**](SlurmdbAPI.md#SlurmdbV0040PostTres) | **Post** /slurmdb/v0.0.40/tres | Set TRES info
[**SlurmdbV0040PostUsers**](SlurmdbAPI.md#SlurmdbV0040PostUsers) | **Post** /slurmdb/v0.0.40/users | Update user
[**SlurmdbV0040PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0040PostUsersAssociation) | **Post** /slurmdb/v0.0.40/users_association | Add users with conditional association
[**SlurmdbV0040PostWckeys**](SlurmdbAPI.md#SlurmdbV0040PostWckeys) | **Post** /slurmdb/v0.0.40/wckeys | Add wckeys



## SlurmdbV0040DeleteAccount

> V0040OpenapiAccountsRemovedResp SlurmdbV0040DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAccount`: V0040OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiAccountsRemovedResp**](V0040OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteAssociation

> V0040OpenapiAssocsRemovedResp SlurmdbV0040DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | CSV format list (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | usage end UNIX timestamp (optional)
	usageStart := "usageStart_example" // string | usage start UNIX timestamp (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | fill in usage (optional)
	withDeleted := "withDeleted_example" // string | return deleted associations (optional)
	withRawQos := "withRawQos_example" // string | return a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | return sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | don't give me parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | don't give me limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAssociation`: V0040OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | CSV format list | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | usage end UNIX timestamp | 
 **usageStart** | **string** | usage start UNIX timestamp | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | fill in usage | 
 **withDeleted** | **string** | return deleted associations | 
 **withRawQos** | **string** | return a raw qos or delta_qos | 
 **withSubAccts** | **string** | return sub acct information also | 
 **withoutParentInfo** | **string** | don&#39;t give me parent id/name | 
 **withoutParentLimits** | **string** | don&#39;t give me limits from parents | 

### Return type

[**V0040OpenapiAssocsRemovedResp**](V0040OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteAssociations

> V0040OpenapiAssocsRemovedResp SlurmdbV0040DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | CSV format list (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | usage end UNIX timestamp (optional)
	usageStart := "usageStart_example" // string | usage start UNIX timestamp (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | fill in usage (optional)
	withDeleted := "withDeleted_example" // string | return deleted associations (optional)
	withRawQos := "withRawQos_example" // string | return a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | return sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | don't give me parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | don't give me limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteAssociations`: V0040OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | CSV format list | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | usage end UNIX timestamp | 
 **usageStart** | **string** | usage start UNIX timestamp | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | fill in usage | 
 **withDeleted** | **string** | return deleted associations | 
 **withRawQos** | **string** | return a raw qos or delta_qos | 
 **withSubAccts** | **string** | return sub acct information also | 
 **withoutParentInfo** | **string** | don&#39;t give me parent id/name | 
 **withoutParentLimits** | **string** | don&#39;t give me limits from parents | 

### Return type

[**V0040OpenapiAssocsRemovedResp**](V0040OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteCluster

> V0040OpenapiClustersRemovedResp SlurmdbV0040DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string |  (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string |  (optional)
	format := "format_example" // string | CSV format list (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end UNIX timestamp (seconds) (optional)
	usageStart := "usageStart_example" // string | Usage start UNIX timestamp (seconds) (optional)
	withDeleted := "withDeleted_example" // string | include deleted clusters (optional)
	withUsage := "withUsage_example" // string | query usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteCluster`: V0040OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** |  | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** |  | 
 **format** | **string** | CSV format list | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end UNIX timestamp (seconds) | 
 **usageStart** | **string** | Usage start UNIX timestamp (seconds) | 
 **withDeleted** | **string** | include deleted clusters | 
 **withUsage** | **string** | query usage | 

### Return type

[**V0040OpenapiClustersRemovedResp**](V0040OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteSingleQos

> V0040OpenapiSlurmdbdQosRemovedResp SlurmdbV0040DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteSingleQos`: V0040OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiSlurmdbdQosRemovedResp**](V0040OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteUser

> V0040OpenapiResp SlurmdbV0040DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteUser`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040DeleteWckey

> V0040OpenapiWckeyRemovedResp SlurmdbV0040DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040DeleteWckey`: V0040OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiWckeyRemovedResp**](V0040OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAccount

> V0040OpenapiAccountsResp SlurmdbV0040GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | include associations (optional)
	withCoords := "withCoords_example" // string | include coordinators (optional)
	withDeleted := "withDeleted_example" // string | include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAccount`: V0040OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | include associations | 
 **withCoords** | **string** | include coordinators | 
 **withDeleted** | **string** | include deleted | 

### Return type

[**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAccounts

> V0040OpenapiAccountsResp SlurmdbV0040GetAccounts(ctx).Description(description).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	withAssocs := "withAssocs_example" // string | include associations (optional)
	withCoords := "withCoords_example" // string | include coordinators (optional)
	withDeleted := "withDeleted_example" // string | include deleted accounts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAccounts(context.Background()).Description(description).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAccounts`: V0040OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **withAssocs** | **string** | include associations | 
 **withCoords** | **string** | include coordinators | 
 **withDeleted** | **string** | include deleted accounts | 

### Return type

[**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAssociation

> V0040OpenapiAssocsResp SlurmdbV0040GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | CSV format list (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | usage end UNIX timestamp (optional)
	usageStart := "usageStart_example" // string | usage start UNIX timestamp (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | fill in usage (optional)
	withDeleted := "withDeleted_example" // string | return deleted associations (optional)
	withRawQos := "withRawQos_example" // string | return a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | return sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | don't give me parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | don't give me limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAssociation`: V0040OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | CSV format list | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | usage end UNIX timestamp | 
 **usageStart** | **string** | usage start UNIX timestamp | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | fill in usage | 
 **withDeleted** | **string** | return deleted associations | 
 **withRawQos** | **string** | return a raw qos or delta_qos | 
 **withSubAccts** | **string** | return sub acct information also | 
 **withoutParentInfo** | **string** | don&#39;t give me parent id/name | 
 **withoutParentLimits** | **string** | don&#39;t give me limits from parents | 

### Return type

[**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetAssociations

> V0040OpenapiAssocsResp SlurmdbV0040GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	format := "format_example" // string | CSV format list (optional)
	id := "id_example" // string | CSV id list (optional)
	onlyDefaults := "onlyDefaults_example" // string | filter to only defaults (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | usage end UNIX timestamp (optional)
	usageStart := "usageStart_example" // string | usage start UNIX timestamp (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | fill in usage (optional)
	withDeleted := "withDeleted_example" // string | return deleted associations (optional)
	withRawQos := "withRawQos_example" // string | return a raw qos or delta_qos (optional)
	withSubAccts := "withSubAccts_example" // string | return sub acct information also (optional)
	withoutParentInfo := "withoutParentInfo_example" // string | don't give me parent id/name (optional)
	withoutParentLimits := "withoutParentLimits_example" // string | don't give me limits from parents (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).Format(format).Id(id).OnlyDefaults(onlyDefaults).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).WithRawQos(withRawQos).WithSubAccts(withSubAccts).WithoutParentInfo(withoutParentInfo).WithoutParentLimits(withoutParentLimits).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetAssociations`: V0040OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **format** | **string** | CSV format list | 
 **id** | **string** | CSV id list | 
 **onlyDefaults** | **string** | filter to only defaults | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | usage end UNIX timestamp | 
 **usageStart** | **string** | usage start UNIX timestamp | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | fill in usage | 
 **withDeleted** | **string** | return deleted associations | 
 **withRawQos** | **string** | return a raw qos or delta_qos | 
 **withSubAccts** | **string** | return sub acct information also | 
 **withoutParentInfo** | **string** | don&#39;t give me parent id/name | 
 **withoutParentLimits** | **string** | don&#39;t give me limits from parents | 

### Return type

[**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetCluster

> V0040OpenapiClustersResp SlurmdbV0040GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string |  (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string |  (optional)
	format := "format_example" // string | CSV format list (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end UNIX timestamp (seconds) (optional)
	usageStart := "usageStart_example" // string | Usage start UNIX timestamp (seconds) (optional)
	withDeleted := "withDeleted_example" // string | include deleted clusters (optional)
	withUsage := "withUsage_example" // string | query usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetCluster`: V0040OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** |  | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** |  | 
 **format** | **string** | CSV format list | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end UNIX timestamp (seconds) | 
 **usageStart** | **string** | Usage start UNIX timestamp (seconds) | 
 **withDeleted** | **string** | include deleted clusters | 
 **withUsage** | **string** | query usage | 

### Return type

[**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetClusters

> V0040OpenapiClustersResp SlurmdbV0040GetClusters(ctx).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	classification := "classification_example" // string |  (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string |  (optional)
	format := "format_example" // string | CSV format list (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end UNIX timestamp (seconds) (optional)
	usageStart := "usageStart_example" // string | Usage start UNIX timestamp (seconds) (optional)
	withDeleted := "withDeleted_example" // string | include deleted clusters (optional)
	withUsage := "withUsage_example" // string | query usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetClusters(context.Background()).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetClusters`: V0040OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classification** | **string** |  | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** |  | 
 **format** | **string** | CSV format list | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end UNIX timestamp (seconds) | 
 **usageStart** | **string** | Usage start UNIX timestamp (seconds) | 
 **withDeleted** | **string** | include deleted clusters | 
 **withUsage** | **string** | query usage | 

### Return type

[**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetConfig

> V0040OpenapiSlurmdbdConfigResp SlurmdbV0040GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetConfig`: V0040OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetConfigRequest struct via the builder pattern


### Return type

[**V0040OpenapiSlurmdbdConfigResp**](V0040OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetDiag

> V0040OpenapiSlurmdbdStatsResp SlurmdbV0040GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetDiag`: V0040OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetDiagRequest struct via the builder pattern


### Return type

[**V0040OpenapiSlurmdbdStatsResp**](V0040OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetInstance

> V0040OpenapiInstancesResp SlurmdbV0040GetInstance(ctx).Cluster(cluster).Extra(extra).InstanceId(instanceId).InstanceType(instanceType).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | Cluster name (optional)
	extra := "extra_example" // string | Arbitrary string (optional)
	instanceId := "instanceId_example" // string | Cloud instance ID (optional)
	instanceType := "instanceType_example" // string | Cloud instance type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetInstance(context.Background()).Cluster(cluster).Extra(extra).InstanceId(instanceId).InstanceType(instanceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetInstance`: V0040OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **extra** | **string** | Arbitrary string | 
 **instanceId** | **string** | Cloud instance ID | 
 **instanceType** | **string** | Cloud instance type | 

### Return type

[**V0040OpenapiInstancesResp**](V0040OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetInstances

> V0040OpenapiInstancesResp SlurmdbV0040GetInstances(ctx).Cluster(cluster).Extra(extra).InstanceId(instanceId).InstanceType(instanceType).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | Cluster name (optional)
	extra := "extra_example" // string | Arbitrary string (optional)
	instanceId := "instanceId_example" // string | Cloud instance ID (optional)
	instanceType := "instanceType_example" // string | Cloud instance type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetInstances(context.Background()).Cluster(cluster).Extra(extra).InstanceId(instanceId).InstanceType(instanceType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetInstances`: V0040OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | Cluster name | 
 **extra** | **string** | Arbitrary string | 
 **instanceId** | **string** | Cloud instance ID | 
 **instanceType** | **string** | Cloud instance type | 

### Return type

[**V0040OpenapiInstancesResp**](V0040OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetJob

> V0040OpenapiSlurmdbdJobsResp SlurmdbV0040GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetJob`: V0040OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiSlurmdbdJobsResp**](V0040OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetJobs

> V0040OpenapiSlurmdbdJobsResp SlurmdbV0040GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).TimelimitMax(timelimitMax).TimelimitMin(timelimitMin).EndTime(endTime).StartTime(startTime).SubmitTime(submitTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	cpusMax := "cpusMax_example" // string | number of cpus high range (optional)
	cpusMin := "cpusMin_example" // string | number of cpus low range (optional)
	schedulerUnset := "schedulerUnset_example" // string |  (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string |  (optional)
	scheduledByMain := "scheduledByMain_example" // string |  (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string |  (optional)
	jobStarted := "jobStarted_example" // string |  (optional)
	exitCode := "exitCode_example" // string | job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string |  (optional)
	skipSteps := "skipSteps_example" // string |  (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string |  (optional)
	wholeHetjob := "wholeHetjob_example" // string |  (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string |  (optional)
	disableWaitForResult := "disableWaitForResult_example" // string |  (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string |  (optional)
	showBatchScript := "showBatchScript_example" // string |  (optional)
	showJobEnvironment := "showJobEnvironment_example" // string |  (optional)
	format := "format_example" // string | CSV format list (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	nodesMax := "nodesMax_example" // string | number of nodes high range (optional)
	nodesMin := "nodesMin_example" // string | number of nodes low range (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	timelimitMax := "timelimitMax_example" // string | maximum timelimit (seconds) (optional)
	timelimitMin := "timelimitMin_example" // string | minimum timelimit (seconds) (optional)
	endTime := "endTime_example" // string | usage end timestamp (optional)
	startTime := "startTime_example" // string | usage start timestamp (optional)
	submitTime := "submitTime_example" // string | submit time timestamp (optional)
	node := "node_example" // string | ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV wckey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).CpusMax(cpusMax).CpusMin(cpusMin).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).NodesMax(nodesMax).NodesMin(nodesMin).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).TimelimitMax(timelimitMax).TimelimitMin(timelimitMin).EndTime(endTime).StartTime(startTime).SubmitTime(submitTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetJobs`: V0040OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **cpusMax** | **string** | number of cpus high range | 
 **cpusMin** | **string** | number of cpus low range | 
 **schedulerUnset** | **string** |  | 
 **scheduledOnSubmit** | **string** |  | 
 **scheduledByMain** | **string** |  | 
 **scheduledByBackfill** | **string** |  | 
 **jobStarted** | **string** |  | 
 **exitCode** | **string** | job exit code (numeric) | 
 **showDuplicates** | **string** |  | 
 **skipSteps** | **string** |  | 
 **disableTruncateUsageTime** | **string** |  | 
 **wholeHetjob** | **string** |  | 
 **disableWholeHetjob** | **string** |  | 
 **disableWaitForResult** | **string** |  | 
 **usageTimeAsSubmitTime** | **string** |  | 
 **showBatchScript** | **string** |  | 
 **showJobEnvironment** | **string** |  | 
 **format** | **string** | CSV format list | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **nodesMax** | **string** | number of nodes high range | 
 **nodesMin** | **string** | number of nodes low range | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **timelimitMax** | **string** | maximum timelimit (seconds) | 
 **timelimitMin** | **string** | minimum timelimit (seconds) | 
 **endTime** | **string** | usage end timestamp | 
 **startTime** | **string** | usage start timestamp | 
 **submitTime** | **string** | submit time timestamp | 
 **node** | **string** | ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV wckey list | 

### Return type

[**V0040OpenapiSlurmdbdJobsResp**](V0040OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetQos

> V0040OpenapiSlurmdbdQosResp SlurmdbV0040GetQos(ctx).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | CSV format list (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string |  (optional)
	withDeleted := "withDeleted_example" // string | Include deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetQos(context.Background()).Description(description).Id(id).Format(format).Name(name).PreemptMode(preemptMode).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetQos`: V0040OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | CSV format list | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** |  | 
 **withDeleted** | **string** | Include deleted QOS | 

### Return type

[**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetSingleQos

> V0040OpenapiSlurmdbdQosResp SlurmdbV0040GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetSingleQos`: V0040OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetTres

> V0040OpenapiTresResp SlurmdbV0040GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetTres`: V0040OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetTresRequest struct via the builder pattern


### Return type

[**V0040OpenapiTresResp**](V0040OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetUser

> V0040OpenapiUsersResp SlurmdbV0040GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include assocations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetUser`: V0040OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include assocations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include wckeys | 

### Return type

[**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetUsers

> V0040OpenapiUsersResp SlurmdbV0040GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default wckey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With wckeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetUsers`: V0040OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default wckey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With wckeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetWckey

> V0040OpenapiWckeyResp SlurmdbV0040GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | wckey id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetWckey`: V0040OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | wckey id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040GetWckeys

> V0040OpenapiWckeyResp SlurmdbV0040GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | CSV format name list (optional)
	id := "id_example" // string | CSV id list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | only query defaults (optional)
	usageEnd := "usageEnd_example" // string | usage end UNIX timestamp (seconds) (optional)
	usageStart := "usageStart_example" // string | usage start UNIX timestamp (seconds) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | include usage with query (optional)
	withDeleted := "withDeleted_example" // string | include deleted wckeys with query (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040GetWckeys`: V0040OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | CSV format name list | 
 **id** | **string** | CSV id list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | only query defaults | 
 **usageEnd** | **string** | usage end UNIX timestamp (seconds) | 
 **usageStart** | **string** | usage start UNIX timestamp (seconds) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | include usage with query | 
 **withDeleted** | **string** | include deleted wckeys with query | 

### Return type

[**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAccounts

> V0040OpenapiResp SlurmdbV0040PostAccounts(ctx).V0040OpenapiAccountsResp(v0040OpenapiAccountsResp).Execute()

Update accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiAccountsResp := *openapiclient.NewV0040OpenapiAccountsResp([]openapiclient.V0040Account{*openapiclient.NewV0040Account("Description_example", "Name_example", "Organization_example")}) // V0040OpenapiAccountsResp | update/create accounts

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAccounts(context.Background()).V0040OpenapiAccountsResp(v0040OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAccounts`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAccountsResp** | [**V0040OpenapiAccountsResp**](V0040OpenapiAccountsResp.md) | update/create accounts | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAccountsAssociation

> V0040OpenapiAccountsAddCondRespStr SlurmdbV0040PostAccountsAssociation(ctx).V0040OpenapiAccountsAddCondResp(v0040OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiAccountsAddCondResp := *openapiclient.NewV0040OpenapiAccountsAddCondResp() // V0040OpenapiAccountsAddCondResp | Create accounts with conditional association

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAccountsAssociation(context.Background()).V0040OpenapiAccountsAddCondResp(v0040OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAccountsAssociation`: V0040OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAccountsAddCondResp** | [**V0040OpenapiAccountsAddCondResp**](V0040OpenapiAccountsAddCondResp.md) | Create accounts with conditional association | 

### Return type

[**V0040OpenapiAccountsAddCondRespStr**](V0040OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostAssociations

> V0040OpenapiResp SlurmdbV0040PostAssociations(ctx).V0040OpenapiAssocsResp(v0040OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiAssocsResp := *openapiclient.NewV0040OpenapiAssocsResp([]openapiclient.V0040Assoc{*openapiclient.NewV0040Assoc("User_example")}) // V0040OpenapiAssocsResp | Add or update associations

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostAssociations(context.Background()).V0040OpenapiAssocsResp(v0040OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostAssociations`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiAssocsResp** | [**V0040OpenapiAssocsResp**](V0040OpenapiAssocsResp.md) | Add or update associations | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostClusters

> V0040OpenapiResp SlurmdbV0040PostClusters(ctx).V0040OpenapiClustersResp(v0040OpenapiClustersResp).Execute()

update clusters

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiClustersResp := *openapiclient.NewV0040OpenapiClustersResp([]openapiclient.V0040ClusterRec{*openapiclient.NewV0040ClusterRec()}) // V0040OpenapiClustersResp | Add or modify clusters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostClusters(context.Background()).V0040OpenapiClustersResp(v0040OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostClusters`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiClustersResp** | [**V0040OpenapiClustersResp**](V0040OpenapiClustersResp.md) | Add or modify clusters | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostConfig

> V0040OpenapiResp SlurmdbV0040PostConfig(ctx).V0040OpenapiSlurmdbdConfigResp(v0040OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiSlurmdbdConfigResp := *openapiclient.NewV0040OpenapiSlurmdbdConfigResp() // V0040OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostConfig(context.Background()).V0040OpenapiSlurmdbdConfigResp(v0040OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostConfig`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiSlurmdbdConfigResp** | [**V0040OpenapiSlurmdbdConfigResp**](V0040OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostQos

> V0040OpenapiResp SlurmdbV0040PostQos(ctx).V0040OpenapiSlurmdbdQosResp(v0040OpenapiSlurmdbdQosResp).Execute()

Set QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiSlurmdbdQosResp := *openapiclient.NewV0040OpenapiSlurmdbdQosResp([]openapiclient.V0040Qos{*openapiclient.NewV0040Qos()}) // V0040OpenapiSlurmdbdQosResp | Add or update QOSs

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostQos(context.Background()).V0040OpenapiSlurmdbdQosResp(v0040OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostQos`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiSlurmdbdQosResp** | [**V0040OpenapiSlurmdbdQosResp**](V0040OpenapiSlurmdbdQosResp.md) | Add or update QOSs | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostTres

> V0040OpenapiResp SlurmdbV0040PostTres(ctx).V0040OpenapiTresResp(v0040OpenapiTresResp).Execute()

Set TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiTresResp := *openapiclient.NewV0040OpenapiTresResp([]openapiclient.V0040Tres{*openapiclient.NewV0040Tres("Type_example")}) // V0040OpenapiTresResp | Add or Update TRES

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostTres(context.Background()).V0040OpenapiTresResp(v0040OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostTres`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiTresResp** | [**V0040OpenapiTresResp**](V0040OpenapiTresResp.md) | Add or Update TRES | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostUsers

> V0040OpenapiResp SlurmdbV0040PostUsers(ctx).V0040OpenapiUsersResp(v0040OpenapiUsersResp).Execute()

Update user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiUsersResp := *openapiclient.NewV0040OpenapiUsersResp([]openapiclient.V0040User{*openapiclient.NewV0040User("Name_example")}) // V0040OpenapiUsersResp | add or update user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostUsers(context.Background()).V0040OpenapiUsersResp(v0040OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostUsers`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiUsersResp** | [**V0040OpenapiUsersResp**](V0040OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostUsersAssociation

> V0040OpenapiUsersAddCondRespStr SlurmdbV0040PostUsersAssociation(ctx).V0040OpenapiUsersAddCondResp(v0040OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiUsersAddCondResp := *openapiclient.NewV0040OpenapiUsersAddCondResp() // V0040OpenapiUsersAddCondResp | Create users with conditional association

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostUsersAssociation(context.Background()).V0040OpenapiUsersAddCondResp(v0040OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostUsersAssociation`: V0040OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiUsersAddCondResp** | [**V0040OpenapiUsersAddCondResp**](V0040OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0040OpenapiUsersAddCondRespStr**](V0040OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0040PostWckeys

> V0040OpenapiResp SlurmdbV0040PostWckeys(ctx).V0040OpenapiWckeyResp(v0040OpenapiWckeyResp).Execute()

Add wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0040OpenapiWckeyResp := *openapiclient.NewV0040OpenapiWckeyResp([]openapiclient.V0040Wckey{*openapiclient.NewV0040Wckey("Cluster_example", "Name_example", "User_example")}) // V0040OpenapiWckeyResp | add wckeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0040PostWckeys(context.Background()).V0040OpenapiWckeyResp(v0040OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0040PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0040PostWckeys`: V0040OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0040PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0040PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0040OpenapiWckeyResp** | [**V0040OpenapiWckeyResp**](V0040OpenapiWckeyResp.md) | add wckeys | 

### Return type

[**V0040OpenapiResp**](V0040OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

