
using System;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.ViewModels
{
    public class ConfirmationDialogViewModel : DomainModelBase
    {
        public ConfirmationDialogViewModel()
        {
            AcceptDialogCommand = new RelayCommand(AcceptDialogHandler);
            RejectDialogCommand = new RelayCommand(RejectDialogHandler);
        }

        private String _informationText = String.Empty;
        public String InformationText
        {
            get { return _informationText; }
            set { SetProperty(ref _informationText, value); }
        }

        private String _title = String.Empty;
        public String Title
        {
            get { return _title; }
            set { SetProperty(ref _title, value); }
        }

        private bool? mDialogResult;
        public bool? DialogResult
        {
            get { return mDialogResult; }
            set { SetProperty(ref mDialogResult, value); }
        }

        public RelayCommand AcceptDialogCommand { get; set; }

        public RelayCommand RejectDialogCommand { get; set; }

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
