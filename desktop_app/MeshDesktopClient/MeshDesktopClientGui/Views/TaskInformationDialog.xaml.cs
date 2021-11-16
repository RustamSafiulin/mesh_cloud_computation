
using System.Windows;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class TaskInformationDialog : Window
    {
        public TaskInformationDialog(TaskInformationViewModel viewModel)
        {
            InitializeComponent();
            this.DataContext = viewModel;
        }
    }
}
