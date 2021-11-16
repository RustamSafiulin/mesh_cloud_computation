
using System;
using System.Linq;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Threading.Tasks;

using MeshDesktopClient;
using MeshDesktopClient.Dto;
using MeshDesktopClient.Service;
using MeshDesktopClient.Models;
using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.ViewModels
{ 
    public class TasksViewModel : DomainModelBase
    {
        public TasksViewModel(ApplicationEnvironment appEnvironment, IDialogService dialogService)
        {
            mAppEnvironment = appEnvironment;
            mDialogService = dialogService;

            mTasks = new ObservableCollection<TaskModel>();

            CreateTaskCommand = new AsyncRelayCommand(CreateTask);
            LoadAllTasksCommand = new AsyncRelayCommand(LoadAllTasks);
            DeleteTaskCommand = new AsyncRelayCommand(DeleteTask, () => SelectedTask != null);
            ShowTaskInfoCommand = new RelayCommand(ShowTaskInfo, () => SelectedTask != null);
            UploadTaskDataCommand = new AsyncRelayCommand(UploadTaskData, () => SelectedTask != null);
            DownloadTaskDataCommand = new AsyncRelayCommand(DownloadTaskData, () => SelectedTask != null);
            StartTaskCommand = new AsyncRelayCommand(StartTask, () => SelectedTask != null);
            StopTaskCommand = new AsyncRelayCommand(StopTask, () => SelectedTask != null);

            Func<Task> initFunctor = async () =>
            {
                await LoadAllTasks();
            };

            InitResult = new NotifyTaskCompletion(initFunctor());
        }

        #region Props and Fields

        public NotifyTaskCompletion InitResult { get; set; }

        private ApplicationEnvironment mAppEnvironment;
        private IDialogService mDialogService;

        public AsyncRelayCommand CreateTaskCommand { get; set; }
        public AsyncRelayCommand LoadAllTasksCommand { get; set; }
        public AsyncRelayCommand DeleteTaskCommand { get; set; }

        public AsyncRelayCommand UploadTaskDataCommand { get; set; }

        public AsyncRelayCommand DownloadTaskDataCommand { get; set; }

        public AsyncRelayCommand StartTaskCommand { get; set; }

        public AsyncRelayCommand StopTaskCommand { get; set; }

        public RelayCommand ShowTaskInfoCommand { get; set; }

        private String mNewTaskDescription = String.Empty;
        public String NewTaskDescription
        {
            get { return mNewTaskDescription; }
            set { SetProperty(ref mNewTaskDescription, value); }
        }

        private ObservableCollection<TaskModel> mTasks;
        public ObservableCollection<TaskModel> Tasks
        {
            get { return mTasks; }
            set { SetProperty(ref mTasks, value); }
        }

        private TaskModel mSelectedTask;
        public TaskModel SelectedTask
        {
            get { return mSelectedTask; }
            set { SetProperty(ref mSelectedTask, value); }
        }

        #endregion

        private async Task CreateTask()
        {
            try
            {
                if (mAppEnvironment.ServiceConnector == null)
                {
                    throw new NullReferenceException($"Null reference of { nameof(mAppEnvironment.ServiceConnector) }");
                }

                if (String.IsNullOrEmpty(NewTaskDescription))
                {
                    bool? dialogResult = mDialogService.ShowDialog<InformationDialogViewModel>(new InformationDialogViewModel()
                    {
                        Title = "Ошибка добавления новой задачи",
                        InformationText = "Введите корректное описание задачи."
                    });

                    if (dialogResult.HasValue && dialogResult.Value)
                    {
                        return;
                    }
                }

                var taskCreationDto = new TaskCreationDto { Description = NewTaskDescription };

                var response = await mAppEnvironment.ServiceConnector.Post<TaskCreationDto, TaskDto, ErrorResponseDto>(new Request<TaskCreationDto>
                {
                    AuthToken = SessionStorage.Instance.AuthInfo.Token,
                    Body = taskCreationDto,
                    OpName = OperationType.CREATE_TASK
                });

                Logger.Log.Debug($"Got HTTP status code after CREATE_TASK operation: {response.StatusCode}");

                if (response.SuccessBody != null)
                {
                    var taskDto = response.SuccessBody;
                    var newTask = new TaskModel()
                    {
                        Id = taskDto.Id,
                        AccountId = taskDto.AccountId,
                        Description = taskDto.Description,
                        CompletedAt = taskDto.CompletedAt,
                        StartedAt = taskDto.StartedAt,
                        State = taskDto.State,
                        StateText = taskDto.StateText
                    };

                    Tasks.Add(newTask);
                }
                else
                {
                    Logger.Log.Debug($"CREATE_TASK operation was failed. Time: { response.ErrorBody.Timestamp}, Message: { response.ErrorBody.Message }, Details: { response.ErrorBody.Details }");
                }
            }
            catch (ApplicationException e)
            {
                Logger.Log.Debug($"Application exception was caused during CREATE_TASK operation. Reason: {e.Message}");
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Exception was caused during CREATE_TASK operation. Reason: {e.Message}");
            }
        }

        private async Task LoadAllTasks()
        {
            try
            {
                if (mAppEnvironment.ServiceConnector == null)
                {
                    throw new NullReferenceException($"Null reference of { nameof(mAppEnvironment.ServiceConnector) }");
                }

                var response = await mAppEnvironment.ServiceConnector.Get<EmptyRequestBodyDto, List<TaskDto>, ErrorResponseDto>(new Request<EmptyRequestBodyDto>
                {
                    AuthToken = SessionStorage.Instance.AuthInfo.Token,
                    OpName = OperationType.GET_ALL_ACCOUNTS_TASKS
                });

                Logger.Log.Debug($"Got HTTP status code after GET_ALL_ACCOUNTS_TASKS operation: {response.StatusCode}");

                if (response.SuccessBody != null)
                {
                    foreach (var taskDto in response.SuccessBody)
                    {
                        var newTask = new TaskModel()
                        {
                            Id = taskDto.Id,
                            AccountId = taskDto.AccountId,
                            Description = taskDto.Description,
                            CompletedAt = taskDto.CompletedAt,
                            StartedAt = taskDto.StartedAt,
                            State = taskDto.State,
                            StateText = taskDto.StateText
                        };

                        Tasks.Add(newTask);
                    }
                }
                else if (response.ErrorBody != null)
                {
                    Logger.Log.Debug($"GET_ALL_ACCOUNTS_TASKS operation was failed. " +
                        $"Time: { response.ErrorBody.Timestamp}, " +
                        $"Message: { response.ErrorBody.Message }, " +
                        $"Details: { response.ErrorBody.Details }"
                    );
                }
            }
            catch (ApplicationException e)
            {
                Logger.Log.Debug($"Application exception was caused during GET_ALL_ACCOUNTS_TASKS operation. Reason: {e.Message}");
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Exception was caused during GET_ALL_ACCOUNTS_TASKS operation. Reason: {e.Message}");
            }
        }

        private async Task DeleteTask()
        {
            try
            {
                if (mAppEnvironment.ServiceConnector == null)
                {
                    throw new NullReferenceException($"Null reference of { nameof(mAppEnvironment.ServiceConnector) }");
                }

                if (SelectedTask != null)
                {
                    bool? dialogResult = mDialogService.ShowDialog<ConfirmationDialogViewModel>(new ConfirmationDialogViewModel()
                    {
                        Title = "Удаление информации о задаче",
                        InformationText = "Вы действительно хотите удалить информацию о задаче?"
                    });

                    if (dialogResult.HasValue)
                    {
                        if (!dialogResult.Value)
                        {
                            return;
                        }
                    }

                    var response = await mAppEnvironment.ServiceConnector.Delete<EmptyRequestBodyDto, EmptyResponseDto, ErrorResponseDto>(new Request<EmptyRequestBodyDto>
                    {
                        AuthToken = SessionStorage.Instance.AuthInfo.Token,
                        UrlSegments = new Dictionary<string, string> { { "id", Convert.ToString(SelectedTask.Id) } },
                        OpName = OperationType.DELETE_TASK
                    });

                    Logger.Log.Debug($"Got HTTP status code after DELETE_TASK operation: {response.StatusCode}");

                    if (response.StatusCode == System.Net.HttpStatusCode.OK)
                    {
                        var taskForRemove = Tasks.Where(t => t.InternalGuid == SelectedTask.InternalGuid).FirstOrDefault();
                        if (taskForRemove != null)
                        {
                            Tasks.Remove(taskForRemove);
                        }
                    }
                    else
                    {
                        Logger.Log.Debug($"DELETE_TASK operation was failed. " +
                            $"Time: { response.ErrorBody.Timestamp}, " +
                            $"Message: { response.ErrorBody.Message }, " +
                            $"Details: { response.ErrorBody.Details }"
                        );
                    }
                }
            }
            catch (ApplicationException e)
            {
                Logger.Log.Debug($"Application exception was caused during DELETE_TASK operation. Reason: {e.Message}");
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Exception was caused during DELETE_TASK operation. Reason: {e.Message}");
            }
        }

        private void ShowTaskInfo()
        {
            try
            {
                var taskInfoViewModel = new TaskInformationViewModel();

                bool? dialogResult = mDialogService.ShowDialog<TaskInformationViewModel>(taskInfoViewModel);

                if (dialogResult.HasValue && dialogResult.Value)
                {

                }
            }
            catch (Exception e)
            { }
        }

        private async Task UploadTaskData()
        {

        }

        private async Task DownloadTaskData()
        { }

        private async Task StartTask()
        {

        }

        private async Task StopTask()
        { }
    }
}