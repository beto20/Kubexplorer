export namespace database {
	
	export class CommonParameterDto {
	    Name: string;
	    Link: string;
	    Icon: string;
	
	    static createFrom(source: any = {}) {
	        return new CommonParameterDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Link = source["Link"];
	        this.Icon = source["Icon"];
	    }
	}
	export class HeadParamsDto {
	    Title: string;
	    Key: string;
	    Align: string;
	    Sortable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new HeadParamsDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Title = source["Title"];
	        this.Key = source["Key"];
	        this.Align = source["Align"];
	        this.Sortable = source["Sortable"];
	    }
	}
	export class K8sObject {
	    Name: string;
	    Link: string;
	    IsVisible: boolean;
	    IsEditable: boolean;
	
	    static createFrom(source: any = {}) {
	        return new K8sObject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Link = source["Link"];
	        this.IsVisible = source["IsVisible"];
	        this.IsEditable = source["IsEditable"];
	    }
	}
	export class ObjectType {
	    Name: string;
	    IsVisible: boolean;
	    IsEditable: boolean;
	    K8sObject: K8sObject[];
	
	    static createFrom(source: any = {}) {
	        return new ObjectType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.IsVisible = source["IsVisible"];
	        this.IsEditable = source["IsEditable"];
	        this.K8sObject = this.convertValues(source["K8sObject"], K8sObject);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace objects {
	
	export class Resource {
	    Cpu: string;
	    Memory: string;
	    Storage: string;
	    StorageEphemeral: string;
	
	    static createFrom(source: any = {}) {
	        return new Resource(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Cpu = source["Cpu"];
	        this.Memory = source["Memory"];
	        this.Storage = source["Storage"];
	        this.StorageEphemeral = source["StorageEphemeral"];
	    }
	}
	export class Container {
	    Limit: Resource;
	    Request: Resource;
	
	    static createFrom(source: any = {}) {
	        return new Container(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Limit = this.convertValues(source["Limit"], Resource);
	        this.Request = this.convertValues(source["Request"], Resource);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class DeploymentDto {
	    Name: string;
	    Namespace: string;
	    Status: string;
	    Age: string;
	
	    static createFrom(source: any = {}) {
	        return new DeploymentDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.Status = source["Status"];
	        this.Age = source["Age"];
	    }
	}
	export class EnvironmentDto {
	    Name: string;
	    Description: string;
	    Env: string;
	    Status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new EnvironmentDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Description = source["Description"];
	        this.Env = source["Env"];
	        this.Status = source["Status"];
	    }
	}
	export class IngressDto {
	
	
	    static createFrom(source: any = {}) {
	        return new IngressDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class NamespaceDto {
	    Name: string;
	    Resource: Resource;
	    Roles: string[];
	    Version: string;
	    Age: string;
	    Status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NamespaceDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Resource = this.convertValues(source["Resource"], Resource);
	        this.Roles = source["Roles"];
	        this.Version = source["Version"];
	        this.Age = source["Age"];
	        this.Status = source["Status"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class NodeDto {
	    Name: string;
	    Resource: Resource;
	    Roles: string[];
	    Version: string;
	    Age: string;
	    Status: boolean;
	
	    static createFrom(source: any = {}) {
	        return new NodeDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Resource = this.convertValues(source["Resource"], Resource);
	        this.Roles = source["Roles"];
	        this.Version = source["Version"];
	        this.Age = source["Age"];
	        this.Status = source["Status"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class PersistentVolumeDto {
	
	
	    static createFrom(source: any = {}) {
	        return new PersistentVolumeDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class PodDto {
	    Name: string;
	    Namespace: string;
	    Replicas: number;
	    Container: Container;
	    Age: string;
	    Status: string;
	
	    static createFrom(source: any = {}) {
	        return new PodDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.Replicas = source["Replicas"];
	        this.Container = this.convertValues(source["Container"], Container);
	        this.Age = source["Age"];
	        this.Status = source["Status"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	export class ServiceDto {
	
	
	    static createFrom(source: any = {}) {
	        return new ServiceDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}

}

