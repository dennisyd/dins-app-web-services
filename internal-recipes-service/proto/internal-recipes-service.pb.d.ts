import * as $protobuf from "protobufjs";
export namespace lasagna {

    namespace srv {

        namespace internal {

            namespace recipes {

                namespace service {

                    class InternalRecipesService extends $protobuf.rpc.Service {
                        constructor(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean);
                        public static create(rpcImpl: $protobuf.RPCImpl, requestDelimited?: boolean, responseDelimited?: boolean): InternalRecipesService;
                        public createRecipe(request: lasagna.srv.internal.recipes.service.IRecipe, callback: lasagna.srv.internal.recipes.service.InternalRecipesService.CreateRecipeCallback): void;
                        public createRecipe(request: lasagna.srv.internal.recipes.service.IRecipe): Promise<lasagna.srv.internal.recipes.service.Response>;
                        public getRecipes(request: lasagna.srv.internal.recipes.service.IGetRequest, callback: lasagna.srv.internal.recipes.service.InternalRecipesService.GetRecipesCallback): void;
                        public getRecipes(request: lasagna.srv.internal.recipes.service.IGetRequest): Promise<lasagna.srv.internal.recipes.service.Response>;
                    }

                    namespace InternalRecipesService {

                        type CreateRecipeCallback = (error: (Error|null), response?: lasagna.srv.internal.recipes.service.Response) => void;

                        type GetRecipesCallback = (error: (Error|null), response?: lasagna.srv.internal.recipes.service.Response) => void;
                    }

                    interface IRecipe {
                        id?: (string|null);
                        name?: (string|null);
                        description?: (string|null);
                    }

                    class Recipe implements IRecipe {
                        constructor(properties?: lasagna.srv.internal.recipes.service.IRecipe);
                        public id: string;
                        public name: string;
                        public description: string;
                        public static create(properties?: lasagna.srv.internal.recipes.service.IRecipe): lasagna.srv.internal.recipes.service.Recipe;
                        public static encode(message: lasagna.srv.internal.recipes.service.IRecipe, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static encodeDelimited(message: lasagna.srv.internal.recipes.service.IRecipe, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): lasagna.srv.internal.recipes.service.Recipe;
                        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): lasagna.srv.internal.recipes.service.Recipe;
                        public static verify(message: { [k: string]: any }): (string|null);
                        public static fromObject(object: { [k: string]: any }): lasagna.srv.internal.recipes.service.Recipe;
                        public static toObject(message: lasagna.srv.internal.recipes.service.Recipe, options?: $protobuf.IConversionOptions): { [k: string]: any };
                        public toJSON(): { [k: string]: any };
                    }

                    interface IGetRequest {
                    }

                    class GetRequest implements IGetRequest {
                        constructor(properties?: lasagna.srv.internal.recipes.service.IGetRequest);
                        public static create(properties?: lasagna.srv.internal.recipes.service.IGetRequest): lasagna.srv.internal.recipes.service.GetRequest;
                        public static encode(message: lasagna.srv.internal.recipes.service.IGetRequest, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static encodeDelimited(message: lasagna.srv.internal.recipes.service.IGetRequest, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): lasagna.srv.internal.recipes.service.GetRequest;
                        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): lasagna.srv.internal.recipes.service.GetRequest;
                        public static verify(message: { [k: string]: any }): (string|null);
                        public static fromObject(object: { [k: string]: any }): lasagna.srv.internal.recipes.service.GetRequest;
                        public static toObject(message: lasagna.srv.internal.recipes.service.GetRequest, options?: $protobuf.IConversionOptions): { [k: string]: any };
                        public toJSON(): { [k: string]: any };
                    }

                    interface IResponse {
                        created?: (boolean|null);
                        recipe?: (lasagna.srv.internal.recipes.service.IRecipe|null);
                        recipes?: (lasagna.srv.internal.recipes.service.IRecipe[]|null);
                    }

                    class Response implements IResponse {
                        constructor(properties?: lasagna.srv.internal.recipes.service.IResponse);
                        public created: boolean;
                        public recipe?: (lasagna.srv.internal.recipes.service.IRecipe|null);
                        public recipes: lasagna.srv.internal.recipes.service.IRecipe[];
                        public static create(properties?: lasagna.srv.internal.recipes.service.IResponse): lasagna.srv.internal.recipes.service.Response;
                        public static encode(message: lasagna.srv.internal.recipes.service.IResponse, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static encodeDelimited(message: lasagna.srv.internal.recipes.service.IResponse, writer?: $protobuf.Writer): $protobuf.Writer;
                        public static decode(reader: ($protobuf.Reader|Uint8Array), length?: number): lasagna.srv.internal.recipes.service.Response;
                        public static decodeDelimited(reader: ($protobuf.Reader|Uint8Array)): lasagna.srv.internal.recipes.service.Response;
                        public static verify(message: { [k: string]: any }): (string|null);
                        public static fromObject(object: { [k: string]: any }): lasagna.srv.internal.recipes.service.Response;
                        public static toObject(message: lasagna.srv.internal.recipes.service.Response, options?: $protobuf.IConversionOptions): { [k: string]: any };
                        public toJSON(): { [k: string]: any };
                    }
                }
            }
        }
    }
}
