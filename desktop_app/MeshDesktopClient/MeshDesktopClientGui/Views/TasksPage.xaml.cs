
using System;
using System.Windows.Controls;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class TasksPage : Page
    {
        public TasksPage(TasksViewModel viewModel)
        {
            InitializeComponent();

            this.DataContext = viewModel;
        }
    }
}
