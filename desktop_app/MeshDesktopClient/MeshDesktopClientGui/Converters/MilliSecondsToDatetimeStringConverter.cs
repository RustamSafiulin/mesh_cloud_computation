
using System;
using System.Windows;
using System.Windows.Data;
using System.Windows.Input;
using System.Globalization;

namespace MeshDesktopClient.Converters
{
    public class MilliSecondsToDatetimeStringConverter : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, System.Globalization.CultureInfo culture)
        {
            var epoch = (Int64)value;
            var dateTime = new DateTime(1970, 1, 1, 0, 0, 0, 0, DateTimeKind.Unspecified).AddMilliseconds(epoch);

            return dateTime.ToLocalTime().ToString("dd.MM.yyyy");
        }

        public object ConvertBack(object value, Type targetType, object parameter,
            System.Globalization.CultureInfo culture)
        {
            throw new NotImplementedException();
        }
    }
}
