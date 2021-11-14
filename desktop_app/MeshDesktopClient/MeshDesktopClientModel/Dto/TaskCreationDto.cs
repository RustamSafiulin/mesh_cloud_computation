using System;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class TaskCreationDto
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "description")]
        public String Description { get; set; }

        #endregion

        public TaskCreationDto()
        { }
    }
}