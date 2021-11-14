
using System;
using System.Net;
using System.Collections.Generic;

namespace MeshDesktopClient.Service
{
    public sealed class Response<TSuccess, TError>
    {
        public TSuccess SuccessBody { get; set; }

        public TError ErrorBody { get; set; }

        public HttpStatusCode StatusCode { get; set; }

        public Response()
        {}
    }
}
