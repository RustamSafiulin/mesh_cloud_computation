
using System;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.ViewModels
{
    public class SettingsViewModel : DomainModelBase
    {
        public SettingsViewModel(ApplicationEnvironment appEnvironment, IDialogService dialogService)
        {
            mAppEnvironment = appEnvironment;
            mDialogService = dialogService;
        }

        #region Props and Fields

        private ApplicationEnvironment mAppEnvironment;
        private IDialogService mDialogService;

        #endregion


    }
}