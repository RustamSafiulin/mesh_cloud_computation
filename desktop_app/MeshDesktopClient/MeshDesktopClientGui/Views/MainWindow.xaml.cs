
using System;
using System.ComponentModel;
using System.Collections.Generic;
using System.Linq;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Navigation;

using MeshDesktopClient;
using MeshDesktopClient.ViewModels;

namespace MeshDesktopClient.Views
{
    public partial class MainWindow : Window
    {
        public MainWindow(ApplicationViewModel viewModel, ApplicationEnvironment appEnvironment, IDialogService dialogService)
        {
            InitializeComponent();

            mViewModel = viewModel;
            mAppEnvironment = appEnvironment;
            mDialogService = dialogService;

            mNavigationService = NavigationFrame.NavigationService;
            mPages = new List<Page>
            {
                new TasksPage(mViewModel.TasksViewModel),
                new SettingsPage(mViewModel.SettingsViewModel)
            };

            this.Closing += MainViewClosing;
            this.SourceInitialized += (obj, sender) =>
            {
                this.HideMinimizeAndMaximizeButtons();
                this.SetForegroundWindow();
            };

            mNavigationService.Navigate(mPages.ElementAt(0));
            lwPageChanger.SelectionChanged += PageChangedHandler;
        }

        #region Props and fields

        private IDialogService mDialogService;
        private ApplicationViewModel mViewModel;
        private ApplicationEnvironment mAppEnvironment;

        private readonly List<Page> mPages;
        private readonly NavigationService mNavigationService;

        #endregion

        private void PageChangedHandler(object sender, RoutedEventArgs eventArgs)
        {
            var listView = sender as System.Windows.Controls.ListView;

            if (listView != null)
            {
                var page = mPages.ElementAt(listView.SelectedIndex);
                mNavigationService.Navigate(page);
            }
        }

        private void ToolBar_Loaded(object sender, RoutedEventArgs e)
        {
            var toolBar = sender as System.Windows.Controls.ToolBar;
            var overflowGrid = toolBar.Template.FindName("OverflowGrid", toolBar) as FrameworkElement;
            if (overflowGrid != null)
            {
                overflowGrid.Visibility = Visibility.Collapsed;
            }
            var mainPanelBorder = toolBar.Template.FindName("MainPanelBorder", toolBar) as FrameworkElement;
            if (mainPanelBorder != null)
            {
                mainPanelBorder.Margin = new Thickness();
            }
        }

        private void MainViewClosing(object sender, CancelEventArgs e)
        {
            bool? dialogResult = mDialogService.ShowDialog<ConfirmationDialogViewModel>(new ConfirmationDialogViewModel()
            {
                Title = this.Title,
                InformationText = "Вы действительно хотите выйти?"
            });

            if (dialogResult.Value)
            {
                e.Cancel = false;
                mViewModel.NotifyNeedAppExit();
            }
            else
            {
                e.Cancel = true;
            }
        }

        private void ListViewCollapsedAnimationEvent(object sender, EventArgs e)
        {
            lwTaskLabel.Visibility = Visibility.Hidden;
            lwSettingsLabel.Visibility = Visibility.Hidden;
        }

        private void ListViewRiseAnimationEvent(object sender, EventArgs e)
        {
            lwTaskLabel.Visibility = Visibility.Visible;
            lwSettingsLabel.Visibility = Visibility.Visible;
        }
    }
}
