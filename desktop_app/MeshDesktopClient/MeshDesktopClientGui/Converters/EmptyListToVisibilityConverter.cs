﻿
using System;
using System.Windows;
using System.Windows.Data;
using System.Globalization;
using System.Collections;

namespace MeshDesktopClient.Converters
{
    public class EmptyListToVisibilityConverter : IValueConverter
    {
        public object Convert(object value, Type targetType, object parameter, CultureInfo culture)
        {
            if (value == null)
                return Visibility.Collapsed;
            else
            {
                ICollection list = value as ICollection;
                if (list != null)
                {
                    if (list.Count == 0)
                        return Visibility.Hidden;
                    else
                        return Visibility.Visible;
                }
                else
                    return Visibility.Visible;
            }
        }
        public object ConvertBack(object value, Type targetType, object parameter, System.Globalization.CultureInfo culture)

        {
            throw new NotImplementedException();
        }
    }
}
