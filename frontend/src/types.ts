export interface Service {
  name: string;
  actions: any[];
  reactions: any;
}

export interface ServerResponse {
  server: {
    services: Service[];
  };
}
