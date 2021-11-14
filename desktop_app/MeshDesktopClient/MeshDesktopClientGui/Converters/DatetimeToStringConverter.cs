
using System;
using System.Windows;
using System.Windows.Data;
using System.Windows.Input;
using System.Globalization;

namespace MeshDesktopClient.Converters
{
    public class DatetimeToStringConverter : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, System.Globalization.CultureInfo culture)
        {
            if (value == null)
                return String.Empty;

            var datetimeObject = (DateTime)value;
            if (datetimeObject == null)
                return String.Empty;

            return datetimeObject.ToLocalTime().ToString("dd.MM.yyyy");
        }

        public object ConvertBack(object value, Type targetType, object parameter,
            System.Globalization.CultureInfo culture)
        {
            throw new NotImplementedException();
        }
    }
}
