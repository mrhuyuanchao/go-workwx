package workwx

// ChatDataSetReceiveCallback 设置专区接收回调事件
func (c *WorkwxApp) ChatDataSetReceiveCallback(programId string, callbackUrl string, token string, encodingAESKey string) error {
	_, err := c.execChatDataSetReceiveCallback(reqChatDataSetReceiveCallback{
		ProgramId: programId,
	})
	return err
}

// ChatDataSyncCallProgramSyncMsg 调用专区默认的获取会话消息
func (c *WorkwxApp) ChatDataSyncCallProgramSyncMsg(programId string, abilityId string, cursor string, token string, limit int) (*SyncMsgDataResult, error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId: programId,
		AbilityId: abilityId,
		RequestData: reqCallProgramRequestData{
			Func: "sync_msg",
			FuncReq: map[string]interface{}{
				"cursor": cursor,
				"token":  token,
				"limit":  limit,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	var obj SyncMsgDataResult
	err = resp.intoResult(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ChatDataSyncDoAsyncJob 调用专区默认的获取回调数据
func (c *WorkwxApp) ChatDataSyncDoAsyncJob(programId string, abilityId string, notifyId string) (*DoSyncJobResult, error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId: programId,
		AbilityId: abilityId,
		NotifyId:  notifyId,
		RequestData: reqCallProgramRequestData{
			Func:    "do_async_job",
			FuncReq: map[string]interface{}{},
		},
	})
	if err != nil {
		return nil, err
	}
	var obj DoSyncJobResult
	err = resp.intoResult(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ChatDataSyncCreateRecommendDialogTask 调用专区默认的创建话术推荐模型
func (c *WorkwxApp) ChatDataSyncCreateRecommendDialogTask(programId string, abilityId string, kbId string, msgList []MsgList) (*CreateRecommendDialogTaskResult, error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId: programId,
		AbilityId: abilityId,
		RequestData: reqCallProgramRequestData{
			Func: "create_recommend_dialog_task",
			FuncReq: map[string]interface{}{
				"kb_id":    kbId,
				"msg_list": msgList,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	var obj CreateRecommendDialogTaskResult
	err = resp.intoResult(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ChatDataSyncGetRecommendDialogResult 调用专区默认的获取话术推荐模型结果
func (c *WorkwxApp) ChatDataSyncGetRecommendDialogResult(programId string, abilityId string, jobId string) (*GetRecommendDialogResult, error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId: programId,
		AbilityId: abilityId,
		RequestData: reqCallProgramRequestData{
			Func: "get_recommend_dialog_result",
			FuncReq: map[string]interface{}{
				"job_id": jobId,
			},
		},
	})
	if err != nil {
		return nil, err
	}
	var obj GetRecommendDialogResult
	err = resp.intoResult(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ChatDataSyncKnowledgeBaseList 调用专区默认的获取企业知识集
func (c *WorkwxApp) ChatDataSyncKnowledgeBaseList(programId string, abilityId string) (*KnowledgeBaseListResult, error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId: programId,
		AbilityId: abilityId,
		RequestData: reqCallProgramRequestData{
			Func:    "knowledge_base_list",
			FuncReq: map[string]interface{}{},
		},
	})
	if err != nil {
		return nil, err
	}
	var obj KnowledgeBaseListResult
	err = resp.intoResult(&obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}

// ChatDataAsyncProgramTask 应用异步调用专区程序
// todo 调用具体的专区程序实现
func (c *WorkwxApp) chatDataAsyncProgramTask(programId string, abilityId string, requestData string) (jobId string, err error) {
	resp, err := c.execChatDataAsyncProgramTask(reqChatDataAsyncProgramTask{
		ProgramId:   programId,
		AbilityId:   abilityId,
		RequestData: requestData,
	})
	if err != nil {
		return "", err
	}
	return resp.JobId, nil
}

// ChatDataAsyncProgramResult 获取专区程序任务结果
func (c *WorkwxApp) ChatDataAsyncProgramResult(jobId string) (*AsyncProgramResult, error) {
	resp, err := c.execChatDataAsyncProgramResult(reqChatDataAsyncProgramResult{
		JobId: jobId,
	})
	if err != nil {
		return nil, err
	}
	return &resp.AsyncProgramResult, nil
}

// ChatDataOpenDebugMode 开启专区调试模式
func (c *WorkwxApp) ChatDataOpenDebugMode(programId string) error {
	_, err := c.execChatDataOpenDebugMode(reqChatDataOpenDebugMode{
		ProgramId: programId,
	})
	return err
}

// ChatDataCloseDebugMode 关闭专区调试模式
func (c *WorkwxApp) ChatDataCloseDebugMode(programId string) error {
	_, err := c.execChatDataCloseDebugMode(reqChatDataCloseDebugMode{
		ProgramId: programId,
	})
	return err
}

// ChatDataCheckDebugMode 获取专区调试模式状态
func (c *WorkwxApp) ChatDataCheckDebugMode(programId string) (DebugModeStatusType, error) {
	resp, err := c.execChatDataCheckDebugMode(reqChatDataCheckDebugMode{
		ProgramId: programId,
	})
	if err != nil {
		return 0, err
	}
	return resp.DebugModeStatus, nil
}
