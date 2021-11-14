
using System;
using System.Linq;
using System.Threading.Tasks;
using System.Collections.Generic;
using System.Collections.ObjectModel;

using MeshDesktopClient.Models;
using MeshDesktopClient.Helpers;
using MeshDesktopClient.EventBus;

namespace MeshDesktopClient.ViewModels
{
    public sealed class ApplicationViewModel : DomainModelBase
    {
        public ApplicationViewModel(ApplicationEnvironment appEnvironment, IDialogService dialogService)
        {
            mAppEnvironment = appEnvironment;
            mDialogService = dialogService;

            CreateEventHubSubscriptions();

            TasksViewModel = new TasksViewModel(mAppEnvironment, mDialogService);
            SettingsViewModel = new SettingsViewModel(mAppEnvironment, mDialogService);
        }

        #region Props and Fields

        private ApplicationEnvironment mAppEnvironment;
        private IDialogService mDialogService;

        public SettingsViewModel SettingsViewModel { get; private set; }
        public TasksViewModel TasksViewModel { get; private set; }

        #endregion

        public void NotifyNeedAppExit()
        {
            mAppEnvironment.EventHub.Publish(new AppExitEvent());
        }
        
        private void CreateEventHubSubscriptions()
        {
        }
    }
}
