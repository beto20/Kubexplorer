// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {model} from '../models';

export function DeletePersistentVolumeByName(arg1:string):Promise<void>;

export function DeletePersistentVolumeClaim(arg1:string,arg2:string):Promise<void>;

export function GetPersistentVolumeByName(arg1:string):Promise<model.PersistentVolumeDto>;

export function GetPersistentVolumeClaim(arg1:string,arg2:string):Promise<model.PersistentVolumeClaimDto>;

export function GetPersistentVolumes():Promise<Array<model.PersistentVolumeDto>>;

export function GetPersistentVolumesClaim(arg1:string):Promise<Array<model.PersistentVolumeClaimDto>>;

export function UpdatePersistentVolumeByName(arg1:string,arg2:model.PersistentVolumeDto):Promise<void>;

export function UpdatePersistentVolumeClaim(arg1:string,arg2:string,arg3:model.PersistentVolumeClaimDto):Promise<void>;
