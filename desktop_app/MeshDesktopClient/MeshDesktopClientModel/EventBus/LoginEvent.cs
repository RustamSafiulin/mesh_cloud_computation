
using System;

using MeshDesktopClient.Dto;

namespace MeshDesktopClient.EventBus
{
    public class LoginEvent : ITinyMessage
    {
        public object Sender { get; private set; }

        public AuthResponseDto AuthInfo { get; set; }
    }
}
