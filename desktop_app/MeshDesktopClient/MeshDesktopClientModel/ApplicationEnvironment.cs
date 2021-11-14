
using System;

using MeshDesktopClient.EventBus;
using MeshDesktopClient.Service;
using MeshDesktopClient.Helpers;

namespace MeshDesktopClient
{
    public class ApplicationEnvironment
    {
        public ApplicationEnvironment()
        { }

        #region Props and Fields

        public Configuration AppConfiguration { get; set; }
        public TinyMessengerHub EventHub { get; set; }
        public ServiceConnector ServiceConnector { get; set; }

        #endregion

        public void SetupApplicationEnvironment()
        {
            Logger.InitLogger();

            InitConfiguration();
            InitEventHub();
            InitServiceConnector();
        }

        public void TearDownApplicationEnvironment()
        {
            AppConfiguration.SaveConfiguration();
        }

        private void InitConfiguration()
        {
            Logger.Log.Info("Read stored configuration parameters");

            AppConfiguration = new Configuration();
            AppConfiguration.ParseConfiguration();
            AppConfiguration.ReadConnectionInfo();
        }

        private void InitEventHub()
        {
            Logger.Log.Info("Initialize Event Hub");

            EventHub = new TinyMessengerHub();
        }


        private void InitServiceConnector()
        {
            Logger.Log.Info("Initialize Web Service connector");
            ServiceConnector = new ServiceConnector(new Options { 
                BaseAddress = AppConfiguration.ServiceBaseAddress, 
                ExchangeFormat = RestSharp.DataFormat.Json 
            });

            ServiceConnector.RegisterRoute(OperationType.LOGIN, new Route { Path = "api/v1/accounts/signin" });

            ServiceConnector.RegisterRoute(OperationType.CREATE_TASK, new Route { Path = "api/v1/tasks" });
            ServiceConnector.RegisterRoute(OperationType.GET_ALL_ACCOUNTS_TASKS, new Route { Path = "api/v1/tasks" });
            ServiceConnector.RegisterRoute(OperationType.UPLOAD_TASK_DATA, new Route { Path = "api/v1/tasks/{id}/upload" });
            ServiceConnector.RegisterRoute(OperationType.DOWNLOAD_TASK_DATA, new Route { Path = "api/v1/tasks/{id}/download" });
            ServiceConnector.RegisterRoute(OperationType.START_TASK, new Route { Path = "api/v1/tasks/{id}/start" });
            ServiceConnector.RegisterRoute(OperationType.STOP_TASK, new Route { Path = "api/v1/tasks/{id}/stop" });
            ServiceConnector.RegisterRoute(OperationType.DELETE_TASK, new Route { Path = "api/v1/tasks/{id}" });
        }
    }
}
