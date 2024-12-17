export type WorkflowResponse = {
  body: string;
  pull_request_url: string;
};

export type Workflow = {
  name: string;
  action_id: number;
  reaction_id: string;
  is_active: boolean;
  created_at: string;
};

export type Action = {
  name: string;
  action_id: number;
  description: string;
};

export type Reaction = {
  name: string;
  reaction_id: number;
  description: string;
};

export type Service = {
  name: string;
  actions: Action[];
  reactions: Reaction[];
};

export type ServerResponse = {
  server: {
    services: Service[];
    workflows: Workflow[];
  };
};
