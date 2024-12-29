export {};

// TODO: Se agrego para Wails events
declare global {
    interface Window {
        runtime: {
            EventsOn(event: string, callback: (data: any) => void): void;
            EventsEmit(event: string, data?: any): void;
            Call(method: string, ...args: any[]): Promise<any>;
        };
    }
}
