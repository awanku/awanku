// tslint:disable

import * as request from "superagent";
import { Response, SuperAgentRequest, SuperAgentStatic } from "superagent";

export type RequestHeaders = {
  [header: string]: string;
};
export type RequestHeadersHandler = (headers: RequestHeaders) => RequestHeaders;

export type ConfigureAgentHandler = (
  agent: SuperAgentStatic
) => SuperAgentStatic;

export type ConfigureRequestHandler = (
  agent: SuperAgentRequest
) => SuperAgentRequest;

export type CallbackHandler = (err: any, res?: request.Response) => void;

export type apihelper_HTTPError = {
  errors?: {} & {
    [key: string]: string;
  };
  type?: string;
} & {
  [key: string]: any;
};

export type apihelper_InternalServerError = {
  error?: string;
} & {
  [key: string]: any;
};

export type auth_getProviderCallbackParam = {
  code?: string;
} & {
  [key: string]: any;
};

export type auth_getProviderConnectParam = {
  redirect_to: string;
  state: string;
  user_id?: string;
} & {
  [key: string]: any;
};

export type auth_postTokenParam = {
  code?: string;
  grant_type?: string;
  refresh_token?: string;
} & {
  [key: string]: any;
};

export type contract_AuthStore = {} & {
  [key: string]: any;
};

export type contract_UserStore = {} & {
  [key: string]: any;
};

export type core_OauthAuthorizationCode = {
  code?: string;
  expiresAt?: string;
  userID?: number;
} & {
  [key: string]: any;
};

export type core_OauthToken = {
  accessToken?: Array<number>;
  accessTokenHash?: Array<number>;
  deletedAt?: string;
  expiresAt?: string;
  id?: number;
  refreshToken?: Array<number>;
  refreshTokenHash?: Array<number>;
  requesterIP?: string;
  requesterUserAgent?: string;
  userID?: number;
} & {
  [key: string]: any;
};

export type core_User = {
  created_at?: string;
  email?: string;
  github_login_username?: string;
  google_login_email?: string;
  id?: number;
  name?: string;
  updated_at?: string;
} & {
  [key: string]: any;
};

export type coreapi_statusResponse = {
  database?: hansip_ClusterHealth;
} & {
  [key: string]: any;
};

export type hansip_ClusterHealth = {
  primary_ok?: boolean;
  replicas_ok?: boolean;
} & {
  [key: string]: any;
};

export type oauth2_Token = {
  access_token?: string;
  expiry?: string;
  refresh_token?: string;
  token_type?: string;
} & {
  [key: string]: any;
};

export type Response_api_v1_auth_provider_callback_301 = string;

export type Response_api_v1_auth_provider_connect_301 = string;

export type Logger = {
  log: (line: string) => any;
};

export interface ResponseWithBody<S extends number, T> extends Response {
  status: S;
  body: T;
}

export type QueryParameters = {
  [param: string]: any;
};

export interface CommonRequestOptions {
  $queryParameters?: QueryParameters;
  $domain?: string;
  $path?: string | ((path: string) => string);
  $retries?: number; // number of retries; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#retrying-requests
  $timeout?: number; // request timeout in milliseconds; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#timeouts
  $deadline?: number; // request deadline in milliseconds; see: https://github.com/visionmedia/superagent/blob/master/docs/index.md#timeouts
}

/**
 *
 * @class API
 * @param {(string)} [domainOrOptions] - The project domain.
 */
export class API {
  private domain: string = "";
  private errorHandlers: CallbackHandler[] = [];
  private requestHeadersHandler?: RequestHeadersHandler;
  private configureAgentHandler?: ConfigureAgentHandler;
  private configureRequestHandler?: ConfigureRequestHandler;

  constructor(domain?: string, private logger?: Logger) {
    if (domain) {
      this.domain = domain;
    }
  }

  getDomain() {
    return this.domain;
  }

  addErrorHandler(handler: CallbackHandler) {
    this.errorHandlers.push(handler);
  }

  setRequestHeadersHandler(handler: RequestHeadersHandler) {
    this.requestHeadersHandler = handler;
  }

  setConfigureAgentHandler(handler: ConfigureAgentHandler) {
    this.configureAgentHandler = handler;
  }

  setConfigureRequestHandler(handler: ConfigureRequestHandler) {
    this.configureRequestHandler = handler;
  }

  private request(
    method: string,
    url: string,
    body: any,
    headers: RequestHeaders,
    queryParameters: QueryParameters,
    form: any,
    reject: CallbackHandler,
    resolve: CallbackHandler,
    opts: CommonRequestOptions
  ) {
    if (this.logger) {
      this.logger.log(`Call ${method} ${url}`);
    }

    const agent = this.configureAgentHandler
      ? this.configureAgentHandler(request.default)
      : request.default;

    let req = agent(method, url);
    if (this.configureRequestHandler) {
      req = this.configureRequestHandler(req);
    }

    req = req.query(queryParameters);

    if (this.requestHeadersHandler) {
      headers = this.requestHeadersHandler({
        ...headers,
      });
    }

    req.set(headers);

    if (body) {
      req.send(body);

      if (typeof body === "object" && !(body.constructor.name === "Buffer")) {
        headers["content-type"] = "application/json";
      }
    }

    if (Object.keys(form).length > 0) {
      req.type("form");
      req.send(form);
    }

    if (opts.$retries && opts.$retries > 0) {
      req.retry(opts.$retries);
    }

    if (
      (opts.$timeout && opts.$timeout > 0) ||
      (opts.$deadline && opts.$deadline > 0)
    ) {
      req.timeout({
        deadline: opts.$deadline,
        response: opts.$timeout,
      });
    }

    req.end((error, response) => {
      // an error will also be emitted for a 4xx and 5xx status code
      // the error object will then have error.status and error.response fields
      // see superagent error handling: https://github.com/visionmedia/superagent/blob/master/docs/index.md#error-handling
      if (error) {
        reject(error);
        this.errorHandlers.forEach((handler) => handler(error));
      } else {
        resolve(response);
      }
    });
  }

  private convertParameterCollectionFormat<T>(
    param: T,
    collectionFormat: string | undefined
  ): T | string {
    if (Array.isArray(param) && param.length >= 2) {
      switch (collectionFormat) {
        case "csv":
          return param.join(",");
        case "ssv":
          return param.join(" ");
        case "tsv":
          return param.join("\t");
        case "pipes":
          return param.join("|");
        default:
          return param;
      }
    }

    return param;
  }

  api_statusURL(parameters: {} & CommonRequestOptions): string {
    let queryParameters: QueryParameters = {};
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/status";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    if (parameters.$queryParameters) {
      queryParameters = {
        ...queryParameters,
        ...parameters.$queryParameters,
      };
    }

    let keys = Object.keys(queryParameters);
    return (
      domain +
      path +
      (keys.length > 0
        ? "?" +
          keys
            .map((key) => key + "=" + encodeURIComponent(queryParameters[key]))
            .join("&")
        : "")
    );
  }

  /**
   * Get API health status
   * @method
   * @name API#api_status
   */
  api_status(
    parameters: {} & CommonRequestOptions
  ): Promise<ResponseWithBody<200, coreapi_statusResponse>> {
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/status";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    let body: any;
    let queryParameters: QueryParameters = {};
    let headers: RequestHeaders = {};
    let form: any = {};
    return new Promise((resolve, reject) => {
      headers["accept"] = "application/json";

      if (parameters.$queryParameters) {
        queryParameters = {
          ...queryParameters,
          ...parameters.$queryParameters,
        };
      }

      this.request(
        "GET",
        domain + path,
        body,
        headers,
        queryParameters,
        form,
        reject,
        resolve,
        parameters
      );
    });
  }

  api_v1_auth_exchangeTokenURL(
    parameters: {
      param: auth_postTokenParam;
    } & CommonRequestOptions
  ): string {
    let queryParameters: QueryParameters = {};
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/token";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    if (parameters.$queryParameters) {
      queryParameters = {
        ...queryParameters,
        ...parameters.$queryParameters,
      };
    }

    queryParameters = {};

    let keys = Object.keys(queryParameters);
    return (
      domain +
      path +
      (keys.length > 0
        ? "?" +
          keys
            .map((key) => key + "=" + encodeURIComponent(queryParameters[key]))
            .join("&")
        : "")
    );
  }

  /**
   * Exchange authorization code for authentication token
   * @method
   * @name API#api_v1_auth_exchangeToken
   * @param {} param - Request body
   */
  api_v1_auth_exchangeToken(
    parameters: {
      param: auth_postTokenParam;
    } & CommonRequestOptions
  ): Promise<
    | ResponseWithBody<200, oauth2_Token>
    | ResponseWithBody<400, apihelper_HTTPError>
    | ResponseWithBody<401, apihelper_HTTPError>
    | ResponseWithBody<500, apihelper_InternalServerError>
  > {
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/token";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    let body: any;
    let queryParameters: QueryParameters = {};
    let headers: RequestHeaders = {};
    let form: any = {};
    return new Promise((resolve, reject) => {
      headers["accept"] = "application/json";
      headers["Content-Type"] = "application/json";

      if (parameters["param"] !== undefined) {
        body = parameters["param"];
      }

      if (parameters["param"] === undefined) {
        reject(new Error("Missing required  parameter: param"));
        return;
      }

      if (parameters.$queryParameters) {
        queryParameters = {
          ...queryParameters,
          ...parameters.$queryParameters,
        };
      }

      form = queryParameters;
      queryParameters = {};

      this.request(
        "POST",
        domain + path,
        body,
        headers,
        queryParameters,
        form,
        reject,
        resolve,
        parameters
      );
    });
  }

  api_v1_auth_provider_callbackURL(
    parameters: {
      provider: "github" | "google";
      code?: string;
    } & CommonRequestOptions
  ): string {
    let queryParameters: QueryParameters = {};
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/{provider}/callback";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    path = path.replace(
      "{provider}",
      `${encodeURIComponent(
        this.convertParameterCollectionFormat(
          parameters["provider"],
          ""
        ).toString()
      )}`
    );
    if (parameters["code"] !== undefined) {
      queryParameters["code"] = this.convertParameterCollectionFormat(
        parameters["code"],
        ""
      );
    }

    if (parameters.$queryParameters) {
      queryParameters = {
        ...queryParameters,
        ...parameters.$queryParameters,
      };
    }

    let keys = Object.keys(queryParameters);
    return (
      domain +
      path +
      (keys.length > 0
        ? "?" +
          keys
            .map((key) => key + "=" + encodeURIComponent(queryParameters[key]))
            .join("&")
        : "")
    );
  }

  /**
   * Auth provider callback
   * @method
   * @name API#api_v1_auth_provider_callback
   * @param {string} provider - Auth provider
   * @param {string} code -
   */
  api_v1_auth_provider_callback(
    parameters: {
      provider: "github" | "google";
      code?: string;
    } & CommonRequestOptions
  ): Promise<
    | ResponseWithBody<301, Response_api_v1_auth_provider_callback_301>
    | ResponseWithBody<400, apihelper_HTTPError>
    | ResponseWithBody<401, apihelper_HTTPError>
    | ResponseWithBody<500, apihelper_InternalServerError>
  > {
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/{provider}/callback";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    let body: any;
    let queryParameters: QueryParameters = {};
    let headers: RequestHeaders = {};
    let form: any = {};
    return new Promise((resolve, reject) => {
      headers["accept"] = "application/json";

      path = path.replace(
        "{provider}",
        `${encodeURIComponent(
          this.convertParameterCollectionFormat(
            parameters["provider"],
            ""
          ).toString()
        )}`
      );

      if (parameters["provider"] === undefined) {
        reject(new Error("Missing required  parameter: provider"));
        return;
      }

      if (parameters["code"] !== undefined) {
        queryParameters["code"] = this.convertParameterCollectionFormat(
          parameters["code"],
          ""
        );
      }

      if (parameters.$queryParameters) {
        queryParameters = {
          ...queryParameters,
          ...parameters.$queryParameters,
        };
      }

      this.request(
        "GET",
        domain + path,
        body,
        headers,
        queryParameters,
        form,
        reject,
        resolve,
        parameters
      );
    });
  }

  api_v1_auth_provider_connectURL(
    parameters: {
      provider: "github" | "google";
      redirectTo: string;
      state: string;
      userId?: string;
    } & CommonRequestOptions
  ): string {
    let queryParameters: QueryParameters = {};
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/{provider}/connect";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    path = path.replace(
      "{provider}",
      `${encodeURIComponent(
        this.convertParameterCollectionFormat(
          parameters["provider"],
          ""
        ).toString()
      )}`
    );
    if (parameters["redirectTo"] !== undefined) {
      queryParameters["redirect_to"] = this.convertParameterCollectionFormat(
        parameters["redirectTo"],
        ""
      );
    }

    if (parameters["state"] !== undefined) {
      queryParameters["state"] = this.convertParameterCollectionFormat(
        parameters["state"],
        ""
      );
    }

    if (parameters["userId"] !== undefined) {
      queryParameters["user_id"] = this.convertParameterCollectionFormat(
        parameters["userId"],
        ""
      );
    }

    if (parameters.$queryParameters) {
      queryParameters = {
        ...queryParameters,
        ...parameters.$queryParameters,
      };
    }

    let keys = Object.keys(queryParameters);
    return (
      domain +
      path +
      (keys.length > 0
        ? "?" +
          keys
            .map((key) => key + "=" + encodeURIComponent(queryParameters[key]))
            .join("&")
        : "")
    );
  }

  /**
   * Auth provider connect
   * @method
   * @name API#api_v1_auth_provider_connect
   * @param {string} provider - Auth provider
   * @param {string} redirectTo -
   * @param {string} state -
   * @param {string} userId -
   */
  api_v1_auth_provider_connect(
    parameters: {
      provider: "github" | "google";
      redirectTo: string;
      state: string;
      userId?: string;
    } & CommonRequestOptions
  ): Promise<
    | ResponseWithBody<301, Response_api_v1_auth_provider_connect_301>
    | ResponseWithBody<400, apihelper_HTTPError>
    | ResponseWithBody<401, apihelper_HTTPError>
    | ResponseWithBody<500, apihelper_InternalServerError>
  > {
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/auth/{provider}/connect";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    let body: any;
    let queryParameters: QueryParameters = {};
    let headers: RequestHeaders = {};
    let form: any = {};
    return new Promise((resolve, reject) => {
      headers["accept"] = "application/json";

      path = path.replace(
        "{provider}",
        `${encodeURIComponent(
          this.convertParameterCollectionFormat(
            parameters["provider"],
            ""
          ).toString()
        )}`
      );

      if (parameters["provider"] === undefined) {
        reject(new Error("Missing required  parameter: provider"));
        return;
      }

      if (parameters["redirectTo"] !== undefined) {
        queryParameters["redirect_to"] = this.convertParameterCollectionFormat(
          parameters["redirectTo"],
          ""
        );
      }

      if (parameters["redirectTo"] === undefined) {
        reject(new Error("Missing required  parameter: redirectTo"));
        return;
      }

      if (parameters["state"] !== undefined) {
        queryParameters["state"] = this.convertParameterCollectionFormat(
          parameters["state"],
          ""
        );
      }

      if (parameters["state"] === undefined) {
        reject(new Error("Missing required  parameter: state"));
        return;
      }

      if (parameters["userId"] !== undefined) {
        queryParameters["user_id"] = this.convertParameterCollectionFormat(
          parameters["userId"],
          ""
        );
      }

      if (parameters.$queryParameters) {
        queryParameters = {
          ...queryParameters,
          ...parameters.$queryParameters,
        };
      }

      this.request(
        "GET",
        domain + path,
        body,
        headers,
        queryParameters,
        form,
        reject,
        resolve,
        parameters
      );
    });
  }

  api_v1_users_getMeURL(parameters: {} & CommonRequestOptions): string {
    let queryParameters: QueryParameters = {};
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/users/me";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    if (parameters.$queryParameters) {
      queryParameters = {
        ...queryParameters,
        ...parameters.$queryParameters,
      };
    }

    let keys = Object.keys(queryParameters);
    return (
      domain +
      path +
      (keys.length > 0
        ? "?" +
          keys
            .map((key) => key + "=" + encodeURIComponent(queryParameters[key]))
            .join("&")
        : "")
    );
  }

  /**
   * Get current user data
   * @method
   * @name API#api_v1_users_getMe
   */
  api_v1_users_getMe(
    parameters: {} & CommonRequestOptions
  ): Promise<
    | ResponseWithBody<200, core_User>
    | ResponseWithBody<400, apihelper_HTTPError>
    | ResponseWithBody<401, apihelper_HTTPError>
    | ResponseWithBody<500, apihelper_InternalServerError>
  > {
    const domain = parameters.$domain ? parameters.$domain : this.domain;
    let path = "/v1/users/me";
    if (parameters.$path) {
      path =
        typeof parameters.$path === "function"
          ? parameters.$path(path)
          : parameters.$path;
    }

    let body: any;
    let queryParameters: QueryParameters = {};
    let headers: RequestHeaders = {};
    let form: any = {};
    return new Promise((resolve, reject) => {
      headers["accept"] = "application/json";

      if (parameters.$queryParameters) {
        queryParameters = {
          ...queryParameters,
          ...parameters.$queryParameters,
        };
      }

      this.request(
        "GET",
        domain + path,
        body,
        headers,
        queryParameters,
        form,
        reject,
        resolve,
        parameters
      );
    });
  }
}

export default API;
