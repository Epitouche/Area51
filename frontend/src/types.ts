export type WorkflowResponse = {
  body: string;
  pull_request_url: string;
};

export type Workflow = {
  name: string;
  action_id: number;
  action_name: string;
  reaction_id: number;
  reaction_name: string;
  is_active: boolean;
  created_at: string;
  checked?: boolean;
  workflow_id: number;
};

export type Action = {
  name: string;
  action_id: number;
  description: string;
  options: string;
};

export type Reaction = {
  name: string;
  reaction_id: number;
  description: string;
  options: string;
};

export type AboutResponse = {
  server: {
    services: [
      {
        name: string;
        description: string;
        actions: Action[];
        reactions: Reaction[];
        image: string;
        is_oauth: boolean;
      }
    ]
  }
};

export type Service = {
  name: string;
  created_at: string;
  description: string;
  id: number;
  image: string;
  updated_at: string;
  actions: Action[] | null;
  reactions: Reaction[] | null;
};

export type OptionWorkflow = {
  name: string;
  input: string;
  type: string;
};