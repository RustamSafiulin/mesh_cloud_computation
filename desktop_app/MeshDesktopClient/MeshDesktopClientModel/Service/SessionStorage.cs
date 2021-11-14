
using System;

using MeshDesktopClient.Dto;

namespace MeshDesktopClient.Service
{
    public sealed class SessionStorage
    {
        private SessionStorage()
        { }

        private static SessionStorage source = null;
        private static readonly object threadlock = new object();

        public static SessionStorage Instance
        {
            get
            {
                lock (threadlock)
                {
                    if (source == null)
                        source = new SessionStorage();

                    return source;
                }
            }
        }

        public AuthResponseDto AuthInfo { get; set; }
    }
}
