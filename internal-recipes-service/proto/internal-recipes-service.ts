
import {createTwirpRequest, throwTwirpError, Fetch} from './twirp';


export interface Recipe {
    id: string;
    name: string;
    description: string;
    
}

interface RecipeJSON {
    id: string;
    name: string;
    description: string;
    
}


const RecipeToJSON = (m: Recipe): RecipeJSON => {
    return {
        id: m.id,
        name: m.name,
        description: m.description,
        
    };
};

const JSONToRecipe = (m: Recipe | RecipeJSON): Recipe => {
    
    return {
        id: m.id,
        name: m.name,
        description: m.description,
        
    };
};

export interface GetRequest {
    
}

interface GetRequestJSON {
    
}


const GetRequestToJSON = (m: GetRequest): GetRequestJSON => {
    return {
        
    };
};

export interface Response {
    created: boolean;
    recipe: Recipe;
    recipes: Recipe[];
    
}

interface ResponseJSON {
    created: boolean;
    recipe: RecipeJSON;
    recipes: RecipeJSON[];
    
}


const JSONToResponse = (m: Response | ResponseJSON): Response => {
    
    return {
        created: m.created,
        recipe: JSONToRecipe(m.recipe),
        recipes: (m.recipes as (Recipe | RecipeJSON)[]).map(JSONToRecipe),
        
    };
};

export interface InternalRecipesService {
    createRecipe: (recipe: Recipe) => Promise<Response>;
    
    getRecipes: (getRequest: GetRequest) => Promise<Response>;
    
}

export class DefaultInternalRecipesService implements InternalRecipesService {
    private hostname: string;
    private fetch: Fetch;
    private writeCamelCase: boolean;
    private pathPrefix = "/lasagna.srv.internal.recipes.service.InternalRecipesService/";

    constructor(hostname: string, fetch: Fetch, writeCamelCase = false) {
        this.hostname = hostname;
        this.fetch = fetch;
        this.writeCamelCase = writeCamelCase;
    }
    createRecipe(recipe: Recipe): Promise<Response> {
        const url = this.hostname + this.pathPrefix + "CreateRecipe";
        let body: Recipe | RecipeJSON = recipe;
        if (!this.writeCamelCase) {
            body = RecipeToJSON(recipe);
        }
        return this.fetch(createTwirpRequest(url, body)).then((resp) => {
            if (!resp.ok) {
                return throwTwirpError(resp);
            }

            return resp.json().then(JSONToResponse);
        });
    }
    
    getRecipes(getRequest: GetRequest): Promise<Response> {
        const url = this.hostname + this.pathPrefix + "GetRecipes";
        let body: GetRequest | GetRequestJSON = getRequest;
        if (!this.writeCamelCase) {
            body = GetRequestToJSON(getRequest);
        }
        return this.fetch(createTwirpRequest(url, body)).then((resp) => {
            if (!resp.ok) {
                return throwTwirpError(resp);
            }

            return resp.json().then(JSONToResponse);
        });
    }
    
}

