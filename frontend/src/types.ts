export type Workflow = {
  body: string;
  pull_request_url: string;
}

export interface Action {
  name: string;
  action_id: number;
  description: string;
}

export interface Reaction {
  name: string;
  reaction_id: number;
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
