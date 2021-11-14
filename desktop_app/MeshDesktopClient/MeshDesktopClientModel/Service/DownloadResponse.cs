
using System;
using System.Net;

namespace MeshDesktopClient.Service
{
    public sealed class DownloadResponse<TError>
    {
        public byte[] SuccessBody { get; set; }

        public TError ErrorBody { get; set; }

        public HttpStatusCode StatusCode { get; set; }

        public DownloadResponse()
        {}
    }
}
