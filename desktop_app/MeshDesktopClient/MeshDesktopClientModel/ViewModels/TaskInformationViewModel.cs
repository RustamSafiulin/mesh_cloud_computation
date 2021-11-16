
using System;

using MeshDesktopClient.Helpers;
using MeshDesktopClient.Models;

namespace MeshDesktopClient.ViewModels
{ 
    public class TaskInformationViewModel : DomainModelBase
    {
        public TaskInformationViewModel()
        {
            AcceptDialogCommand = new RelayCommand(AcceptDialogHandler);
            RejectDialogCommand = new RelayCommand(RejectDialogHandler);
        }

        #region Props and Fields

        private bool? mDialogResult;
        public bool? DialogResult
        {
            get { return mDialogResult; }
            set { SetProperty(ref mDialogResult, value); }
        }

        public RelayCommand AcceptDialogCommand { get; set; }

        public RelayCommand RejectDialogCommand { get; set; }

        #endregion

        private void AcceptDialogHandler()
        {
            DialogResult = true;
        }

        private void RejectDialogHandler()
        {
            DialogResult = false;
        }
    }
}