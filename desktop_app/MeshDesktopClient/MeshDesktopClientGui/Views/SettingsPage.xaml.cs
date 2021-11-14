
using System;
using System.Windows.Controls;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class SettingsPage : Page
    {
        public SettingsPage(SettingsViewModel viewModel)
        {
            InitializeComponent();
            this.DataContext = viewModel;
        }
    }
}
