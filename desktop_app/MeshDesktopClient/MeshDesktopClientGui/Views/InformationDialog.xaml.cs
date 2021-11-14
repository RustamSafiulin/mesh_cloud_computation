
using System;
using System.Windows;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class InformationDialog : Window
    {
        public InformationDialog(InformationDialogViewModel viewModel)
        {
            InitializeComponent();

            this.DataContext = viewModel;
            this.SetForegroundWindow();
            this.Topmost = true;
        }
    }
}
