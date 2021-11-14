
using System.ComponentModel;

namespace MeshDesktopClient.Helpers
{
    public class OnSortCommandParams
    {
        public ListSortDirection SortDirection { get; set; }
        public string SortMemberPath { get; set; }

        public OnSortCommandParams()
        { }
    }
}
