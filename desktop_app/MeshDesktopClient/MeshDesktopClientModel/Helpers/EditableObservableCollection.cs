
using System;
using System.Linq;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.ComponentModel;
using System.Collections.Specialized;
using System.Runtime.CompilerServices;

namespace MeshDesktopClient.Helpers
{
    public class EditableObservableCollection<T> : ObservableCollection<T>, ITrackable, INotifyPropertyChanged
        where T : EditableObject<T>
    {
        internal enum ItemState
        {
            Unknown = 0,
            AddedNotCommited = 1,
            DeletedNotCommited = 2,
            Commited = 3
        }

        internal class ItemBackupInfo<TItem> where TItem : EditableObject<T>
        {
            public ItemBackupInfo(TItem originalObject, ItemState itemState)
            {
                Item = originalObject;
                ItemState = itemState;
            }

            public TItem Item { get; private set; }

            public ItemState ItemState { get; set; }
        }

        public EditableObservableCollection(IEnumerable<String> itemNoWatchableProps = null)
        {
            mItemNoWatchableProps = (itemNoWatchableProps != null) ? itemNoWatchableProps.ToList() : new List<string>();
            mBackupItems = new List<ItemBackupInfo<T>>();
            IsTrackable = true;

            base.CollectionChanged += CollectionChanged_Handler;
        }

        private readonly List<String> mItemNoWatchableProps;

        private readonly List<ItemBackupInfo<T>> mBackupItems;

        public Boolean IsTrackable { get; set; }

        public event PropertyChangedEventHandler PropertyChanged;

        private Boolean _isDirty = false;
        public Boolean IsDirty
        {
            get { return _isDirty; }
            set
            {
                if (_isDirty != value)
                {
                    _isDirty = value;
                    OnPropertyChanged();
                }
            }
        }

        public void AcceptChanges()
        {
            using (UntrackedContext<EditableObservableCollection<T>>.Untrack(this))
            {
                foreach (var item in Items)
                {
                    if (item.IsDirty)
                    {
                        item.AcceptChanges();
                    }
                }

                foreach (var item in mBackupItems)
                {
                    item.ItemState = ItemState.Commited;
                }
            }

            IsDirty = false;
        }

        public void RejectChanges()
        {
            using (UntrackedContext<EditableObservableCollection<T>>.Untrack(this))
            {
                foreach (var item in Items)
                {
                    if (item.IsDirty)
                    {
                        item.RejectChanges();
                    }
                }

                var newDeletedItems = mBackupItems.FindAll(bi => bi.ItemState == ItemState.DeletedNotCommited);
                var newAddedItems = mBackupItems.FindAll(bi => bi.ItemState == ItemState.AddedNotCommited);

                foreach (var item in newDeletedItems)
                {
                    Add(item.Item);
                }

                foreach (var item in newAddedItems)
                {
                    Remove(item.Item);
                }

                mBackupItems.RemoveAll(bi => newAddedItems.Find(ni => ni.Item.InternalGuid == bi.Item.InternalGuid) != null);

                foreach (var item in mBackupItems)
                {
                    item.ItemState = ItemState.Commited;
                }
            }

            IsDirty = false;
        }

        private void CollectionChanged_Handler(object sender, NotifyCollectionChangedEventArgs e)
        {
            if (e.OldItems != null && e.Action == NotifyCollectionChangedAction.Remove)
            {
                foreach (T x in e.OldItems)
                {
                    var founded = mBackupItems.Find(b => b.Item.InternalGuid == x.InternalGuid);
                    if (founded != null)
                    {
                        if (IsTrackable)
                        {
                            if (founded.ItemState == ItemState.Commited)
                            {
                                founded.ItemState = ItemState.DeletedNotCommited;
                            }
                        }
                    }

                    x.PropertyChanged -= ItemChanged;
                }

                if (IsTrackable)
                {
                    IsDirty = true;
                }
            }

            if (e.NewItems != null && e.Action == NotifyCollectionChangedAction.Add)
            {
                foreach (T x in e.NewItems)
                {
                    var founded = mBackupItems.Find(b => b.Item.InternalGuid == x.InternalGuid);

                    if (founded == null)
                    {
                        var state = IsTrackable ? ItemState.AddedNotCommited : ItemState.Commited;
                        mBackupItems.Add(new ItemBackupInfo<T>(x, state));
                        x.PropertyChanged += ItemChanged;
                    }
                }

                if (IsTrackable)
                {
                    IsDirty = true;
                }
            }
        }

        private void ItemChanged(object sender, PropertyChangedEventArgs e)
        {
            if (mItemNoWatchableProps.FirstOrDefault(prop => prop == e.PropertyName) != null)
            {
                return;
            }

            var item = sender as EditableObject<T>;
            if (item != null)
            {
                if (!item.IsTrackable)
                    return;

                IsDirty = true;
            }
        }

        protected override void OnCollectionChanged(NotifyCollectionChangedEventArgs e)
        {
            base.OnCollectionChanged(e);
        }

        protected void OnPropertyChanged([CallerMemberName] string propertyName = "")
        {
            var handler = PropertyChanged;
            if (handler != null)
            {
                handler(this, new PropertyChangedEventArgs(propertyName));
            }
        }
    }
}
