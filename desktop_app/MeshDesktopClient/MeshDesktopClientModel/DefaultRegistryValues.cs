
using System;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.Models
{
    public static class DefaultRegistryValues
    {
        private static readonly string AppRootKeyPath = @"HKEY_LOCAL_MACHINE\Software\Wow6432Node\EkanSpb\EkanDB";

        public static string GetLogsDir()
        {
            var result = RegistryHelpers.GetValue<string>(AppRootKeyPath, "LogsDir");
            return result;
        }

        public static string GetConfigsDir()
        {
            var result = RegistryHelpers.GetValue<string>(AppRootKeyPath, "ConfigsDir");
            return result;
        }
    }
}
