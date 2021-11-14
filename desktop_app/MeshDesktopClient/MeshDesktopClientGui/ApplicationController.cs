
using System;
using System.Collections.Generic;

using System.Windows;

using MeshDesktopClient;
using MeshDesktopClient.Service;
using MeshDesktopClient.Helpers;
using MeshDesktopClient.EventBus;
using MeshDesktopClient.Views;
using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient
{
    public sealed class ApplicationController
    {
        #region Props and Fields

        private EntryWindow mEntryWindow;
        private EntryViewModel mEntryViewModel;

        private MainWindow mMainWindow;
        private ApplicationViewModel mAppViewModel;

        private IDialogService mDialogService;
        private ApplicationEnvironment mAppEnvironment;

        #endregion

        public ApplicationController()
        { }

        public void Startup()
        {
            mAppEnvironment = new ApplicationEnvironment();
            mAppEnvironment.SetupApplicationEnvironment();
            mAppEnvironment.EventHub.Subscribe<AppExitEvent>(OnAppShutdownEventHandler);
            mAppEnvironment.EventHub.Subscribe<LoginEvent>(OnLoginEventHandler);

            mEntryViewModel = new EntryViewModel(mAppEnvironment);
            mEntryWindow = new EntryWindow(mEntryViewModel);
        }

        private void OnAppShutdownEventHandler(AppExitEvent exitEvent)
        {
            mAppEnvironment.TearDownApplicationEnvironment();
            Application.Current.Shutdown();
        }

        private void OnLoginEventHandler(LoginEvent loginEvent)
        {
            try
            {
                SessionStorage.Instance.AuthInfo = loginEvent.AuthInfo;

                Logger.Log.Debug("Init dialog service and dialogs.");
                mDialogService = new DialogService();

                Logger.Log.Info("Init view models.");
                mAppViewModel = new ApplicationViewModel(mAppEnvironment, mDialogService);
                mMainWindow = new MainWindow(mAppViewModel, mAppEnvironment, mDialogService);
                mMainWindow.Show(); //!Important: For user as owner for settings and reports windows (without show() obtain app crash)

                Logger.Log.Debug("Register modal dialogs.");
                mDialogService.RegisterDialog<ConfirmationDialogViewModel, ConfirmationDialog>();
                mDialogService.RegisterDialog<InformationDialogViewModel, InformationDialog>();
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Error was caused during call OnLoginEventHandler. Reason: {e.Message}");
            }
        }
    }
}
