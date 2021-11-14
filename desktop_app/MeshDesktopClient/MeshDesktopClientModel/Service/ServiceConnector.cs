
using System;
using System.IO;
using System.Net;
using System.Collections.Generic;
using System.Threading.Tasks;

using RestSharp;

using Newtonsoft.Json;

using MeshDesktopClient.Helpers;

namespace MeshDesktopClient.Service
{
    public sealed class ServiceConnector
    {
        public ServiceConnector(Options options)
        {
            _options = options;
            _routes = new Dictionary<OperationType, Route>();

            _restClient = new RestClient(_options.BaseAddress);
            _jsonSerializer = new Helpers.JsonSerializer();
        }

        #region Props and Fields

        private Options _options;
        private Dictionary<OperationType, Route> _routes;
        private RestClient _restClient;
        private Helpers.JsonSerializer _jsonSerializer;

        #endregion

        #region Public methods

        public void RegisterRoute(OperationType name, Route route)
        {
            if (_routes.ContainsKey(name))
                throw new ApplicationException("Route already registered");

            _routes[name] = route;
        }

        public async Task<Response<TSuccessBody, TErrorBody>> UploadFile<TRequestBody, TSuccessBody, TErrorBody>(Request<TRequestBody> request, KeyValuePair<string, string> formFile)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.POST);
            restRequest.AddHeader("Content-Type", "multipart/form-data");

            if (!File.Exists(formFile.Value))
                throw new FileNotFoundException($"Uploaded file {formFile.Value} doesn't exist");

            restRequest.AddFile(formFile.Key, formFile.Value);

            return await Execute<TSuccessBody, TErrorBody>(restRequest);
        }

        public async Task<DownloadResponse<TErrorBody>> DownloadFile<TRequestBody, TErrorBody>(Request<TRequestBody> request)
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.GET);

            var response = await _restClient.ExecuteAsync(restRequest);
            if (response.ErrorException != null)
            {
                throw new ApplicationException($"Error retrieving response. Details: {response.ErrorException}");
            }

            var responseResult = new DownloadResponse<TErrorBody>();
            responseResult.StatusCode = response.StatusCode;
            if (response.StatusCode == HttpStatusCode.OK)
            {
                responseResult.SuccessBody = response.RawBytes;
            }
            else
            {
                responseResult.ErrorBody = JsonConvert.DeserializeObject<TErrorBody>(response.Content);
            }

            return responseResult;
        }

        public async Task<Response<TSuccessBody, TErrorBody>> Delete<TRequestBody, TSuccessBody, TErrorBody>(Request<TRequestBody> request)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.DELETE);

            return await Execute<TSuccessBody, TErrorBody>(restRequest);
        }

        public async Task<Response<TSuccessBody, TErrorBody>> Put<TRequestBody, TSuccessBody, TErrorBody>(Request<TRequestBody> request)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.PUT);
            restRequest.AddJsonBody(request.Body);

            return await Execute<TSuccessBody, TErrorBody>(restRequest);
        }

        public async Task<Response<TSuccessBody, TErrorBody>> Post<TRequestBody, TSuccessBody, TErrorBody>(Request<TRequestBody> request, Dictionary<String, String> requestParams = null)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.POST);

            if (request.Body != null)
            {
                restRequest.AddJsonBody(request.Body);
            }

            if (requestParams != null && requestParams.Count > 0)
            {
                foreach (var param in requestParams)
                {
                    restRequest.AddQueryParameter(param.Key, param.Value);
                }
            }

            return await Execute<TSuccessBody, TErrorBody>(restRequest);
        }

        public async Task<Response<TSuccessBody, TErrorBody>> Get<TRequestBody, TSuccessBody, TErrorBody>(Request<TRequestBody> request, Dictionary<String, String> requestParams = null)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var restRequest = ConfigureRestRequest(request, RestSharp.Method.GET);

            if (requestParams != null && requestParams.Count > 0)
            {
                foreach (var param in requestParams)
                {
                    restRequest.AddParameter(param.Key, param.Value);
                }
            }

            return await Execute<TSuccessBody, TErrorBody>(restRequest);
        }

        #endregion

        #region Private methods

        private Route GetRegisteredRoute(OperationType name)
        {
            if (!_routes.ContainsKey(name))
                throw new ApplicationException("Operation doesn't registered");

            return _routes[name];
        }

        private RestSharp.RestRequest ConfigureRestRequest<TRequestBody>(Request<TRequestBody> request, RestSharp.Method httpMethod)
        {
            var route = GetRegisteredRoute(request.OpName);
            var restRequest = new RestRequest(route.Path, httpMethod, DataFormat.Json);
            restRequest.JsonSerializer = _jsonSerializer;

            if (!String.IsNullOrEmpty(request.AuthToken))
            {
                restRequest.AddHeader("Authorization", String.Concat("Bearer_", request.AuthToken));
            }

            if (request.UrlSegments != null)
            {
                foreach (var urlSegment in request.UrlSegments)
                {
                    restRequest.AddUrlSegment(urlSegment.Key, urlSegment.Value);
                }
            }

            return restRequest;
        }

        private async Task<Response<TSuccessBody, TErrorBody>> Execute<TSuccessBody, TErrorBody>(RestSharp.RestRequest restRequest)
            where TSuccessBody : new()
            where TErrorBody : new()
        {
            var response = await _restClient.ExecuteAsync(restRequest);

            if (response.ErrorException != null)
            {
                throw new ApplicationException($"Error retrieving response. Details: {response.ErrorException}");
            }

            var responseResult = new Response<TSuccessBody, TErrorBody>();
            responseResult.StatusCode = response.StatusCode;
            if (response.StatusCode == HttpStatusCode.OK)
            {
                responseResult.SuccessBody = JsonConvert.DeserializeObject<TSuccessBody>(response.Content);
            }
            else
            {
                responseResult.ErrorBody = JsonConvert.DeserializeObject<TErrorBody>(response.Content);
            }

            return responseResult;
        }

        #endregion
    }
}
