import { AboutJsonParse, Action, Reaction } from "./aboutJson";

export type AuthParamList = {
  Login: undefined;
  Register: undefined;
  Home: undefined;
};

type ActionOrReactionProps = {
  isAction: boolean;
  setAction?: (isAction: Action) => void;
  setReaction?: (isReaction: Reaction) => void;
}

export type AppStackList = {
  App: undefined;
  ActionOrReaction: ActionOrReactionProps;
  Home: undefined;
  Workflows: undefined;
  Service: undefined;
  Auth: undefined;
};
