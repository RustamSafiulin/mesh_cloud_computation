package model

import "strconv"

type TaskStartInfo struct {
	FileID        string `json:"file_id,omitempty"`
	TaskID        string `json:"task_id,omitempty"`
	FileName      string `json:"file_name,omitempty"`
	ServiceHostIP string `json:"service_ip,omitempty"`
	ServicePort   uint16 `json:"service_port,omitempty"`
}

func (ts *TaskStartInfo) ToString() string {
	var resultString string
	resultString = "TaskStartInfo: [TaskID:" +
		ts.TaskID + ", FileID: " +
		ts.FileID + ", FileName: " +
		ts.FileName + ", ServiceHostIP: " +
		ts.ServiceHostIP + ", ServicePort: " +
		strconv.Itoa(int(ts.ServicePort)) + "]"

	return resultString
}