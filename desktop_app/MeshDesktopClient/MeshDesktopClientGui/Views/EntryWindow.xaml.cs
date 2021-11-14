
using System.Windows;

using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class EntryWindow : Window
    {
        public EntryWindow(EntryViewModel entryViewModel)
        {
            InitializeComponent();

            mViewModel = entryViewModel;
            this.DataContext = mViewModel;
            this.Closing += (sender, e) => { mViewModel.NotifyNeedAppExit(); };
        }

        #region Props and fields

        private EntryViewModel mViewModel;

        #endregion
    }
}
