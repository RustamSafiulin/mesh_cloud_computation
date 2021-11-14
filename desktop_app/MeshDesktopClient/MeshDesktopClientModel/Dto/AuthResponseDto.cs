
using System;
using System.Collections.Generic;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class AuthResponseDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "account_id")]
        public String AccountId { get; set; }

        [JsonProperty(PropertyName = "username")]
        public String UserName { get; set; }

        [JsonProperty(PropertyName = "email")]
        public String Email { get; set; }

        [JsonProperty(PropertyName = "session_token")]
        public String Token { get; set; }

        #endregion

        public AuthResponseDto()
        { }
    }
}
