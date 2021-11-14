
using System;
using System.Windows;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class ConfirmationDialog : Window
    {
        public ConfirmationDialog(ConfirmationDialogViewModel viewModel)
        {
            InitializeComponent();

            this.DataContext = viewModel;
            this.SetForegroundWindow();
            this.Topmost = true;
        }
    }
}
