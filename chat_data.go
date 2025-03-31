package workwx

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
