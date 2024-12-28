export interface Environment {
    name: string;
    env: string;
    status: boolean;
    options: MenuItems[];
}

export interface MenuItems {
    name: string;
    link: string;
    icon: string;
}
