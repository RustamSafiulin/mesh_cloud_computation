
using System;
using System.Collections.ObjectModel;
using System.Threading.Tasks;

using MeshDesktopClient;
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

        }

        private async Task LoadAllTasks()
        {

        }

        private async Task DeleteTask()
        {

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