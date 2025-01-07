import { AboutJsonParse } from "./aboutJson";

export type AuthParamList = {
  Login: undefined;
  Register: undefined;
  Home: undefined;
};

type ActionOrReactionProps = {
  isAction: boolean;
}

export type AppStackList = {
  App: undefined;
  ActionOrReaction: ActionOrReactionProps;
  Home: undefined;
  Dashboard: undefined;
  Service: undefined;
  Auth: undefined;
};
