
using System;

using RestSharp;

namespace MeshDesktopClient.Service
{
    public sealed class Options
    {
        public DataFormat ExchangeFormat { get; set; }
        public string BaseAddress { get; set; }

        public Options()
        { }
    }
}
