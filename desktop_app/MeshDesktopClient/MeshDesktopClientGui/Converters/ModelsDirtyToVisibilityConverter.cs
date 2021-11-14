
using System;
using System.Windows;
using System.Windows.Data;
using System.Windows.Input;
using System.Globalization;

namespace MeshDesktopClient.Converters
{
    public class ModelsDirtyToVisibilityConverter : IMultiValueConverter
    {
        public object Convert(object[] values, Type targetType, object parameter, CultureInfo culture)
        {
            bool result = false;
            foreach (var value in values)
            {
                if (value is bool)
                {
                    result |= (bool)value;
                }
            }

            return result ? Visibility.Visible : Visibility.Hidden;
        }

        public object[] ConvertBack(object value, Type[] targetTypes, object parameter, CultureInfo culture)
        {
            throw new NotImplementedException();
        }
    }
}