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

export namespace model {
	
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
	export class Host {
	    Path: string;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Host(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Type = source["Type"];
	    }
	}
	export class RuleDto {
	    Host: string;
	    Path: string;
	
	    static createFrom(source: any = {}) {
	        return new RuleDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Host = source["Host"];
	        this.Path = source["Path"];
	    }
	}
	export class IngressDto {
	    Name: string;
	    Namespace: string;
	    Rules: RuleDto[];
	    Creation: string;
	    Labels: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new IngressDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.Rules = this.convertValues(source["Rules"], RuleDto);
	        this.Creation = source["Creation"];
	        this.Labels = source["Labels"];
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
	export class Local {
	    Path: string;
	    FSType: string;
	
	    static createFrom(source: any = {}) {
	        return new Local(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.FSType = source["FSType"];
	    }
	}
	export class NFS {
	    Path: string;
	    Server: string;
	
	    static createFrom(source: any = {}) {
	        return new NFS(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Path = source["Path"];
	        this.Server = source["Server"];
	    }
	}
	export class NamespaceDto {
	    Name: string;
	    Version: string;
	    CreationTime: string;
	    Labels: Record<string, string>;
	    Status: string;
	
	    static createFrom(source: any = {}) {
	        return new NamespaceDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Version = source["Version"];
	        this.CreationTime = source["CreationTime"];
	        this.Labels = source["Labels"];
	        this.Status = source["Status"];
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
	export class NodeDtoV2 {
	    Name: string;
	    Namespace: string;
	    Resource: Resource;
	    Version: string;
	    CreationTimestamp: string;
	    Labels: Record<string, string>;
	
	    static createFrom(source: any = {}) {
	        return new NodeDtoV2(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.Resource = this.convertValues(source["Resource"], Resource);
	        this.Version = source["Version"];
	        this.CreationTimestamp = source["CreationTimestamp"];
	        this.Labels = source["Labels"];
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
	export class VolumeClaimSpec {
	    VolumeName: string;
	    VolumeMode: string;
	    AccessModes: string[];
	    DataSourceName: string;
	    StorageClass: string;
	    VolumeAttributesClassName: string;
	    Limit: Resource;
	    Request: Resource;
	
	    static createFrom(source: any = {}) {
	        return new VolumeClaimSpec(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.VolumeName = source["VolumeName"];
	        this.VolumeMode = source["VolumeMode"];
	        this.AccessModes = source["AccessModes"];
	        this.DataSourceName = source["DataSourceName"];
	        this.StorageClass = source["StorageClass"];
	        this.VolumeAttributesClassName = source["VolumeAttributesClassName"];
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
	export class PersistentVolumeClaimDto {
	    Name: string;
	    Namespace: string;
	    CreationTimestamp: string;
	    Labels: Record<string, string>;
	    VolumeClaimSpec: VolumeClaimSpec;
	
	    static createFrom(source: any = {}) {
	        return new PersistentVolumeClaimDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.CreationTimestamp = source["CreationTimestamp"];
	        this.Labels = source["Labels"];
	        this.VolumeClaimSpec = this.convertValues(source["VolumeClaimSpec"], VolumeClaimSpec);
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
	export class VolumeSpec {
	    Local: Local;
	    VolumeMode: string;
	    AccessModes: string[];
	    StorageClass: string;
	    VolumeAttributesClassName: string;
	    PersistentVolumeReclaimPolicy: string;
	    MountOptions: string[];
	    Capacity: Resource;
	    Host: Host;
	    NFS: NFS;
	
	    static createFrom(source: any = {}) {
	        return new VolumeSpec(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Local = this.convertValues(source["Local"], Local);
	        this.VolumeMode = source["VolumeMode"];
	        this.AccessModes = source["AccessModes"];
	        this.StorageClass = source["StorageClass"];
	        this.VolumeAttributesClassName = source["VolumeAttributesClassName"];
	        this.PersistentVolumeReclaimPolicy = source["PersistentVolumeReclaimPolicy"];
	        this.MountOptions = source["MountOptions"];
	        this.Capacity = this.convertValues(source["Capacity"], Resource);
	        this.Host = this.convertValues(source["Host"], Host);
	        this.NFS = this.convertValues(source["NFS"], NFS);
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
	    Name: string;
	    Namespace: string;
	    CreationTimestamp: string;
	    Labels: Record<string, string>;
	    VolumeSpec: VolumeSpec;
	
	    static createFrom(source: any = {}) {
	        return new PersistentVolumeDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.CreationTimestamp = source["CreationTimestamp"];
	        this.Labels = source["Labels"];
	        this.VolumeSpec = this.convertValues(source["VolumeSpec"], VolumeSpec);
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
	    Name: string;
	    Namespace: string;
	    Labels: Record<string, string>;
	    Status: string;
	    CreationTimestamp: string;
	    Spec: string;
	
	    static createFrom(source: any = {}) {
	        return new ServiceDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.Namespace = source["Namespace"];
	        this.Labels = source["Labels"];
	        this.Status = source["Status"];
	        this.CreationTimestamp = source["CreationTimestamp"];
	        this.Spec = source["Spec"];
	    }
	}
	

}

