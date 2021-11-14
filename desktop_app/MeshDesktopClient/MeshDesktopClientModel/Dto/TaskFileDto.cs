using System;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class TaskFileDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "email")]
        public String Id { get; set; }

        [JsonProperty(PropertyName = "task_id")]
        public String TaskId { get; set; }

        [JsonProperty(PropertyName = "name")]
        public String Name { get; set; }

        [JsonProperty(PropertyName = "size")]
        public Int64 Size;

        #endregion

        public TaskFileDto()
        { }
    }
}