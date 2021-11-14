
using System;

namespace MeshDesktopClient
{
    public class AppAlreadyStartedException : Exception
    {
        public AppAlreadyStartedException(String message)
            : base(message)
        { }
    }
}
