
using System;
using System.IO;
using System.Diagnostics;

using MeshDesktopClient.Helpers;
using MeshDesktopClient.Models;

using Newtonsoft.Json;

namespace MeshDesktopClient
{
    public sealed class Configuration
    {
        public Configuration()
        { }

        private readonly String _configFilePath = "app_params.conf";
        private readonly String _connectionInfoFilePath = "connection_info.json";

        internal class ConnectionInfo
        {
            [JsonProperty(PropertyName = "service_address")]
            public String ServiceAddress;

            public ConnectionInfo()
            {}
        }

        private readonly AppStoredParameters DefaultStoredParams = new AppStoredParameters()
        {
            RememberMe = false,
            Email = "",
            Password = ""
        };

        private readonly String DefaultServiceBaseAddress = @"http://127.0.0.1:8081";
        public String ServiceBaseAddress { get; private set; }

        public AppStoredParameters StoredParameters { get; set; }

        public void ReadConnectionInfo()
        {
            bool isError = false;

            try
            {
                using (var fileStream = File.OpenText(_connectionInfoFilePath))
                {
                    var serializer = new Newtonsoft.Json.JsonSerializer();
                    var connectionInfo = (ConnectionInfo)serializer.Deserialize(fileStream, typeof(ConnectionInfo));

                    ServiceBaseAddress = connectionInfo.ServiceAddress.Trim();
                }

                if (String.IsNullOrEmpty(ServiceBaseAddress))
                {
                    ServiceBaseAddress = DefaultServiceBaseAddress;
                }
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Error during reading connection info file. Error text: {e.Message}");
                isError = true;
            }

            if (isError)
            {
                ServiceBaseAddress = DefaultServiceBaseAddress;
            }
        }

        public void ParseConfiguration()
        {
            AppStoredParameters appStoredParams = default(AppStoredParameters);
            bool isError = false;

            try
            {
                if (!File.Exists(_configFilePath))
                {
                    throw new FileNotFoundException($"Config file {_configFilePath} doesn't exist");
                }

                appStoredParams = SerializationHelpers.ReadFromBinaryFile<AppStoredParameters>(_configFilePath);
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Error during reading configuration from file. Error text: {e.Message}");
                isError = true;
            }

            if (isError)
            {
                StoredParameters = DefaultStoredParams;
                return;
            }

            if (appStoredParams != null)
            {
                StoredParameters = appStoredParams;
                return;
            }

            StoredParameters = DefaultStoredParams;
        }

        public void SaveConfiguration()
        {
            try
            {
                StoredParameters.WriteToBinaryFile(_configFilePath);
            }
            catch (Exception e)
            {
                Logger.Log.Debug($"Error during save configuration to file. Error text: {e.Message}");
            }
        }
    }
}
