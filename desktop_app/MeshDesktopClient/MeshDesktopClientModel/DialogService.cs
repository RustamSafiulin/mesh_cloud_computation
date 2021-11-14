
using System;
using System.Windows;
using System.ComponentModel;
using System.Collections.Generic;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient
{
    public sealed class DialogService : IDialogService
    {
        #region Props and Fields

        internal class DialogInfo
        {
            public Type DialogType { get; private set; }
            public Window DialogOwner { get; private set; }

            public List<object> AdditionalCtorArgs { get; private set; }

            public DialogInfo(Type dialogType, Window dialogOwner, List<object> additionalCtorArgs)
            {
                DialogType = dialogType;
                DialogOwner = dialogOwner;
                AdditionalCtorArgs = additionalCtorArgs;
            }
        }

        private Dictionary<Type, DialogInfo> mRegisteredModelDialogCache;

        #endregion

        public void RegisterDialog<TModel, TDialog>(Window owner, List<object> additionalCtorArgs)
            where TModel : INotifyPropertyChanged
            where TDialog : Window
        {
            AddDialogInfo<TModel, TDialog>(owner, additionalCtorArgs);
        }

        public void RegisterDialog<TModel, TDialog>()
            where TModel : INotifyPropertyChanged
            where TDialog : Window
        {
            AddDialogInfo<TModel, TDialog>(null, null);
        }

        public bool? ShowDialog<TModel>(TModel dataContextVm)
            where TModel : INotifyPropertyChanged
        {
            if (!mRegisteredModelDialogCache.ContainsKey(typeof(TModel)))
            {
                throw new InvalidOperationException($"Dialog {nameof(TModel)} doesn't registered");
            }

            var dialogInfo = mRegisteredModelDialogCache[typeof(TModel)];
            List<object> arguments = new List<object>();
            arguments.Add(dataContextVm);

            if (dialogInfo.AdditionalCtorArgs != null && dialogInfo.AdditionalCtorArgs.Count != 0)
                arguments.AddRange(dialogInfo.AdditionalCtorArgs);

            var dialog = CreateDialog(dialogInfo.DialogType, arguments);
            dialog.Owner = dialogInfo.DialogOwner;

            return dialog.ShowDialog();
        }

        private void AddDialogInfo<TModel, TDialog>(Window owner, List<object> additionalCtorArgs)
            where TModel : INotifyPropertyChanged
            where TDialog : Window
        {
            var modelType = typeof(TModel);
            if (!mRegisteredModelDialogCache.ContainsKey(modelType))
            {
                mRegisteredModelDialogCache[modelType] = new DialogInfo(typeof(TDialog), owner, additionalCtorArgs);
                return;
            }

            throw new InvalidOperationException($"Dialog {nameof(TModel)} has already been registered");
        }

        private Window CreateDialog(Type dialogType, List<object> arguments)
        {
            if (dialogType == null)
                throw new ArgumentNullException(dialogType.Name);

            var instance = Activator.CreateInstance(dialogType, arguments.ToArray());
            var dialog = instance as Window;
            if (dialog != null)
            {
                return dialog;
            }

            throw new ArgumentException($"Dialog {dialogType.Name} doesn't supported");
        }

        public DialogService()
        {
            mRegisteredModelDialogCache = new Dictionary<Type, DialogInfo>();
        }
    }
}
