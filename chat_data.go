package workwx

// ChatDataSetReceiveCallback 设置专区接收回调事件
func (c *WorkwxApp) ChatDataSetReceiveCallback(programId string, callbackUrl string, token string, encodingAESKey string) error {
	_, err := c.execChatDataSetReceiveCallback(reqChatDataSetReceiveCallback{
		ProgramId: programId,
	})
	return err
}

// ChatDataSyncCallProgram 应用同步调用专区程序
func (c *WorkwxApp) ChatDataSyncCallProgram(programId string, abilityId string, notifyId string, requestData string) (responseData string, err error) {
	resp, err := c.execChatDataSyncCallProgram(reqChatDataSyncCallProgram{
		ProgramId:   programId,
		AbilityId:   abilityId,
		NotifyId:    notifyId,
		RequestData: requestData,
	})
	if err != nil {
		return "", err
	}
	return resp.ResponseData, nil
}

// ChatDataAsyncProgramTask 应用异步调用专区程序
func (c *WorkwxApp) ChatDataAsyncProgramTask(programId string, abilityId string, requestData string) (jobId string, err error) {
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
func (c *WorkwxApp) ChatDataAsyncProgramResult(jobId string) (result *AsyncProgramResult, err error) {
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
func (c *WorkwxApp) ChatDataCheckDebugMode(programId string) (isDebugMode DebugModeStatusType, err error) {
	resp, err := c.execChatDataCheckDebugMode(reqChatDataCheckDebugMode{
		ProgramId: programId,
	})
	if err != nil {
		return 0, err
	}
	return resp.DebugModeStatus, nil
}
