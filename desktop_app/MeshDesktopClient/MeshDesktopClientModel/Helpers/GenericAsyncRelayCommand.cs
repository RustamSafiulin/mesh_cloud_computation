
using System;
using System.Windows.Input;
using System.Threading.Tasks;

namespace MeshDesktopClient.Helpers
{
    public class GenericAsyncRelayCommand<T> : ICommand
    {
        private readonly Func<T, Task> _execute = null;
        private readonly Func<bool> _canExecute = null;
        private bool _isExecuting;

        public GenericAsyncRelayCommand(Func<T, Task> execute) : this(execute, null)
        {
        }

        public GenericAsyncRelayCommand(Func<T, Task> execute, Func<bool> canExecute)
        {
            if (execute == null)
                throw new ArgumentNullException("execute");

            _execute = execute;
            _canExecute = canExecute;
        }

        public bool CanExecute(object parameter)
        {
            return !(_isExecuting && _canExecute());
        }

        public event EventHandler CanExecuteChanged;

        public async void Execute(object parameter)
        {
            _isExecuting = true;
            OnCanExecuteChanged();
            try
            {
                await _execute((T)parameter);
            }
            finally
            {
                _isExecuting = false;
                OnCanExecuteChanged();
            }
        }

        protected virtual void OnCanExecuteChanged()
        {
            if (CanExecuteChanged != null) CanExecuteChanged(this, new EventArgs());
        }
    }
}
