import { AboutJsonParse, Action, Reaction, Workflow } from "./aboutJson";

export type AuthParamList = {
  Login: undefined;
  Register: undefined;
  Home: undefined;
};

type OptionsProps = {
  isAction: boolean;
  setAction?: (isAction: Action) => void;
  setReaction?: (isReaction: Reaction) => void;
}

type WorkflowDetailsProps = {
  workflow: Workflow;
};

export type AppStackList = {
  App: undefined;
  Options: OptionsProps;
  Home: undefined;
  Workflows: undefined;
  Service: undefined;
  Auth: undefined;
  Parameters: undefined;
  'Workflow Details': WorkflowDetailsProps;
};
