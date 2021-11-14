
using System;
using System.Windows;
using System.Windows.Data;
using System.Windows.Input;
using System.Globalization;

namespace MeshDesktopClient.Converters
{
    public class BusyToOpacityConverter : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, System.Globalization.CultureInfo culture)
        {
            if (value is Boolean && (bool)value)
            {
                return 0.5;
            }

            return 1.0;
        }

        public object ConvertBack(object value, Type targetType, object parameter, System.Globalization.CultureInfo culture)
        {
            if (value is int && (double)value < 1.0)
            {
                return true;
            }
            
            return false;
        }
    }
}
