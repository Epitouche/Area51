import { ConnectedService } from "./servicesModals";

export type Action = {
  action_id: number;
  name: string;
  description: string;
};

export type Reaction = {
  reaction_id: number;
  name: string;
  description: string;
};

export type Service = {
  name: string;
  description: string;
  actions: Action[];
  reactions: Reaction[];
  image: string;
};

export type Server = {
  current_time: string;
  services: Service[];
};

export type Client = {
  host: string;
};

export type AboutJson = {
  client: Client;
  server: Server;
};

export type Workflow = {
  action_id: number;
  action_name: string;
  created_at: string;
  is_active: boolean;
  name: string;
  reaction_id: number;
  reaction_name: string;
  workflow_id: number;
};

export type PullRequestComment = {
  body: string;
  pull_request_url: string;
};


export type ServicesParse = {
  name: string;
  isConnected: boolean;
  actions: Action[];
  reactions: Reaction[];
}

export type AboutJsonParse = {
  services: ServicesParse[];
};

export interface GetConnectedServiceProps {
  apiEndpoint: string;
  token: string;
  setConnectedService: (connectedService: ConnectedService[]) => void;
};

export interface ParseConnectedServicesProps {
  aboutjson: AboutJson;
  apiEndpoint: string;
  token: string;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
}

export interface ParseServicesProps {
  aboutJson: AboutJson;
  serverIp: string;
  setServicesConnected: (servicesConnected: AboutJsonParse) => void;
}
