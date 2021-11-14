
using System;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.ViewModels
{
    public class InformationDialogViewModel : DomainModelBase
    {
        public InformationDialogViewModel()
        {
            AcceptDialogCommand = new RelayCommand(AcceptDialogHandler);
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

        private void AcceptDialogHandler()
        {
            DialogResult = true;
        }
    }
}
