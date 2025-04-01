# 数据与智能专区

## Models

### `AsyncProgramResult` 客服账号

 Name              | JSON                         | Type     | Doc                                                      
:------------------|:-----------------------------|:---------|:---------------------------------------------------------
 `ResponseErrCode`        | `response_errcode`                  | `string` | 上报异步任务结果中上报的errcode。代表专区程序返回的错误码
 `ResponseData`            | `response_data`                       | `string` | 上报异步任务结果中上报的result。代表专区程序的输出结果，为自定义的JSON字符串，要求与管理端配置的输出协议格式匹配

```go
// DebugModeStatusType 程序当前的调试模式状态
//
// 1为关闭
// 2为开启
type DebugModeStatusType int

const (
	// ExternalUserTypeWeChat 微信用户
	DebugModeStatusTypeOff DebugModeStatusType = 1
	// ExternalUserTypeWorkWeChat 企业微信用户
	DebugModeStatusTypeOn DebugModeStatusType = 2
)

```
