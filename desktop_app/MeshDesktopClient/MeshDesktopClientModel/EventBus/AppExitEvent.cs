
using System;

namespace MeshDesktopClient.EventBus
{
    public class AppExitEvent : ITinyMessage
    {
        public object Sender { get; private set; }
    }
}
