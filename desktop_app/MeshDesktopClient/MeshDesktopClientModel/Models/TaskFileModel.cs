
using System;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.Models
{
    public class TaskFileModel : DomainModelBase
    {
        public TaskFileModel()
        {}

        #region Props and fields

        private String mName = String.Empty;
        public String Name
        {
            get { return mName; }
            set { SetProperty(ref mName, value); }
        }

        private Int64 mSize = 0;
        public Int64 Size
        {
            get { return mSize; }
            set { SetProperty(ref mSize, value); }
        }

        #endregion
    }
}