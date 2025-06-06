export interface Menu {
  name: string
  link: string
  icon: string
}

export interface Environment {
  name: string
  description: string
  env: string
  status: boolean
}

export interface Setting {
  name: string
  link: string
  icon: string
}

export interface EnvironmentLayout {
  name: string
  description: string
  env: string
  status: boolean
  options: Menu[]
}
