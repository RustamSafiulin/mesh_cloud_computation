
using System;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class ErrorResponseDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "timestamp")]
        public String Timestamp { get; set; }

        [JsonProperty(PropertyName = "message")]
        public String Message { get; set; }

        [JsonProperty(PropertyName = "error_details")]
        public String Details { get; set; }

        [JsonProperty(PropertyName = "error_code")]
        public int ErrorCode { get; set; }

        #endregion

        public ErrorResponseDto()
        { }
    }
}
