
using System;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class AuthRequestDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "email")]
        public String Email { get; set; }

        [JsonProperty(PropertyName = "password")]
        public String Password { get; set; }

        #endregion

        public AuthRequestDto()
        { }
    }
}
