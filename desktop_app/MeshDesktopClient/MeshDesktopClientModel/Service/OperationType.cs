
using System;

namespace MeshDesktopClient.Service
{
    public enum OperationType
    {
        LOGIN,

        CREATE_TASK,
        GET_ALL_ACCOUNTS_TASKS,
        UPLOAD_TASK_DATA,
        DOWNLOAD_TASK_DATA,
        START_TASK,
        STOP_TASK,
        DELETE_TASK
    }
}
