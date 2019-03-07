/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.lasagna = (function() {

    /**
     * Namespace lasagna.
     * @exports lasagna
     * @namespace
     */
    var lasagna = {};

    lasagna.srv = (function() {

        /**
         * Namespace srv.
         * @memberof lasagna
         * @namespace
         */
        var srv = {};

        srv.internal = (function() {

            /**
             * Namespace internal.
             * @memberof lasagna.srv
             * @namespace
             */
            var internal = {};

            internal.recipes = (function() {

                /**
                 * Namespace recipes.
                 * @memberof lasagna.srv.internal
                 * @namespace
                 */
                var recipes = {};

                recipes.service = (function() {

                    /**
                     * Namespace service.
                     * @memberof lasagna.srv.internal.recipes
                     * @namespace
                     */
                    var service = {};

                    service.InternalRecipesService = (function() {

                        /**
                         * Constructs a new InternalRecipesService service.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @classdesc Represents an InternalRecipesService
                         * @extends $protobuf.rpc.Service
                         * @constructor
                         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                         */
                        function InternalRecipesService(rpcImpl, requestDelimited, responseDelimited) {
                            $protobuf.rpc.Service.call(this, rpcImpl, requestDelimited, responseDelimited);
                        }

                        (InternalRecipesService.prototype = Object.create($protobuf.rpc.Service.prototype)).constructor = InternalRecipesService;

                        /**
                         * Creates new InternalRecipesService service using the specified rpc implementation.
                         * @function create
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @static
                         * @param {$protobuf.RPCImpl} rpcImpl RPC implementation
                         * @param {boolean} [requestDelimited=false] Whether requests are length-delimited
                         * @param {boolean} [responseDelimited=false] Whether responses are length-delimited
                         * @returns {InternalRecipesService} RPC service. Useful where requests and/or responses are streamed.
                         */
                        InternalRecipesService.create = function create(rpcImpl, requestDelimited, responseDelimited) {
                            return new this(rpcImpl, requestDelimited, responseDelimited);
                        };

                        /**
                         * Callback as used by {@link lasagna.srv.internal.recipes.service.InternalRecipesService#createRecipe}.
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @typedef CreateRecipeCallback
                         * @type {function}
                         * @param {Error|null} error Error, if any
                         * @param {lasagna.srv.internal.recipes.service.Response} [response] Response
                         */

                        /**
                         * Calls CreateRecipe.
                         * @function createRecipe
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @instance
                         * @param {lasagna.srv.internal.recipes.service.IRecipe} request Recipe message or plain object
                         * @param {lasagna.srv.internal.recipes.service.InternalRecipesService.CreateRecipeCallback} callback Node-style callback called with the error, if any, and Response
                         * @returns {undefined}
                         * @variation 1
                         */
                        Object.defineProperty(InternalRecipesService.prototype.createRecipe = function createRecipe(request, callback) {
                            return this.rpcCall(createRecipe, $root.lasagna.srv.internal.recipes.service.Recipe, $root.lasagna.srv.internal.recipes.service.Response, request, callback);
                        }, "name", { value: "CreateRecipe" });

                        /**
                         * Calls CreateRecipe.
                         * @function createRecipe
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @instance
                         * @param {lasagna.srv.internal.recipes.service.IRecipe} request Recipe message or plain object
                         * @returns {Promise<lasagna.srv.internal.recipes.service.Response>} Promise
                         * @variation 2
                         */

                        /**
                         * Callback as used by {@link lasagna.srv.internal.recipes.service.InternalRecipesService#getRecipes}.
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @typedef GetRecipesCallback
                         * @type {function}
                         * @param {Error|null} error Error, if any
                         * @param {lasagna.srv.internal.recipes.service.Response} [response] Response
                         */

                        /**
                         * Calls GetRecipes.
                         * @function getRecipes
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @instance
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest} request GetRequest message or plain object
                         * @param {lasagna.srv.internal.recipes.service.InternalRecipesService.GetRecipesCallback} callback Node-style callback called with the error, if any, and Response
                         * @returns {undefined}
                         * @variation 1
                         */
                        Object.defineProperty(InternalRecipesService.prototype.getRecipes = function getRecipes(request, callback) {
                            return this.rpcCall(getRecipes, $root.lasagna.srv.internal.recipes.service.GetRequest, $root.lasagna.srv.internal.recipes.service.Response, request, callback);
                        }, "name", { value: "GetRecipes" });

                        /**
                         * Calls GetRecipes.
                         * @function getRecipes
                         * @memberof lasagna.srv.internal.recipes.service.InternalRecipesService
                         * @instance
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest} request GetRequest message or plain object
                         * @returns {Promise<lasagna.srv.internal.recipes.service.Response>} Promise
                         * @variation 2
                         */

                        return InternalRecipesService;
                    })();

                    service.Recipe = (function() {

                        /**
                         * Properties of a Recipe.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @interface IRecipe
                         * @property {string|null} [id] Recipe id
                         * @property {string|null} [name] Recipe name
                         * @property {string|null} [description] Recipe description
                         */

                        /**
                         * Constructs a new Recipe.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @classdesc Represents a Recipe.
                         * @implements IRecipe
                         * @constructor
                         * @param {lasagna.srv.internal.recipes.service.IRecipe=} [properties] Properties to set
                         */
                        function Recipe(properties) {
                            if (properties)
                                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                                    if (properties[keys[i]] != null)
                                        this[keys[i]] = properties[keys[i]];
                        }

                        /**
                         * Recipe id.
                         * @member {string} id
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @instance
                         */
                        Recipe.prototype.id = "";

                        /**
                         * Recipe name.
                         * @member {string} name
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @instance
                         */
                        Recipe.prototype.name = "";

                        /**
                         * Recipe description.
                         * @member {string} description
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @instance
                         */
                        Recipe.prototype.description = "";

                        /**
                         * Creates a new Recipe instance using the specified properties.
                         * @function create
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IRecipe=} [properties] Properties to set
                         * @returns {lasagna.srv.internal.recipes.service.Recipe} Recipe instance
                         */
                        Recipe.create = function create(properties) {
                            return new Recipe(properties);
                        };

                        /**
                         * Encodes the specified Recipe message. Does not implicitly {@link lasagna.srv.internal.recipes.service.Recipe.verify|verify} messages.
                         * @function encode
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IRecipe} message Recipe message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        Recipe.encode = function encode(message, writer) {
                            if (!writer)
                                writer = $Writer.create();
                            if (message.id != null && message.hasOwnProperty("id"))
                                writer.uint32(/* id 1, wireType 2 =*/10).string(message.id);
                            if (message.name != null && message.hasOwnProperty("name"))
                                writer.uint32(/* id 2, wireType 2 =*/18).string(message.name);
                            if (message.description != null && message.hasOwnProperty("description"))
                                writer.uint32(/* id 3, wireType 2 =*/26).string(message.description);
                            return writer;
                        };

                        /**
                         * Encodes the specified Recipe message, length delimited. Does not implicitly {@link lasagna.srv.internal.recipes.service.Recipe.verify|verify} messages.
                         * @function encodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IRecipe} message Recipe message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        Recipe.encodeDelimited = function encodeDelimited(message, writer) {
                            return this.encode(message, writer).ldelim();
                        };

                        /**
                         * Decodes a Recipe message from the specified reader or buffer.
                         * @function decode
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @param {number} [length] Message length if known beforehand
                         * @returns {lasagna.srv.internal.recipes.service.Recipe} Recipe
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        Recipe.decode = function decode(reader, length) {
                            if (!(reader instanceof $Reader))
                                reader = $Reader.create(reader);
                            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.lasagna.srv.internal.recipes.service.Recipe();
                            while (reader.pos < end) {
                                var tag = reader.uint32();
                                switch (tag >>> 3) {
                                case 1:
                                    message.id = reader.string();
                                    break;
                                case 2:
                                    message.name = reader.string();
                                    break;
                                case 3:
                                    message.description = reader.string();
                                    break;
                                default:
                                    reader.skipType(tag & 7);
                                    break;
                                }
                            }
                            return message;
                        };

                        /**
                         * Decodes a Recipe message from the specified reader or buffer, length delimited.
                         * @function decodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @returns {lasagna.srv.internal.recipes.service.Recipe} Recipe
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        Recipe.decodeDelimited = function decodeDelimited(reader) {
                            if (!(reader instanceof $Reader))
                                reader = new $Reader(reader);
                            return this.decode(reader, reader.uint32());
                        };

                        /**
                         * Verifies a Recipe message.
                         * @function verify
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {Object.<string,*>} message Plain object to verify
                         * @returns {string|null} `null` if valid, otherwise the reason why it is not
                         */
                        Recipe.verify = function verify(message) {
                            if (typeof message !== "object" || message === null)
                                return "object expected";
                            if (message.id != null && message.hasOwnProperty("id"))
                                if (!$util.isString(message.id))
                                    return "id: string expected";
                            if (message.name != null && message.hasOwnProperty("name"))
                                if (!$util.isString(message.name))
                                    return "name: string expected";
                            if (message.description != null && message.hasOwnProperty("description"))
                                if (!$util.isString(message.description))
                                    return "description: string expected";
                            return null;
                        };

                        /**
                         * Creates a Recipe message from a plain object. Also converts values to their respective internal types.
                         * @function fromObject
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {Object.<string,*>} object Plain object
                         * @returns {lasagna.srv.internal.recipes.service.Recipe} Recipe
                         */
                        Recipe.fromObject = function fromObject(object) {
                            if (object instanceof $root.lasagna.srv.internal.recipes.service.Recipe)
                                return object;
                            var message = new $root.lasagna.srv.internal.recipes.service.Recipe();
                            if (object.id != null)
                                message.id = String(object.id);
                            if (object.name != null)
                                message.name = String(object.name);
                            if (object.description != null)
                                message.description = String(object.description);
                            return message;
                        };

                        /**
                         * Creates a plain object from a Recipe message. Also converts values to other types if specified.
                         * @function toObject
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.Recipe} message Recipe
                         * @param {$protobuf.IConversionOptions} [options] Conversion options
                         * @returns {Object.<string,*>} Plain object
                         */
                        Recipe.toObject = function toObject(message, options) {
                            if (!options)
                                options = {};
                            var object = {};
                            if (options.defaults) {
                                object.id = "";
                                object.name = "";
                                object.description = "";
                            }
                            if (message.id != null && message.hasOwnProperty("id"))
                                object.id = message.id;
                            if (message.name != null && message.hasOwnProperty("name"))
                                object.name = message.name;
                            if (message.description != null && message.hasOwnProperty("description"))
                                object.description = message.description;
                            return object;
                        };

                        /**
                         * Converts this Recipe to JSON.
                         * @function toJSON
                         * @memberof lasagna.srv.internal.recipes.service.Recipe
                         * @instance
                         * @returns {Object.<string,*>} JSON object
                         */
                        Recipe.prototype.toJSON = function toJSON() {
                            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                        };

                        return Recipe;
                    })();

                    service.GetRequest = (function() {

                        /**
                         * Properties of a GetRequest.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @interface IGetRequest
                         */

                        /**
                         * Constructs a new GetRequest.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @classdesc Represents a GetRequest.
                         * @implements IGetRequest
                         * @constructor
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest=} [properties] Properties to set
                         */
                        function GetRequest(properties) {
                            if (properties)
                                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                                    if (properties[keys[i]] != null)
                                        this[keys[i]] = properties[keys[i]];
                        }

                        /**
                         * Creates a new GetRequest instance using the specified properties.
                         * @function create
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest=} [properties] Properties to set
                         * @returns {lasagna.srv.internal.recipes.service.GetRequest} GetRequest instance
                         */
                        GetRequest.create = function create(properties) {
                            return new GetRequest(properties);
                        };

                        /**
                         * Encodes the specified GetRequest message. Does not implicitly {@link lasagna.srv.internal.recipes.service.GetRequest.verify|verify} messages.
                         * @function encode
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest} message GetRequest message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        GetRequest.encode = function encode(message, writer) {
                            if (!writer)
                                writer = $Writer.create();
                            return writer;
                        };

                        /**
                         * Encodes the specified GetRequest message, length delimited. Does not implicitly {@link lasagna.srv.internal.recipes.service.GetRequest.verify|verify} messages.
                         * @function encodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IGetRequest} message GetRequest message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        GetRequest.encodeDelimited = function encodeDelimited(message, writer) {
                            return this.encode(message, writer).ldelim();
                        };

                        /**
                         * Decodes a GetRequest message from the specified reader or buffer.
                         * @function decode
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @param {number} [length] Message length if known beforehand
                         * @returns {lasagna.srv.internal.recipes.service.GetRequest} GetRequest
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        GetRequest.decode = function decode(reader, length) {
                            if (!(reader instanceof $Reader))
                                reader = $Reader.create(reader);
                            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.lasagna.srv.internal.recipes.service.GetRequest();
                            while (reader.pos < end) {
                                var tag = reader.uint32();
                                switch (tag >>> 3) {
                                default:
                                    reader.skipType(tag & 7);
                                    break;
                                }
                            }
                            return message;
                        };

                        /**
                         * Decodes a GetRequest message from the specified reader or buffer, length delimited.
                         * @function decodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @returns {lasagna.srv.internal.recipes.service.GetRequest} GetRequest
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        GetRequest.decodeDelimited = function decodeDelimited(reader) {
                            if (!(reader instanceof $Reader))
                                reader = new $Reader(reader);
                            return this.decode(reader, reader.uint32());
                        };

                        /**
                         * Verifies a GetRequest message.
                         * @function verify
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {Object.<string,*>} message Plain object to verify
                         * @returns {string|null} `null` if valid, otherwise the reason why it is not
                         */
                        GetRequest.verify = function verify(message) {
                            if (typeof message !== "object" || message === null)
                                return "object expected";
                            return null;
                        };

                        /**
                         * Creates a GetRequest message from a plain object. Also converts values to their respective internal types.
                         * @function fromObject
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {Object.<string,*>} object Plain object
                         * @returns {lasagna.srv.internal.recipes.service.GetRequest} GetRequest
                         */
                        GetRequest.fromObject = function fromObject(object) {
                            if (object instanceof $root.lasagna.srv.internal.recipes.service.GetRequest)
                                return object;
                            return new $root.lasagna.srv.internal.recipes.service.GetRequest();
                        };

                        /**
                         * Creates a plain object from a GetRequest message. Also converts values to other types if specified.
                         * @function toObject
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.GetRequest} message GetRequest
                         * @param {$protobuf.IConversionOptions} [options] Conversion options
                         * @returns {Object.<string,*>} Plain object
                         */
                        GetRequest.toObject = function toObject() {
                            return {};
                        };

                        /**
                         * Converts this GetRequest to JSON.
                         * @function toJSON
                         * @memberof lasagna.srv.internal.recipes.service.GetRequest
                         * @instance
                         * @returns {Object.<string,*>} JSON object
                         */
                        GetRequest.prototype.toJSON = function toJSON() {
                            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                        };

                        return GetRequest;
                    })();

                    service.Response = (function() {

                        /**
                         * Properties of a Response.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @interface IResponse
                         * @property {boolean|null} [created] Response created
                         * @property {lasagna.srv.internal.recipes.service.IRecipe|null} [recipe] Response recipe
                         * @property {Array.<lasagna.srv.internal.recipes.service.IRecipe>|null} [recipes] Response recipes
                         */

                        /**
                         * Constructs a new Response.
                         * @memberof lasagna.srv.internal.recipes.service
                         * @classdesc Represents a Response.
                         * @implements IResponse
                         * @constructor
                         * @param {lasagna.srv.internal.recipes.service.IResponse=} [properties] Properties to set
                         */
                        function Response(properties) {
                            this.recipes = [];
                            if (properties)
                                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                                    if (properties[keys[i]] != null)
                                        this[keys[i]] = properties[keys[i]];
                        }

                        /**
                         * Response created.
                         * @member {boolean} created
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @instance
                         */
                        Response.prototype.created = false;

                        /**
                         * Response recipe.
                         * @member {lasagna.srv.internal.recipes.service.IRecipe|null|undefined} recipe
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @instance
                         */
                        Response.prototype.recipe = null;

                        /**
                         * Response recipes.
                         * @member {Array.<lasagna.srv.internal.recipes.service.IRecipe>} recipes
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @instance
                         */
                        Response.prototype.recipes = $util.emptyArray;

                        /**
                         * Creates a new Response instance using the specified properties.
                         * @function create
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IResponse=} [properties] Properties to set
                         * @returns {lasagna.srv.internal.recipes.service.Response} Response instance
                         */
                        Response.create = function create(properties) {
                            return new Response(properties);
                        };

                        /**
                         * Encodes the specified Response message. Does not implicitly {@link lasagna.srv.internal.recipes.service.Response.verify|verify} messages.
                         * @function encode
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IResponse} message Response message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        Response.encode = function encode(message, writer) {
                            if (!writer)
                                writer = $Writer.create();
                            if (message.created != null && message.hasOwnProperty("created"))
                                writer.uint32(/* id 1, wireType 0 =*/8).bool(message.created);
                            if (message.recipe != null && message.hasOwnProperty("recipe"))
                                $root.lasagna.srv.internal.recipes.service.Recipe.encode(message.recipe, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
                            if (message.recipes != null && message.recipes.length)
                                for (var i = 0; i < message.recipes.length; ++i)
                                    $root.lasagna.srv.internal.recipes.service.Recipe.encode(message.recipes[i], writer.uint32(/* id 3, wireType 2 =*/26).fork()).ldelim();
                            return writer;
                        };

                        /**
                         * Encodes the specified Response message, length delimited. Does not implicitly {@link lasagna.srv.internal.recipes.service.Response.verify|verify} messages.
                         * @function encodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.IResponse} message Response message or plain object to encode
                         * @param {$protobuf.Writer} [writer] Writer to encode to
                         * @returns {$protobuf.Writer} Writer
                         */
                        Response.encodeDelimited = function encodeDelimited(message, writer) {
                            return this.encode(message, writer).ldelim();
                        };

                        /**
                         * Decodes a Response message from the specified reader or buffer.
                         * @function decode
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @param {number} [length] Message length if known beforehand
                         * @returns {lasagna.srv.internal.recipes.service.Response} Response
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        Response.decode = function decode(reader, length) {
                            if (!(reader instanceof $Reader))
                                reader = $Reader.create(reader);
                            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.lasagna.srv.internal.recipes.service.Response();
                            while (reader.pos < end) {
                                var tag = reader.uint32();
                                switch (tag >>> 3) {
                                case 1:
                                    message.created = reader.bool();
                                    break;
                                case 2:
                                    message.recipe = $root.lasagna.srv.internal.recipes.service.Recipe.decode(reader, reader.uint32());
                                    break;
                                case 3:
                                    if (!(message.recipes && message.recipes.length))
                                        message.recipes = [];
                                    message.recipes.push($root.lasagna.srv.internal.recipes.service.Recipe.decode(reader, reader.uint32()));
                                    break;
                                default:
                                    reader.skipType(tag & 7);
                                    break;
                                }
                            }
                            return message;
                        };

                        /**
                         * Decodes a Response message from the specified reader or buffer, length delimited.
                         * @function decodeDelimited
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
                         * @returns {lasagna.srv.internal.recipes.service.Response} Response
                         * @throws {Error} If the payload is not a reader or valid buffer
                         * @throws {$protobuf.util.ProtocolError} If required fields are missing
                         */
                        Response.decodeDelimited = function decodeDelimited(reader) {
                            if (!(reader instanceof $Reader))
                                reader = new $Reader(reader);
                            return this.decode(reader, reader.uint32());
                        };

                        /**
                         * Verifies a Response message.
                         * @function verify
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {Object.<string,*>} message Plain object to verify
                         * @returns {string|null} `null` if valid, otherwise the reason why it is not
                         */
                        Response.verify = function verify(message) {
                            if (typeof message !== "object" || message === null)
                                return "object expected";
                            if (message.created != null && message.hasOwnProperty("created"))
                                if (typeof message.created !== "boolean")
                                    return "created: boolean expected";
                            if (message.recipe != null && message.hasOwnProperty("recipe")) {
                                var error = $root.lasagna.srv.internal.recipes.service.Recipe.verify(message.recipe);
                                if (error)
                                    return "recipe." + error;
                            }
                            if (message.recipes != null && message.hasOwnProperty("recipes")) {
                                if (!Array.isArray(message.recipes))
                                    return "recipes: array expected";
                                for (var i = 0; i < message.recipes.length; ++i) {
                                    var error = $root.lasagna.srv.internal.recipes.service.Recipe.verify(message.recipes[i]);
                                    if (error)
                                        return "recipes." + error;
                                }
                            }
                            return null;
                        };

                        /**
                         * Creates a Response message from a plain object. Also converts values to their respective internal types.
                         * @function fromObject
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {Object.<string,*>} object Plain object
                         * @returns {lasagna.srv.internal.recipes.service.Response} Response
                         */
                        Response.fromObject = function fromObject(object) {
                            if (object instanceof $root.lasagna.srv.internal.recipes.service.Response)
                                return object;
                            var message = new $root.lasagna.srv.internal.recipes.service.Response();
                            if (object.created != null)
                                message.created = Boolean(object.created);
                            if (object.recipe != null) {
                                if (typeof object.recipe !== "object")
                                    throw TypeError(".lasagna.srv.internal.recipes.service.Response.recipe: object expected");
                                message.recipe = $root.lasagna.srv.internal.recipes.service.Recipe.fromObject(object.recipe);
                            }
                            if (object.recipes) {
                                if (!Array.isArray(object.recipes))
                                    throw TypeError(".lasagna.srv.internal.recipes.service.Response.recipes: array expected");
                                message.recipes = [];
                                for (var i = 0; i < object.recipes.length; ++i) {
                                    if (typeof object.recipes[i] !== "object")
                                        throw TypeError(".lasagna.srv.internal.recipes.service.Response.recipes: object expected");
                                    message.recipes[i] = $root.lasagna.srv.internal.recipes.service.Recipe.fromObject(object.recipes[i]);
                                }
                            }
                            return message;
                        };

                        /**
                         * Creates a plain object from a Response message. Also converts values to other types if specified.
                         * @function toObject
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @static
                         * @param {lasagna.srv.internal.recipes.service.Response} message Response
                         * @param {$protobuf.IConversionOptions} [options] Conversion options
                         * @returns {Object.<string,*>} Plain object
                         */
                        Response.toObject = function toObject(message, options) {
                            if (!options)
                                options = {};
                            var object = {};
                            if (options.arrays || options.defaults)
                                object.recipes = [];
                            if (options.defaults) {
                                object.created = false;
                                object.recipe = null;
                            }
                            if (message.created != null && message.hasOwnProperty("created"))
                                object.created = message.created;
                            if (message.recipe != null && message.hasOwnProperty("recipe"))
                                object.recipe = $root.lasagna.srv.internal.recipes.service.Recipe.toObject(message.recipe, options);
                            if (message.recipes && message.recipes.length) {
                                object.recipes = [];
                                for (var j = 0; j < message.recipes.length; ++j)
                                    object.recipes[j] = $root.lasagna.srv.internal.recipes.service.Recipe.toObject(message.recipes[j], options);
                            }
                            return object;
                        };

                        /**
                         * Converts this Response to JSON.
                         * @function toJSON
                         * @memberof lasagna.srv.internal.recipes.service.Response
                         * @instance
                         * @returns {Object.<string,*>} JSON object
                         */
                        Response.prototype.toJSON = function toJSON() {
                            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
                        };

                        return Response;
                    })();

                    return service;
                })();

                return recipes;
            })();

            return internal;
        })();

        return srv;
    })();

    return lasagna;
})();

module.exports = $root;
