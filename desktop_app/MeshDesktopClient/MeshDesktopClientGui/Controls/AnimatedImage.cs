﻿
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;
using System.Drawing;
using System.Windows.Threading;
using System.Diagnostics;

namespace MeshDesktopClient.Controls
{
    public class AnimatedImage : System.Windows.Controls.Image
    {
        private BitmapSource[] _BitmapSources = null;
        private int _nCurrentFrame = 0;

        private bool _bIsAnimating = false;

        public bool IsAnimating
        {
            get { return _bIsAnimating; }
        }

        static AnimatedImage()
        {
            DefaultStyleKeyProperty.OverrideMetadata(typeof(AnimatedImage), new FrameworkPropertyMetadata(typeof(AnimatedImage)));
        }

        public Bitmap AnimatedBitmap
        {
            get { return (Bitmap)GetValue(AnimatedBitmapProperty); }
            set { StopAnimate(); SetValue(AnimatedBitmapProperty, value); }
        }

        public static readonly DependencyProperty AnimatedBitmapProperty =
            DependencyProperty.Register(
                "AnimatedBitmap", typeof(Bitmap), typeof(AnimatedImage),
                new FrameworkPropertyMetadata(null, new PropertyChangedCallback(OnAnimatedBitmapChanged)));

        private static void OnAnimatedBitmapChanged(DependencyObject obj, DependencyPropertyChangedEventArgs args)
        {
            AnimatedImage control = (AnimatedImage)obj;

            control.UpdateAnimatedBitmap();

            RoutedPropertyChangedEventArgs<Bitmap> e = new RoutedPropertyChangedEventArgs<Bitmap>(
                (Bitmap)args.OldValue, (Bitmap)args.NewValue, AnimatedBitmapChangedEvent);
            control.OnAnimatedBitmapChanged(e);
        }

        public static readonly RoutedEvent AnimatedBitmapChangedEvent = EventManager.RegisterRoutedEvent(
            "AnimatedBitmapChanged", RoutingStrategy.Bubble,
            typeof(RoutedPropertyChangedEventHandler<Bitmap>), typeof(AnimatedImage));

        public event RoutedPropertyChangedEventHandler<Bitmap> AnimatedBitmapChanged
        {
            add { AddHandler(AnimatedBitmapChangedEvent, value); }
            remove { RemoveHandler(AnimatedBitmapChangedEvent, value); }
        }

        protected virtual void OnAnimatedBitmapChanged(RoutedPropertyChangedEventArgs<Bitmap> args)
        {
            RaiseEvent(args);
        }

        private void UpdateAnimatedBitmap()
        {
            int nTimeFrames = AnimatedBitmap.GetFrameCount(System.Drawing.Imaging.FrameDimension.Time);
            _nCurrentFrame = 0;
            if (nTimeFrames > 0)
            {
                _BitmapSources = new BitmapSource[nTimeFrames];

                for (int i = 0; i < nTimeFrames; i++)
                {
                    AnimatedBitmap.SelectActiveFrame(System.Drawing.Imaging.FrameDimension.Time, i);
                    Bitmap bitmap = new Bitmap(AnimatedBitmap);
                    bitmap.MakeTransparent();

                    _BitmapSources[i] = System.Windows.Interop.Imaging.CreateBitmapSourceFromHBitmap(
                        bitmap.GetHbitmap(),
                        IntPtr.Zero,
                        Int32Rect.Empty,
                        System.Windows.Media.Imaging.BitmapSizeOptions.FromEmptyOptions());
                }

                StartAnimate();
            }
        }

        private delegate void VoidDelegate();

        private void OnFrameChanged(object o, EventArgs e)
        {
            Dispatcher.BeginInvoke(DispatcherPriority.Render, new VoidDelegate(delegate { ChangeSource(); }));

        }
        void ChangeSource()
        {
            Source = _BitmapSources[_nCurrentFrame++];
            _nCurrentFrame = _nCurrentFrame % _BitmapSources.Length;
            ImageAnimator.UpdateFrames();

        }

        public void StopAnimate()
        {
            if (_bIsAnimating)
            {
                ImageAnimator.StopAnimate(AnimatedBitmap, new EventHandler(this.OnFrameChanged));
                _bIsAnimating = false;
            }
        }

        public void StartAnimate()
        {
            if (!_bIsAnimating)
            {

                ImageAnimator.Animate(AnimatedBitmap, new EventHandler(this.OnFrameChanged));
                _bIsAnimating = true;
            }
        }
    }
}