
using System;
using System.Collections.ObjectModel;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.Models
{
    public class TaskModel : DomainModelBase
    {
        public TaskModel()
        {
            mTaskFiles = new ObservableCollection<TaskFileModel>();
        }

        #region Props and fields

        private String mId = String.Empty;
        public String Id
        {
            get { return mId; }
            set { SetProperty(ref mId, value); }
        }

        private String mAccountId = String.Empty;
        public String AccountId
        { 
            get { return mAccountId; }
            set { SetProperty(ref mAccountId, value); }
        }

        private String mDescription = String.Empty;
        public String Description
        { 
            get { return mDescription; }
            set { SetProperty(ref mDescription, value); }
        }

        private Int64 mStartedAt = 0;
        public Int64 StartedAt
        {
            get { return mStartedAt; }
            set { SetProperty(ref mStartedAt, value); }
        }

        private Int64 mCompletedAt = 0;
        public Int64 CompletedAt
        {
            get { return mCompletedAt; }
            set { SetProperty(ref mCompletedAt, value); }
        }

        private Int32 mState = 0;
        public Int32 State
        {
            get { return mState; }
            set { SetProperty(ref mState, value); }
        }

        private String mStateText = String.Empty;
        public String StateText
        {
            get { return mStateText; }
            set { SetProperty(ref mStateText, value); }
        }

        private ObservableCollection<TaskFileModel> mTaskFiles;
        public ObservableCollection<TaskFileModel> TaskFiles
        {
            get { return mTaskFiles; }
            set { SetProperty(ref mTaskFiles, value); }
        }

        #endregion
    }
}