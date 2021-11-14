
using System;
using System.Collections.Generic;

using Newtonsoft.Json;

namespace MeshDesktopClient.Dto
{
    public class GenericPageResponseDto<T> where T : new()
    {
        #region Props and Fields

        [JsonProperty(PropertyName = "page_total")]
        public int PageTotal { get; set; }

        [JsonProperty(PropertyName = "page_number")]
        public int PageNumber { get; set; }

        [JsonProperty(PropertyName = "page_size")]
        public int RequestedPageSize { get; set; }

        [JsonProperty(PropertyName = "items")]
        public List<T> Items { get; set; }

        #endregion

        public GenericPageResponseDto()
        { }
    }
}
