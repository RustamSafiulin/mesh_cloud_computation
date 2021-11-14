
using System;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class TaskDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "email")]
        public String Id { get; set; }

        [JsonProperty(PropertyName = "account_id")]
        public String AccountId { get; set; }

        [JsonProperty(PropertyName = "description")]
        public String Description { get; set; }

        [JsonProperty(PropertyName = "started_at")]
        public Int64 StartedAt;

        [JsonProperty(PropertyName = "completed_at")]
        public Int64 CompletedAt;

        [JsonProperty(PropertyName = "state")]
        public Int32 State;

        [JsonProperty(PropertyName = "state_text")]
        public String StateText;

        #endregion

        public TaskDto()
        { }
    }
}