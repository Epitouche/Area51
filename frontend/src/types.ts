export interface Action {
  name: string;
  description: string;
}

export interface Reaction {
  name: string;
  description: string;
}

export interface Service {
  name: string;
  actions: Action[];
  reactions: Reaction[];
}

export interface ServerResponse {
  server: {
    services: Service[];
  };
}
