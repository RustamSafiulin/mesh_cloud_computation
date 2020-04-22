package model

import (
	"strconv"
)

const (
	TaskResultCompleted = 0
	TaskResultFailed = 1
)

type TaskResult int

type TaskResultInfo struct {
	TaskID 		 string     `json:"task_id,omitempty"`
	Result 		 TaskResult `json:"result,omitempty"`
	WorkerHostIP string		`json:"worker_ip,omitempty"`
	WorkerPort   uint16		`json:"worker_port,omitempty"`
}

func (tr *TaskResultInfo) ToString() string {
	var resultString string
	resultString = "TaskResultInfo: [TaskID:" +
 					tr.TaskID + ", Result: " + strconv.Itoa(int(tr.Result)) +
					", WorkerHostIP: " + tr.WorkerHostIP +
					", WorkerPort: " + strconv.Itoa(int(tr.WorkerPort))  + "]"

	return resultString
}