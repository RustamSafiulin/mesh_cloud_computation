
using System;
using System.Threading.Tasks;
using System.Diagnostics;

using MeshDesktopClient.Models;
using MeshDesktopClient.Helpers;
using MeshDesktopClient.Dto;
using MeshDesktopClient.Service;
using MeshDesktopClient.EventBus;

namespace MeshDesktopClient.ViewModels
{
    public sealed class EntryViewModel : DomainModelBase
    {
        public EntryViewModel(ApplicationEnvironment appEnvironment)
        {
            mAppEnvironment = appEnvironment;

            LoginCommand = new AsyncRelayCommand(LoginHandler, () => !String.IsNullOrEmpty(Email) && !String.IsNullOrEmpty(Password));
            CreateEventHubSubscriptions();

            RememberMe = mAppEnvironment.AppConfiguration.StoredParameters.RememberMe;
            Email = mAppEnvironment.AppConfiguration.StoredParameters.Email;
            Password = mAppEnvironment.AppConfiguration.StoredParameters.Password;
        }

        #region Props and Fields

        private ApplicationEnvironment mAppEnvironment;

        public AsyncRelayCommand LoginCommand { get; set; }

        private String mEmail = String.Empty;
        public String Email
        {
            get { return mEmail; }
            set { SetProperty(ref mEmail, value); }
        }

        private String mPassword = String.Empty;
        public String Password
        {
            get { return mPassword; }
            set { SetProperty(ref mPassword, value); }
        }

        private bool mRememberMe;
        public bool RememberMe
        {
            get { return mRememberMe; }
            set { SetProperty(ref mRememberMe, value); }
        }

        private bool mIsViewVisible = true;
        public bool IsViewVisible
        {
            get { return mIsViewVisible; }
            set { SetProperty(ref mIsViewVisible, value); }
        }

        #endregion

        public void NotifyNeedAppExit()
        {
            if (RememberMe)
            {
                mAppEnvironment.AppConfiguration.StoredParameters.RememberMe = true;
                mAppEnvironment.AppConfiguration.StoredParameters.Email = Email;
                mAppEnvironment.AppConfiguration.StoredParameters.Password = Password;
            }

            mAppEnvironment.EventHub.Publish(new AppExitEvent());
        }

        private void CreateEventHubSubscriptions()
        { }

        private async Task LoginHandler()
        {
            try
            {
                if (mAppEnvironment.ServiceConnector == null)
                {
                    throw new NullReferenceException($"Null reference of { nameof(mAppEnvironment.ServiceConnector) }");
                }

                var authRequestDto = new AuthRequestDto() { Email = this.Email, Password = this.Password };
                var response = await mAppEnvironment.ServiceConnector.Post<AuthRequestDto, AuthResponseDto, ErrorResponseDto>(new Request<AuthRequestDto>
                { 
                    Body = authRequestDto,
                    OpName = OperationType.LOGIN
                });

                Logger.Log.Debug($"Got HTTP status code after LOGIN operation: {response.StatusCode}");

                if (response.SuccessBody != null)
                {
                    Logger.Log.Debug($"Success login at { DateTime.Now.ToString("HH:mm:ss tt") }");
                    IsViewVisible = false;
                    mAppEnvironment.EventHub.Publish(new LoginEvent { AuthInfo = response.SuccessBody });
                }
                else if (response.ErrorBody != null)
                {
                    Logger.Log.Debug($"Login operation was failed. Time: { response.ErrorBody.Timestamp}, Message: { response.ErrorBody.Message }, Details: { response.ErrorBody.Details }");
                }
            }
            catch (ApplicationException e)
            {
                Logger.Log.Debug($"Application exception was caused during login operation. Reason: {e.Message}");
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Exception was caused during login operation. Reason: {e.Message}");
            }
        }
    }
}
