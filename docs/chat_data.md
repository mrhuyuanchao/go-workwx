# 数据与智能专区

## Models

### `SyncMsgDataResult` 会话记录结果

 Name         | JSON          | Type        | Doc                                                                                   
:-------------|:--------------|:------------|:--------------------------------------------------------------------------------------
 `HasMore`    | `has_more`    | `int`       | 是否还有更多数据。0-否；1-是。                                                                     
 `NextCursor` | `next_cursor` | `string`    | 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取。强烈建议对该字段入库保存，每次请求读取带上，请求结束后更新。避免因意外丢，导致必须从头开始拉取，引起消息延迟。 
| `MsgList`    | `msg_list`    | `[]SyncMsg` | 消息列表，按消息发送时间升序排序                                                                      

### `SyncMsg` 消息列表

 Name                 | JSON                   | Type                 | Doc                                    
:---------------------|:-----------------------|:---------------------|:---------------------------------------
| `MsgId`              | `msgid`                | `string`             | 每条消息对应的msgid                           
| `Sender`             | `sender`               | `Sender`             | 	消息发送者                                 
| `ChatId`             | `chatid`               | `string`             | 群ID，当消息是群消息的时候会返回该字段                   
| `ReceiverList`       | `receiver_list`        | `[]ReceiverList`     | 消息接收者列表。当自己发给自己消息时该字段为发送者ID，其他情况不包含发送者 
| `SendTime`           | `send_time`            | `int64`              | 消息发送时间对应的unix时间戳                       
| `MsgType`            | `msgtype`              | `MsgType`            | 消息类型                                   
| `ServiceEncryptInfo` | `service_encrypt_info` | `ServiceEncryptInfo` | 加密内容                                   |
| `ExtraInfo`          | `extra_info`           | `ExtraInfo`          | 额外信息                                   

### `Sender` 消息发送人

 Name   | JSON   | Type         | Doc                                                             
:-------|:-------|:-------------|:----------------------------------------------------------------
| `Id`   | `id`   | `string`     | 消息发送者的id，当消息发送者为员工时，该字段为员工的userid；当消息发送者的身份为外部联系人时，该字段为外部联系人的id 
| `Type` | `type` | `MsgObjType` | 消息发送者身份类型。1：员工；2：外部联系人; 3：机器人                                   

### `ReceiverList` 消息接收者列表

 Name   | JSON   | Type         | Doc                                                             
:-------|:-------|:-------------|:----------------------------------------------------------------
| `Id`   | `id`   | `string`     | 消息发送者的id，当消息发送者为员工时，该字段为员工的userid；当消息发送者的身份为外部联系人时，该字段为外部联系人的id 
| `Type` | `type` | `MsgObjType` | 消息发送者身份类型。1：员工；2：外部联系人; 3：机器人                                   

### `ServiceEncryptInfo` 加密内容

 Name                 | JSON                   | Type     | Doc                                                      
:---------------------|:-----------------------|:---------|:---------------------------------------------------------
| `EncryptedSecretKey` | `encrypted_secret_key` | `string` | 加密后的密钥，使用设置公钥设置的公钥进行加密，需要应用后台用私钥解密后，才可在其他接口使用，例如模型分析接口等) 
| `PublicKeyVer`       | `public_key_ver`       | `int`    | 公钥版本号                                                    

### `ExtraInfo` 额外信息

 Name           | JSON            | Type  | Doc                               
:---------------|:----------------|:------|:----------------------------------
| `CallDuration` | `call_duration` | `int` | 通话时长，单位秒。仅当消息类型为“音视频通话”或“音频存档”时返回 

### `DoSyncJobResult` 回调数据结果

 Name                        | JSON                                     | Type                             | Doc              
:----------------------------|:-----------------------------------------|:---------------------------------|:-----------------
 `EventType`                 | `event_type`                             | `ZoneEventType`                  | 事件类型             
 `Timestamp`                 | `timestamp`                              | `int64`                          | 时间戳(单位秒)         
 `ChatArchiveAuditApproved`  | `chat_archive_audit_approved,omitempty`  | `*ChatArchiveAuditApprovedData`  | 客户同意进行聊天内容存档事件数据 
 `ConversationNewMessage`    | `conversation_new_message,omitempty`     | `*ConversationNewMessageData`    | 产生会话回调通知数据       
 `HitKeyword`                | `hit_keyword,omitempty`                  | `*HitKeywordData`                | 命中关键词规则通知数据      
 `AuthKnowledgeBase`         | `auth_knowledge_base,omitempty`          | `*AuthKnowledgeBaseData`         | 知识集管理回调数据        
 `ChatArchiveExportFinished` | `chat_archive_export_finished,omitempty` | `*ChatArchiveExportFinishedData` | 会话内容导出完成通知数据     

### `ChatArchiveAuditApprovedData` 客户同意进行聊天内容存档事件数据

 Name             | JSON              | Type     | Doc         
:-----------------|:------------------|:---------|:------------
 `UserID`         | `userid`          | `string` | 企业成员UserID  
 `ExternalUserID` | `external_userid` | `string` | 外部联系人UserID 
 `ChatID`         | `chatid`          | `string` | 群聊ID        

### `ConversationNewMessageData` 产生会话回调通知数据

 Name    | JSON    | Type     | Doc     
:--------|:--------|:---------|:--------
 `Token` | `token` | `string` | 会话token 

### `HitKeywordData` 命中关键词规则通知数据

 Name    | JSON    | Type     | Doc     
:--------|:--------|:---------|:--------
 `Token` | `token` | `string` | 会话token 

### `AuthKnowledgeBaseData` 知识集管理回调数据

 Name                | JSON                  | Type     | Doc   
:--------------------|:----------------------|:---------|:------
 `KnowledgeBaseID`   | `knowledge_base_id`   | `string` | 知识库ID 
 `KnowledgeBaseName` | `knowledge_base_name` | `string` | 知识库名称 

### `ChatArchiveExportFinishedData` 会话内容导出完成通知数据

 Name    | JSON    | Type     | Doc  
:--------|:--------|:---------|:-----
 `JobID` | `jobid` | `string` | 任务ID 

### `CreateRecommendDialogTaskResult` 创建话术推荐模型结果

 Name       | JSON                  | Type          | Doc   
:-----------|:----------------------|:--------------|:------
 `JobID`    | `jobid`               | `string`      | 任务ID  
 `FailList` | `fail_list,omitempty` | `*[]FailItem` | 失败项列表 

### `FailItem` 失败项

 Name          | JSON           | Type           | Doc  
:--------------|:---------------|:---------------|:-----
 `ErrCode`     | `errcode`      | `int`          | 错误码  
 `ErrMsg`      | `errmsg`       | `string`       | 错误信息 
 `MsgID`       | `msgid`        | `string`       | 消息ID 
 `EncryptInfo` | `encrypt_info` | `*EncryptInfo` | 加密信息 

### `GetRecommendDialogResult` 获取话术推荐模型结果

 Name        | JSON                  | Type          | Doc     
:------------|:----------------------|:--------------|:--------
 `MessageId` | `response_data`       | `string`      | 话术推荐的结果 
 `FailList`  | `fail_list,omitempty` | `*[]FailItem` | 失败项列表   

### `KnowledgeBaseListResult` 知识集列表结果

 Name         | JSON           | Type                  | Doc     
:-------------|:---------------|:----------------------|:--------
 `KBInfoList` | `kb_info_list` | `[]KnowledgeBaseInfo` | 知识库信息列表 

### `KnowledgeBaseInfo` 知识库信息

 Name     | JSON      | Type     | Doc   
:---------|:----------|:---------|:------
 `KBID`   | `kb_id`   | `string` | 知识库ID 
 `KBName` | `kb_name` | `string` | 知识库名称 

### `AsyncProgramResult` 异步任务结果

 Name              | JSON               | Type     | Doc                                                           
:------------------|:-------------------|:---------|:--------------------------------------------------------------
 `ResponseErrCode` | `response_errcode` | `string` | 上报异步任务结果中上报的errcode。代表专区程序返回的错误码                              
 `ResponseData`    | `response_data`    | `string` | 上报异步任务结果中上报的result。代表专区程序的输出结果，为自定义的JSON字符串，要求与管理端配置的输出协议格式匹配 

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

```go
// MsgType 	消息类型
//
// 0	暂不支持的消息类型
// 1	文本
// 2	图片
// 3	表情
// 4	链接
// 5	小程序
// 6	语音
// 7	视频
// 8	文件
// 9	名片
// 10	转发消息
// 11	视频号
// 12	日程
// 13	红包
// 14	地理位置
// 15	快速会议
// 16	待办
// 17	投票
// 18	在线文档
// 19	图文消息
// 20	图文混合消息
// 21	音频存档
// 22	音视频通话
// 23	微盘文件
// 24	同意会话存档
// 25	拒绝会话存档
// 26	群接龙
// 27	markdown
// 28	笔记
type MsgType int

const (
// MsgTypeUnsupported 暂不支持的消息类型
MsgTypeUnsupported MsgType = 0
// MsgTypeText 文本
MsgTypeText MsgType = 1
// MsgTypeImage 图片
MsgTypeImage MsgType = 2
// MsgTypeEmoji 表情
MsgTypeEmoji MsgType = 3
// MsgTypeLink 链接
MsgTypeLink MsgType = 4
// MsgTypeMiniProgram 小程序
MsgTypeMiniProgram MsgType = 5
// MsgTypeVoice 语音
MsgTypeVoice MsgType = 6
// MsgTypeVideo 视频
MsgTypeVideo MsgType = 7
// MsgTypeFile 文��
MsgTypeFile MsgType = 8
// MsgTypeCard 名片
MsgTypeCard MsgType = 9
// MsgTypeForward 转发消息
MsgTypeForward MsgType = 10
// MsgTypeChannels 视频号
MsgTypeChannels MsgType = 11
// MsgTypeSchedule 日程
MsgTypeSchedule MsgType = 12
// MsgTypeRedPacket 红包
MsgTypeRedPacket MsgType = 13
// MsgTypeLocation 地理位置
MsgTypeLocation MsgType = 14
// MsgTypeQuickMeeting 快速会议
MsgTypeQuickMeeting MsgType = 15
// MsgTypeTodo 待办
MsgTypeTodo MsgType = 16
// MsgTypeVote 投票
MsgTypeVote MsgType = 17
// MsgTypeOnlineDoc 在线文档
MsgTypeOnlineDoc MsgType = 18
// MsgTypeRichText 图文消息
MsgTypeRichText MsgType = 19
// MsgTypeMixedContent 图文混合消息
MsgTypeMixedContent MsgType = 20
// MsgTypeAudioArchive 音频存档
MsgTypeAudioArchive MsgType = 21
// MsgTypeAudioVideoCall 音视频通话
MsgTypeAudioVideoCall MsgType = 22
// MsgTypeWedriveFile 微盘文件
MsgTypeWedriveFile MsgType = 23
// MsgTypeAcceptArchive 同意会话存档
MsgTypeAcceptArchive MsgType = 24
// MsgTypeRejectArchive 拒绝会话存档
MsgTypeRejectArchive MsgType = 25
// MsgTypeGroupChain 群接龙
MsgTypeGroupChain MsgType = 26
// MsgTypeMarkdown markdown
MsgTypeMarkdown MsgType = 27
// MsgTypeNote 笔记
MsgTypeNote MsgType = 28
)

type MsgObjType int

const (
// MsgObjTypeEmployee 员工
MsgObjTypeEmployee = 1
// MsgObjTypeExternalContact 外部联系人
MsgObjTypeExternalContact = 2
// MsgObjTypeRobot 机器人
MsgObjTypeRobot = 3
)

```

```go
// ZoneEventType 专区程序接收事件通知

type ZoneEventType string

const (
// ZoneEventTypeChatArchiveAuditApprovedSingle 客户同意进行聊天内容存档事件回调
ZoneEventTypeChatArchiveAuditApprovedSingle ZoneEventType = "chat_archive_audit_approved_single"

// ZoneEventTypeConversationNewMessage 产生会话回调通知
ZoneEventTypeConversationNewMessage ZoneEventType = "conversation_new_message"

// ZoneEventTypeHitKeyword 命中关键词规则通知
ZoneEventTypeHitKeyword ZoneEventType = "hit_keyword"

// ZoneEventTypeAuthKnowledgeBase 知识集管理回调
ZoneEventTypeAuthKnowledgeBase ZoneEventType = "auth_knowledge_base"

// ZoneEventTypeChatArchiveExportFinished 会话内容导出完成通知
ZoneEventTypeChatArchiveExportFinished ZoneEventType = "chat_archive_export_finished"
)

```
