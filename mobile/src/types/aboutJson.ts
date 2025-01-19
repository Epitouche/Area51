import { ConnectedService } from './servicesModals';

export type ServiceAction = {
  name: string;
  description: string;
  options: { [key: string]: any };
  action_id: number;
};

export type ServiceReaction = {
  name: string;
  description: string;
  options: { [key: string]: any };
  reaction_id: number;
};

export type Service = {
  name: string;
  description: string;
  actions: ServiceAction[] | null;
  reactions: ServiceReaction[] | null;
  image: string;
  is_oauth: boolean;
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

export type ActionParse = {
  action_id: number;
  name: string;
  description: string;
  options: { [key: string]: any };
};

export type ReactionParse = {
  reaction_id: number;
  name: string;
  description: string;
  options: { [key: string]: any };
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
  description: string;
  actions: ActionParse[] | null;
  reactions: ReactionParse[] | null;
  image: string;
  is_oauth: boolean;
  isConnected: boolean;
};

export type AboutJsonParse = {
  services: ServicesParse[];
};

export interface GetConnectedServiceProps {
  apiEndpoint: string;
  token: string;
  setConnectedService: (connectedService: ConnectedService[]) => void;
}

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
