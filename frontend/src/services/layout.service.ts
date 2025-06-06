import {
  GetCommonParameters,
  GetHeaderParams,
  GetKubernetesParameters,
} from '../../wailsjs/go/middleware/ParameterMiddleware'
import { GetAllEnvironment } from '../../wailsjs/go/middleware/EnvironmentMiddleware'

export const fetchCommonParameters = async () => GetCommonParameters()
export const fetchKubernetesParameters = async () => GetKubernetesParameters()
export const fetchAllEnvironments = async () => GetAllEnvironment()
export const fetchHeaderParams = async (k8sObject: string) => GetHeaderParams(k8sObject)
