
using System;
using System.Collections.Generic;

namespace MeshDesktopClient.Service
{
    public sealed class Request<T>
    {
        public T Body { get; set; }

        public Dictionary<string, string> UrlSegments { get; set; }

        public String AuthToken { get; set; }

        public OperationType OpName { get; set; }

        public Request()
        { }
    }
}
