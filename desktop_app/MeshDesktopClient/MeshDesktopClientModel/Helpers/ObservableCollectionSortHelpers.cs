
using System;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;

namespace MeshDesktopClient.Helpers
{
    public static class ObservableCollectionSortHelpers
    {
        public static void Sort<TSource, TKey>(this ObservableCollection<TSource> source, Func<TSource, TKey> keySelector)
        {
            List<TSource> sortedList = source.OrderBy(keySelector).ToList();
            source.Clear();
            foreach (var sortedItem in sortedList)
            {
                source.Add(sortedItem);
            }
        }
    }
}
