
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Windows;

namespace MeshDesktopClient
{
    public interface IDialogService
    {
        void RegisterDialog<TModel, TDialog>(Window owner, List<object> additionalCtorArgs)
            where TModel : INotifyPropertyChanged
            where TDialog : Window;

        void RegisterDialog<TModel, TDialog>()
            where TModel : INotifyPropertyChanged
            where TDialog : Window;

        bool? ShowDialog<TModel>(TModel modalDialogModel)
            where TModel : INotifyPropertyChanged;
    }
}
