﻿
using System;
using System.Threading;
using System.Windows.Controls;
using System.Windows.Controls.Primitives;

namespace MeshDesktopClient.Controls
{
    public class CustomDatePicker : DatePicker
    {
        protected DatePickerTextBox _datePickerTextBox;
        protected readonly String _shortDatePattern;

        public override void OnApplyTemplate()
        {
            base.OnApplyTemplate();

            _datePickerTextBox = this.Template.FindName("PART_TextBox", this) as DatePickerTextBox;
            if (_datePickerTextBox != null)
            {
                _datePickerTextBox.TextChanged += dptb_TextChanged;
            }
        }

        private void dptb_TextChanged(object sender, TextChangedEventArgs e)
        {
            DateTime dt;
            if (DateTime.TryParseExact(_datePickerTextBox.Text, _shortDatePattern, Thread.CurrentThread.CurrentCulture,
                System.Globalization.DateTimeStyles.None, out dt))
            {
                this.SelectedDate = dt;
            }

        }

        public CustomDatePicker()
            : base()
        {
            _shortDatePattern = Thread.CurrentThread.CurrentCulture.DateTimeFormat.ShortDatePattern;
        }
    }
}
